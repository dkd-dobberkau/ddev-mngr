# Homebrew Tap Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Make ddev-mngr installable via `brew tap dkd-dobberkau/tap && brew install ddev-mngr`

**Architecture:** GoReleaser builds multi-platform binaries on git tag push, creates GitHub Release, and auto-updates Homebrew Formula in separate tap repository.

**Tech Stack:** GoReleaser, GitHub Actions, Homebrew

---

### Task 1: Update Module Path

**Files:**
- Modify: `go.mod`
- Modify: all files importing internal packages

**Step 1: Update go.mod module path**

Change the module line in `go.mod` from:
```
module github.com/oliverguenther/ddev-mngr
```
to:
```
module github.com/dkd-dobberkau/ddev-mngr
```

**Step 2: Update all import paths**

Update imports in these files:
- `main.go`
- `cmd/root.go`
- `cmd/list.go`
- `cmd/start.go`
- `cmd/stop.go`
- `cmd/status.go`
- `internal/tui/model.go`
- `internal/tui/update.go`

Change all occurrences of:
```go
"github.com/oliverguenther/ddev-mngr/..."
```
to:
```go
"github.com/dkd-dobberkau/ddev-mngr/..."
```

**Step 3: Verify build**

Run: `go build -o ddev-mngr`
Expected: Build succeeds

**Step 4: Commit**

```bash
git add -A
git commit -m "chore: update module path to dkd-dobberkau"
```

---

### Task 2: Add Apache 2.0 License

**Files:**
- Create: `LICENSE`

**Step 1: Create LICENSE file**

