# Kalco - Kubernetes Analysis & Lifecycle Control

**Professional CLI tool for Kubernetes cluster management, analysis, and lifecycle control**

[![Release](https://img.shields.io/github/v/release/graz-dev/kalco)](https://github.com/graz-dev/kalco/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Documentation](https://img.shields.io/badge/docs-available-brightgreen)](https://graz-dev.github.io/kalco)

*Extract, analyze, and version control your entire Kubernetes cluster with Git integration*

[Quick Start](#quick-start) • [Documentation](https://graz-dev.github.io/kalco) • [Examples](#examples) • [Contributing](#contributing)

---

## Overview

Kalco is a CLI tool designed for comprehensive Kubernetes cluster analysis, resource extraction, and lifecycle management. Kalco provides automated cluster backup and change tracking capabilities through a clean interface.

## Core Features

### Context Management
Manage multiple Kubernetes clusters through a unified interface:

- **Multi-Cluster Support** - Handle dev, staging, and production environments
- **Configuration Persistence** - Store cluster settings and output directories
- **Team Collaboration** - Share and import context configurations
- **Environment Isolation** - Separate configurations for different clusters

### Resource Export
Export cluster resources with professional organization:

- **Structured Output** - Intuitive directory organization by namespace and resource type
- **Complete Coverage** - Native Kubernetes resources and Custom Resource Definitions (CRDs)
- **Clean YAML** - Metadata optimization for re-application
- **Context-Based Configuration** - Output directory automatically derived from active context

### Git Integration
Automatic version control for cluster changes:

- **Repository Initialization** - Automatic Git setup for new export directories
- **Change Tracking** - Commit history with timestamp-based or custom messages
- **Remote Push** - Optional automatic push to remote repositories
- **Branch Management** - Support for main/master branch conventions

### Report Generation
Professional change analysis reports:

- **Change Tracking** - Detailed analysis of resource modifications
- **Git Integration** - Complete commit history and diff information
- **Professional Formatting** - Clean Markdown reports with actionable insights

## Quick Start

### Installation

**Package Managers (recommended):**
```bash
# Homebrew
brew install graz-dev/tap/kalco

# Build from source
git clone https://github.com/graz-dev/kalco.git
cd kalco && go build -o kalco
```

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
```

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.ps1 | iex
```

### Basic Usage

1. **Create a Context:**
   ```bash
   kalco context set production \
     --kubeconfig ~/.kube/prod-config \
     --output ./prod-exports \
     --description "Production cluster for customer workloads"
   ```

2. **Export Cluster Resources:**
   ```bash
   kalco export --git-push --commit-message "Weekly backup"
   ```

3. **Load Existing Context:**
   ```bash
   kalco context load ./existing-kalco-export
   ```

## Command Reference

### Core Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `kalco context` | Manage cluster contexts | `kalco context set/list/use/load` |
| `kalco export` | Export cluster resources | `kalco export [flags]` |
| `kalco completion` | Shell completion | `kalco completion bash\|zsh\|fish\|powershell` |
| `kalco version` | Version information | `kalco version` |

Read the [docs](https://graz-dev.github.io/kalco) for more details.

## Contributing

We welcome contributions to improve Kalco:

- **Bug Reports**: [Create an issue](https://github.com/graz-dev/kalco/issues/new)
- **Feature Requests**: [Create an issue](https://github.com/graz-dev/kalco/issues/new)
- **Code Contributions**: Fork, develop, and submit pull requests
- **Documentation**: Help improve guides and examples

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Documentation**: [https://graz-dev.github.io/kalco](https://graz-dev.github.io/kalco)
- **Issues**: [GitHub Issues](https://github.com/graz-dev/kalco/issues)

---

<div align="center">

**Built with ❤️ for the Kubernetes community**

[Star us on GitHub](https://github.com/graz-dev/kalco) • [Read the Docs](https://graz-dev.github.io/kalco)

</div>
