#!/usr/bin/env bash
set -euo pipefail

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

error() { echo -e "${RED}Error: $1${NC}" >&2; exit 1; }
info() { echo -e "${GREEN}$1${NC}"; }
warn() { echo -e "${YELLOW}$1${NC}"; }

# Check for uncommitted changes
if ! git diff --quiet || ! git diff --cached --quiet; then
    error "Uncommitted changes detected. Please commit or stash them first."
fi

# Get current version from latest tag
CURRENT_VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
info "Current version: $CURRENT_VERSION"

# Parse version components
VERSION_REGEX="^v([0-9]+)\.([0-9]+)\.([0-9]+)$"
if [[ $CURRENT_VERSION =~ $VERSION_REGEX ]]; then
    MAJOR="${BASH_REMATCH[1]}"
    MINOR="${BASH_REMATCH[2]}"
    PATCH="${BASH_REMATCH[3]}"
else
    MAJOR=0; MINOR=0; PATCH=0
fi

# Suggest next versions
NEXT_PATCH="v${MAJOR}.${MINOR}.$((PATCH + 1))"
NEXT_MINOR="v${MAJOR}.$((MINOR + 1)).0"
NEXT_MAJOR="v$((MAJOR + 1)).0.0"

echo ""
echo "Select version bump:"
echo "  1) Patch: $NEXT_PATCH (bug fixes)"
echo "  2) Minor: $NEXT_MINOR (new features)"
echo "  3) Major: $NEXT_MAJOR (breaking changes)"
echo "  4) Custom version"
echo ""
read -rp "Choice [1-4]: " CHOICE

case $CHOICE in
    1) NEW_VERSION="$NEXT_PATCH" ;;
    2) NEW_VERSION="$NEXT_MINOR" ;;
    3) NEW_VERSION="$NEXT_MAJOR" ;;
    4)
        read -rp "Enter version (e.g., v1.2.3): " NEW_VERSION
        if [[ ! $NEW_VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
            error "Invalid version format. Use vX.Y.Z"
        fi
        ;;
    *) error "Invalid choice" ;;
esac

echo ""
warn "This will:"
echo "  - Create tag: $NEW_VERSION"
echo "  - Push tag to origin"
echo "  - Trigger GitHub Actions release workflow"
echo "  - Build binaries for linux/darwin (amd64/arm64)"
echo "  - Create GitHub release"
echo "  - Update Homebrew formula"
echo ""
read -rp "Continue? [y/N]: " CONFIRM

if [[ ! $CONFIRM =~ ^[Yy]$ ]]; then
    info "Aborted."
    exit 0
fi

# Create and push tag
info "Creating tag $NEW_VERSION..."
git tag -a "$NEW_VERSION" -m "Release $NEW_VERSION"

info "Pushing tag to origin..."
git push origin "$NEW_VERSION"

echo ""
info "Release $NEW_VERSION initiated!"
echo ""
echo "Monitor progress at:"
echo "  https://github.com/dkd-dobberkau/ddev-mngr/actions"
echo ""
echo "After release completes, install via:"
echo "  brew tap dkd-dobberkau/tap"
echo "  brew install ddev-mngr"
