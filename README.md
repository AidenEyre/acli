# Aiden CLI

ACLI is a Golang CLI tool build using the Cobra library. It provides a collection of useful
commands to make your life easier on the command line.

**On This Page**
- [Aiden CLI](#aiden-cli)
  - [Installation](#installation)
    - [Homebrew](#homebrew)
    - [Chocolatey](#chocolatey)
    - [Golang](#golang)
  - [Usage](#usage)
  - [Commands](#commands)
    - [kch](#kch)
    - [Config](#config)
    - [Connect](#connect)
  - [Roadmap](#roadmap)
  - [Versioning](#versioning)
  - [Distributing](#distributing)
  - [Authors](#authors)

## Installation

### Homebrew

```
brew tap AidenEyre/homebrew-aideneyre
brew install acli
```

### Chocolatey

> :memo: PENDING REVIEW - not yet available.

```
choco install acli --version=1.4.0
```

### Golang

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

![kch-demo](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExbGp3em9tdjNpeWcxMnlxbTcwaHVhbjN4ZDFxdWc0NWw2bjZzb2czbyZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/R6YePjt3jLzz0rJva0/giphy.gif)

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

![connect-demo](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExYXp4dzRtcTMzc3JlMWFweml0bmt1aGl1cWt5YTU0NWxidXkwMDFpNiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/mSehGiCOWTlWWRHNWN/giphy.gif)

The `connect` command allows you to connect to remote servers using SSH. This
command relies on set configs through the `config` command. If you just run,
`acli connect`, you will be presented with an interactive list of the
configured servers in your config. After choosing one, it will attempt to
connect. If you know the alias you want to connect to, feel free to tack it on,
`acli connect <alias>. Pem file is not required for this connection.

## Roadmap

- [ ] Config
  - [ ] Allow users to specify default kubeconfig path to use.
- [ ] kch
  - [ ] If K9s isn't installed, allow the `-n`/`-r` flags to just run a get on the resources.
- [ ] tests
  - [ ] Make more tests.
  - [ ] Refactor kube package to better handle unit tests.
- [ ] Create pipelines
  - [ ] Lint, and, enforce standards, etc. on MRs.
  - [ ] Implement a changelog tool for better release notes automation.

## Versioning

We use [SemVer](SemVer) for versioning on the root command and each sub-command.

## Distributing

We use [GoReleaser](https://goreleaser.com/) to distribute the CLI.

Currently deploying packages to: `Homebrew` and `Chocolatey`.

## Authors

- Aiden Eyre
