#!/bin/bash

# Kalco Installation Script
# This script downloads and installs the latest version of kalco

set -e

# Configuration
REPO="graz-dev/kalco"
BINARY_NAME="kalco"
INSTALL_DIR="/usr/local/bin"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Helper functions
log_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

log_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

log_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

log_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Detect OS and architecture
detect_platform() {
    local os arch
    
    case "$(uname -s)" in
        Linux*)     os="Linux" ;;
        Darwin*)    os="Darwin" ;;
        CYGWIN*|MINGW*|MSYS*) os="Windows" ;;
        *)          log_error "Unsupported operating system: $(uname -s)" && exit 1 ;;
    esac
    
    case "$(uname -m)" in
        x86_64|amd64)   arch="x86_64" ;;
        arm64|aarch64)  arch="arm64" ;;
        *)              log_error "Unsupported architecture: $(uname -m)" && exit 1 ;;
    esac
    
    echo "${os}_${arch}"
}

# Get latest release version
get_latest_version() {
    log_info "Fetching latest release information..."
    curl -s "https://api.github.com/repos/${REPO}/releases/latest" | \
        grep '"tag_name":' | \
        sed -E 's/.*"([^"]+)".*/\1/'
}

# Download and install kalco
install_kalco() {
    local version platform download_url temp_dir
    
    version=$(get_latest_version)
    if [ -z "$version" ]; then
        log_error "Failed to get latest version"
        exit 1
    fi
    
    platform=$(detect_platform)
    log_info "Detected platform: $platform"
    log_info "Latest version: $version"
    
    # Construct download URL
    if [[ "$platform" == "Windows"* ]]; then
        download_url="https://github.com/${REPO}/releases/download/${version}/kalco_${platform}.zip"
    else
        download_url="https://github.com/${REPO}/releases/download/${version}/kalco_${platform}.tar.gz"
    fi
    
    log_info "Downloading from: $download_url"
    
    # Create temporary directory
    temp_dir=$(mktemp -d)
    cd "$temp_dir"
    
    # Download the archive
    if ! curl -L -o "kalco_archive" "$download_url"; then
        log_error "Failed to download kalco"
        rm -rf "$temp_dir"
        exit 1
    fi
    
    # Extract the archive
    log_info "Extracting archive..."
    if [[ "$platform" == "Windows"* ]]; then
        unzip -q "kalco_archive"
    else
        tar -xzf "kalco_archive"
    fi
    
    # Find the binary
    binary_path=$(find . -name "$BINARY_NAME*" -type f | head -1)
    if [ -z "$binary_path" ]; then
        log_error "Binary not found in archive"
        rm -rf "$temp_dir"
        exit 1
    fi
    
    # Install the binary
    log_info "Installing kalco to $INSTALL_DIR..."
    
    # Check if we need sudo
    if [ -w "$INSTALL_DIR" ]; then
        cp "$binary_path" "$INSTALL_DIR/$BINARY_NAME"
    else
        sudo cp "$binary_path" "$INSTALL_DIR/$BINARY_NAME"
    fi
    
    # Make it executable
    if [ -w "$INSTALL_DIR/$BINARY_NAME" ]; then
        chmod +x "$INSTALL_DIR/$BINARY_NAME"
    else
        sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
    fi
    
    # Cleanup
    rm -rf "$temp_dir"
    
    log_success "kalco $version installed successfully!"
    log_info "Run 'kalco --help' to get started"
}

# Check dependencies
check_dependencies() {
    local missing_deps=()
    
    command -v curl >/dev/null 2>&1 || missing_deps+=("curl")
    command -v tar >/dev/null 2>&1 || missing_deps+=("tar")
    
    if [ ${#missing_deps[@]} -ne 0 ]; then
        log_error "Missing required dependencies: ${missing_deps[*]}"
        log_info "Please install the missing dependencies and try again"
        exit 1
    fi
}

# Main installation process
main() {
    echo "🚀 Kalco Installation Script"
    echo "============================"
    echo
    
    check_dependencies
    install_kalco
    
    echo
    log_success "Installation complete! 🎉"
    echo
    echo "Next steps:"
    echo "  1. Run 'kalco --help' to see available commands"
    echo "  2. Run 'kalco' to dump your current Kubernetes cluster"
    echo "  3. Check out the documentation at: https://github.com/${REPO}"
}

# Run main function
main "$@"