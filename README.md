# ☸️ Kalco

> **Kubernetes Cluster Resource Dumper** - Extract, organize, and version control your entire cluster with Git integration

## 📖 Overview

**Kalco** (Kubernetes Analysis and Lifecycle Control) is a powerful, production-ready Go CLI tool that performs comprehensive dumps of all resources from your Kubernetes cluster into beautifully organized YAML files. It automatically discovers and exports every available API resource - including native Kubernetes resources and Custom Resources (CRDs) - with zero configuration required.

### 🎯 **What Kalco Does**

Kalco transforms your Kubernetes cluster into a **version-controlled, validated, and organized** backup that you can:
- 🔄 **Reapply** to any cluster with confidence
- 📊 **Audit** for compliance and security
- 🧹 **Clean up** by identifying orphaned resources
- 📚 **Document** your infrastructure as code
- 🚀 **Migrate** between environments safely

### 🚀 **Key Features**

- 🎯 **Complete Resource Discovery** - Automatically finds ALL available API resources
- 🔍 **Comprehensive Coverage** - Includes both native K8s resources and Custom Resources (CRDs)
- 📁 **Structured Output** - Creates intuitive directory structures for easy navigation
- 🌐 **Flexible Configuration** - Works seamlessly both in-cluster and out-of-cluster
- 🧹 **Clean YAML** - Intelligently removes metadata fields that aren't useful for re-application
- ⚡ **Lightning Fast** - Optimized for speed and efficiency in production environments
- 🚀 **Git Integration** - Automatic version control with commit history and change tracking
- 📊 **Smart Reporting** - Generates detailed change reports with before/after comparisons and specific field modifications
- 🔍 **Cross-Reference Validation** - Analyzes exported resources for broken references that could cause issues when reapplying
- 🗑️ **Orphaned Resource Detection** - Identifies resources no longer managed by higher-level controllers for cleanup

---

## ✨ **Core Features Deep Dive**

### 🎯 **Resource Discovery & Export**
- **🔍 Automatic API Discovery** - Dynamically finds all available resource types
- **📦 Complete Resource Coverage** - Exports every resource in every namespace
- **🌐 Cluster-Scoped Resources** - Handles both namespaced and cluster-wide resources
- **🔧 Custom Resource Support** - Full CRD compatibility with zero configuration
- **📁 Intelligent Organization** - Creates logical directory structures by namespace and type

### 🚀 **Git Integration & Version Control**
- **🆕 Automatic Repository Setup** - Creates Git repos for new directories
- **📝 Smart Commit Messages** - Timestamp-based or custom commit messages
- **🔄 Change Detection** - Only commits when there are actual changes
- **🌐 Remote Integration** - Automatic push to remote origin if available
- **📊 Complete History** - Full audit trail of all cluster changes

### 📊 **Advanced Reporting & Validation**
- **🔍 Cross-Reference Validation** - Identifies broken resource dependencies
- **🗑️ Orphaned Resource Detection** - Finds resources without owners or references
- **📋 Detailed Change Reports** - Before/after comparisons with specific field changes
- **💡 Actionable Insights** - Clear recommendations for fixing issues
- **🛡️ Reliability Assurance** - Ensures resources can be safely reapplied

---

## 🎯 **Use Cases & Scenarios**

### 🏢 **Enterprise & Production**
- **🔄 Disaster Recovery** - Complete cluster backups with validation
- **📊 Compliance Auditing** - Track infrastructure changes over time
- **🚀 Environment Migration** - Move workloads between clusters safely
- **🛡️ Security Analysis** - Audit resource configurations and permissions
- **📈 Capacity Planning** - Analyze resource usage patterns

### 🧪 **Development & Testing**
- **🔄 Environment Replication** - Clone production setups for testing
- **📚 Infrastructure Documentation** - Generate living documentation
- **🧹 Development Cleanup** - Identify and remove test resources
- **🔍 Configuration Debugging** - Validate resource dependencies
- **📋 Resource Cataloging** - Maintain inventory of all resources

### 🚀 **DevOps & SRE**
- **📊 Change Tracking** - Monitor infrastructure evolution
- **🛠️ Troubleshooting** - Validate cluster configurations
- **🧹 Maintenance** - Clean up orphaned and unused resources
- **📋 Onboarding** - Document cluster setups for new team members
- **🔄 Rollback Support** - Quick recovery from configuration issues

---

## 💡 **Examples & Common Workflows**

### 🔄 **Daily Cluster Backup**
```bash
# Create daily backup with timestamp
./kalco --output-dir ./daily-backups/$(date +%Y%m%d) \
        --commit-message "Daily backup $(date)"

# This creates: ./daily-backups/20250813/ with Git history
```

### 🚀 **Production Migration**
```bash
# Export production cluster
./kalco --output-dir ./prod-cluster \
        --commit-message "Production cluster export"

# Validate the export
cd ./prod-cluster
# Review the validation report in kalco-reports/

# Apply to staging cluster
kubectl apply -f ./
```

