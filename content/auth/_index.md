---
title: Auth
layout: single
date: 2026-01-01
---

<!-- markdownlint-disable-file MD010 -->

## Enabling Basic Authentication

Basic auth sends the base64 encoded login with every request.
It can be enabled with:

```shell
olareg serve --auth-basic auth.yaml
```

## Enabling Opaque Token Authentication

Opaque token based authentication performs a login to the token endpoint and receives an opaque token that is used for all further requests.
This token expires after 5 minutes requiring clients to periodically reauthenticate.
However it is much faster in most scenarios since the password hashing is only performed once per token lifetime rather than once per request.
It can be enabled with:

```shell
olareg serve --auth-token-opaque auth.yaml --external-url https://registry.example.org
```

Note that `--external-url` is strongly encouraged and must specify the URL clients can use to reach the registry server.
If an external URL is not available, olareg will respond to clients with a relative token URL, which may not be supported by the client.

## Auth Config File Syntax

The config file contains the following yaml syntax:

```yaml
users:
  $user_name: cred: "$crypted_pass"
  ...

groups:
  $group_name: members: ["$user1", "$user2"] # both users and groups may be listed
  ...

acls:
  - repo: "$path1"                 # path with an optional trailing * wildcard, an empty string matches /v2/ ping requests
    access: ["read", "write"]      # possible values are "read", "write", "delete", and "*" (matching all access)
    members: ["$user1", "$group2"] # list of user names and/or group names
    anonymous: false               # whether anonymous users are permitted
  - repo: "$path2"                 # see below for how an ACL is selected
    ...
```

Users are defined with a hashed password (using bcrypt).
Hashes may be generated using `olareg hash --pass "$pass_to_hash"`.
This password is hashed with a salt, meaning the same password may hash to multiple values.
Taking steps to prevent this from being stored in your shell history are outside of the scope of olareg, however, a leading space may work depending on your shell settings.

To simplify management of users and ACLs, users may be added to one or more groups.

Authentication is then determined by finding a matching ACL entry per request.
When searching for an ACL entry, the longest matching path containing the user or one of their groups is used.
For the request to be accepted:

- The type of access must be listed in the access.
- If the user is logged in, they or a group they belong to must be listed in the members.
- If the user is anonymous, the anonymous flag must be set to true.

All other requests will be rejected when authentication is enabled.

## Example Configuration File

```yaml
users:
  alice:
    cred: "$2a$10$AeIxYk02nNYLrmkEIQRSse4DsFH0M9exGec0FbSDSY0fPSZ9chPoa" # password1
  bob:
    cred: "$2a$10$4iTFUSDqPMFRdG0ukcoNzePmjmblKtVCQF2Q50aoRymIat5TM/mXy" # password2

groups:
  admin:
    members: ["alice"]
  guest:
    members: ["bob"]
  all:
    members: ["admin", "guest"]

acls:
  - repo: "*" # full access for admin
    members: ["admin"]
    access: ["read", "write", "delete"]
  - repo: "" # v2 ping requests
    members: ["all"]
    access: ["read"]
  - repo: "guest/*"
    members: ["guest"]
    access: ["read"]
  - repo: "public/*"
    anonymous: true
    access: ["read"]
  - repo: "any/*"
    members: ["all"]
    access: ["*"]
```
