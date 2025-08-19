# Kalco - Kubernetes Analysis & Lifecycle Control

**Professional CLI tool for Kubernetes cluster management, analysis, and lifecycle control**

[![Release](https://img.shields.io/github/v/release/graz-dev/kalco)](https://github.com/graz-dev/kalco/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Documentation](https://img.shields.io/badge/docs-available-brightgreen)](https://graz-dev.github.io/kalco)

*Extract, validate, analyze, and version control your entire Kubernetes cluster with comprehensive validation and Git integration*

[Quick Start](#quick-start) • [Documentation](https://graz-dev.github.io/kalco) • [Examples](#examples) • [Contributing](#contributing)

---

## Overview

Kalco is a professional-grade CLI tool designed for comprehensive Kubernetes cluster analysis, resource extraction, validation, and lifecycle management. Built with enterprise needs in mind, Kalco provides automated cluster backup, change tracking, and validation capabilities through a clean, professional interface.

### Key Benefits

- **Automated Cluster Management** - Streamlined workflows for cluster operations
- **Professional CLI Experience** - Clean, emoji-free interface designed for production use
- **Git Integration** - Automatic version control with commit history and change tracking
- **Comprehensive Validation** - Cross-reference checking and orphaned resource detection
- **Enterprise Ready** - Designed for production environments and team collaboration

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
- **Flexible Filtering** - Namespace, resource type, and exclusion options

### Git Integration
Automatic version control for cluster changes:

- **Repository Initialization** - Automatic Git setup for new export directories
- **Change Tracking** - Commit history with timestamp-based or custom messages
- **Remote Push** - Optional automatic push to remote repositories
- **Branch Management** - Support for main/master branch conventions

### Report Generation
Professional change analysis and validation reports:

- **Change Tracking** - Detailed analysis of resource modifications
- **Cross-Reference Validation** - Detection of broken resource references
- **Orphaned Resource Detection** - Identification of unused resources
- **Professional Formatting** - Clean Markdown reports with actionable insights

## Quick Start

### Installation

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
```

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.ps1 | iex
```

**Package Managers:**
```bash
# Homebrew
brew install graz-dev/tap/kalco

# Build from source
git clone https://github.com/graz-dev/kalco.git
cd kalco && go build -o kalco
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

### Context Management

```bash
# Create context
kalco context set <name> --kubeconfig <path> --output <dir>

# List contexts
kalco context list

# Switch context
kalco context use <name>

# Load from existing directory
kalco context load <directory>

# Show context details
kalco context show <name>

# Delete context
kalco context delete <name>
```

### Export Options

```bash
# Basic export
kalco export

# Custom output directory
kalco export --output ./cluster-backup

# Specific namespaces
kalco export --namespaces default,kube-system

# Resource filtering
kalco export --resources pods,services,deployments

# Exclude resources
kalco export --exclude events,replicasets

# Git integration
kalco export --git-push --commit-message "Daily backup"

# Skip Git operations
kalco export --no-commit
```

## Output Structure

Kalco creates a professional, organized directory structure:

```
<output_dir>/
├── <namespace>/
│   ├── <resource_kind>/
│   │   ├── <resource_name>.yaml
│   │   └── ...
│   └── ...
├── _cluster/
│   ├── <resource_kind>/
│   │   ├── <resource_name>.yaml
│   │   └── ...
│   └── ...
├── kalco-reports/
│   └── <timestamp>-<commit-message>.md
└── kalco-config.json
```

- **Namespaced Resources**: `<output>/<namespace>/<kind>/<name>.yaml`
- **Cluster Resources**: `<output>/_cluster/<kind>/<name>.yaml`
- **Reports**: `<output>/kalco-reports/<timestamp>-<commit-message>.md`
- **Configuration**: `<output>/kalco-config.json`

## Examples

### Quickstart Demo

Run the comprehensive example to see Kalco in action:

```bash
# Make executable and run
chmod +x examples/quickstart.sh
./examples/quickstart.sh

# Keep demo directory for inspection
./examples/quickstart.sh --keep
```

The quickstart demonstrates:
- Context creation and management
- Resource export and organization
- Git repository initialization
- Report generation and validation

### Production Workflow

```bash
# Set up production context
kalco context set production \
  --kubeconfig ~/.kube/prod-config \
  --output ./prod-exports \
  --labels env=prod,team=platform

# Export with Git integration
kalco export --git-push --commit-message "Production backup $(date)"

# Load context on another machine
kalco context load ./prod-exports
```

### CI/CD Integration

```bash
# Automated backup in pipeline
kalco export \
  --namespaces production \
  --exclude events,pods \
  --output ./gitops-repo \
  --commit-message "Automated backup $(date)"
```

## Configuration

### Context Configuration

Contexts store cluster-specific settings:

```yaml
# Example context configuration
name: production
kubeconfig: ~/.kube/prod-config
output_dir: ./prod-exports
description: Production cluster for customer workloads
labels:
  env: production
  team: platform
  customer: enterprise
```

### Export Configuration

Customize export behavior through flags:

- `--output`: Specify output directory
- `--namespaces`: Filter by namespace
- `--resources`: Filter by resource type
- `--exclude`: Exclude specific resources
- `--git-push`: Enable remote push
- `--no-commit`: Skip Git operations

## Architecture

### Core Components

- **Context Manager**: Handles cluster configurations and settings
- **Resource Exporter**: Discovers and exports Kubernetes resources
- **Git Integration**: Manages version control operations
- **Report Generator**: Creates change analysis and validation reports
- **Validation Engine**: Performs cross-reference and orphaned resource checks

### Design Principles

- **Professional Interface**: Clean, emoji-free CLI design
- **Minimal Dependencies**: Focused functionality without bloat
- **Enterprise Ready**: Production-grade reliability and performance
- **Team Collaboration**: Shared configurations and context sharing
- **Automation First**: Designed for CI/CD and automated workflows

## Use Cases

### DevOps Teams
- **Automated Backups**: Regular cluster snapshots with Git history
- **Change Tracking**: Monitor cluster modifications over time
- **Disaster Recovery**: Quick cluster restoration from exports
- **Environment Replication**: Copy configurations between clusters

### Platform Engineers
- **Infrastructure as Code**: Version-controlled cluster configurations
- **Compliance Auditing**: Track and validate cluster changes
- **Team Onboarding**: Share standardized cluster contexts
- **Migration Support**: Export and import cluster configurations

### Security Teams
- **Configuration Auditing**: Track security-related changes
- **Compliance Reporting**: Generate audit reports for compliance
- **Access Control**: Manage cluster access through contexts
- **Security Validation**: Check for security misconfigurations

## Contributing

We welcome contributions to improve Kalco:

- **Bug Reports**: [Create an issue](https://github.com/graz-dev/kalco/issues/new)
- **Feature Requests**: [Start a discussion](https://github.com/graz-dev/kalco/discussions)
- **Code Contributions**: Fork, develop, and submit pull requests
- **Documentation**: Help improve guides and examples

### Development Setup

```bash
# Clone repository
git clone https://github.com/graz-dev/kalco.git
cd kalco

# Install dependencies
go mod tidy

# Build and test
go build -o kalco
go test ./...

# Run locally
./kalco --help
```

## Documentation

- **User Guide**: [https://graz-dev.github.io/kalco](https://graz-dev.github.io/kalco)
- **API Reference**: [https://graz-dev.github.io/kalco/api](https://graz-dev.github.io/kalco/api)
- **Examples**: [https://graz-dev.github.io/kalco/examples](https://graz-dev.github.io/kalco/examples)
- **Contributing**: [https://graz-dev.github.io/kalco/contributing](https://graz-dev.github.io/kalco/contributing)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Documentation**: [https://graz-dev.github.io/kalco](https://graz-dev.github.io/kalco)
- **Issues**: [GitHub Issues](https://github.com/graz-dev/kalco/issues)
- **Discussions**: [GitHub Discussions](https://github.com/graz-dev/kalco/discussions)
- **Community**: Join our community discussions

---

<div align="center">

**Built with ❤️ for the Kubernetes community**

[Star us on GitHub](https://github.com/graz-dev/kalco) • [Read the Docs](https://graz-dev.github.io/kalco) • [Join Discussions](https://github.com/graz-dev/kalco/discussions)

</div>
