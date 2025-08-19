---
layout: default
title: Installation
nav_order: 1
parent: Getting Started
---

# Installation

This guide covers installing Kalco on various operating systems and platforms.

## Prerequisites

Before installing Kalco, ensure you have:

- **Kubernetes Access** - Valid kubeconfig or in-cluster access
- **Git** (optional) - For version control functionality
- **Go 1.21+** (if building from source) - [Download here](https://golang.org/dl/)

## Quick Install

### Linux/macOS

Install Kalco with a single command:

```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
```

The script will:
- Detect your operating system and architecture
- Download the appropriate binary
- Install it to `/usr/local/bin/kalco`
- Make it executable

### Windows (PowerShell)

Install Kalco on Windows using PowerShell:

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

### Manual Download

Download the appropriate binary for your platform:

1. Visit the [releases page](https://github.com/graz-dev/kalco/releases)
2. Download the binary for your OS and architecture
3. Extract and move to your PATH

**Linux:**
```bash
# Download and extract
wget https://github.com/graz-dev/kalco/releases/latest/download/kalco_Linux_x86_64.tar.gz
tar -xzf kalco_Linux_x86_64.tar.gz

# Move to PATH
sudo mv kalco /usr/local/bin/
```

**macOS:**
```bash
# Download and extract
curl -L https://github.com/graz-dev/kalco/releases/latest/download/kalco_Darwin_x86_64.tar.gz | tar -xz

# Move to PATH
sudo mv kalco /usr/local/bin/
```

**Windows:**
```powershell
# Download and extract
Invoke-WebRequest -Uri "https://github.com/graz-dev/kalco/releases/latest/download/kalco_Windows_x86_64.zip" -OutFile "kalco.zip"
Expand-Archive -Path "kalco.zip" -DestinationPath "kalco"

# Add to PATH (requires admin)
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\kalco", [EnvironmentVariableTarget]::Machine)
```

## Build from Source

### Prerequisites

- **Go 1.21+** - [Download and install](https://golang.org/dl/)
- **Git** - For cloning the repository

### Build Steps

```bash
# Clone the repository
git clone https://github.com/graz-dev/kalco.git
cd kalco

# Install dependencies
go mod tidy

# Build the binary
go build -o kalco

# Make it available system-wide (optional)
sudo mv kalco /usr/local/bin/
```

### Development Build

For development and testing:

```bash
# Clone and build
git clone https://github.com/graz-dev/kalco.git
cd kalco
go build -o kalco

# Run locally
./kalco --help
```

## Verification

After installation, verify Kalco is working correctly:

```bash
# Check version
kalco version

# Show help
kalco --help

# Test context command
kalco context --help
```

Expected output should show:
- Version information
- Available commands (context, export, completion, version)
- Global flags and options

## Configuration

### Initial Setup

Kalco creates its configuration directory on first run:

```bash
# First run creates ~/.kalco/
kalco context list
```

The configuration directory structure:
```
~/.kalco/
├── contexts.yaml      # Context configurations
├── current-context    # Currently active context
└── config.json        # Global configuration
```

### Environment Variables

Kalco respects these environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `KUBECONFIG` | Path to kubeconfig file | `~/.kube/config` |
| `KALCO_CONFIG_DIR` | Configuration directory | `~/.kalco` |
| `NO_COLOR` | Disable colored output | `false` |

## Shell Completion

Enable shell completion for better user experience:

### Bash

```bash
# Generate completion script
kalco completion bash > /etc/bash_completion.d/kalco

# Or for current user
kalco completion bash > ~/.bash_completion.d/kalco

# Source in current shell
source <(kalco completion bash)
```

### Zsh

```bash
# Generate completion script
kalco completion zsh > ~/.zsh/completion/_kalco

# Add to .zshrc
echo 'fpath=(~/.zsh/completion $fpath)' >> ~/.zshrc
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

# Source in current session
. .\kalco-completion.ps1
```

## Troubleshooting

### Common Installation Issues

#### Permission Denied

```bash
Error: permission denied
```

**Solutions:**
- Use `sudo` for system-wide installation
- Check file permissions on the binary
- Verify PATH configuration

#### Binary Not Found

```bash
kalco: command not found
```

**Solutions:**
- Verify the binary is in your PATH
- Check if the installation completed successfully
- Restart your terminal session

#### Go Version Issues

```bash
go: version go1.20 is not supported
```

**Solutions:**
- Update to Go 1.21 or later
- Use the pre-built binary instead of building from source

#### Architecture Mismatch

```bash
cannot execute binary file: Exec format error
```

**Solutions:**
- Download the correct binary for your architecture
- Verify you're using the right OS/architecture combination

### Getting Help

- **Installation issues**: [GitHub Issues](https://github.com/graz-dev/kalco/issues)
- **Build problems**: Check Go version and dependencies
- **Binary issues**: Try downloading the pre-built binary

## Next Steps

After successful installation:

1. **Read the [First Run](first-run.md) guide** to get started
2. **Set up your first context** with `kalco context set`
3. **Export your cluster** with `kalco export`
4. **Explore the [Commands Reference](../commands/index.md)**

## Uninstallation

### Remove Binary

```bash
# Remove from system PATH
sudo rm /usr/local/bin/kalco

# Or if installed in user directory
rm ~/bin/kalco
```

### Remove Configuration

```bash
# Remove configuration directory
rm -rf ~/.kalco
```

### Remove Shell Completion

```bash
# Bash
sudo rm /etc/bash_completion.d/kalco

# Zsh
rm ~/.zsh/completion/_kalco

# Fish
rm ~/.config/fish/completions/kalco.fish
```

---

*For more installation help, see [GitHub Issues](https://github.com/graz-dev/kalco/issues) or [Discussions](https://github.com/graz-dev/kalco/discussions).*
