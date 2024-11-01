# Usage Instructions for Development and Testing Configurations

This guide provides instructions on temporarily setting up and using dummy configuration files in
the `testdata` directory for testing purposes.

---

## Table of Contents
- [Usage Instructions for Development and Testing Configurations](#usage-instructions-for-development-and-testing-configurations)
  - [Table of Contents](#table-of-contents)
  - [1. Kubeconfig Setup](#1-kubeconfig-setup)
    - [Setup](#setup)
    - [Cleanup](#cleanup)
  - [2. ACLI Config Setup](#2-acli-config-setup)
    - [Setup](#setup-1)
    - [Cleanup](#cleanup-1)

---

## 1. Kubeconfig Setup

To use the dummy kubeconfig file for testing without affecting your default configuration, set the
`KUBECONFIG` environment variable temporarily.

### Setup

Run the following command in your terminal, replacing `<path_to_repo>` with the path to your
local repository:

```bash
export KUBECONFIG=<path_to_repo>/acli/testdata/dummy_kubeconfig
```

### Cleanup

To revert any temporary settings and restore your default configurations:

**Unset Kubeconfig**: Run the following command to remove the `KUBECONFIG` environment variable,
which will make `kubectl` revert to using the default kubeconfig (usually located at
`~/.kube/config`):

```bash
unset KUBECONFIG
```

> **Note**
> If you had a custom `KUBECONFIG` value before testing, set it back with the export command above.

## 2. ACLI Config Setup

To use the dummy ACLI config file located in the `testdata` directory for testing, follow these
steps to temporarily replace your default config. 

### Setup

1. **Back up the Original Config** (if it exists) to prevent data loss:

   ```bash
   mv ~/.acli ~/.acli.bak
   ```
2. **Copy the Dummy Config** into the expected location:

   ```bash
   cp <path_to_repo>/acli/testdata/dummy_config.yaml ~/.acli
   ```
3. **Run Your Tests**. With the dummy configuration in place, run ACLI commands to test the
   application as needed.

### Cleanup

Once testing is complete, restore the original config to avoid leaving the dummy configuration in
place:

1. **Restore the Original Config**:

   ```bash
   mv ~/.acli.bak ~/.acli
   ```
