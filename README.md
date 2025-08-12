# 🚀 Kalco

> **Kubernetes Cluster Resource Dumper** - Extract, organize, and backup your entire cluster with style! 🎯

## 🎯 Overview

**Kalco** is a powerful Go-based CLI tool that performs comprehensive dumps of all resources from your Kubernetes cluster into beautifully organized YAML files. It automatically discovers and exports every available API resource - including native Kubernetes resources and Custom Resources (CRDs) - with zero configuration required.

## ✨ Features

- 🎯 **Complete Resource Discovery**: Automatically discovers ALL available API resources in your cluster
- 🔍 **Comprehensive Coverage**: Includes both native Kubernetes resources and Custom Resources (CRDs)
- 📁 **Structured Output**: Creates beautifully organized directory structures for easy navigation
- 🌐 **Flexible Configuration**: Works seamlessly both in-cluster and out-of-cluster
- 🧹 **Clean YAML**: Intelligently removes metadata fields that aren't useful for re-application
- ⚡ **Lightning Fast**: Optimized for speed and efficiency in production environments

## 📁 Output Structure

The tool creates a beautifully structured directory layout that makes navigation intuitive:

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

## 🚀 Installation

### 📋 Prerequisites

- 🐹 Go 1.21 or later
- ☸️ Access to a Kubernetes cluster

### 🔨 Build from Source

```bash
git clone <repository-url>
cd kalco
go mod tidy
go build -o kalco
```

## 🚀 Quick Start: Complete Git Workflow Demo

