---
layout: home
title: Home
nav_order: 1
---

# ☸️ Kalco - Kubernetes Analysis & Lifecycle Control

**Extract, validate, analyze, and version control your entire Kubernetes cluster with comprehensive validation and Git integration**

{: .fs-6 .fw-300 }

---

## 🚀 What is Kalco?

Kalco transforms your Kubernetes cluster management experience by providing a **comprehensive, automated, and intelligent** approach to cluster analysis and lifecycle control. Whether you're managing production workloads, ensuring compliance, or planning migrations, Kalco has you covered.

## ✨ Key Features

<div class="row">
<div class="col-12 col-md-6">

### 🔍 **Intelligent Discovery**
- **Zero Configuration** - Works out of the box
- **Complete Coverage** - Native K8s + CRDs
- **Smart Filtering** - Namespace, resource, and label-based
- **Real-time Analysis** - Live cluster insights

### 🛡️ **Enterprise Ready**
- **Git Integration** - Automatic version control
- **Validation Engine** - Cross-reference checking
- **Security Analysis** - Compliance and best practices
- **Scalable Architecture** - Handles clusters of any size

</div>
<div class="col-12 col-md-6">

### 📊 **Actionable Insights**
- **Orphaned Resources** - Identify cleanup opportunities
- **Broken References** - Find configuration issues
- **Usage Analytics** - Resource utilization analysis
- **Change Tracking** - Historical cluster evolution

### 🎨 **Developer Experience**
- **Modern CLI** - Intuitive commands with rich output
- **Multiple Formats** - JSON, YAML, HTML reports
- **Shell Completion** - Bash, Zsh, Fish, PowerShell
- **Extensive Documentation** - Comprehensive guides and examples

</div>
</div>

## 🚀 Quick Start

### ⚡ **One-Line Install**

**Linux/macOS:**
```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
```

**Windows (PowerShell):**
```powershell
iwr -useb https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.ps1 | iex
```

### 📦 **Package Managers**

**Homebrew (macOS/Linux):**
```bash
brew install graz-dev/tap/kalco
```

**Debian/Ubuntu:**
```bash
wget https://github.com/graz-dev/kalco/releases/latest/download/kalco_Linux_x86_64.deb
sudo dpkg -i kalco_Linux_x86_64.deb
```

### 🔧 **Build from Source**

```bash
git clone https://github.com/graz-dev/kalco.git
cd kalco
go mod tidy
go build -o kalco ./cmd
```

## 💡 Common Use Cases

### 🎯 **Cluster Analysis & Backup**
```bash
# Complete cluster export with Git versioning
kalco export --git-push --commit-message "Weekly backup"

# Validate cluster health
kalco validate --output json

# Find cleanup opportunities  
kalco analyze orphaned --detailed
```

### 🛡️ **Security & Compliance**
```bash
# Security posture analysis
kalco analyze security --output html

# Export security-related resources
kalco export --resources "roles,rolebindings,networkpolicies,podsecuritypolicies"

# Generate compliance report
kalco report --types security,validation --output-file compliance-report.html
```

### 🚀 **DevOps & Automation**
```bash
# CI/CD integration
kalco export --namespaces production --exclude events,pods --output ./gitops-repo

# Environment replication
kalco export --namespaces staging --resources deployments,services,configmaps

# Resource inventory
kalco resources list --detailed --output json > inventory.json
```

## 🌟 Why Choose Kalco?

<div class="row">
<div class="col-12 col-md-6">

### 🎯 **Perfect for DevOps Teams**
- **Site Reliability Engineers** - Automated cluster backups and disaster recovery
- **Platform Engineers** - Infrastructure as Code and GitOps workflows  
- **Security Teams** - Compliance auditing and security posture analysis
- **Developers** - Environment replication and configuration management

</div>
<div class="col-12 col-md-6">

### 📊 **Project Stats**
- **0 Configuration Required** - Works out of the box
- **100% Resource Coverage** - Native K8s + CRDs
- **⚡ Lightning Fast** - Optimized for speed
- **🔒 Enterprise Ready** - Production-grade reliability

</div>
</div>

## 🚀 Next Steps

1. **[Install Kalco]({{ site.baseurl }}/docs/getting-started/installation)** - Get up and running quickly
2. **[Explore Commands]({{ site.baseurl }}/docs/commands/)** - Learn all available options
3. **[Configuration Guide]({{ site.baseurl }}/docs/configuration/)** - Customize for your environment
4. **[Use Cases]({{ site.baseurl }}/docs/use-cases/)** - Real-world examples and workflows

---

<div class="text-center">
**Made with ❤️ for the Kubernetes community**

[🌟 Star us on GitHub](https://github.com/graz-dev/kalco){: .btn .btn-primary .btn-lg } [📖 Read the Docs]({{ site.baseurl }}/docs/getting-started/){: .btn .btn-outline .btn-lg }
</div>