---
layout: default
title: GitOps Validation Pipeline
nav_order: 1
has_children: false
parent: Tutorials & Use Cases
---

# Tutorial: GitOps Validation Pipeline

This tutorial walks you through setting up a robust **GitOps Validation Pipeline** using Kalco.

**The Problem**: You use ArgoCD to deploy from Git. But what happens if someone manually changes the cluster? Or if you make a change in Git that conflicts with the *actual* cluster state in a way you didn't expect?
**The Solution**: A CI check that compares your Pull Request not just against the *codebase*, but against the *live cluster state*.

## Scenario

We have a "Guestbook" application managed by ArgoCD. We want to ensure that every PR to our repository is validated against the actual running cluster to detect **Drift**.

If a PR proposes a change, Kalco checks:
1.  Is this a new change? (Intent) -> **Allowed**.
2.  Is there a difference in the cluster that is NOT in my PR? (Drift) -> **Blocked**.

## Directory Structure

We'll organize our repo like this:

```
.
├── manifests/              # Kubernetes Manifests
│   └── frontend.yaml
├── argocd/                 # ArgoCD Application Definition
│   └── application.yaml
└── .github/
    └── workflows/
        └── validation.yaml # CI Pipeline
```

## Step 1: The Workload (Manifests)

Creating a simple Guestbook Frontend (`manifests/frontend.yaml`):

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
  namespace: guestbook
  labels:
    app: guestbook
    tier: frontend
spec:
  replicas: 3
  selector:
    matchLabels:
        app: guestbook
        tier: frontend
  template:
    metadata:
      labels:
        app: guestbook
        tier: frontend
    spec:
      containers:
      - name: php-redis
        image: us-docker.pkg.dev/google-samples/containers/gke/gb-frontend:v5
        env:
        - name: GET_HOSTS_FROM
          value: "dns"
        resources:
          requests:
            cpu: 100m
            memory: 100Mi
        ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: frontend
  namespace: guestbook
  labels:
    app: guestbook
    tier: frontend
spec:
  ports:
  - port: 80
  selector:
    app: guestbook
    tier: frontend
```

## Step 2: ArgoCD Application

Define how ArgoCD syncs this app (`argocd/application.yaml`).
*Note: We disable SelfHeal for this demo so we can simulate drift manually.*

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: guestbook
  namespace: argocd
spec:
  project: default
  source:
    repoURL: https://github.com/YOUR_USERNAME/your-repo  # Update this!
    targetRevision: HEAD
    path: manifests
  destination:
    server: https://kubernetes.default.svc
    namespace: guestbook
  syncPolicy:
    automated: 
      selfHeal: false  # Disabled to allow drift testing
      prune: true
```

## Step 3: The Validation Pipeline

This GitHub Action runs on every Pull Request. It uses **Kalco** to:
1.  Snapshot the live cluster state.
2.  Validate changed manifests against that snapshot.

File: `.github/workflows/validation.yaml`

```yaml
name: Kalco GitOps Validation

on:
  pull_request:
    branches: [ master ]
    paths:
      - 'manifests/**'

jobs:
  validate-drift:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      pull-requests: write
    steps:
      - name: Checkout Source Code
        uses: actions/checkout@v3
        with:
          fetch-depth: 0 # Important for git diff

      # 1. Install Kalco
      - name: Install Kalco
        run: |
          curl -sL https://github.com/graz-dev/kalco/releases/download/v0.1.21/kalco_Linux_x86_64.tar.gz | tar xz
          sudo mv kalco /usr/local/bin/

      # 2. Configure Context (Connect to Cluster)
      - name: Configure Kalco Context
        env:
          KUBECONFIG_DATA: ${{ secrets.KUBECONFIG }}
        run: |
          mkdir -p ~/.kube
          echo "$KUBECONFIG_DATA" | base64 -d > ~/.kube/config
          mkdir -p tmp-snapshot
          
          kalco context set validation-ctx \
            --kubeconfig ~/.kube/config \
            --output $(pwd)/tmp-snapshot
          kalco context use validation-ctx

      # 3. Snapshot Live State
      - name: Snapshot Live State
        run: kalco export

      # 4. Run Validation
      - name: Kalco Validate
        id: kalco-validate
        run: |
          git fetch origin ${{ github.base_ref }}
          
          kalco validate \
            --base origin/${{ github.base_ref }} \
            --dirs manifests \
            --snapshot-dir tmp-snapshot \
            --namespace guestbook \
            --report-file report.md \
            || echo "has_drift=true" >> $GITHUB_OUTPUT
            
          # Prepare output for PR Comment
          EOF=$(dd if=/dev/urandom bs=15 count=1 status=none | base64)
          echo "diff_output<<$EOF" >> $GITHUB_OUTPUT
          cat report.md >> $GITHUB_OUTPUT || echo "No changes." >> $GITHUB_OUTPUT
          echo "" >> $GITHUB_OUTPUT
          echo "$EOF" >> $GITHUB_OUTPUT

      # 5. Comment on PR
      - name: Comment PR & Label
        uses: actions/github-script@v6
        env:
          DIFF_OUTPUT: ${{ steps.kalco-validate.outputs.diff_output }}
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          script: |
            const output = process.env.DIFF_OUTPUT;
            if (output && output.trim() !== "") {
              await github.rest.issues.createComment({
                issue_number: context.issue.number,
                owner: context.repo.owner,
                repo: context.repo.repo,
                body: `### 🛡️ Kalco Validation Report\n\n${output}`
              });
            }
```

## How it works in practice

1.  **You open a PR** changing `replicas: 3` to `replicas: 5`.
2.  **Kalco** sees `frontend.yaml` changed.
3.  它 extracts the base version (`replicas: 3`) from master.
4.  It compares `replicas: 5` (PR) vs `replicas: 3` (Snapshot).
5.  It knows `3 -> 5` is your **Intent**. Result: **✅ No Drift**.

**BUT**, if someone manually edited the cluster to add `annotation: "manual-fix"`, and your PR *doesn't* have that:
4.  It compares your PR (Clean) vs Snapshot (Has annotation).
5.  It sees the annotation was NOT in the Base branch either.
6.  Result: **🔴 Drift Detected**. The cluster has something extra that isn't in Git.

This protects your GitOps workflow from reconciling over unknown manual changes or missing critical configuration.
