# GitOps Validation Demos

This directory contains step-by-step tutorials for implementing **GitOps Validation** using Kalco.

**Philosophy**: These demos are designed to be "Drop-in". They assume you want validation *without* maintaining a persistent "Reality Repo". The pipeline creates a temporary snapshot during the check.

## Scenarios

1.  **[Raw Resources](./raw-resources/README.md)**
    *   **Best for:** Deployments using plain Kubernetes YAML manifests (Deployment, Service, etc.).
    *   **How to use:** Copy the folder contents (`manifests`, `.github`) to your repo.

2.  **[Helm Charts](./helm/README.md)**
    *   **Best for:** Deployments managed via Helm Charts.
    *   **How to use:** Copy the folder contents (`chart`, `.github`) to your repo.

All you need is to set the `KUBECONFIG` secret in your repository.
