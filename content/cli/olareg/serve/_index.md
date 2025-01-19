---
title: olareg serve
layout: single
warning: Auto generated content
---

## Synopsis

Run a registry server

```shell
olareg serve [flags]
```

## Examples

```shell
# run a server listening on localhost, port 5000, serving content from the current directory
olareg serve --addr 127.0.0.1

# run a read-only server on port 5050 serving content from the mirror directory
olareg serve --port 5050 --store-ro --dir mirror/

# run an ephemeral server from memory
olareg serve --store-type mem

# disable garbage collection
olareg serve --gc-frequency -1

# run an HTTPS server
olareg serve --tls-cert host.pem --tls-key host.key --port 443
```

## Options

```text
      --addr string                listener interface or address
      --api-blob-delete            enable blob delete API
      --api-delete                 enable delete APIs
      --api-push                   enable push APIs (default true)
      --api-referrer               enable referrer API (default true)
      --dir string                 root directory for storage (default ".")
      --gc-frequency duration      garbage collection frequency (default 15m0s)
      --gc-grace-period duration   garbage collection grace period (default 1h0m0s)
      --gc-referrer-dangling       garbage collect dangling referrers
      --gc-referrer-subject        garbage collect referrers when subject is deleted (default true)
      --gc-untagged                garbage collect untagged manifests
      --port int                   listener port (default 5000)
      --rate-limit int             limit requests per second per source IP
      --store-ro                   restrict storage as read-only
      --store-type string          storage type (dir, mem) (default "dir")
      --tls-cert string            TLS certificate filename for HTTPS
      --tls-key string             TLS key filename for HTTPS
      --warning stringArray        warning headers to include with all responses
```

## Options from parent commands

```text
  -v, --verbosity string   Log level (trace, debug, info, warn, error) (default "warn")
```
