# GitOps Validation with Kalco: Helm Charts

This tutorial shows how to validate Helm Chart changes against your live cluster using Kalco.

## Prerequisites

- [ ] Kubernetes Cluster & ArgoCD.
- [ ] **One GitHub Repository**: This contains your chart and pipeline.
- [ ] **Secrets** configured: `KUBECONFIG`.

## Setup

### 1. Install ArgoCD (if not installed)

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Wait for ArgoCD to be ready
kubectl wait --for=condition=available deployment -l "app.kubernetes.io/name=argocd-server" -n argocd --timeout=300s
```

### 2. Copy Files
Copy the contents of this folder (`chart/`, `.github/`, `argocd/`) to the **root** of your repository.

### 3. Configure Secret
Add `KUBECONFIG` to your repo secrets.

### 4. Deploy App
Apply `argocd/application.yaml` (edit the repoURL first!).

## Workflow Overview

1.  **Snapshot**: Kalco exports the live state to a temporary directory (`./tmp-snapshot`).
2.  **Render**: `helm template` generates manifests from your local chart code.
3.  **Split & Diff**: The rendered manifests are split and compared against the temporary snapshot.

## Command Reference

```bash
# Workflow logic roughly equivalent to:
kalco context set validation-ctx --output ./tmp-snapshot --kubeconfig ...
kalco export
helm template . > all.yaml
kalco diff --target <split-file> --snapshot-dir ./tmp-snapshot
```

## Usage

1.  Change a value in `chart/values.yaml` (e.g., replica count).
2.  Open a PR.
3.  The pipeline will detect if your change conflicts with the live state.
