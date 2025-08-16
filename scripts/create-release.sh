#!/bin/bash

# Script to help create a new release for kalco

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

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

# Check if we're in a git repository
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    log_error "This script must be run from within a git repository"
    exit 1
fi

# Check if working directory is clean
if ! git diff-index --quiet HEAD --; then
    log_warning "Working directory is not clean. Please commit or stash your changes."
    git status --porcelain
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

# Get current version from git tags
current_version=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
log_info "Current version: $current_version"

# Suggest next version
if [[ $current_version =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]; then
    major=${BASH_REMATCH[1]}
    minor=${BASH_REMATCH[2]}
    patch=${BASH_REMATCH[3]}
    
    next_patch="v$major.$minor.$((patch + 1))"
    next_minor="v$major.$((minor + 1)).0"
    next_major="v$((major + 1)).0.0"
    
    echo
    echo "Suggested versions:"
    echo "  1. Patch: $next_patch (bug fixes)"
    echo "  2. Minor: $next_minor (new features)"
    echo "  3. Major: $next_major (breaking changes)"
    echo "  4. Custom version"
    echo
    
    read -p "Select version type (1-4): " -n 1 -r
    echo
    
    case $REPLY in
        1) new_version=$next_patch ;;
        2) new_version=$next_minor ;;
        3) new_version=$next_major ;;
        4) 
            read -p "Enter custom version (e.g., v1.2.3): " new_version
            if [[ ! $new_version =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
                log_error "Invalid version format. Use vX.Y.Z format."
                exit 1
            fi
            ;;
        *) log_error "Invalid selection"; exit 1 ;;
    esac
else
    read -p "Enter version (e.g., v1.0.0): " new_version
    if [[ ! $new_version =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
        log_error "Invalid version format. Use vX.Y.Z format."
        exit 1
    fi
fi

# Check if tag already exists
if git rev-parse "$new_version" >/dev/null 2>&1; then
    log_error "Tag $new_version already exists"
    exit 1
fi

log_info "Creating release $new_version"

# Get release notes
echo
echo "Enter release notes (press Ctrl+D when finished):"
release_notes=$(cat)

# Create and push tag
log_info "Creating git tag..."
if [ -n "$release_notes" ]; then
    git tag -a "$new_version" -m "$release_notes"
else
    git tag -a "$new_version" -m "Release $new_version"
fi

log_info "Pushing tag to origin..."
git push origin "$new_version"

log_success "Release $new_version created successfully!"
log_info "GitHub Actions will now build and publish the release automatically."
log_info "Check the progress at: https://github.com/$(git config --get remote.origin.url | sed 's/.*github.com[:/]\([^.]*\).*/\1/')/actions"

echo
echo "🎉 Release process initiated!"
echo
echo "What happens next:"
echo "  1. GitHub Actions will build binaries for all platforms"
echo "  2. Create a GitHub release with the binaries"
echo "  3. Update Homebrew tap (if configured)"
echo "  4. Publish packages to package managers"
echo
echo "The release will be available at:"
echo "  https://github.com/$(git config --get remote.origin.url | sed 's/.*github.com[:/]\([^.]*\).*/\1/')/releases/tag/$new_version"