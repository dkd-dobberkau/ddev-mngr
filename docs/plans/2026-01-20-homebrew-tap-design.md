# Homebrew Tap Design

## Overview

Make ddev-mngr installable via Homebrew.

## Installation (End Result)

```bash
brew tap dkd-dobberkau/tap
brew install ddev-mngr
```

## Components

### 1. Main Repository (github.com/dkd-dobberkau/ddev-mngr)

Current code plus:
- `.goreleaser.yaml` - Build configuration
- `.github/workflows/release.yaml` - CI/CD workflow
- `LICENSE` - Apache 2.0

### 2. Tap Repository (github.com/dkd-dobberkau/homebrew-tap)

- Contains Homebrew Formula
- Auto-updated by GoReleaser on release

## Release Flow

```
git tag v1.0.0 && git push --tags
        ↓
GitHub Actions triggered
        ↓
GoReleaser builds binaries (macOS arm64/amd64, Linux arm64/amd64)
        ↓
GoReleaser creates GitHub Release with binaries
        ↓
GoReleaser updates Formula in homebrew-tap
```

## Build Targets

| OS | Arch | Binary |
|----|------|--------|
| macOS | arm64 | ddev-mngr_darwin_arm64 |
| macOS | amd64 | ddev-mngr_darwin_amd64 |
| Linux | amd64 | ddev-mngr_linux_amd64 |
| Linux | arm64 | ddev-mngr_linux_arm64 |

## Files to Create/Modify

| File | Action |
|------|--------|
| `go.mod` | Change module path to github.com/dkd-dobberkau/ddev-mngr |
| `LICENSE` | Add Apache 2.0 license |
| `.goreleaser.yaml` | Add GoReleaser config |
| `.github/workflows/release.yaml` | Add release workflow |

## GitHub Setup Required

1. Create repository `dkd-dobberkau/ddev-mngr`
2. Create repository `dkd-dobberkau/homebrew-tap`
3. Create Personal Access Token with `repo` scope
4. Add secret `HOMEBREW_TAP_TOKEN` to ddev-mngr repository

## License

Apache 2.0 (consistent with DDEV)