### 🧹 **Cluster Cleanup**
```bash
# Export current state
./kalco --output-dir ./cleanup-audit

# Review orphaned resources in the report
# Manually remove identified orphaned resources

# Export again to verify cleanup
./kalco --output-dir ./cleanup-audit \
        --commit-message "Post-cleanup verification"
```

### 📊 **Change Tracking Over Time**
```bash
# Initial export
./kalco --output-dir ./cluster-evolution \
        --commit-message "Initial cluster state"

# After making changes
./kalco --output-dir ./cluster-evolution \
        --commit-message "Added monitoring stack"

# View change history
cd ./cluster-evolution
git log --oneline
git show HEAD --name-only
```

---

## 🔧 **Troubleshooting & Best Practices**

### 🚨 **Common Issues & Solutions**

#### **Permission Denied Errors**
```bash
# Ensure you have proper RBAC permissions
kubectl auth can-i get pods --all-namespaces
kubectl auth can-i get deployments --all-namespaces

# If using service account, verify it has necessary roles
kubectl get serviceaccount -n default
kubectl get rolebinding -n default
```

#### **Resource Export Failures**
```bash
# Check cluster connectivity
kubectl cluster-info
kubectl get nodes

# Verify API server accessibility
kubectl get apiservices

# Check for resource quota issues
kubectl describe resourcequota -A
```

#### **Git Integration Issues**
```bash
# Verify Git is installed and configured
git --version
git config --list

# Check repository status
cd ./output-directory
git status
git remote -v
```

### 💡 **Best Practices**

#### **🔒 Security**
- **Use dedicated service accounts** with minimal required permissions
- **Review exported resources** before committing to version control
- **Secure your Git repositories** with proper access controls
- **Regularly rotate credentials** and service account tokens

#### **📊 Performance**
- **Schedule exports during low-traffic periods** for production clusters
- **Use separate output directories** for different environments
- **Clean up old exports** to save disk space
- **Monitor resource usage** during large exports

#### **🔄 Workflow**
- **Always validate exports** before reapplying to other clusters
- **Use descriptive commit messages** for better change tracking
- **Review validation reports** to catch issues early
- **Test in staging** before applying to production

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

## 🗑️ Orphaned Resource Detection

Kalco automatically detects resources that are no longer managed by higher-level controllers and may be consuming unnecessary resources:

### **What Gets Detected:**

- **🔗 Orphaned ReplicaSets**: ReplicaSets not owned by Deployments
- **📦 Orphaned Pods**: Pods without controller owners (excluding static/mirror pods)
- **📋 Orphaned ConfigMaps**: ConfigMaps not referenced by any Pod/Deployment
- **🔐 Orphaned Secrets**: Secrets not referenced by any Pod/Deployment
- **🌐 Orphaned Services**: Services not referenced by any Pod/Deployment
- **💾 Orphaned PVCs**: PersistentVolumeClaims not referenced by any Pod

### **Detection Results:**

- **🗑️ Orphaned Resources**: Resources that can be safely cleaned up
- **📊 Resource Breakdown**: Counts by resource type
- **🔍 Detailed Analysis**: File locations and reasons for orphaned status
- **💡 Cleanup Guidance**: Step-by-step recommendations for safe removal

### **Benefits:**

- **🧹 Cluster Cleanup**: Identify and remove unnecessary resources
- **💰 Resource Savings**: Free up cluster resources and reduce costs
- **📚 Better Organization**: Maintain clean, well-managed clusters
- **🛡️ Safe Cleanup**: Clear guidance on what can be safely removed

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

## 🏗️ **Architecture & Design**

### 🔧 **Core Components**

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Kubernetes    │    │      Kalco       │    │   Output & Git  │
│     Cluster     │◄──►│   Core Engine    │───►│   Integration   │
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │                       │
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│  API Discovery  │    │  Resource Dumper │    │  Report Gen.    │
│  & Enumeration  │    │  & Processing    │    │  & Validation   │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

### 📦 **Package Structure**

- **`cmd/root.go`** - CLI command definitions and flag handling
- **`pkg/kube/`** - Kubernetes client creation and connection management
- **`pkg/dumper/`** - Core resource discovery and export logic
- **`pkg/git/`** - Git repository management and version control
- **`pkg/reports/`** - Change report generation and formatting
- **`pkg/validation/`** - Cross-reference validation engine
- **`pkg/orphaned/`** - Orphaned resource detection system

### 🔄 **Data Flow**

1. **🔍 Discovery Phase** - Enumerate all available API resources
2. **📦 Export Phase** - Dump resources to organized YAML files
3. **🔍 Validation Phase** - Analyze cross-references and dependencies
4. **🗑️ Detection Phase** - Identify orphaned and unused resources
5. **📊 Reporting Phase** - Generate comprehensive change reports
6. **🚀 Git Phase** - Commit changes and manage version history

---

## 🤝 **Contributing & Community**

### 🚀 **Getting Started**

We welcome contributions from the community! Here's how to get started:

