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
- **Git** - For version control functionality
- **Go 1.21+** (if building from source) - [Download here](https://golang.org/dl/)

## Quick Install

## Package Managers (recommended)

### Homebrew (macOS/Linux)

```bash
# Add the tap
brew tap graz-dev/tap

# Install Kalco
brew install kalco
```

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

## Verification

After installation, verify Kalco is working correctly:

```bash
# Check version
kalco version

# Show help
kalco --help
```

Expected output should show:
- Version information
- Available commands (context, export, version)
- Global flags and options


### Getting Help

- **Installation issues**: [GitHub Issues](https://github.com/graz-dev/kalco/issues)
- **Build problems**: Check Go version and dependencies
- **Binary issues**: Try downloading the pre-built binary

## Next Steps

After successful installation:

1. **Read the [First Run]({{ site.baseurl }}/getting-started/first-run) guide** to get started
2. **Set up your first context** with `kalco context set`
3. **Export your cluster** with `kalco export`
4. **Explore the [Command Reference]({{ site.baseurl }}/commands/index.md)**

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

---

*For more installation help, see [GitHub Issues](https://github.com/graz-dev/kalco/issues).*
