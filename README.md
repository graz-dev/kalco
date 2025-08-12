# 🚀 Kalco

> **Kubernetes Cluster Resource Dumper** - Extract, organize, and version control your entire cluster with Git integration

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/kalco)](https://goreportcard.com/report/github.com/yourusername/kalco)

## 📖 Overview

**Kalco** is a powerful, production-ready Go CLI tool that performs comprehensive dumps of all resources from your Kubernetes cluster into beautifully organized YAML files. It automatically discovers and exports every available API resource - including native Kubernetes resources and Custom Resources (CRDs) - with zero configuration required.

**Key Features:**
- 🎯 **Complete Resource Discovery** - Automatically finds ALL available API resources
- 🔍 **Comprehensive Coverage** - Includes both native K8s resources and Custom Resources (CRDs)
- 📁 **Structured Output** - Creates intuitive directory structures for easy navigation
- 🌐 **Flexible Configuration** - Works seamlessly both in-cluster and out-of-cluster
- 🧹 **Clean YAML** - Intelligently removes metadata fields that aren't useful for re-application
- ⚡ **Lightning Fast** - Optimized for speed and efficiency in production environments
- 🚀 **Git Integration** - Automatic version control with commit history and change tracking
- 📊 **Smart Reporting** - Generates detailed change reports for every snapshot

## 🚀 Quick Start

### Prerequisites

- 🐹 **Go 1.21+** - [Download here](https://golang.org/dl/)
- ☸️ **Kubernetes Access** - Valid kubeconfig or in-cluster access
- 🔑 **Git** - For version control functionality (optional but recommended)

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/kalco.git
cd kalco

# Install dependencies and build
go mod tidy
go build -o kalco

# Make it available system-wide (optional)
sudo mv kalco /usr/local/bin/
```

### Basic Usage

```bash
# Dump all resources to default output directory
./kalco

# Specify custom output directory
./kalco --output-dir ./my-cluster-dump

# Use specific kubeconfig file
./kalco --kubeconfig ~/.kube/config --output-dir ./cluster-backup
```

## 🎯 Git Integration & Version Control

Kalco automatically sets up Git version control for your cluster snapshots, providing a complete history of changes over time.

### Automatic Git Workflow

1. **🆕 Repository Initialization** - Automatically creates new Git repos for new directories
2. **🔄 Change Detection** - Only commits when there are actual changes to track
3. **📅 Smart Committing** - Uses timestamp-based commit messages or custom messages
4. **🌐 Remote Integration** - Automatically pushes to remote origin if available

### Git Usage Examples

```bash
# Basic export with Git version control
./kalco --output-dir ./cluster-backup

# Custom commit message
./kalco --output-dir ./cluster-backup --commit-message "Production cluster backup"

# Auto-push to remote origin
./kalco --output-dir ./cluster-backup --git-push

# Full customization
./kalco --output-dir ./cluster-backup --commit-message "Monthly audit" --git-push
```

### Git Repository Structure

```
cluster-backup/
├── .git/                    # Git repository
├── .gitignore              # Auto-generated ignore file
├── kalco-reports/          # Change reports for each snapshot
│   ├── Initial-snapshot.md # First export report
│   └── Updated-resources.md # Change tracking report
├── default/                 # Namespace resources
├── kube-system/            # System resources
├── _cluster/               # Cluster-scoped resources
└── README.md               # Repository documentation
```

## 📊 Output Structure

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

- 🏷️ **Namespaced resources**: `<output_dir>/<namespace>/<resource_kind>/<resource_name>.yaml`
- 🌐 **Cluster-scoped resources**: `<output_dir>/_cluster/<resource_kind>/<resource_name>.yaml`

## 📋 Command Line Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--output-dir` | `-o` | Output directory path | `./kalco-dump-<timestamp>` |
| `--kubeconfig` | | Path to kubeconfig file | Auto-detected |
| `--commit-message` | | Custom Git commit message | Timestamp-based |
| `--git-push` | | Auto-push to remote origin | `false` |
| `--help` | `-h` | Show help information | |

## 🧪 Examples & Testing

### Complete Git Workflow Demo

We provide a comprehensive test script that demonstrates kalco's capabilities:

```bash
# Run the automated Git workflow test
./examples/test-git-workflow.sh
```

This script demonstrates:
- 🏗️ KIND cluster creation and setup
- 📦 Automatic Git repository initialization
- 🔄 Resource modification and change tracking
- 📊 Git history analysis and verification
- 📋 Automatic change report generation
- 🌐 Remote integration guidance
- 🧹 Proper cleanup and summary

### Manual Testing Example

```bash
# Create test cluster (requires KIND)
kind create cluster --name kalco-test

# Export resources (auto-creates Git repo)
./kalco --output-dir ./test-backup --commit-message "Initial snapshot"

# Modify cluster resources
kubectl create namespace test-apps
kubectl create configmap app-config --from-literal=env=dev -n test-apps

# Export again (updates existing Git repo)
./kalco --output-dir ./test-backup --commit-message "Added test resources"

# Cleanup
kind delete cluster --name kalco-test
```

## 🔧 How It Works

1. **🚀 Client Creation** - Creates Kubernetes clients (clientset, discovery client, dynamic client)
2. **🔍 Resource Discovery** - Uses discovery client to get all server resources
3. **🏷️ Namespace Enumeration** - Lists all namespaces for namespaced resources
4. **📊 Resource Dumping** - For each resource type:
   - If namespaced: Lists all instances across all namespaces
   - If cluster-scoped: Lists all instances at cluster level
5. **📄 YAML Export** - Converts each resource to clean YAML and writes to appropriate directory
6. **🧹 Metadata Cleanup** - Removes fields like `uid`, `resourceVersion`, `managedFields`, `status`, etc.
7. **🚀 Git Integration** - Initializes repository, commits changes, and optionally pushes to remote

## 🛡️ Error Handling & Resilience

Kalco is designed to be production-ready:
- ⚡ **Continues Processing** - Handles individual resource failures gracefully
- ⚠️ **Clear Warnings** - Provides informative messages for failed operations
- 🚀 **Graceful Degradation** - Manages API errors and permission issues
- 📊 **Progress Reporting** - Shows real-time status of operations

## 🛠️ Development

### Project Structure

```
kalco/
├── 📂 cmd/
│   └── 🎯 root.go          # Main CLI command definition
├── 📂 pkg/
│   ├── 🌐 kube/
│   │   └── 🔌 client.go    # Kubernetes client creation
│   ├── 📊 dumper/
│   │   └── 🚀 dumper.go    # Core resource dumping logic
│   ├── 🚀 git/
│   │   └── 🔄 git.go       # Git integration logic
│   └── 📋 reports/
│       └── 📊 reports.go    # Change report generation
├── 🚀 main.go              # Application entry point
├── 📦 go.mod               # Go module definition
└── 📖 README.md            # This file
```

### Dependencies

- 🎯 `github.com/spf13/cobra` - CLI framework
- 🌐 `k8s.io/client-go` - Kubernetes client library
- ⚙️ `k8s.io/apimachinery` - Kubernetes API machinery
- 📄 `gopkg.in/yaml.v3` - YAML processing

### Building & Testing

```bash
# Build the application
go build -o kalco

# Run tests
go test ./...

# Run with race detection
go test -race ./...

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o kalco-linux
GOOS=darwin GOARCH=amd64 go build -o kalco-darwin
```

## 📚 Use Cases

### 🏢 Production Environments
- **Daily Backups** - Automated cluster snapshots with version control
- **Deployment Tracking** - Before/after snapshots for deployments
- **Compliance Auditing** - Maintain complete cluster change history
- **Disaster Recovery** - Quick cluster state restoration from YAML

### 🧪 Development & Testing
- **Environment Comparison** - Compare cluster states across environments
- **Resource Templates** - Extract working configurations for reuse
- **Debugging** - Export cluster state for offline analysis
- **Documentation** - Generate cluster resource documentation

### 🔄 CI/CD Integration
- **Pipeline Snapshots** - Capture cluster state at each pipeline stage
- **Rollback Verification** - Confirm rollback to previous cluster state
- **Change Tracking** - Monitor cluster modifications over time

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

### Development Setup

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes and add tests
4. Commit your changes: `git commit -m 'Add amazing feature'`
5. Push to the branch: `git push origin feature/amazing-feature`
6. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Kubernetes community for the excellent client libraries
- Cobra team for the powerful CLI framework
- All contributors who help improve this tool

## 📞 Support

- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/yourusername/kalco/issues)
- 💡 **Feature Requests**: [GitHub Discussions](https://github.com/yourusername/kalco/discussions)
- 📖 **Documentation**: [GitHub Wiki](https://github.com/yourusername/kalco/wiki)
- 💬 **Community**: [GitHub Discussions](https://github.com/yourusername/kalco/discussions)

---

**Made with ❤️ for the Kubernetes community**