```
                                 Apache License
                           Version 2.0, January 2004
                        http://www.apache.org/licenses/

   TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION

   1. Definitions.

      "License" shall mean the terms and conditions for use, reproduction,
      and distribution as defined by Sections 1 through 9 of this document.

      "Licensor" shall mean the copyright owner or entity authorized by
      the copyright owner that is granting the License.

      "Legal Entity" shall mean the union of the acting entity and all
      other entities that control, are controlled by, or are under common
      control with that entity. For the purposes of this definition,
      "control" means (i) the power, direct or indirect, to cause the
      direction or management of such entity, whether by contract or
      otherwise, or (ii) ownership of fifty percent (50%) or more of the
      outstanding shares, or (iii) beneficial ownership of such entity.

      "You" (or "Your") shall mean an individual or Legal Entity
      exercising permissions granted by this License.

      "Source" form shall mean the preferred form for making modifications,
      including but not limited to software source code, documentation
      source, and configuration files.

      "Object" form shall mean any form resulting from mechanical
      transformation or translation of a Source form, including but
      not limited to compiled object code, generated documentation,
      and conversions to other media types.

      "Work" shall mean the work of authorship, whether in Source or
      Object form, made available under the License, as indicated by a
      copyright notice that is included in or attached to the work
      (an example is provided in the Appendix below).

      "Derivative Works" shall mean any work, whether in Source or Object
      form, that is based on (or derived from) the Work and for which the
      editorial revisions, annotations, elaborations, or other modifications
      represent, as a whole, an original work of authorship. For the purposes
      of this License, Derivative Works shall not include works that remain
      separable from, or merely link (or bind by name) to the interfaces of,
      the Work and Derivative Works thereof.

      "Contribution" shall mean any work of authorship, including
      the original version of the Work and any modifications or additions
      to that Work or Derivative Works thereof, that is intentionally
      submitted to the Licensor for inclusion in the Work by the copyright owner
      or by an individual or Legal Entity authorized to submit on behalf of
      the copyright owner. For the purposes of this definition, "submitted"
      means any form of electronic, verbal, or written communication sent
      to the Licensor or its representatives, including but not limited to
      communication on electronic mailing lists, source code control systems,
      and issue tracking systems that are managed by, or on behalf of, the
      Licensor for the purpose of discussing and improving the Work, but
      excluding communication that is conspicuously marked or otherwise
      designated in writing by the copyright owner as "Not a Contribution."

      "Contributor" shall mean Licensor and any individual or Legal Entity
      on behalf of whom a Contribution has been received by Licensor and
      subsequently incorporated within the Work.

   2. Grant of Copyright License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      copyright license to reproduce, prepare Derivative Works of,
      publicly display, publicly perform, sublicense, and distribute the
      Work and such Derivative Works in Source or Object form.

   3. Grant of Patent License. Subject to the terms and conditions of
      this License, each Contributor hereby grants to You a perpetual,
      worldwide, non-exclusive, no-charge, royalty-free, irrevocable
      (except as stated in this section) patent license to make, have made,
      use, offer to sell, sell, import, and otherwise transfer the Work,
      where such license applies only to those patent claims licensable
      by such Contributor that are necessarily infringed by their
      Contribution(s) alone or by combination of their Contribution(s)
      with the Work to which such Contribution(s) was submitted. If You
      institute patent litigation against any entity (including a
      cross-claim or counterclaim in a lawsuit) alleging that the Work
      or a Contribution incorporated within the Work constitutes direct
      or contributory patent infringement, then any patent licenses
      granted to You under this License for that Work shall terminate
      as of the date such litigation is filed.

   4. Redistribution. You may reproduce and distribute copies of the
      Work or Derivative Works thereof in any medium, with or without
      modifications, and in Source or Object form, provided that You
      meet the following conditions:

      (a) You must give any other recipients of the Work or
          Derivative Works a copy of this License; and

      (b) You must cause any modified files to carry prominent notices
          stating that You changed the files; and

      (c) You must retain, in the Source form of any Derivative Works
          that You distribute, all copyright, patent, trademark, and
          attribution notices from the Source form of the Work,
          excluding those notices that do not pertain to any part of
          the Derivative Works; and

      (d) If the Work includes a "NOTICE" text file as part of its
          distribution, then any Derivative Works that You distribute must
          include a readable copy of the attribution notices contained
          within such NOTICE file, excluding those notices that do not
          pertain to any part of the Derivative Works, in at least one
          of the following places: within a NOTICE text file distributed
          as part of the Derivative Works; within the Source form or
          documentation, if provided along with the Derivative Works; or,
          within a display generated by the Derivative Works, if and
          wherever such third-party notices normally appear. The contents
          of the NOTICE file are for informational purposes only and
          do not modify the License. You may add Your own attribution
          notices within Derivative Works that You distribute, alongside
          or as an addendum to the NOTICE text from the Work, provided
          that such additional attribution notices cannot be construed
          as modifying the License.

      You may add Your own copyright statement to Your modifications and
      may provide additional or different license terms and conditions
      for use, reproduction, or distribution of Your modifications, or
      for any such Derivative Works as a whole, provided Your use,
      reproduction, and distribution of the Work otherwise complies with
      the conditions stated in this License.

   5. Submission of Contributions. Unless You explicitly state otherwise,
      any Contribution intentionally submitted for inclusion in the Work
      by You to the Licensor shall be under the terms and conditions of
      this License, without any additional terms or conditions.
      Notwithstanding the above, nothing herein shall supersede or modify
      the terms of any separate license agreement you may have executed
      with Licensor regarding such Contributions.

   6. Trademarks. This License does not grant permission to use the trade
      names, trademarks, service marks, or product names of the Licensor,
      except as required for reasonable and customary use in describing the
      origin of the Work and reproducing the content of the NOTICE file.

   7. Disclaimer of Warranty. Unless required by applicable law or
      agreed to in writing, Licensor provides the Work (and each
      Contributor provides its Contributions) on an "AS IS" BASIS,
      WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
      implied, including, without limitation, any warranties or conditions
      of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
      PARTICULAR PURPOSE. You are solely responsible for determining the
      appropriateness of using or redistributing the Work and assume any
      risks associated with Your exercise of permissions under this License.

   8. Limitation of Liability. In no event and under no legal theory,
      whether in tort (including negligence), contract, or otherwise,
      unless required by applicable law (such as deliberate and grossly
      negligent acts) or agreed to in writing, shall any Contributor be
      liable to You for damages, including any direct, indirect, special,
      incidental, or consequential damages of any character arising as a
      result of this License or out of the use or inability to use the
      Work (including but not limited to damages for loss of goodwill,
      work stoppage, computer failure or malfunction, or any and all
      other commercial damages or losses), even if such Contributor
      has been advised of the possibility of such damages.

   9. Accepting Warranty or Additional Liability. While redistributing
      the Work or Derivative Works thereof, You may choose to offer,
      and charge a fee for, acceptance of support, warranty, indemnity,
      or other liability obligations and/or rights consistent with this
      License. However, in accepting such obligations, You may act only
      on Your own behalf and on Your sole responsibility, not on behalf
      of any other Contributor, and only if You agree to indemnify,
      defend, and hold each Contributor harmless for any liability
      incurred by, or claims asserted against, such Contributor by reason
      of your accepting any such warranty or additional liability.

   END OF TERMS AND CONDITIONS

   Copyright 2026 dkd-dobberkau

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```

