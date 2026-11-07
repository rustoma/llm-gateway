---
id: GW-012
title: CI with GitHub Actions
milestone: GW-M1 · Go basics
labels: ci
size: S (about 45 min)
depends_on: GW-011
---

# GW-012 · CI with GitHub Actions

## Goal
Every push and pull request runs the checks on GitHub, and `main` only accepts green pull requests.

## Why it matters
From now on, every ticket ends with a green pull request. That is what reviewers and employers look at.

## Steps
1. Create `.github/workflows/ci.yml` that runs on `push` to `main` and on `pull_request`:
   - `actions/checkout@v7`
   - `actions/setup-go@v7` with `go-version-file: go.mod`
   - `go vet ./...` and `go test -race ./...`
   - `golangci/golangci-lint-action@v9` with `version: v2.15.0`
2. Push on a branch, open a pull request, and watch it go green.
3. Protect `main`: Settings → Rules → Rulesets → require the CI check to pass before merging.
4. Prove it works: push a commit with a failing test, see the red check, then fix it.

## Acceptance criteria
- [ ] CI runs on every pull request and passes.
- [ ] A failing test turns the check red and blocks merging.

## Explain it back
1. Why run the race detector in CI even if tests pass locally?
