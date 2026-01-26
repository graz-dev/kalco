---
layout: home
title: Home
nav_order: 1
---

# Kalco Documentation

Welcome to the official documentation for **Kalco** (Kubernetes Analysis & Lifecycle Control).

Kalco is the CLI tool that bridges the gap between your Git configuration and your actual Kubernetes cluster state.

## 🌟 Why Kalco?

- **Backup & Audit**: Snapshot your entire cluster to Git-friendly YAML.
- **GitOps Validation**: Detect "drift" between what you *think* is deployed (Git) and what *is* deployed (Cluster).
- **Environment Management**: Switch context configurations easily between Prod, Staging, and Dev.

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

## 📚 Command Reference

| Command | Purpose |
|---------|---------|
| **[context]({{ site.baseurl }}/commands/context)** | Manage cluster connections and output settings. |
| **[export]({{ site.baseurl }}/commands/export)** | Dump cluster state to local files. |
| **[diff]({{ site.baseurl }}/commands/diff)** | **[NEW]** Compare local YAML vs Snapshot/Live state. |
| **version** | Check version information. |

## 🤝 Support & Contributing

- **Issues**: [GitHub Issues](https://github.com/graz-dev/kalco/issues)
- **Repo**: [github.com/graz-dev/kalco](https://github.com/graz-dev/kalco)

Is something missing in these docs? [Edit this page on GitHub](https://github.com/graz-dev/kalco/edit/master/docs/index.md).
