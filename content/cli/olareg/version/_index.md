---
title: olareg version
layout: single
warning: Auto generated content
---

## Synopsis

Show the version of olareg

```shell
olareg version [flags]
```

## Examples

```shell
# display full version details
olareg version

# retrieve the version number
olareg version --format '{{.VCSTag}}'
```

## Options

```text
      --format string   Format output with go template syntax (default "{{printPretty .}}")
```

## Options from parent commands

```text
  -v, --verbosity string   Log level (trace, debug, info, warn, error) (default "warn")
```
