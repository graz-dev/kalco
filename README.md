<div align="center">

# ☸️ Kalco - Kubernetes Analysis & Lifecycle Control

**🚀 The ultimate CLI tool for Kubernetes cluster management, analysis, and lifecycle control**

[![Release](https://img.shields.io/github/v/release/graz-dev/kalco)](https://github.com/graz-dev/kalco/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Documentation](https://img.shields.io/badge/docs-available-brightgreen)](https://graz-dev.github.io/kalco)

*Extract, validate, analyze, and version control your entire Kubernetes cluster with comprehensive validation and Git integration*

[🚀 Quick Start](#quick-start) • [📖 Documentation](https://graz-dev.github.io/kalco) • [💡 Examples](#examples) • [🤝 Contributing](#contributing)

</div>

---

## 🌟 Why Kalco? 

Kalco transforms your Kubernetes cluster management experience by providing a **comprehensive, automated, and intelligent** approach to cluster analysis and lifecycle control. Whether you're managing production workloads, ensuring compliance, or planning migrations, Kalco has you covered.

### 🎯 **Perfect for DevOps Teams**
- **Site Reliability Engineers** - Automated cluster backups and disaster recovery
- **Platform Engineers** - Infrastructure as Code and GitOps workflows  
- **Security Teams** - Compliance auditing and security posture analysis
- **Developers** - Environment replication and configuration management

## 🚀 What Makes Kalco Special?

<table>
<tr>
<td width="50%">

### 🔍 **Intelligent Discovery**
- **Zero Configuration** - Works out of the box
- **Complete Coverage** - Native K8s + CRDs
- **Smart Filtering** - Namespace, resource, and label-based
- **Real-time Analysis** - Live cluster insights

</td>
<td width="50%">

### 🛡️ **Enterprise Ready**
- **Git Integration** - Automatic version control
- **Validation Engine** - Cross-reference checking
- **Security Analysis** - Compliance and best practices
- **Scalable Architecture** - Handles clusters of any size

</td>
</tr>
<tr>
<td width="50%">

### 📊 **Actionable Insights**
- **Orphaned Resources** - Identify cleanup opportunities
- **Broken References** - Find configuration issues
- **Usage Analytics** - Resource utilization analysis
- **Change Tracking** - Historical cluster evolution

</td>
<td width="50%">

### 🎨 **Developer Experience**
- **Modern CLI** - Intuitive commands with rich output
- **Multiple Formats** - JSON, YAML, HTML reports
- **Shell Completion** - Bash, Zsh, Fish, PowerShell
- **Extensive Documentation** - Comprehensive guides and examples

</td>
</tr>
</table>

## ✨ Key Features

<div align="center">

| 🔍 **Discovery** | 🛡️ **Validation** | 📊 **Analysis** | 🚀 **Automation** |
|:---:|:---:|:---:|:---:|
| Complete resource discovery | Cross-reference validation | Orphaned resource detection | Git integration |
| Native K8s + CRDs | Broken reference detection | Security posture analysis | Automated reporting |
| Smart filtering | Configuration validation | Resource usage analytics | CI/CD integration |
| Real-time insights | Schema validation | Dependency analysis | Shell completion |

</div>

### 🎯 **Core Capabilities**

- 🔍 **Complete Resource Discovery** - Automatically finds ALL available API resources
- 📁 **Structured Output** - Creates intuitive directory structures for easy navigation  
- 🧹 **Clean YAML** - Intelligently removes metadata fields for re-application
- ⚡ **Lightning Fast** - Optimized for speed and efficiency
- 🔀 **Git Integration** - Automatic version control with commit history and change tracking
- 📊 **Smart Reporting** - Detailed change reports with before/after comparisons
- ✅ **Cross-Reference Validation** - Analyzes resources for broken references
- 🧹 **Orphaned Resource Detection** - Identifies cleanup opportunities
- 🎨 **Modern CLI Experience** - Rich styling, progress indicators, and helpful output
- ⚙️ **Flexible Configuration** - Project and global configuration support

## 🚀 Quick Start

### ⚡ **One-Line Install**

<table>
<tr>
<td><strong>Linux/macOS</strong></td>
<td>

```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
```

</td>
</tr>
<tr>
<td><strong>Windows</strong></td>
<td>

```powershell
iwr -useb https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.ps1 | iex
```

</td>
</tr>
</table>

### 📋 **Prerequisites**

- ☸️ **Kubernetes Access** - Valid kubeconfig or in-cluster access
- 🐹 **Go 1.21+** (if building from source) - [Download here](https://golang.org/dl/)
- 🔀 **Git** (optional) - For version control functionality

### Installation

#### Quick Install (Recommended)

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
```

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.ps1 | iex
```

#### Package Managers

**Homebrew (macOS/Linux):**
```bash
brew install graz-dev/tap/kalco
```

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

## 💡 Examples

### 🎯 **Common Workflows**

<details>
<summary><strong>🔍 Cluster Analysis & Backup</strong></summary>

```bash
# Complete cluster export with Git versioning
kalco export --git-push --commit-message "Weekly backup"

# Validate cluster health
kalco validate --output json

# Find cleanup opportunities  
kalco analyze orphaned --detailed
```

</details>

<details>
<summary><strong>🛡️ Security & Compliance</strong></summary>

```bash
# Security posture analysis
kalco analyze security --output html

# Export security-related resources
kalco export --resources "roles,rolebindings,networkpolicies,podsecuritypolicies"

# Generate compliance report
kalco report --types security,validation --output-file compliance-report.html
```

</details>

<details>
<summary><strong>🚀 DevOps & Automation</strong></summary>

```bash
# CI/CD integration
kalco export --namespaces production --exclude events,pods --output ./gitops-repo

# Environment replication
kalco export --namespaces staging --resources deployments,services,configmaps

# Resource inventory
kalco resources list --detailed --output json > inventory.json
```

</details>

### 🎨 **Beautiful CLI Experience**

```bash
# Rich, colorful output with progress indicators
kalco export --verbose

# Multiple output formats
kalco validate --output table  # Human-readable (default)
kalco validate --output json   # Machine-readable
kalco validate --output yaml   # Configuration format

# Shell completion for faster workflows
kalco completion bash > /etc/bash_completion.d/kalco
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

## 📖 Documentation

<div align="center">

### 🎯 **[Complete Documentation](https://graz-dev.github.io/kalco)**

| 📚 **Guide** | 🔗 **Link** | 📝 **Description** |
|:---:|:---:|:---|
| 🚀 | [Getting Started](https://graz-dev.github.io/kalco/getting-started) | Installation and first steps |
| 📖 | [Commands Reference](https://graz-dev.github.io/kalco/commands/) | Complete command documentation |
| ⚙️ | [Configuration](https://graz-dev.github.io/kalco/configuration) | Configuration options and examples |
| 💡 | [Use Cases](https://graz-dev.github.io/kalco/use-cases) | Real-world scenarios and examples |
| ❓ | [FAQ](https://graz-dev.github.io/kalco/faq) | Frequently asked questions |

</div>

### 🎯 **Command Overview**

<table>
<tr>
<td width="50%">

#### 🔧 **Core Operations**
- `kalco export` - Export cluster resources
- `kalco validate` - Validate resources
- `kalco analyze` - Cluster analysis
- `kalco report` - Generate reports

</td>
<td width="50%">

#### ⚙️ **Management & Setup**  
- `kalco resources` - Resource inspection
- `kalco config` - Configuration management
- `kalco completion` - Shell completion
- `kalco version` - Version information

</td>
</tr>
</table>


## 🤝 Contributing

We welcome contributions! Here's how you can help:

- 🐛 **Report Bugs** - [Create an issue](https://github.com/graz-dev/kalco/issues/new)
- 💡 **Request Features** - [Start a discussion](https://github.com/graz-dev/kalco/discussions)
- 📖 **Improve Docs** - Submit documentation improvements
- 🔧 **Submit Code** - Fork, develop, and create a pull request

### 🛠️ **Development Setup**

```bash
# Clone the repository
git clone https://github.com/graz-dev/kalco.git
cd kalco

# Install dependencies
go mod tidy

# Build and test
make build
make test

# Run locally
./kalco --help
```

## 📊 **Project Stats**

<div align="center">

![GitHub stars](https://img.shields.io/github/stars/graz-dev/kalco?style=social)
![GitHub forks](https://img.shields.io/github/forks/graz-dev/kalco?style=social)
![GitHub issues](https://img.shields.io/github/issues/graz-dev/kalco)
![GitHub pull requests](https://img.shields.io/github/issues-pr/graz-dev/kalco)

</div>

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- The Kubernetes community for inspiration and feedback
- All contributors who help make Kalco better
- The Go ecosystem for excellent tooling and libraries

---

<div align="center">

**Made with ❤️ for the Kubernetes community**

[🌟 Star us on GitHub](https://github.com/graz-dev/kalco) • [📖 Read the Docs](https://graz-dev.github.io/kalco) • [💬 Join Discussions](https://github.com/graz-dev/kalco/discussions)

</div>
