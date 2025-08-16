---
layout: default
title: First Run
nav_order: 2
parent: Getting Started
---

# First Run

Learn how to export your first Kubernetes cluster with Kalco.

## 🎯 Prerequisites

Before you begin, ensure you have:

- ✅ Kalco installed and working
- ✅ `kubectl` configured and connected to a cluster
- ✅ Access to a Kubernetes cluster (local or remote)

## 🔍 Verify Cluster Access

First, verify that you can access your cluster:

```bash
kubectl cluster-info
kubectl get nodes
```

## 🚀 Basic Export

Start with a simple export to see Kalco in action:

```bash
# Export to a timestamped directory
kalco export

# Or specify a custom output directory
kalco export --output ./my-cluster-backup
```

## 📁 Understanding the Output

Kalco creates an organized directory structure:

```
my-cluster-backup/
├── _cluster/           # Cluster-scoped resources
│   ├── ClusterRole/
│   ├── ClusterRoleBinding/
│   ├── StorageClass/
│   └── ...
├── default/            # Default namespace
│   ├── ConfigMap/
│   ├── Service/
│   └── ...
├── kube-system/        # System namespace
│   ├── Deployment/
│   ├── Service/
│   └── ...
└── kalco-reports/      # Analysis reports
    ├── Cluster-snapshot-2025-08-16-14-55-34.md
    └── ...
```

## 🔍 View the Report

Check the generated report to understand your cluster:

```bash
# View the latest report
ls -la ./my-cluster-backup/kalco-reports/
cat ./my-cluster-backup/kalco-reports/*.md | head -50
```

## 📊 What the Report Shows

The report includes:

- **Resource Summary** - Count of each resource type
- **Validation Results** - Cross-reference checks
- **Orphaned Resources** - Unmanaged resources
- **Detailed Changes** - Since previous snapshot (if any)

## 🔄 Git Integration

Initialize Git tracking for version control:

```bash
cd ./my-cluster-backup

# Initialize Git repository
git init
git add .
git commit -m "Initial cluster export - $(date)"

# Add remote origin (optional)
git remote add origin <your-repo-url>
git push -u origin main
```

## 🎯 Next Steps

Now that you've completed your first export:

1. **[Explore Commands]({{ site.baseurl }}/docs/commands/)** - Learn about all available options
2. **[Configure Kalco]({{ site.baseurl }}/docs/getting-started/configuration)** - Customize for your workflow
3. **[Use Cases]({{ site.baseurl }}/docs/use-cases/)** - Common scenarios and workflows

## 🐛 Troubleshooting

### Common Issues

**Permission denied**: Ensure you have cluster access
```bash
kubectl auth can-i get pods --all-namespaces
```

**Empty export**: Check if resources exist
```bash
kubectl get all --all-namespaces
```

**Report not generated**: Verify output directory permissions
```bash
ls -la ./my-cluster-backup/
```

Congratulations! You've successfully exported your first Kubernetes cluster with Kalco. 🎉
