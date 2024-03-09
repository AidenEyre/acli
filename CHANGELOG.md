# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


***Below is the template for how you should record changes under any of the sub commands in the release***
```
- `Subcommand`| `Version` | `Explanation...`
```

## [1.2.0]

### Added
- `acli` | `1.3.0` | Added `~/.acli.yaml` file managed by viper for persistent
  configuration.

### Changed

### Fixed
- `kch` | `2.2.0` | Don't error when there are no contexts configured.
- `kch` | `2.2.0` | Exit program when CLI exits early (don't open K9s on
  ctrl+c).
- `kch` | `2.2.0` | Handle current context not existing in available contexts.
- `kch` | `2.2.0` | Allow `-n`/`-r` flags without a value to open K9s without
  defining a resource/namespace.

### Removed

## [1.2.1]

### Added

### Changed
- `kch` | `2.1.0` | Improved output when there are no context configured.

### Fixed
- `kch` | `2.1.0` | Fixed issue where `acli kch` would always use
  `~/.kube/config` when setting context.

### Removed
