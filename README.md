# Aiden CLI

ACLI is a Golang CLI tool build using the Cobra library. It provides a collection of useful
commands to make your life easier on the command line.

**On This Page**
- [Aiden CLI](#aiden-cli)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Commands](#commands)
    - [kch](#kch)
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
# Get help page with an interactive prompt
acli idk
```

## Commands

### kch

Kube change (kch) is a command that can be run to set the current context. The
command provides an interactive selection prompt or the context can be
directly set by adding it after the command (e.g. `acli kch my-cluster`). The
`-r`/`--resource` flag can be used to open K9s after setting the contents to the
specified resource. The `-n`/`--namespace` flag can be used to specify the
namespace used when opening K9s (default namespace: `default`).

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
