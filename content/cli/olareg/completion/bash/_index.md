---
title: olareg completion bash
layout: single
warning: Auto generated content
---

## Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(olareg completion bash)

To load completions for every new session, execute once:

#### Linux:

	olareg completion bash > /etc/bash_completion.d/olareg

#### macOS:

	olareg completion bash > $(brew --prefix)/etc/bash_completion.d/olareg

You will need to start a new shell for this setup to take effect.

```shell
olareg completion bash
```

## Options

```text
      --no-descriptions   disable completion descriptions
```

## Options from parent commands

```text
  -v, --verbosity string   Log level (trace, debug, info, warn, error) (default "warn")
```
