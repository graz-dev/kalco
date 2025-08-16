---
layout: default
title: Troubleshooting
parent: Getting Started
nav_order: 4
---

# Troubleshooting

Common issues and their solutions when using Kalco.

## Quick Diagnosis

Start with these diagnostic commands:

```bash
# Check Kalco version and status
kalco --version
kalco --help

# Verify Kubernetes access
kubectl get nodes
kubectl config current-context

# Check system resources
df -h
free -h
```

## Common Issues

### Installation Problems

#### Permission Denied

**Error:**
```bash
./kalco: Permission denied
```

**Solution:**
```bash
# Make binary executable
chmod +x kalco

# Check file permissions
ls -la kalco

# Install to system directory
sudo mv kalco /usr/local/bin/
```

#### Command Not Found

**Error:**
```bash
kalco: command not found
```

**Solution:**
```bash
# Check if binary is in PATH
echo $PATH

# Verify binary location
which kalco

# Add to PATH if needed
export PATH=$PATH:/path/to/kalco

# Or move to system PATH
sudo mv kalco /usr/local/bin/
```

#### Version Compatibility

**Error:**
```bash
kalco: version compatibility error
```

**Solution:**
```bash
# Check Go version
go version

# Ensure Go 1.21+ is installed
# Download from https://golang.org/dl/

# Rebuild from source
go mod tidy
go build -o kalco ./cmd
```

### Kubernetes Connection Issues

#### No Configuration Provided

**Error:**
```bash
error: no configuration has been provided
```

**Solution:**
```bash
# Check kubeconfig location
echo $KUBECONFIG

# Set kubeconfig path
export KUBECONFIG="$HOME/.kube/config"

# Or use --kubeconfig flag
kalco export --kubeconfig ~/.kube/config

# Verify context
kubectl config current-context
```

#### Context Not Found

**Error:**
```bash
error: context "my-cluster" does not exist
```

**Solution:**
```bash
# List available contexts
kubectl config get-contexts

# Switch to existing context
kubectl config use-context existing-context

# Or create new context
kubectl config set-context my-cluster --cluster=my-cluster --user=my-user
```

#### Access Denied

**Error:**
```bash
error: You must be logged in to the server (Unauthorized)
```

**Solution:**
```bash
# Check authentication
kubectl auth can-i get pods

# Re-authenticate
kubectl login

# Check RBAC permissions
kubectl auth can-i get nodes
kubectl auth can-i list deployments
```

### Export Issues

#### Resource Discovery Failed

**Error:**
```bash
error: failed to discover resources
```

**Solution:**
```bash
# Check API server access
kubectl api-resources

# Verify cluster health
kubectl get componentstatuses

# Check for API server issues
kubectl get events --all-namespaces
```

#### Output Directory Issues

**Error:**
```bash
error: failed to create output directory
```

**Solution:**
```bash
# Check directory permissions
ls -la /path/to/output

# Create directory with proper permissions
mkdir -p /path/to/output
chmod 755 /path/to/output

# Use absolute path
kalco export --output-dir /absolute/path/to/output
```

#### Memory Issues

**Error:**
```bash
error: out of memory
```

**Solution:**
```bash
# Check available memory
free -h

# Export with resource limits
kalco export --max-memory 2GB

# Export specific namespaces only
kalco export --namespaces default,production

# Exclude large resource types
kalco export --exclude events,pods,logs
```

### Git Integration Issues

#### Not a Git Repository

**Error:**
```bash
error: not a git repository
```

**Solution:**
```bash
# Initialize Git repository
git init

# Or use --no-git flag
kalco export --no-git

# Check Git status
git status
```

#### Git Authentication Failed

**Error:**
```bash
error: authentication failed
```

**Solution:**
```bash
# Configure Git credentials
git config --global user.name "Your Name"
git config --global user.email "your.email@example.com"

# Set up SSH keys or personal access token
# For GitHub: https://docs.github.com/en/authentication

# Test Git access
git ls-remote origin
```

#### Git Push Failed

**Error:**
```bash
error: failed to push to remote
```

**Solution:**
```bash
# Check remote configuration
git remote -v

# Verify remote access
git ls-remote origin

# Force push if needed (use with caution)
git push --force-with-lease origin main

# Or skip push
kalco export --git-push=false
```

### Performance Issues

#### Slow Export

**Symptoms:**
- Export takes a very long time
- High CPU/memory usage
- Timeout errors

**Solutions:**
```bash
# Enable verbose output to identify bottlenecks
kalco export --verbose

# Export specific namespaces only
kalco export --namespaces default,production

# Exclude resource types that cause delays
kalco export --exclude events,pods,logs

# Use resource limits
kalco export --max-memory 4GB --max-cpu 2
```

#### High Resource Usage

**Solutions:**
```bash
# Limit concurrent operations
kalco export --max-concurrency 5

# Set memory limits
kalco export --max-memory 2GB

# Use batch processing
kalco export --batch-size 100

# Enable progress monitoring
kalco export --progress
```

## Debug Mode

Enable detailed debugging:

```bash
# Enable debug output
kalco export --debug

# Set log level
export KALCO_LOG_LEVEL=debug

# Enable trace logging
export KALCO_TRACE=true

# Check debug information
kalco debug info
```

## Getting Help

### Self-Service

1. **Check Documentation** - This troubleshooting guide
2. **Review Logs** - Enable verbose/debug output
3. **Test Commands** - Try with minimal options first
4. **Verify Environment** - Check system requirements

### Community Support

- **GitHub Issues** - [Report bugs](https://github.com/graz-dev/kalco/issues)
- **Discussions** - [Ask questions](https://github.com/graz-dev/kalco/discussions)
- **Documentation** - [Browse docs]({{ site.baseurl }}/docs/)

### Reporting Issues

When reporting an issue, include:

```bash
# System information
kalco --version
kubectl version --client
go version
uname -a

# Error details
kalco export --verbose 2>&1 | tee kalco-debug.log

# Configuration
cat .kalco.yml 2>/dev/null || echo "No config file"
echo "KUBECONFIG: $KUBECONFIG"
```

## Prevention

### Best Practices

1. **Regular Updates** - Keep Kalco updated
2. **Resource Monitoring** - Monitor system resources
3. **Backup Strategy** - Regular cluster exports
4. **Testing** - Test in non-production first
5. **Documentation** - Keep configuration documented

### Monitoring

```bash
# Check Kalco health
kalco health

# Monitor resource usage
kalco export --monitor

# Set up alerts for failures
# Integrate with your monitoring system
```

## Next Steps

After resolving issues:

1. **[Review Configuration]({{ site.baseurl }}/docs/getting-started/configuration)** - Optimize your setup
2. **[Explore Commands]({{ site.baseurl }}/docs/commands/)** - Learn advanced features
3. **[Automate Workflows]({{ site.baseurl }}/docs/use-cases/automation)** - Prevent future issues
4. **[Join Community](https://github.com/graz-dev/kalco/discussions)** - Share solutions