**Step 2: Commit**

```bash
git add LICENSE
git commit -m "chore: add Apache 2.0 license"
```

---

### Task 3: Add GoReleaser Configuration

**Files:**
- Create: `.goreleaser.yaml`

**Step 1: Create GoReleaser config**

```yaml
version: 2

project_name: ddev-mngr

before:
  hooks:
    - go mod tidy

builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - darwin
    goarch:
      - amd64
      - arm64
    ldflags:
      - -s -w -X main.version={{.Version}}

archives:
  - format: tar.gz
    name_template: >-
      {{ .ProjectName }}_
      {{- .Version }}_
      {{- .Os }}_
      {{- .Arch }}

checksum:
  name_template: 'checksums.txt'

changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'
      - '^chore:'

brews:
  - name: ddev-mngr
    repository:
      owner: dkd-dobberkau
      name: homebrew-tap
      token: "{{ .Env.HOMEBREW_TAP_TOKEN }}"
    directory: Formula
    homepage: "https://github.com/dkd-dobberkau/ddev-mngr"
    description: "CLI tool to manage DDEV projects with interactive TUI"
    license: "Apache-2.0"
    install: |
      bin.install "ddev-mngr"
    test: |
      system "#{bin}/ddev-mngr", "--help"
```

**Step 2: Commit**

```bash
git add .goreleaser.yaml
git commit -m "chore: add GoReleaser configuration"
```

---

### Task 4: Add GitHub Actions Release Workflow

**Files:**
- Create: `.github/workflows/release.yaml`

**Step 1: Create workflow directory**

Run: `mkdir -p .github/workflows`

**Step 2: Create release workflow**

```yaml
name: Release

on:
  push:
    tags:
      - 'v*'

permissions:
  contents: write

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'

      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v6
        with:
          distribution: goreleaser
          version: latest
          args: release --clean
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          HOMEBREW_TAP_TOKEN: ${{ secrets.HOMEBREW_TAP_TOKEN }}
```

**Step 3: Commit**

```bash
git add .github/
git commit -m "ci: add GitHub Actions release workflow"
```

---

### Task 5: Update README with Installation Instructions

**Files:**
- Modify: `README.md`

**Step 1: Update README**

Add Homebrew installation section at the top of the Installation section:

```markdown
## Installation

### Homebrew (macOS/Linux)

```bash
brew tap dkd-dobberkau/tap
brew install ddev-mngr
```

### Go Install

```bash
go install github.com/dkd-dobberkau/ddev-mngr@latest
```

### Build from Source

```bash
git clone https://github.com/dkd-dobberkau/ddev-mngr.git
cd ddev-mngr
go build -o ddev-mngr
```
```

**Step 2: Commit**

```bash
git add README.md
git commit -m "docs: add Homebrew installation instructions"
```

---

### Task 6: Manual GitHub Setup (Instructions)

This task requires manual steps in the GitHub UI.

**Step 1: Create main repository**

1. Go to https://github.com/new
2. Repository name: `ddev-mngr`
3. Owner: `dkd-dobberkau`
4. Public repository
5. Don't initialize with README (we'll push existing code)

**Step 2: Create tap repository**

1. Go to https://github.com/new
2. Repository name: `homebrew-tap`
3. Owner: `dkd-dobberkau`
4. Public repository
5. Initialize with README

**Step 3: Create Personal Access Token**

1. Go to https://github.com/settings/tokens
2. Generate new token (classic)
3. Name: `HOMEBREW_TAP_TOKEN`
4. Scope: `repo` (full control)
5. Copy the token

**Step 4: Add secret to ddev-mngr repository**

1. Go to `https://github.com/dkd-dobberkau/ddev-mngr/settings/secrets/actions`
2. New repository secret
3. Name: `HOMEBREW_TAP_TOKEN`
4. Value: (paste the token)

**Step 5: Push code to GitHub**

```bash
git remote add origin git@github.com:dkd-dobberkau/ddev-mngr.git
git push -u origin main
```

**Step 6: Create first release**

```bash
git tag v0.1.0
git push origin v0.1.0
```

Watch GitHub Actions run and verify:
- Release created with binaries
- Formula pushed to homebrew-tap repository

---

## Summary

5 code tasks + 1 manual setup task:

1. Update module path
2. Add Apache 2.0 license
3. Add GoReleaser configuration
4. Add GitHub Actions workflow
5. Update README
6. Manual GitHub setup (repositories, secrets, push)
