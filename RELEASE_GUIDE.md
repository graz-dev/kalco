# 🚀 Kalco Release Guide

This guide explains how to use the automated GitHub Actions workflows to create releases and packages for Kalco.

## 📋 Overview

Kalco now has comprehensive GitHub Actions workflows that automatically:

- 🏗️ **Build** binaries for all platforms (Linux, macOS, Windows, AMD64, ARM64)
- 📦 **Package** releases with proper archives (.tar.gz for Unix, .zip for Windows)
- 🚀 **Release** on GitHub with detailed release notes
- 🔒 **Scan** for security vulnerabilities
- 🧪 **Test** the quickstart script and all functionality
- 📋 **Generate** release notes automatically

## 🔄 Workflow Types

### 1. **CI Workflow** (`.github/workflows/ci.yml`)
- **Triggers**: Push to master, Pull requests
- **Purpose**: Continuous integration testing
- **Actions**:
  - Run tests with race detection
  - Build for multiple platforms
  - Lint code with golangci-lint
  - Upload build artifacts

### 2. **Release Workflow** (`.github/workflows/release.yml`)
- **Triggers**: Push tags (e.g., `v1.0.0`)
- **Purpose**: Create official releases
- **Actions**:
  - Build binaries for all platforms
  - Create GitHub release with assets

  - Generate comprehensive release notes





### 3. **Dependencies Workflow** (`.github/workflows/dependencies.yml`)
- **Triggers**: Weekly (Monday 9 AM UTC), Manual
- **Purpose**: Keep dependencies updated
- **Actions**:
  - Check for dependency updates
  - Security vulnerability scanning
  - Create automated PRs

### 4. **Release Drafter Workflow** (`.github/workflows/release-drafter.yml`)
- **Triggers**: Push to master, Pull requests
- **Purpose**: Generate release notes
- **Actions**:
  - Auto-generate release notes
  - Categorize changes
  - Create release drafts

## 🚀 Creating a Release

### **Step 1: Prepare Your Changes**
```bash
# Ensure all changes are committed and pushed
git add .
git commit -m "feat: add new feature"
git push origin master
```

### **Step 2: Create and Push a Tag**
```bash
# Create a new version tag
git tag v1.0.0

# Push the tag to trigger the release workflow
git push origin v1.0.0
```

### **Step 3: Monitor the Workflow**
The release workflow will automatically:
1. Build binaries for all platforms
2. Create a GitHub release
3. Upload all platform-specific packages
4. Generate comprehensive release notes

## 📦 What Gets Created

### **Binary Packages**
- `kalco-linux-amd64.tar.gz` - Linux x86_64
- `kalco-linux-arm64.tar.gz` - Linux ARM64
- `kalco-darwin-amd64.tar.gz` - macOS Intel
- `kalco-darwin-arm64.tar.gz` - macOS Apple Silicon
- `kalco-windows-amd64.zip` - Windows x86_64
- `kalco-windows-arm64.zip` - Windows ARM64



### **Release Assets**
- All binary packages
- Comprehensive release notes
- Installation instructions
- Quick start guide

## 🔧 Local Development

### **Using the Makefile**
```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run tests
make test

# Test quickstart script
make quickstart

# Create local release packages
make release-local


```

### **Testing Workflows Locally**
```bash
# Install act (GitHub Actions local runner)
brew install act  # macOS
# or download from: https://github.com/nektos/act

# Run specific workflow
act push -W .github/workflows/ci.yml
```

## 📊 Workflow Status

### **Required Status Checks**
Before merging to master, ensure:
- ✅ **CI** - All tests pass
- ✅ **Lint** - Code quality checks pass
- ✅ **Security** - No vulnerabilities detected

### **Release Requirements**
For a successful release:
- 🏷️ Valid semantic version tag (e.g., `v1.0.0`)
- ✅ All CI checks passing
- 🔒 Security scan clean



## 🔒 Security Features

### **Automated Scanning**
- **Dependencies**: Weekly vulnerability checks
- **Code**: Trivy security scanning on every PR
- **Images**: Container vulnerability scanning
- **Updates**: Automated dependency update PRs

### **Security Best Practices**
- Non-root user in containers
- Minimal base images (Alpine Linux)
- Regular dependency updates
- Vulnerability scanning in CI/CD

## 📋 Troubleshooting

### **Common Issues**

#### **Release Workflow Fails**
- Check that the tag follows semantic versioning (`v1.0.0`)
- Ensure all CI checks are passing
- Verify repository permissions for releases

#### **Docker Build Fails**
- Check Dockerfile syntax
- Verify multi-platform build support
- Ensure GitHub Container Registry access

#### **Binary Build Fails**
- Verify Go version compatibility
- Check for platform-specific build issues
- Review build flags and environment variables

### **Debug Commands**
```bash
# Check workflow runs
gh run list

# View workflow logs
gh run view <run-id>

# Download artifacts
gh run download <run-id>

# Rerun failed workflow
gh run rerun <run-id>
```

## 🎯 Best Practices

### **Version Management**
- Use semantic versioning (`v1.0.0`, `v1.1.0`, `v2.0.0`)
- Create tags from master branch only
- Use conventional commit messages for auto-release notes

### **Release Process**
- Test locally before tagging
- Review generated release notes
- Verify all platform binaries
- Test Docker images locally

### **Workflow Maintenance**
- Keep actions up to date
- Monitor workflow execution times
- Review and optimize build steps
- Update supported platforms as needed

## 📚 Additional Resources

- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)
- [Release Drafter](https://github.com/release-drafter/release-drafter)
- [Docker Multi-platform Builds](https://docs.docker.com/build/building/multi-platform/)

---

## 🚀 Quick Release Checklist

Before creating a release:

- [ ] All changes committed and pushed to master
- [ ] CI workflow passing
- [ ] Security scan clean
- [ ] Quickstart test successful
- [ ] Documentation updated
- [ ] Release notes reviewed
- [ ] Version number decided
- [ ] Release notes prepared

**Then simply:**
```bash
git tag v1.0.0
git push origin v1.0.0
```

**And watch the magic happen! ✨**
