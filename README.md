```
██╗  ██╗ █████╗ ██╗      ██████╗ ██████╗ 
██║ ██╔╝██╔══██╗██║     ██╔════╝██╔═══██╗
█████╔╝ ███████║██║     ██║     ██║   ██║
██╔═██╗ ██╔══██║██║     ██║     ██║   ██║
██║  ██╗██║  ██║███████╗╚██████╗╚██████╔╝
╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝ ╚═════╝ ╚═════╝ 
```

> Extract, validate, analyze, and version control your entire cluster with comprehensive validation and Git integration

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

## Use Cases & Scenarios

### Enterprise & Production
- **Disaster Recovery** - Complete cluster backups with validation
- **Compliance Auditing** - Track infrastructure changes over time
- **Environment Migration** - Move workloads between clusters safely
- **Security Analysis** - Audit resource configurations and permissions
- **Capacity Planning** - Analyze resource usage patterns

### Development & Testing
- **Environment Replication** - Clone production setups for testing
- **Infrastructure Documentation** - Generate living documentation
- **Development Cleanup** - Identify and remove test resources
- **Configuration Debugging** - Validate resource dependencies
- **Resource Cataloging** - Maintain inventory of all resources

### DevOps & SRE
- **Change Tracking** - Monitor infrastructure evolution
- **Troubleshooting** - Validate cluster configurations
- **Maintenance** - Clean up orphaned and unused resources
- **Onboarding** - Document cluster setups for new team members
- **Rollback Support** - Quick recovery from configuration issues

## Quick Start

### Prerequisites

- **Go 1.21+** - [Download here](https://golang.org/dl/)
- **Kubernetes Access** - Valid kubeconfig or in-cluster access
- **Git** - For version control functionality (optional but recommended)
- **KIND** - For local testing (optional) - [Installation Guide](https://kind.sigs.k8s.io/docs/user/quick-start/)

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

## Features

### Git Integration & Version Control

Kalco automatically sets up Git version control for your cluster snapshots, providing a complete history of changes over time.

#### Automatic Git Workflow

1. **Repository Initialization** - Automatically creates new Git repos for new directories
2. **Change Detection** - Only commits when there are actual changes to track
3. **Smart Committing** - Uses timestamp-based commit messages or custom messages
4. **Remote Integration** - Automatically pushes to remote origin if available

#### Git Usage Examples

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

### Enhanced Change Reports

Kalco automatically generates comprehensive change reports for every cluster snapshot, providing detailed insights into what actually changed in each resource.

#### What Reports Include

##### **Initial Snapshot Reports**
- **Resource Summary** - Complete overview of all exported resources
- **Namespace Coverage** - List of all namespaces and resource types
- **Timestamp Information** - When the snapshot was taken
- **Git Setup** - Confirmation of repository initialization

##### **Change Tracking Reports**
- **Change Summary** - Total files changed, namespaces affected, resource types modified
- **Detailed Changes** - File-by-file breakdown of modifications
- **Namespace Grouping** - Changes organized by namespace and resource type
- **Resource Statistics** - Counts of new, modified, and deleted resources
- **Git Commands** - Reference commands for further investigation

##### **Enhanced Resource Details**
- **New Resources** - Complete YAML content of newly created resources
- **Deleted Resources** - Full YAML content of removed resources
- **Modified Resources** - Git diff output showing exact changes with before/after comparisons
- **Change Analysis** - Human-readable summary of what sections and fields were modified
- **Field-Level Tracking** - Identification of specific YAML sections that changed

#### Report File Naming

Reports are automatically named based on your commit messages:
- **Custom Message**: `Production-backup-2025-08-13.md`
- **Timestamp Default**: `Cluster-snapshot-2025-08-13-15-04-05.md`
- **Special Characters**: Automatically cleaned for valid filenames

### Cross-Reference Validation

Kalco automatically validates cross-references between exported resources to identify potential issues:

#### **What Gets Validated:**

- **Service Selectors**: Services targeting non-existent Pods/Deployments
- **RoleBinding Subjects**: ServiceAccount references in RBAC
- **Network Policies**: Pod selector references
- **Ingress Backends**: Service references in Ingress rules
- **HPA Targets**: Scale target references
- **PDB Selectors**: Pod selector references

#### **Validation Results:**

- **Valid References**: Properly configured cross-references
- **Broken References**: Missing target resources (will cause errors)
- **Warning References**: External references requiring manual verification

#### **Benefits:**

- **Prevents Errors**: Catch issues before reapplying resources
- **Silent Failures**: Find configuration problems kubectl apply won't detect
- **Actionable Insights**: Clear recommendations for fixing issues
- **Reliability**: Ensure cluster resources can be safely reapplied

### Orphaned Resource Detection

Kalco automatically detects resources that are no longer managed by higher-level controllers and may be consuming unnecessary resources:

#### **What Gets Detected:**

- **Orphaned ReplicaSets**: ReplicaSets not owned by Deployments
- **Orphaned Pods**: Pods without controller owners (excluding static/mirror pods)
- **Orphaned ConfigMaps**: ConfigMaps not referenced by any Pod/Deployment
- **Orphaned Secrets**: Secrets not referenced by any Pod/Deployment
- **Orphaned Services**: Services not referenced by any Pod/Deployment
- **Orphaned PVCs**: PersistentVolumeClaims not referenced by any Pod

#### **Detection Results:**

- **Orphaned Resources**: Resources that can be safely cleaned up
- **Resource Breakdown**: Counts by resource type
- **Detailed Analysis**: File locations and reasons for orphaned status
- **Cleanup Guidance**: Step-by-step recommendations for safe removal

#### **Benefits:**

- **Cluster Cleanup**: Identify and remove unnecessary resources
- **Resource Savings**: Free up cluster resources and reduce costs
- **Better Organization**: Maintain clean, well-managed clusters
- **Safe Cleanup**: Clear guidance on what can be safely removed

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
