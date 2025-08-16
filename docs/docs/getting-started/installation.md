---
layout: default
title: Installation
parent: Getting Started
nav_order: 1
---

# Installation

Choose the installation method that best fits your environment and get Kalco running in minutes.

## Quick Install (Recommended)

The fastest way to get started with Kalco:

### Linux/macOS

```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
```

This script automatically:
- Downloads the latest version
- Installs to `/usr/local/bin/`
- Sets up shell completion
- Verifies the installation

### Windows (PowerShell)

```powershell
iwr -useb https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.ps1 | iex
```

## Package Managers

### Homebrew (macOS/Linux)

```bash
# Add the tap
brew tap graz-dev/tap

# Install Kalco
brew install kalco
```

### Debian/Ubuntu

```bash
# Download the .deb package
wget https://github.com/graz-dev/kalco/releases/latest/download/kalco_Linux_x86_64.deb

# Install
sudo dpkg -i kalco_Linux_x86_64.deb
```

### RHEL/CentOS/Fedora

```bash
# Download the .rpm package
wget https://github.com/graz-dev/kalco/releases/latest/download/kalco_Linux_x86_64.rpm

# Install
sudo rpm -i kalco_Linux_x86_64.rpm
```

## Build from Source

For the latest development version or custom builds:

### Prerequisites

- **Go 1.21+** - [Download here](https://golang.org/dl/)
- **Git** - For cloning the repository

### Build Steps

```bash
# Clone the repository
git clone https://github.com/graz-dev/kalco.git
cd kalco

# Install dependencies
go mod tidy

# Build the binary
go build -o kalco ./cmd

# Make it available system-wide (optional)
sudo mv kalco /usr/local/bin/
```

## Manual Installation

1. **Download Binary**
   - Visit the [releases page](https://github.com/graz-dev/kalco/releases)
   - Download the appropriate binary for your platform

2. **Extract and Install**
   ```bash
   # Extract the archive
   tar -xzf kalco-*.tar.gz
   
   # Move to a directory in your PATH
   sudo mv kalco /usr/local/bin/
   
   # Verify installation
   kalco --version
   ```

## Verify Installation

After installation, verify that Kalco is working:

```bash
# Check version
kalco --version

# Check help
kalco --help

# Verify binary location
which kalco
```

## Shell Completion

Kalco provides shell completion for faster workflows:

### Bash

```bash
# Generate completion script
kalco completion bash > /etc/bash_completion.d/kalco

# Reload shell or source the file
source /etc/bash_completion.d/kalco
```

### Zsh

```bash
# Generate completion script
kalco completion zsh > ~/.zsh/completion/_kalco

# Add to .zshrc
echo 'autoload -U compinit && compinit' >> ~/.zshrc
```

### Fish

```bash
# Generate completion script
kalco completion fish > ~/.config/fish/completions/kalco.fish
```

### PowerShell

```powershell
# Generate completion script
kalco completion powershell > kalco-completion.ps1

# Source the file
. .\kalco-completion.ps1
```

## Next Steps

Once Kalco is installed, proceed to [First Run]({{ site.baseurl }}/docs/getting-started/first-run) to export your first cluster.

## Troubleshooting

### Permission Denied

```bash
# Make binary executable
chmod +x kalco

# Check file permissions
ls -la kalco
```

### Command Not Found

```bash
# Check if binary is in PATH
echo $PATH

# Verify binary location
which kalco

# Add to PATH if needed
export PATH=$PATH:/path/to/kalco
```

### Version Issues

```bash
# Check Go version
go version

# Ensure Go 1.21+ is installed
# Download from https://golang.org/dl/
```
