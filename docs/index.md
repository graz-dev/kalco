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

Kalco is a powerful Kubernetes cluster management tool that provides comprehensive resource extraction, validation, and version control capabilities. It's designed to help DevOps engineers, SREs, and platform teams maintain clean, validated, and version-controlled cluster configurations.

## ✨ Key Features

- **🔍 Complete Resource Extraction** - Export all cluster resources including CRDs
- **✅ Smart Validation** - Cross-reference validation and orphaned resource detection  
- **📝 Git Integration** - Automatic version control with commit and push capabilities
- **🎯 Flexible Filtering** - Export specific namespaces, resources, or exclude noisy types
- **📊 Detailed Reporting** - Comprehensive change analysis and resource summaries
- **🔄 Incremental Updates** - Track changes between cluster snapshots

## 🚀 Quick Start

```bash
# Install Kalco
go install github.com/graz-dev/kalco/cmd/kalco@latest

# Export your cluster
kalco export --output ./cluster-backup

# Export with Git integration
kalco export --git-push --commit-message "Cluster snapshot $(date)"
```

## 📚 Documentation

- **[Getting Started]({{ site.baseurl }}/getting-started/)** - Installation and first steps
- **[Commands Reference]({{ site.baseurl }}/commands/)** - Complete command documentation
- **[Use Cases]({{ site.baseurl }}/use-cases/)** - Common scenarios and workflows
- **[Configuration]({{ site.baseurl }}/configuration/)** - Customization options

## 🔗 Useful Links

- [GitHub Repository](https://github.com/graz-dev/kalco)
- [Issues & Bug Reports](https://github.com/graz-dev/kalco/issues)
- [Discussions](https://github.com/graz-dev/kalco/discussions)
- [Releases](https://github.com/graz-dev/kalco/releases)

---

*Built with ❤️ for the Kubernetes community*
