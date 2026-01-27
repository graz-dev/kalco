# GitOps Validation with Kalco: Raw Resources

This tutorial guides you through setting up a **GitOps Validation Workflow** using **Kalco**.

## Prerequisites

- [ ] Kubernetes Cluster & ArgoCD.
- [ ] **One GitHub Repository**: This contains your manifests and pipeline.
- [ ] **Secrets** configured in this repo:
    -   `KUBECONFIG`: Base64 encoded kubeconfig.

## Setup

### 1. Install ArgoCD (if not installed)

```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Wait for ArgoCD to be ready
kubectl wait --for=condition=available deployment -l "app.kubernetes.io/name=argocd-server" -n argocd --timeout=300s
```

### 2. Copy Files
Copy the contents of this folder (`manifests/`, `.github/`, `argocd/`) to the **root** of your repository.
    *   Your repo should look like:
        ```
        .github/workflows/validation.yaml
        manifests/frontend.yaml
        manifests/redis-leader.yaml
        ...
        argocd/application.yaml
        ```

### 3. Configure Secret
Add `KUBECONFIG` to your repo secrets.

### 4. Deploy App
Apply `argocd/application.yaml` (edit the repoURL first!).

## Workflow Overview

The GitHub Action (`.github/workflows/validation.yaml`) performs the following steps:

1.  **Context Setup**: Configures Kalco to use a **temporary directory** (`./tmp-snapshot`) as the output.
2.  **Export**: Exports the current live state of the cluster to this temporary directory.
3.  **Diff**: Iterates over changed `.yaml` files in `manifests/` and compares them against the temporary snapshot.

## Usage

1.  Create a PR modifying a file in `manifests/`.
2.  The pipeline will catch any drift between your PR intent and the actual live cluster state.