### 📋 Prerequisites
- 🐹 Go 1.21 or later
- ☸️ KIND (Kubernetes in Docker) - [Installation Guide](https://kind.sigs.k8s.io/docs/user/quick-start/)
- 🔑 Sufficient permissions to list and read resources
- 📦 Git (for version control functionality)

### 🎯 What You'll Learn

This comprehensive quick start demonstrates kalco's complete Git workflow:
- **🆕 Automatic Git Setup**: New directories become Git repositories automatically
- **🔄 Change Tracking**: See exactly what changed between cluster snapshots
- **📈 Version History**: Maintain complete audit trail of cluster changes
- **🌐 Remote Integration**: Push changes to remote repositories automatically
- **💡 Best Practices**: Learn how to use kalco in production environments

### 🧪 Complete Git Workflow Test

We've created a comprehensive test script that demonstrates everything step-by-step:

```bash
# Run the complete Git workflow test
./examples/test-git-workflow.sh
```

This script will:
1. 🏗️ Create a KIND cluster with test resources
2. 📦 Export resources (auto-creates Git repo)
3. 🔄 Modify cluster resources
4. 📦 Export again (updates existing Git repo)
5. 📊 Show Git history and changes
6. 🧹 Clean up test environment

### 🚀 Manual Step-by-Step Guide

If you prefer to run the steps manually, here's the complete workflow:

#### **Step 1: Create Test Environment**
```bash
# Create KIND cluster
kind create cluster --name kalco-git-demo

# Create namespace and resources
kubectl create namespace demo-apps
kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
  namespace: demo-apps
data:
  environment: "development"
  log-level: "info"
  version: "1.0.0"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  namespace: demo-apps
  labels:
    app: nginx
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.21
        ports:
        - containerPort: 80
EOF
```

**What This Does:**
- Creates a local Kubernetes cluster using KIND
- Sets up a namespace called `demo-apps`
- Creates a ConfigMap with application configuration
- Deploys an nginx application with 1 replica

#### **Step 2: First Export (Creates Git Repo + Report)**
```bash
# Export to new directory - kalco will auto-create Git repo and generate report
./kalco --output-dir ./cluster-history --commit-message "Initial snapshot: $(date)"
```

**What This Does:**
- Creates `./cluster-history` directory (doesn't exist yet)
- Exports ALL cluster resources (system + user resources)
- Automatically initializes a Git repository
- Creates `.gitignore` file for Kubernetes dumps
- Commits all resources with your custom message
- **📊 Generates detailed change report** in `kalco-reports/` folder
- Shows progress, Git status, and report generation

**Expected Output:**
```
🚀 Starting Kubernetes cluster resource dump...
🔍 Discovering resources and building directory structure...
  📡 Discovering API resources...
  ✅ Found 21 API resource groups
  🏷️  Enumerating namespaces...
  ✅ Found 7 namespaces
  🔄 Processing resource groups...
📦 Setting up Git repository for version control...
  🆕 New directory detected, initializing Git repository...
  ✅ Git repository initialized successfully
  📝 Committed changes: Initial snapshot: 2025-08-13 10:30:00
  ℹ️  No remote origin found
📊 Generating cluster change report...
  📊 Generated change report: Initial-snapshot-2025-08-13-10-30-00.md
```

#### **Step 3: Verify Git Repository and Reports**
```bash
cd ./cluster-history

# Check Git status
git status

# View commit history
git log --oneline

# Explore directory structure
ls -la

# Check the generated reports
ls -la kalco-reports/
cat kalco-reports/*.md

cd ..
```

**What This Shows:**
- Git repository was automatically created
- All resources are committed
- Clean working directory (no uncommitted changes)
- Organized structure by namespace and resource type
- **📊 Detailed change report** automatically generated
- **📋 Markdown documentation** of the initial snapshot

#### **Step 4: Modify Cluster Resources**
```bash
# Update ConfigMap
kubectl apply -f - <<EOF
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
  namespace: demo-apps
data:
  environment: "staging"
  log-level: "debug"
  version: "1.1.0"
  feature-flags: "new-feature=true"
EOF

# Scale deployment
kubectl scale deployment nginx-deployment --namespace demo-apps --replicas=3

# Create new Secret
kubectl apply -f - <<EOF
apiVersion: v1
kind: Secret
metadata:
  name: app-secret
  namespace: demo-apps
type: Opaque
data:
  api-key: YXBpLWtleS1zdGFnaW5n
  password: cGFzc3dvcmQtc3RhZ2luZw==
EOF
```

**What This Does:**
- Changes environment from "development" to "staging"
- Updates log level from "info" to "debug"
- Adds new version and feature flags
- Scales nginx from 1 to 3 replicas
- Creates a new Secret resource

#### **Step 5: Second Export (Updates Git Repo + Generates Change Report)**
```bash
# Export again to same directory - kalco will update existing Git repo and generate change report
./kalco --output-dir ./cluster-history --commit-message "Updated resources: $(date)"
```

**What This Does:**
- Detects existing Git repository
- Exports updated cluster resources
- Shows what changed since last export
- Creates new commit with changes
- Maintains complete history
- **📊 Generates detailed change report** comparing with previous snapshot
- **🔄 Tracks all modifications** with file-by-file breakdown

**Expected Output:**
```
📦 Setting up Git repository for version control...
  📦 Using existing Git repository
  📝 Committed changes: Updated resources: 2025-08-13 10:35:00
  ℹ️  No remote origin found
📊 Generating cluster change report...
  📊 Generated change report: Updated-resources-2025-08-13-10-35-00.md
```

#### **Step 6: Analyze Git History**
```bash
cd ./cluster-history

# View all commits
git log --oneline

# See what files changed
git diff HEAD~1 HEAD --name-only

# View specific changes
git diff HEAD~1 HEAD -- "demo-apps/ConfigMap/app-config.yaml"

cd ..
```

**What This Shows:**
- Two commits: initial snapshot and updates
- Changed files: ConfigMap, Deployment, new Secret
- Detailed diff of what changed in each resource
- Complete audit trail of cluster evolution

#### **Step 7: Test Remote Integration (Optional)**
```bash
cd ./cluster-history

# Add remote origin (replace with your repo URL)
git remote add origin https://github.com/yourusername/cluster-backups.git

# Push to remote
git push -u origin main

cd ..

# Now kalco can auto-push
./kalco --output-dir ./cluster-history --git-push
```

**What This Does:**
- Connects local Git repo to remote repository
- Pushes cluster history to remote storage
- Enables automatic pushing with `--git-push` flag
- Provides backup and collaboration capabilities

#### **Step 8: Cleanup**
```bash
# Delete KIND cluster
kind delete cluster --name kalco-git-demo

# Keep cluster-history directory for future use
echo "Your cluster history is preserved in ./cluster-history/"
```

### 🎯 What You've Accomplished

By following this guide, you've learned:

1. **🆕 Automatic Git Setup**: kalco creates Git repos automatically for new directories
2. **🔄 Incremental Versioning**: Each export creates a new commit with changes
3. **📊 Change Tracking**: See exactly what changed between snapshots
4. **🌐 Remote Integration**: Push cluster history to remote repositories
5. **📈 Audit Trail**: Maintain complete history of cluster evolution
6. **💡 Best Practices**: Use kalco effectively in production environments

### 🚀 Production Usage

Now you can use kalco in production:

```bash
# Daily backups with Git history
./kalco --output-dir ./production-backups --commit-message "Daily backup $(date)"

# Before/after deployments
./kalco --output-dir ./deployment-history --commit-message "Pre-deployment snapshot"
# ... deploy changes ...
./kalco --output-dir ./deployment-history --commit-message "Post-deployment snapshot"

# Auto-push to remote backup
./kalco --output-dir ./production-backups --git-push
```

Your cluster changes are now fully version-controlled and backed up! 🎉

### 💻 Basic Usage

```bash
# Dump all resources to default output directory
./kalco

# Specify custom output directory
./kalco --output-dir ./my-dump

# Use specific kubeconfig file
./kalco --kubeconfig ~/.kube/config --output-dir ./cluster-dump
```

### 🧪 Automated Testing Examples

We provide comprehensive examples to help you learn and test kalco:

#### **Complete Git Workflow Demo**
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

Perfect for learning kalco's Git capabilities, reporting features, and testing your setup!

### 📊 Test Results

Our quick start test successfully demonstrated kalco's capabilities:

- **✅ 21 API Resource Groups** processed successfully
- **✅ 7 Namespaces** discovered and organized
- **✅ 15+ Resource Types** exported (Pods, Deployments, Services, ConfigMaps, etc.)
- **✅ Clean YAML Output** with metadata properly cleaned
- **✅ Hierarchical Organization** by namespace and resource type
- **✅ Cluster-scoped Resources** properly separated in `_cluster/` directory
- **✅ Automatic Git Repository** initialization and management
- **✅ Comprehensive Change Reports** generated for every snapshot

The tool successfully exported resources from:
- `test-apps` namespace (our test resources)
- `kube-system` namespace (system components)
- `default` namespace (default resources)
- `_cluster` directory (cluster-wide resources like Nodes, ClusterRoles, etc.)

> **💡 Tip**: The test cluster is automatically cleaned up after the demo. In production, you can use kalco on any Kubernetes cluster without affecting the running workloads.

## 📦 Git Version Control

Kalco automatically sets up Git version control for your cluster snapshots, providing a complete history of changes over time.

### 🔄 Automatic Git Workflow

1. **Repository Initialization**: Automatically initializes a new Git repository if none exists
2. **Change Detection**: Only commits when there are actual changes to track
3. **Smart Committing**: Uses timestamp-based commit messages or custom messages
4. **Remote Integration**: Automatically pushes to remote origin if available

### 📝 Git Features

- **🆕 New Repositories**: Automatically initialized with proper `.gitignore`
- **🔄 Existing Repositories**: Seamlessly works with previously created repos
- **📅 Timestamp Tracking**: Each snapshot gets a unique timestamp-based commit
- **✏️ Custom Messages**: Override default messages with `--commit-message`
- **🚀 Auto Push**: Use `--git-push` to automatically push to remote origin

### 💻 Git Usage Examples

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

### 📁 Git Repository Structure

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

## 📊 Automatic Change Reports

Kalco automatically generates detailed markdown reports for every cluster snapshot, providing comprehensive change tracking and documentation.

### 🔍 What Reports Include

#### **Initial Snapshot Reports**
- **📋 Resource Summary**: Complete overview of all exported resources
- **🏷️ Namespace Coverage**: List of all namespaces and resource types
- **📅 Timestamp Information**: When the snapshot was taken
- **🔧 Git Setup**: Confirmation of repository initialization

#### **Change Tracking Reports**
- **📊 Change Summary**: Total files changed, namespaces affected, resource types modified
- **🔄 Detailed Changes**: File-by-file breakdown of modifications
- **🌐 Namespace Grouping**: Changes organized by namespace and resource type
- **📈 Resource Statistics**: Counts of new, modified, and deleted resources
- **💻 Git Commands**: Reference commands for further investigation

### 📁 Report File Naming

Reports are automatically named based on your commit messages:
- **Custom Message**: `Production-backup-2025-08-13.md`
- **Timestamp Default**: `Cluster-snapshot-2025-08-13-15-04-05.md`
- **Special Characters**: Automatically cleaned for valid filenames

### 📋 Report Content Example

```markdown
# Cluster Change Report

**Generated**: 2025-08-13 15:04:05 UTC
**Commit Message**: Production backup

## Changes Since Previous Snapshot

**Previous Commit**: `abc1234`

### Change Summary
- **Total Files Changed**: 15
- **Namespaces Affected**: 3
- **Resource Types Changed**: 4
- **New Resources**: 2
- **Modified Resources**: 13
- **Deleted Resources**: 0

### Detailed Changes

#### 📁 Namespace: `production`
**ConfigMap**:
- ✏️ `app-config.yaml`
- 🆕 `feature-flags.yaml`

**Deployment**:
- ✏️ `web-app.yaml`

#### 🌐 Cluster-Scoped Resources
**StorageClass**:
- ✏️ `fast-storage.yaml`

### Git Commands for Reference
```bash
# View this commit
git show def5678

# Compare with previous snapshot
git diff abc1234..def5678

# View file changes
git diff --name-status abc1234..def5678
```

### 🎯 Benefits of Git Integration

- **📈 Complete History**: Track every cluster change over time
- **🔄 Rollback Capability**: Easily revert to previous cluster states
- **👥 Collaboration**: Share cluster snapshots via Git remotes
- **📊 Change Tracking**: See exactly what changed between snapshots
- **🔒 Version Control**: Maintain audit trail for compliance
- **📋 Detailed Reports**: Automatic markdown reports for every snapshot

### ⚙️ Command Line Options

- `--kubeconfig`: Path to the kubeconfig file (optional)
- `--output-dir, -o`: Path to the output directory (default: `./kalco-dump-<timestamp>`)
- `--commit-message`: Custom commit message (default: timestamp-based message)
- `--git-push`: Automatically push changes to remote origin if available

### 📝 Examples

```bash
# Dump to a specific directory
./kalco -o ./production-backup

# Use a specific kubeconfig
./kalco --kubeconfig /path/to/kubeconfig -o ./staging-backup

# In-cluster execution (when running inside a pod)
./kalco -o /tmp/cluster-dump

# Git version control with custom commit message
./kalco -o ./cluster-backup --commit-message "Production backup $(date)"

# Auto-push to remote origin
./kalco -o ./cluster-backup --git-push

# Full customization
./kalco -o ./cluster-backup --commit-message "Monthly audit" --git-push
```

## 🔧 How It Works

1. 🚀 **Client Creation**: Creates Kubernetes clients (clientset, discovery client, and dynamic client)
2. 🔍 **Resource Discovery**: Uses the discovery client to get all server resources
3. 🏷️ **Namespace Enumeration**: Lists all namespaces for namespaced resources
4. 📊 **Resource Dumping**: For each resource type:
   - If namespaced: Lists all instances across all namespaces
   - If cluster-scoped: Lists all instances at cluster level
5. 📄 **YAML Export**: Converts each resource to clean YAML and writes to appropriate directory
6. 🧹 **Metadata Cleanup**: Removes fields like `uid`, `resourceVersion`, `managedFields`, `status`, etc.

## 🛡️ Error Handling

The tool is designed to be resilient and production-ready:
- ⚡ Continues processing even if individual resources fail to dump
- ⚠️ Provides clear warning messages for failed operations
- 🚀 Gracefully handles API errors and permission issues

## 🛠️ Development

### 📁 Project Structure

```
kalco/
├── 📂 cmd/
│   └── 🎯 root.go          # Main CLI command definition
├── 📂 pkg/
│   ├── 🌐 kube/
│   │   └── 🔌 client.go    # Kubernetes client creation
│   └── 📊 dumper/
│       └── 🚀 dumper.go    # Core resource dumping logic
├── 🚀 main.go              # Application entry point
├── 📦 go.mod               # Go module definition
└── 📖 README.md            # This file
```

### 📦 Dependencies

- 🎯 `github.com/spf13/cobra`: CLI framework
- 🌐 `k8s.io/client-go`: Kubernetes client library
- ⚙️ `k8s.io/apimachinery`: Kubernetes API machinery
- 📄 `gopkg.in/yaml.v3`: YAML processing

## License

[Add your license information here]

## Contributing

[Add contribution guidelines here]
