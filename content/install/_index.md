---
title: Installing
layout: single
date: 2025-01-01
---

<!-- markdownlint-disable-file MD010 -->

## Running as a Container

With persistent storage:

```shell
docker run -p 5000:5000 \
  -v "olareg-data:/home/appuser/registry" \
  ghcr.io/olareg/olareg serve --dir /home/appuser/registry
```

For an ephemeral registry:

```shell
docker run -p 5000:5000 --rm \
  ghcr.io/olareg/olareg serve --store-type mem
```

## Installing as a Binary

Binaries can be downloaded from the [releases page](https://github.com/olareg/olareg/releases).

Downloading the latest release for linux/amd64 can be done with curl:

```shell
curl -L https://github.com/olareg/olareg/releases/latest/download/olareg-linux-amd64 >olareg
chmod 755 olareg
```

For binaries downloaded on MacOS, the quarantine attribute can be removed with:

```shell
xattr -d com.apple.quarantine olareg
```

## Using in Go Unit Tests

One of the goals of olareg was to serve as a simple registry implementation to test tooling that works with container registries.
Integrating it into unit tests with the httptest based server involves the following setup:

```go
func TestMethod(t *testing.T) {
	rh := olareg.New(config.Config{
		Storage: config.ConfigStorage{
			StoreType: config.StoreMem,
			RootDir:   "./testdata", // serve content from testdata, writes only apply to memory
		},
	})
	ts := httptest.NewServer(rh)
	t.Cleanup(func() {
		ts.Close()
		_ = rh.Close()
	})
	tsURL, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatal("failed to parse url %s: %v", ts.URL, err)
	}
	// send requests to tsURL.Host registry server ...
}
```

Go docs may be found at <https://pkg.go.dev/github.com/olareg/olareg>.

## Verifying Signatures

Binaries and images have been signed with cosign.

For images:

```shell
cosign verify \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  --certificate-identity-regexp https://github.com/olareg/olareg/.github/workflows/ \
  ghcr.io/olareg/olareg:latest
```

For binaries:

```shell
curl -L https://github.com/olareg/olareg/releases/latest/download/olareg-linux-amd64 >olareg
chmod 755 olareg
curl -L https://github.com/olareg/olareg/releases/latest/download/olareg-linux-amd64.pem >olareg-linux-amd64.pem
curl -L https://github.com/olareg/olareg/releases/latest/download/olareg-linux-amd64.sig >olareg-linux-amd64.sig
cosign verify-blob \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com \
  --certificate-identity-regexp https://github.com/olareg/olareg/.github/workflows/ \
  --certificate olareg-linux-amd64.pem \
  --signature olareg-linux-amd64.sig \
  olareg
rm olareg-linux-amd64.pem olareg-linux-amd64.sig
```
