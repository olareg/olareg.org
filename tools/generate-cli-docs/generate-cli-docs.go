package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args
	// TODO(bmitch): handle flags
	if len(args) <= 1 {
		fmt.Fprintf(os.Stderr, "usage: %s [cmd]\n", args[0])
		os.Exit(1)
	}
	err := genDocs(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}

//go:embed docs.tmpl
var format string

func genDocs(cliName string) error {
	cmd := exec.Command(cliName, "cli-doc", "--list")
	cmdOut, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to list commands for %s: %v", cliName, err)
	}
	cliList := strings.Split(string(cmdOut), "\n")
	for _, cliCmd := range cliList {
		cliCmd = strings.TrimSpace(cliCmd)
		if cliCmd == "" {
			continue
		}
		fmt.Printf("Processing: %s\n", cliCmd)
		cliSplit := strings.Split(cliCmd, " ")[1:]
		docArgs := []string{"cli-doc", "--format", format}
		docArgs = append(docArgs, cliSplit...)
		cmd = exec.Command(cliName, docArgs...)
		out, err := cmd.Output()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				return fmt.Errorf("failed to generate output for %s %v: %s", cliName, docArgs, exitErr.Stderr)
			}
			return fmt.Errorf("failed to generate output for %s %v: %v", cliName, docArgs, err)
		}
		cliPath := filepath.Join(cliSplit...)
		if len(cliSplit) > 0 {
			err = os.MkdirAll(cliPath, 0755)
			if err != nil {
				return fmt.Errorf("failed to create directory %s: %v", cliPath, err)
			}
		}
		cliPath = filepath.Join(cliPath, "_index.md")
		fp, err := os.Create(cliPath)
		if err != nil {
			return fmt.Errorf("failed to create %s: %v", cliPath, err)
		}
		_, err = fp.Write(out)
		if err != nil {
			_ = fp.Close()
			return fmt.Errorf("failed to write %s: %v", cliPath, err)
		}
		err = fp.Close()
		if err != nil {
			return fmt.Errorf("failed to close %s: %v", cliPath, err)
		}
	}
	return nil
}
