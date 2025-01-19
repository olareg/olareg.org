---
title: olareg completion zsh
layout: single
warning: Auto generated content
---

## Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(olareg completion zsh)

To load completions for every new session, execute once:

#### Linux:

	olareg completion zsh > "${fpath[1]}/_olareg"

#### macOS:

	olareg completion zsh > $(brew --prefix)/share/zsh/site-functions/_olareg

You will need to start a new shell for this setup to take effect.

```shell
olareg completion zsh [flags]
```

## Options

```text
      --no-descriptions   disable completion descriptions
```

## Options from parent commands

```text
  -v, --verbosity string   Log level (trace, debug, info, warn, error) (default "warn")
```
