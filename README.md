# ☸️ Kalco

> **Kubernetes Cluster Resource Dumper** - Extract, organize, and version control your entire cluster with Git integration

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
- 📊 **Smart Reporting** - Generates detailed change reports with before/after comparisons and specific field modifications
- 🔍 **Cross-Reference Validation** - Analyzes exported resources for broken references that could cause issues when reapplying

---

## 🚀 Quick Start

### Prerequisites

- 🐹 **Go 1.21+** - [Download here](https://golang.org/dl/)
- ☸️ **Kubernetes Access** - Valid kubeconfig or in-cluster access
- 🔑 **Git** - For version control functionality (optional but recommended)
- 🏗️ **KIND** - For local testing (optional) - [Installation Guide](https://kind.sigs.k8s.io/docs/user/quick-start/)

### Installation

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

Want to see kalco in action? Run our comprehensive quickstart:

```bash
# Run the complete quickstart demo
./examples/quickstart.sh
```

This will:
- 🏗️ Create a test Kubernetes cluster
- 📦 Export resources with automatic Git setup
- 🔄 Modify cluster resources
- 📊 Generate enhanced change reports
- 🧹 Clean up the test environment

---

### Basic Usage

```bash
# Dump all resources to default output directory
./kalco

# Specify custom output directory
./kalco --output-dir ./my-cluster-dump

# Use specific kubeconfig file
./kalco --kubeconfig ~/.kube/config --output-dir ./cluster-backup
```

---

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

---

## 📊 Enhanced Change Reports

Kalco automatically generates comprehensive change reports for every cluster snapshot, providing detailed insights into what actually changed in each resource.

### 🔍 What Reports Include

#### **Initial Snapshot Reports**
- **📋 Resource Summary** - Complete overview of all exported resources
- **🏷️ Namespace Coverage** - List of all namespaces and resource types
- **📅 Timestamp Information** - When the snapshot was taken
- **🔧 Git Setup** - Confirmation of repository initialization

#### **Change Tracking Reports**
- **📊 Change Summary** - Total files changed, namespaces affected, resource types modified
- **🔄 Detailed Changes** - File-by-file breakdown of modifications
- **🌐 Namespace Grouping** - Changes organized by namespace and resource type
- **📈 Resource Statistics** - Counts of new, modified, and deleted resources
- **💻 Git Commands** - Reference commands for further investigation

#### **Enhanced Resource Details**
- **🆕 New Resources** - Complete YAML content of newly created resources
- **🗑️ Deleted Resources** - Full YAML content of removed resources
- **✏️ Modified Resources** - Git diff output showing exact changes with before/after comparisons
- **📋 Change Analysis** - Human-readable summary of what sections and fields were modified
- **🔍 Field-Level Tracking** - Identification of specific YAML sections that changed

### 📁 Report File Naming

Reports are automatically named based on your commit messages:
- **Custom Message**: `Production-backup-2025-08-13.md`
- **Timestamp Default**: `Cluster-snapshot-2025-08-13-15-04-05.md`
- **Special Characters**: Automatically cleaned for valid filenames

---

## 🔍 Cross-Reference Validation

Kalco automatically validates cross-references between exported resources to identify potential issues:

### **What Gets Validated:**

- **🔗 Service Selectors**: Services targeting non-existent Pods/Deployments
- **👥 RoleBinding Subjects**: ServiceAccount references in RBAC
- **🌐 Network Policies**: Pod selector references
- **🚪 Ingress Backends**: Service references in Ingress rules
- **📈 HPA Targets**: Scale target references
- **🛡️ PDB Selectors**: Pod selector references

### **Validation Results:**

- **✅ Valid References**: Properly configured cross-references
- **❌ Broken References**: Missing target resources (will cause errors)
- **⚠️  Warning References**: External references requiring manual verification

### **Benefits:**

- **🚫 Prevents Errors**: Catch issues before reapplying resources
- **🔍 Silent Failures**: Find configuration problems kubectl apply won't detect
- **📋 Actionable Insights**: Clear recommendations for fixing issues
- **🛡️ Reliability**: Ensure cluster resources can be safely reapplied

---

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

---

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
├── 🔧 Makefile             # Development and build commands
├── 📖 README.md            # This file

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

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o kalco-linux
GOOS=darwin GOARCH=amd64 go build -o kalco-darwin
```

---

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



**Made with ❤️ for the Kubernetes community**
