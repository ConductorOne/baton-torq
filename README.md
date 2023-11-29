# `baton-torq` [![Go Reference](https://pkg.go.dev/badge/github.com/conductorone/baton-torq.svg)](https://pkg.go.dev/github.com/conductorone/baton-torq) ![main ci](https://github.com/conductorone/baton-torq/actions/workflows/main.yaml/badge.svg)

`baton-torq` is a connector for Torq built using the [Baton SDK](https://github.com/conductorone/baton-sdk). It communicates with the Torq API to sync data about users and their roles in the workspace.
Check out [Baton](https://github.com/conductorone/baton) to learn more about the project in general.

# Getting Started

## Prerequisites

- Access to the Torq Workspace.
- Generate API token for your workspace. API token consists of a client_id and client_secret. To generate a token click on your avatar -> API Keys -> Create API Key. Store the client_id & client_secret as the client_secret won't be visible again. 

## brew

```
brew install conductorone/baton/baton conductorone/baton/baton-torq

BATON_TORQ_CLIENT_ID=torqClientId BATON_TORQ_CLIENT_SECRET=torqClientSecret baton-torq
baton resources
```

## docker

```
docker run --rm -v $(pwd):/out -e BATON_TORQ_CLIENT_ID=torqClientId BATON_TORQ_CLIENT_SECRET=torqClientSecret ghcr.io/conductorone/baton-torq:latest -f "/out/sync.c1z"
docker run --rm -v $(pwd):/out ghcr.io/conductorone/baton:latest -f "/out/sync.c1z" resources
```

## source

```
go install github.com/conductorone/baton/cmd/baton@main
go install github.com/conductorone/baton-torq/cmd/baton-torq@main

BATON_TORQ_CLIENT_ID=torqClientId BATON_TORQ_CLIENT_SECRET=torqClientSecret baton-torq
baton resources
```

# Data Model

`baton-torq` will pull down information about the following Torq resources:

- Users
- Roles

# Contributing, Support and Issues

We started Baton because we were tired of taking screenshots and manually building spreadsheets. We welcome contributions, and ideas, no matter how small -- our goal is to make identity and permissions sprawl less painful for everyone. If you have questions, problems, or ideas: Please open a Github Issue!

See [CONTRIBUTING.md](https://github.com/ConductorOne/baton/blob/main/CONTRIBUTING.md) for more details.

# `baton-torq` Command Line Usage

```
baton-torq

Usage:
  baton-torq [flags]
  baton-torq [command]

Available Commands:
  completion         Generate the autocompletion script for the specified shell
  help               Help about any command

Flags:
      --client-id string            The client ID used to authenticate with ConductorOne ($BATON_CLIENT_ID)
      --client-secret string        The client secret used to authenticate with ConductorOne ($BATON_CLIENT_SECRET)
  -f, --file string                 The path to the c1z file to sync with ($BATON_FILE) (default "sync.c1z")
  -h, --help                        help for baton-torq
      --log-format string           The output format for logs: json, console ($BATON_LOG_FORMAT) (default "json")
      --log-level string            The log level: debug, info, warn, error ($BATON_LOG_LEVEL) (default "info")
  -p, --provisioning                This must be set in order for provisioning actions to be enabled. ($BATON_PROVISIONING)
      --torq-client-id string       Client ID used to authenticate to the Torq API. ($BATON_TORQ_CLIENT_ID)
      --torq-client-secret string   Client Secret used to authenticate to the Torq API. ($BATON_TORQ_CLIENT_SECRET)
  -v, --version                     version for baton-torq

Use "baton-torq [command] --help" for more information about a command.
```
