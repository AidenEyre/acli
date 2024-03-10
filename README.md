# Aiden CLI

ACLI is a Golang CLI tool build using the Cobra library. It provides a collection of useful
commands to make your life easier on the command line.

**On This Page**
- [Aiden CLI](#aiden-cli)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Commands](#commands)
    - [kch](#kch)
    - [Config](#config)
    - [Connect](#connect)
  - [Roadmap](#roadmap)
  - [Versioning](#versioning)
  - [Authors](#authors)

## Installation

> :memo: Ensure that Golang is installed on your machine.

1. Run `make install` in the root of the repository.
2. You may need to add the Go bin to your path:
   1. Add the following to your .bashrc, .zshrc, etc.
      ```
      export GOPATH=$HOME/go
      export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
      ```
    2. Refresh your terminal by restarting or running `source <rc-file>`.

> :information_source: You can also install directly from the repo with
> `go install github.com/aideneyre/acli@latest`

## Usage

```
# Set context with an interactive prompt and open K9s to specified resource/namespace
acli kch -r pod -n kube-system
# Configure persistent config for the CLI
acli config
# View persistent config for the CLI
acli config view
# Connect to a server with SSH and an interactive prompt
acli connect
```

## Commands

### kch

![kch-demo](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExbHF3cDZxazN4NHVmNmF1aWZ1NHNwZHBhbHJqd292bXV3NWhyZGVqYyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/GyAYfkc4UWfVw2mrJo/giphy.gif)

Kube change (kch) is a command that can be run to set the current context. The
command provides an interactive selection prompt or the context can be
directly set by adding it after the command (e.g. `acli kch my-cluster`). The
`-r`/`--resource` flag can be used to open K9s after setting the contents to the
specified resource. The `-n`/`--namespace` flag can be used to specify the
namespace used when opening K9s (default namespace: `default`).

> **Note**
> Much of this functionality assumes you have K9s installed and use it.

### Config

The config command allows you to define persistent configs for the CLI. The CLI
will attempt to create the folder at `~/.acli`. If you choose, you can directly
modify this file rather than using the command (see options below). When
running the command, you will be given an interactive prompt. If you would like
to print the existing configs, run `acli config view`.
```yaml
# the connect section configures SSH connection settings and aliases for quick
# access.
connect:
  ## Aliases for SSH connections
  aliases:
    - ip: "" # IP address of the SSH server, e.g., "192.168.1.127"
      name: "" # Unique name for this alias, e.g., "rpi"
      pemfile: "" # Path to PEM file for key authentication, if used
      port: "" # SSH port, default is "22"
      user: "" # Username for SSH connection, e.g., "aideneyre"
  sshdir: "" # Default directory for SSH keys

# Configuration for Kubernetes command-line tools and preferences.
kch:
  alwaysopenk9s: false # Whether to always open K9s on start (true/false)
  defaults:
    namespace: "" # Default namespace, e.g., "argocd"
    resource: "" # Default resource type, e.g., "service"

```

### Connect

![connect-demo](https://media.giphy.com/media/swYUWOADqUfzA1NzNN/giphy.gif)

The `connect` command allows you to connect to remote servers using SSH. This
command relies on set configs through the `config` command. If you just run,
`acli connect`, you will be presented with an interactive list of the
configured servers in your config. After choosing one, it will attempt to
connect. If you know the alias you want to connect to, feel free to tack it on,
`acli connect <alias>. Pem file is not required for this connection.

## Roadmap

- [ ] Implement an easy way to distribute this (brew, etc).
- [ ] Config
  - [ ] Allow users to specify default kubeconfig path to use.
- [ ] tests
  - [ ] Make more tests.
  - [ ] Refactor kube package to better handle unit tests.
- [ ] Create pipelines
  - [ ] Lint, and, enforce standards, etc. on MRs.
  - [ ] On tags, I want to build the binary and store it in the package registry.

## Versioning

We use [SemVer](SemVer) for versioning on the root command and each sub-command.

## Authors

- Aiden Eyre
