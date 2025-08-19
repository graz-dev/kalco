# Scripts Directory

This directory contains utility scripts for Kalco development and testing.

## Available Scripts

### `create-guestbook-cluster.sh`

Creates a local Kubernetes cluster using Kind with the Guestbook application deployed. This script is perfect for testing Kalco's export functionality with a real application.

#### What it does

1. **Creates a Kind cluster** with 3 nodes (1 control-plane + 2 workers)
2. **Deploys Redis database** with leader-follower architecture
3. **Deploys Guestbook frontend** application
4. **Creates a Kalco context** for easy cluster management
5. **Keeps the cluster running** for ongoing testing

#### Prerequisites

- **Docker** - Running and accessible
- **Kind** - Kubernetes in Docker tool
- **kubectl** - Kubernetes command-line tool
- **Kalco** (optional) - For context management

#### Installation of Prerequisites

**Install Kind:**
```bash
# Linux/macOS
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind

# macOS with Homebrew
brew install kind
```

**Install kubectl:**
```bash
# Linux/macOS
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/kubectl

# macOS with Homebrew
brew install kubectl
```

#### Usage

**Basic usage:**
```bash
./scripts/create-guestbook-cluster.sh
```

**With custom cluster name:**
```bash
./scripts/create-guestbook-cluster.sh my-custom-cluster
```

#### What gets deployed

The script deploys the complete Guestbook application as described in the [Kubernetes Guestbook Tutorial](https://kubernetes.io/docs/tutorials/stateless-application/guestbook/):

- **Redis Leader**: Single Redis instance for write operations
- **Redis Followers**: 2 Redis instances for read operations
- **Frontend**: 3 PHP web servers serving the Guestbook application
- **Services**: Internal networking for all components

#### Cluster Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Control Plane │    │   Worker Node 1 │    │   Worker Node 2 │
│   (Port 8080)  │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                    ┌─────────────────┐
                    │  Guestbook App  │
                    │  (Frontend)     │
                    └─────────────────┘
                                 │
                    ┌─────────────────┐
                    │  Redis Cluster  │
                    │  (Leader + 2    │
                    │   Followers)    │
                    └─────────────────┘
```

#### Accessing the Application

After the script completes successfully:

1. **Port forward the frontend service:**
   ```bash
   kubectl port-forward svc/frontend 8080:80
   ```

2. **Open your browser** and navigate to `http://localhost:8080`

3. **Test the application** by adding guestbook entries

#### Using with Kalco

The script automatically creates a Kalco context for easy management:

```bash
# Switch to the guestbook cluster context
kalco context use guestbook-cluster

# Export the cluster resources
kalco export --git-push --commit-message "Initial guestbook cluster export"

# View the exported resources
ls -la ./guestbook-exports/
```

#### Cluster Management

**View cluster status:**
```bash
kubectl get pods
kubectl get services
kubectl get nodes
```

**Scale the frontend:**
```bash
kubectl scale deployment frontend --replicas=5
```

**View logs:**
```bash
kubectl logs -l app=guestbook,tier=frontend
kubectl logs -l app=redis,tier=backend
```

#### Cleanup

**The cluster is designed to persist** for ongoing testing. To manually remove it:

```bash
kind delete cluster --name guestbook-cluster
```

**Remove Kalco context:**
```bash
kalco context delete guestbook-cluster
```

#### Troubleshooting

**Common issues and solutions:**

1. **Docker not running:**
   ```bash
   # Start Docker Desktop or Docker daemon
   sudo systemctl start docker  # Linux
   # Or start Docker Desktop on macOS/Windows
   ```

2. **Port 8080 already in use:**
   ```bash
   # Use a different port
   kubectl port-forward svc/frontend 8081:80
   ```

3. **Pods not starting:**
   ```bash
   # Check pod status
   kubectl describe pods
   kubectl logs <pod-name>
   ```

4. **Cluster creation fails:**
   ```bash
   # Check available resources
   docker system df
   # Ensure enough disk space and memory
   ```

#### Customization

You can modify the script to:

- **Change cluster size** by modifying the node configuration
- **Use different Redis versions** by changing the image tag
- **Add more applications** by extending the deployment functions
- **Modify resource limits** for different testing scenarios

#### Example Modifications

**Add more worker nodes:**
```bash
# In the create_cluster function, add more worker nodes
- role: worker
- role: worker
- role: worker  # Add this line
```

**Change Redis version:**
```bash
# In deploy_redis function, change the image
image: redis:7.0-alpine
```

**Modify resource limits:**
```bash
# Adjust CPU and memory limits as needed
resources:
  requests:
    cpu: 200m
    memory: 256Mi
  limits:
    cpu: 500m
    memory: 512Mi
```

#### Integration with CI/CD

This script can be integrated into CI/CD pipelines for:

- **Automated testing** of Kalco functionality
- **Development environment** setup
- **Demo environment** creation
- **Integration testing** with real Kubernetes clusters

#### Support

For issues with the script:

1. Check the prerequisites are installed
2. Ensure Docker has sufficient resources
3. Check the troubleshooting section above
4. Review the script output for error messages
5. Verify your system meets the requirements

---

*This script is based on the [Kubernetes Guestbook Tutorial](https://kubernetes.io/docs/tutorials/stateless-application/guestbook/) and adapted for Kalco testing purposes.*
