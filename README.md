# ☸️ Kalco - Kubernetes Analysis & Lifecycle Control

> Extract, validate, analyze, and version control your entire cluster with comprehensive validation and Git integration

https://github.com/user-attachments/assets/26f41f6a-6fa6-42fc-a0c9-f2dd9f5016e5

## Overview

**Kalco** (Kubernetes Analysis and Lifecycle Control) is a powerful Go CLI tool that performs comprehensive analysis, validation, and export of all resources from your Kubernetes cluster into organized YAML files. It automatically discovers and exports every available API resource - including native Kubernetes resources and Custom Resources (CRDs) - with zero configuration required, while providing intelligent validation, orphaned resource detection, and lifecycle management capabilities.

### What Kalco Does

Kalco transforms your Kubernetes cluster into a **comprehensive, validated, and lifecycle-managed** infrastructure snapshot that you can:
- **Reapply** to any cluster with confidence
- **Audit** for compliance and security
- **Clean up** by identifying orphaned resources
- **Document** your infrastructure as code
- **Migrate** between environments safely
- **Analyze** resource dependencies and relationships
- **Optimize** by identifying unused and orphaned resources

### Key Features

- **Complete Resource Discovery** - Automatically finds ALL available API resources
- **Comprehensive Coverage** - Includes both native K8s resources and Custom Resources (CRDs)
- **Structured Output** - Creates intuitive directory structures for easy navigation
- **Clean YAML** - Intelligently removes metadata fields that aren't useful for re-application
- **Lightning Fast** - Optimized for speed and efficiency
- **Git Integration** - Automatic version control with commit history and change tracking
- **Smart Reporting** - Generates detailed change reports with before/after comparisons and specific field modifications
- **Cross-Reference Validation** - Analyzes exported resources for broken references that could cause issues when reapplying
- **Orphaned Resource Detection** - Identifies resources no longer managed by higher-level controllers for cleanup

## Quick Start

### Prerequisites

- **Go 1.21+** - [Download here](https://golang.org/dl/)
- **Kubernetes Access** - Valid kubeconfig or in-cluster access
- **Git** - For version control functionality (optional but recommended)
- **KIND** - For local testing (optional) - [Installation Guide](https://kind.sigs.k8s.io/docs/user/quick-start/)

### Installation

#### Quick Install (Recommended)

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/main/scripts/install.sh | bash
```

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/graz-dev/kalco/main/scripts/install.ps1 | iex
```

#### Package Managers

**Debian/Ubuntu:**
```bash
# Download the .deb file from releases and install
wget https://github.com/graz-dev/kalco/releases/latest/download/kalco_Linux_x86_64.deb
sudo dpkg -i kalco_Linux_x86_64.deb
```

**RHEL/CentOS/Fedora:**
```bash
# Download the .rpm file from releases and install
wget https://github.com/graz-dev/kalco/releases/latest/download/kalco_Linux_x86_64.rpm
sudo rpm -i kalco_Linux_x86_64.rpm
```

#### Manual Installation

1. Download the appropriate binary for your platform from the [releases page](https://github.com/graz-dev/kalco/releases)
2. Extract the archive and move the binary to your PATH

#### Build from Source

```bash
# Clone the repository
git clone https://github.com/graz-dev/kalco.git
cd kalco

# Install dependencies and build
go mod tidy
go build -o kalco

# Make it available system-wide (optional)
sudo mv kalco /usr/local/bin/
```

### Quick Demo

Want to see kalco in action? Run the comprehensive quickstart:

```bash
# Run the complete quickstart demo
./examples/quickstart.sh
```

This will:
- Create a test Kubernetes cluster
- Export resources with automatic Git setup
- Modify cluster resources
- Generate enhanced change reports
- Clean up the test environment

### Basic Usage

```bash
# Dump all resources to default output directory
./kalco

# Specify custom output directory
./kalco --output-dir ./my-cluster-dump

# Use specific kubeconfig file
./kalco --kubeconfig ~/.kube/config --output-dir ./cluster-backup
```

### Output Structure

Kalco creates an intuitive directory layout that makes navigation simple:

```
<output_dir>/
├── <namespace>/
│   ├── <resource_kind>/
│   │   ├── <resource_name>.yaml
│   │   └── ...
│   └── ...
└── _cluster/
    ├── <resource_kind>/
    │   ├── <resource_name>.yaml
    │   └── ...
    └── ...
```

- **Namespaced resources**: `<output_dir>/<namespace>/<resource_kind>/<resource_name>.yaml`
- **Cluster-scoped resources**: `<output_dir>/_cluster/<resource_kind>/<resource_name>.yaml`

### Command Line Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--output-dir` | `-o` | Output directory path | `./kalco-dump-<timestamp>` |
| `--kubeconfig` | | Path to kubeconfig file | Auto-detected |
| `--commit-message` | | Custom Git commit message | Timestamp-based |
| `--git-push` | | Auto-push to remote origin | `false` |
| `--help` | `-h` | Show help information | |


---

*Made with ❤️ for the Kubernetes community*
