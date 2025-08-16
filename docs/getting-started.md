# Getting Started with Kalco

This guide will help you get up and running with Kalco in minutes.

## Prerequisites

- **Kubernetes Access** - Valid kubeconfig or in-cluster access
- **Go 1.21+** (if building from source) - [Download here](https://golang.org/dl/)
- **Git** (optional) - For version control functionality

## Installation

### Quick Install (Recommended)

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/main/scripts/install.sh | bash
```

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/graz-dev/kalco/main/scripts/install.ps1 | iex
```

### Package Managers

**Debian/Ubuntu:**
```bash
wget https://github.com/graz-dev/kalco/releases/latest/download/kalco_Linux_x86_64.deb
sudo dpkg -i kalco_Linux_x86_64.deb
```

**RHEL/CentOS/Fedora:**
```bash
wget https://github.com/graz-dev/kalco/releases/latest/download/kalco_Linux_x86_64.rpm
sudo rpm -i kalco_Linux_x86_64.rpm
```

### Manual Installation

1. Download the appropriate binary for your platform from the [releases page](https://github.com/graz-dev/kalco/releases)
2. Extract the archive and move the binary to your PATH

### Build from Source

```bash
git clone https://github.com/graz-dev/kalco.git
cd kalco
go mod tidy
go build -o kalco
sudo mv kalco /usr/local/bin/
```

## First Steps

### 1. Verify Installation

```bash
kalco version
```

### 2. Check Available Commands

```bash
kalco --help
```

### 3. Export Your First Cluster

```bash
# Export entire cluster to timestamped directory
kalco export

# Export to specific directory
kalco export --output ./my-cluster-backup
```

### 4. Validate Your Cluster

```bash
# Check for broken references and issues
kalco validate
```

### 5. Analyze for Optimization

```bash
# Find orphaned resources
kalco analyze orphaned

# Get detailed analysis
kalco analyze orphaned --detailed
```

## Basic Workflow

Here's a typical workflow for using Kalco:

```bash
# 1. Export cluster resources
kalco export --output ./cluster-backup --git-push

# 2. Validate the exported resources
kalco validate

# 3. Analyze for cleanup opportunities
kalco analyze orphaned --output json > orphaned-resources.json

# 4. Generate comprehensive report
kalco report --output-file cluster-report.html
```

## Configuration

### Initialize Configuration

```bash
# Create local configuration
kalco config init

# Create global configuration
kalco config init --global
```

### Set Default Options

```bash
# Set default output directory
kalco config set output.directory ./backups

# Set default exclusions
kalco config set filters.exclude "events,replicasets"
```

## Shell Completion

Enable tab completion for faster workflows:

**Bash:**
```bash
kalco completion bash > /etc/bash_completion.d/kalco
```

**Zsh:**
```bash
kalco completion zsh > "${fpath[1]}/_kalco"
```

**Fish:**
```bash
kalco completion fish > ~/.config/fish/completions/kalco.fish
```

**PowerShell:**
```powershell
kalco completion powershell > kalco.ps1
# Add to your PowerShell profile
```

## Next Steps

- Explore the [Commands Reference](commands/index.md) for detailed command documentation
- Learn about [Configuration](configuration.md) options
- Check out [Use Cases](use-cases.md) for real-world examples
- Read the [FAQ](faq.md) for common questions

## Getting Help

If you need help:

1. Use `kalco [command] --help` for command-specific help
2. Check the [FAQ](faq.md) for common issues
3. Search existing [GitHub Issues](https://github.com/graz-dev/kalco/issues)
4. Create a new issue if you can't find an answer

---

[← Back to Documentation Home](index.md) | [Commands Reference →](commands/index.md)