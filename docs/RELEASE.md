# Release Workflow

## Übersicht

Releases werden automatisch via GitHub Actions und GoReleaser erstellt.

## Neuen Release erstellen

```bash
# 1. Sicherstellen dass alles committed und gepusht ist
git status
git push origin main

# 2. Tag erstellen und pushen
git tag v0.2.0
git push origin v0.2.0
```

Das war's! GitHub Actions:
- Baut Binaries für macOS (arm64/amd64) und Linux (arm64/amd64)
- Erstellt GitHub Release mit allen Binaries
- Aktualisiert automatisch die Homebrew Formula

## Versionierung

Wir verwenden [Semantic Versioning](https://semver.org/):

- `v1.0.0` → Major (breaking changes)
- `v1.1.0` → Minor (neue Features)
- `v1.1.1` → Patch (Bugfixes)

## Release prüfen

```bash
# GitHub Release anzeigen
gh release view v0.2.0

# Workflow-Status prüfen
gh run list --limit 1

# Homebrew Formula prüfen
gh api repos/dkd-dobberkau/homebrew-tap/contents/Formula/ddev-mngr.rb --jq '.name'
```

## Homebrew Installation testen

```bash
# Tap hinzufügen (einmalig)
brew tap dkd-dobberkau/tap

# Installieren
brew install ddev-mngr

# Updaten nach neuem Release
brew upgrade ddev-mngr
```

## Repositories

| Repository | Zweck |
|------------|-------|
| [dkd-dobberkau/ddev-mngr](https://github.com/dkd-dobberkau/ddev-mngr) | Hauptrepository mit Code |
| [dkd-dobberkau/homebrew-tap](https://github.com/dkd-dobberkau/homebrew-tap) | Homebrew Formulas |

## Konfigurationsdateien

| Datei | Zweck |
|-------|-------|
| `.goreleaser.yaml` | GoReleaser Build-Konfiguration |
| `.github/workflows/release.yaml` | GitHub Actions Workflow |

## Secrets

Das Repository benötigt folgendes Secret:

| Name | Zweck |
|------|-------|
| `HOMEBREW_TAP_TOKEN` | Personal Access Token mit `repo` Scope für Formula-Updates |

Token erneuern: https://github.com/settings/tokens
