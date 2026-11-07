---
id: GW-009
title: Create the Go module and project layout
milestone: GW-M1 · Go basics
labels: go, setup
size: S (about 45 min)
depends_on: GW-008
---

# GW-009 · Create the Go module and project layout

## Goal
The repo is a Go module with the standard layout, and `go run ./cmd/gateway` prints a version.

## Why it matters
Every later package lives in this structure.

## You will learn
- `go.mod`: the module's name and Go version (≈ `package.json`, without scripts).
- Module path: `github.com/<you>/llm-gateway`, also the import prefix of your packages.
- One package per directory. `cmd/<name>/` holds programs; `internal/` holds packages that
  **cannot** be imported from outside this module (the compiler enforces it).
- Injecting a version at build time with `-ldflags`.

## Steps
1. `go mod init github.com/<your-github-user>/llm-gateway`
2. Create `internal/version/version.go` with `var Version = "dev"`.
3. Create `cmd/gateway/main.go` that logs `llm-gateway starting` and the version.
4. Run: `go run ./cmd/gateway`
5. Build with a version:
   `go build -ldflags "-X github.com/<you>/llm-gateway/internal/version.Version=0.0.1" -o bin/gateway ./cmd/gateway`
   then run `./bin/gateway`.
6. Check `go build ./...` builds everything, including the playground.

## Acceptance criteria
- [ ] `go run ./cmd/gateway` prints the start message with version `dev`.
- [ ] The built binary prints version `0.0.1`.
- [ ] `bin/` is ignored by git.

## Explain it back
1. What stops another project from importing `internal/version`?
2. Why inject the version at build time instead of hard-coding it?
