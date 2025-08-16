# Release Process

This document describes how to create and manage releases for kalco.

## Automated Release System

Kalco uses an automated release system powered by:
- **GitHub Actions** - Automated CI/CD pipeline
- **GoReleaser** - Cross-platform builds and release management
- **Package Managers** - Automatic publishing to Homebrew, APT, RPM, etc.

## Creating a Release

### Method 1: Using the Release Script (Recommended)

```bash
./scripts/create-release.sh
```

This interactive script will:
1. Check your git status
2. Suggest version numbers based on semantic versioning
3. Allow you to enter release notes
4. Create and push the git tag
5. Trigger the automated build process

### Method 2: Manual Release

1. **Ensure your code is ready:**
   ```bash
   git status  # Should be clean
   go test ./...  # All tests should pass
   go build  # Should build successfully
   ```

2. **Create a git tag:**
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

3. **GitHub Actions will automatically:**
   - Build binaries for all supported platforms
   - Create checksums
   - Generate a GitHub release
   - Upload all artifacts

## What Gets Built

The release process creates:

### Binaries
- **Linux**: amd64, arm64 (tar.gz)
- **macOS**: amd64, arm64 (tar.gz) 
- **Windows**: amd64 (zip)

### Packages
- **Debian/Ubuntu**: .deb packages
- **RHEL/CentOS/Fedora**: .rpm packages
- **Arch Linux**: .pkg.tar.xz packages
- **Alpine**: .apk packages
- **Homebrew**: Formula (if tap is configured)

### Additional Files
- Checksums file for verification
- Source code archives
- Release notes

## Installation Methods for Users

After a release is published, users can install kalco using:

### Quick Install Scripts
```bash
# Linux/macOS
curl -fsSL https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.sh | bash

# Windows PowerShell
iwr -useb https://raw.githubusercontent.com/graz-dev/kalco/master/scripts/install.ps1 | iex
```

### Package Managers
```bash
# Homebrew
brew install graz-dev/tap/kalco

# Debian/Ubuntu
wget https://github.com/graz-dev/kalco/releases/latest/download/kalco_Linux_x86_64.deb
sudo dpkg -i kalco_Linux_x86_64.deb

# RHEL/CentOS/Fedora
wget https://github.com/graz-dev/kalco/releases/latest/download/kalco_Linux_x86_64.rpm
sudo rpm -i kalco_Linux_x86_64.rpm
```

### Direct Download
Users can download binaries directly from the [releases page](https://github.com/graz-dev/kalco/releases).

## Version Numbering

We follow [Semantic Versioning](https://semver.org/):

- **MAJOR** version (v2.0.0): Incompatible API changes
- **MINOR** version (v1.1.0): New functionality in a backwards compatible manner
- **PATCH** version (v1.0.1): Backwards compatible bug fixes

## Pre-release Versions

For pre-release versions, append a suffix:
- `v1.0.0-alpha.1` - Alpha release
- `v1.0.0-beta.1` - Beta release
- `v1.0.0-rc.1` - Release candidate

## Troubleshooting

### Release Failed
1. Check the [GitHub Actions page](https://github.com/graz-dev/kalco/actions)
2. Look for error messages in the workflow logs
3. Common issues:
   - Missing `GITHUB_TOKEN` permissions
   - Invalid GoReleaser configuration
   - Build failures on specific platforms

### Missing Packages
If certain package formats aren't being generated:
1. Check the `.goreleaser.yml` configuration
2. Ensure the `nfpms` section includes the desired formats
3. Verify the GitHub Actions workflow has the necessary permissions

### Homebrew Tap Issues
If Homebrew publishing fails:
1. Ensure you have a separate repository for your Homebrew tap
2. Update the `brews` section in `.goreleaser.yml` with correct repository details
3. Make sure the bot has write access to the tap repository

## Configuration Files

- `.github/workflows/release.yml` - GitHub Actions workflow
- `.goreleaser.yml` - GoReleaser configuration
- `scripts/install.sh` - Linux/macOS installation script
- `scripts/install.ps1` - Windows installation script

## Security Considerations

- All releases are signed with checksums
- Installation scripts verify checksums before installation
- Use HTTPS for all downloads
- GitHub Actions uses secure tokens for publishing