1. **🔍 Find an Issue** - Check our [Issues](https://github.com/graz-dev/kalco/issues) for bugs or feature requests
2. **🍴 Fork & Clone** - Fork the repository and clone your fork locally
3. **🌿 Create a Branch** - Create a feature branch for your changes
4. **💻 Make Changes** - Implement your changes with tests
5. **🧪 Test Thoroughly** - Ensure all tests pass and add new ones
6. **📝 Submit PR** - Create a pull request with clear description

### 🧪 **Development Setup**

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/kalco.git
cd kalco

# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build the binary
go build -o kalco

# Run linting
golangci-lint run
```

### 📋 **Contribution Guidelines**

- **🎯 Focus on Value** - Ensure changes provide real user value
- **🧪 Test Coverage** - Maintain or improve test coverage
- **📚 Documentation** - Update docs for new features
- **🔍 Code Quality** - Follow Go best practices and project style
- **📝 Clear Commits** - Use descriptive commit messages

### 🐛 **Reporting Issues**

When reporting issues, please include:
- **Kubernetes version** and cluster type
- **Kalco version** and command used
- **Error messages** and stack traces
- **Expected vs actual behavior**
- **Steps to reproduce**

---

## ❓ **Frequently Asked Questions**

### 🔍 **General Questions**

#### **Q: How is Kalco different from `kubectl get --export`?**
**A:** Kalco provides comprehensive resource discovery, validation, Git integration, and intelligent reporting. Unlike basic kubectl exports, Kalco automatically finds all resources, validates dependencies, and creates organized, version-controlled backups.

#### **Q: Does Kalco work with Custom Resource Definitions (CRDs)?**
**A:** Yes! Kalco automatically discovers and exports all CRDs and their instances with zero configuration required. It handles both native Kubernetes resources and custom resources seamlessly.

#### **Q: Can I use Kalco in a CI/CD pipeline?**
**A:** Absolutely! Kalco is designed for automation. You can use it in CI/CD to create cluster snapshots, validate configurations, and track infrastructure changes over time.

### 🚀 **Performance & Scalability**

#### **Q: How long does it take to export a large cluster?**
**A:** Export time depends on cluster size and API server performance. A typical production cluster (1000+ resources) usually exports in 2-5 minutes. Kalco is optimized for speed and efficiency.

#### **Q: Does Kalco consume cluster resources during export?**
**A:** Kalco runs as a client and only reads from the API server. It doesn't create or modify cluster resources, making it safe for production use.

#### **Q: Can I export specific namespaces or resource types?**
**A:** Currently, Kalco exports all resources for comprehensive coverage. This ensures you have complete cluster state for backup and validation purposes.

### 🔒 **Security & Permissions**

#### **Q: What permissions does Kalco need?**
**A:** Kalco needs read access to all resources you want to export. This typically means `get`, `list`, and `watch` permissions on resources across namespaces.

#### **Q: Is it safe to run Kalco in production?**
**A:** Yes, Kalco is read-only and designed for production use. It only reads cluster state and doesn't modify any resources.

#### **Q: How do I handle sensitive data in exports?**
**A:** Kalco exports resources as-is. For sensitive data, consider using Git filters, external secret management, or reviewing exports before committing to version control.

### 🔄 **Git Integration**

#### **Q: What if I don't want Git integration?**
**A:** Git integration is optional. You can use Kalco just for resource export and validation without version control.

#### **Q: Can I use existing Git repositories?**
**A:** Yes! Kalco will use existing Git repositories if present, or create new ones for new directories.

#### **Q: How do I handle large binary files in Git?**
**A:** Kalco exports YAML files which are text-based and Git-friendly. For large clusters, consider using Git LFS or regular cleanup of old exports.

---

## 📚 **Support & Resources**

### 🔗 **Official Resources**
- **📖 Documentation** - This README and inline code documentation
- **🐛 Issue Tracker** - [GitHub Issues](https://github.com/graz-dev/kalco/issues)
- **💬 Discussions** - [GitHub Discussions](https://github.com/graz-dev/kalco/discussions)
- **📋 Roadmap** - [Project Milestones](https://github.com/graz-dev/kalco/milestones)

### 🎯 **Getting Help**
- **🔍 Search Issues** - Check if your question has already been answered
- **📝 Create Issue** - Provide detailed information about your problem
- **💡 Feature Request** - Suggest new features or improvements
- **🐛 Bug Report** - Report bugs with reproduction steps

### 📖 **Additional Documentation**
- **📋 API Reference** - Detailed package documentation
- **🔧 Configuration Guide** - Advanced configuration options
- **🚀 Deployment Guide** - Production deployment best practices
- **🧪 Testing Guide** - How to test and validate exports

### 🌟 **Show Your Support**
- **⭐ Star the Repository** - Show your appreciation
- **🔗 Share with Others** - Help spread the word
- **💻 Contribute Code** - Submit pull requests
- **📚 Improve Docs** - Help make documentation better

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
