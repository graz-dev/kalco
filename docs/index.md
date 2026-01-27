---
layout: home
title: Home
nav_order: 1
---

# Kalco Documentation

Welcome to the official documentation for **Kalco** (Kubernetes Analysis & Lifecycle Control).

Kalco is a powerful CLI tool designed to bridge the gap between your Git configuration (Intent) and your running Kubernetes cluster (Reality). It provides a robust workflow for **Snapshotting**, **Diffing**, and **Validating** cluster state.

## 🌟 Why Kalco?

In complex Kubernetes environments, "Drift" is inevitable. Someone patches a deployment manually, a controller adds default fields, or a Helm chart renders differently than expected. Kalco solves this by:

- **📸 Snapshotting**: Exporting your entire cluster state into clean, Git-friendly YAML files.
- **🔍 Diffing**: performing "Smart Diffs" that ignore noise (like timestamps, `managedFields`, and dynamic status) to show only *real* changes.
- **🛡️ Validating**: Integrating into CI/CD pipelines to prevent drift *before* it merges.

## 🏗️ How It Works

Kalco operates on a simple but powerful loop:

1.  **Export**: `kalco export` connects to your cluster and downloads resources, stripping away runtime noise to create a "Golden Snapshot".
2.  **Compare**: `kalco diff` compares your local manifests (or Helm charts) against this snapshot.
3.  **Validate**: `kalco validate` automates this process in CI, fetching the base branch, detecting changed files, and verifying that your PR matches the cluster state (or intended changes).

## 🚀 Getting Started

### Installation

**Homebrew (Recommended)**
```bash
brew install graz-dev/tap/kalco
```

**Manual Install**
```bash
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash
```

### First Steps

1.  **[Configure Context]({{ site.baseurl }}/commands/context)**: Tell Kalco which cluster to manage and where to save data.
2.  **[Export Resources]({{ site.baseurl }}/commands/export)**: Take your first snapshot.
3.  **[Verify Drift]({{ site.baseurl }}/commands/diff)**: Compare your local manifests against the snapshot.

## 📚 Documentation Sections

### Command Reference
Deep dive into every command and flag.

- **[context]({{ site.baseurl }}/commands/context)**: Manage cluster connections.
- **[export]({{ site.baseurl }}/commands/export)**: Snapshot cluster state.
- **[diff]({{ site.baseurl }}/commands/diff)**: Compare local files vs cluster.
- **[validate]({{ site.baseurl }}/commands/validate)**: **[NEW]** Automated CI/CD drift detection.

### Tutorials & Use Cases
Real-world examples of Kalco in action.

- **[GitOps Validation Pipeline]({{ site.baseurl }}/tutorials/gitops-validation)**: A complete walkthrough of setting up a Drift Detection Pipeline for raw manifests using GitHub Actions and ArgoCD.

## 🤝 Support & Contributing

- **Issues**: [GitHub Issues](https://github.com/graz-dev/kalco/issues)
- **Repo**: [github.com/graz-dev/kalco](https://github.com/graz-dev/kalco)
