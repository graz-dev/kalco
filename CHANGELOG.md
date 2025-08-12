# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- 🚀 GitHub Actions workflows for automated releases
- 🐳 Docker container support with multi-platform builds
- 📦 Automated package creation for all platforms (Linux, macOS, Windows)
- 🔍 Enhanced change reports with detailed resource diffs
- 📊 Comprehensive testing workflows
- 🔒 Security scanning with Trivy vulnerability detection
- 📋 Release drafter for automatic changelog generation

### Changed
- 📖 Improved README with comprehensive documentation
- 🔧 Enhanced Makefile with development and release commands
- 🧪 Consolidated testing into single comprehensive quickstart script

### Fixed
- 🐛 Enhanced error handling in change report generation
- 🔍 Improved Git diff analysis for resource changes

## [0.2.0] - 2025-08-13

### Added
- 🚀 Git integration with automatic repository initialization
- 📊 Enhanced change reports with detailed resource tracking
- 🔍 Field-level change identification in reports
- 📋 Automatic markdown report generation for each snapshot
- 🌐 Remote origin detection and push automation

### Changed
- 📁 Improved output structure organization
- 🔧 Better error handling and resilience
- 📖 Comprehensive documentation and examples

### Fixed
- 🐛 Enhanced metadata cleanup for exported resources
- 🔍 Better resource discovery and enumeration

## [0.1.0] - 2025-08-13

### Added
- 🎯 Complete Kubernetes resource discovery and export
- 📁 Structured output organization by namespace and resource type
- 🧹 Intelligent metadata cleanup for re-application
- 🌐 Support for both in-cluster and out-of-cluster execution
- ⚡ Optimized performance for production environments
- 📊 Support for Custom Resources (CRDs)
- 🏷️ Namespace-aware resource organization

### Features
- **Resource Discovery**: Automatically finds all available API resources
- **Comprehensive Coverage**: Includes native K8s resources and CRDs
- **Structured Output**: Creates intuitive directory layouts
- **Clean YAML**: Removes unnecessary metadata fields
- **Multi-Platform**: Works on Linux, macOS, and Windows
- **Production Ready**: Optimized for speed and reliability

---

## Release Process

### Creating a Release

1. **Prepare Changes**: Ensure all changes are committed and pushed to main
2. **Create Tag**: `git tag v1.0.0`
3. **Push Tag**: `git push origin v1.0.0`
4. **Automated Build**: GitHub Actions will automatically:
   - Build binaries for all platforms
   - Create GitHub release
   - Push Docker images to GitHub Container Registry
   - Generate release notes

### Versioning

- **Major**: Breaking changes or major feature additions
- **Minor**: New features or enhancements
- **Patch**: Bug fixes and minor improvements

### Supported Platforms

- **Linux**: AMD64, ARM64
- **macOS**: AMD64, ARM64 (Apple Silicon)
- **Windows**: AMD64, ARM64

### Installation Methods

- **Binary**: Download platform-specific binary from releases
- **Docker**: `docker pull ghcr.io/yourusername/kalco:latest`
- **Source**: `go install github.com/yourusername/kalco@latest`

---

*For detailed information about each release, see the [GitHub releases page](https://github.com/yourusername/kalco/releases).*
