# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


***Below is the template for how you should record changes under any of the sub commands in the release***
```
- `Subcommand`| `Version` | `Explanation...`
```

## [1.4.0]

### Added

### Changed

### Fixed
- `kch` | `2.2.1` | Fully exit program if no kubeconfig is found.
- `kch` | `2.2.1` | Gracefully handle K9s not being installed.

### Removed

## [1.3.0]

### Added
- `acli` | `1.3.0` | Added `~/.acli.yaml` file managed by viper for persistent
  configuration.
- `config` | `0.1.0` | Added `config` command to open an interactive prompt to
  configure persistent CLI configurations. Added the following configurations:
- `connect` | `0.1.0` | Added `connect` command to connect to different servers
  using SSH and a pre-defined set of servers that are modified with the config
  command.

### Changed
- `kch` | `2.2.0` | If opening K9s without specifying resource/namespace, use
  defaults defined in `~/.acli` configuration file.
- `kch` | `2.2.0` | If "kch.alwaysopenk9s" is set true in the `~/.acli`
  configuration file, K9s will always open by default.

### Fixed
- `kch` | `2.2.0` | Don't error when there are no contexts configured.
- `kch` | `2.2.0` | Exit program when CLI exits early (don't open K9s on
  ctrl+c).
- `kch` | `2.2.0` | Handle current context not existing in available contexts.

### Removed
- `idk` | `1.0.0` | Removed the `idk` command.

## [1.2.1]

### Added

### Changed
- `kch` | `2.1.0` | Improved output when there are no context configured.

### Fixed
- `kch` | `2.1.0` | Fixed issue where `acli kch` would always use
  `~/.kube/config` when setting context.

### Removed
