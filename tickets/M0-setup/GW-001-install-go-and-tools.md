---
id: GW-001
title: Install Go and your tools
milestone: GW-M0 · Setup
labels: setup, go
size: S (about 30 min)
depends_on: none
---

# GW-001 · Install Go and your tools

## Goal
Go 1.27, Git, the GitHub CLI and Go support in your editor are installed and working.

## Why it matters
Everything in this repo is built, tested and run with the `go` command.

## You will learn
- What the `go` tool does: it builds, runs, tests, formats and manages dependencies. In the
  JavaScript world that is node + npm + tsc + jest + prettier in one binary.
- What `gopls` is: the Go language server behind autocomplete and "go to definition".

## Steps
1. Install Go 1.27:
   - macOS: `brew install go`
   - Windows: `winget install GoLang.Go`
   - Linux: download from https://go.dev/dl/ and follow the instructions there.
2. Check: `go version` prints `go1.27` or newer.
3. Run `go env GOPATH GOMODCACHE` and ask the AI what these two folders are for.
4. Check the GitHub CLI: `gh --version`, then `gh auth login` if needed.
5. Editor: in VS Code install the **Go** extension (by the Go team). It installs `gopls` for you.

## Acceptance criteria
- [ ] `go version` shows Go 1.27 or newer.
- [ ] `gh auth status` shows you are logged in.
- [ ] Hovering over a Go function in your editor shows its documentation.

## Explain it back
1. Which tools from your JavaScript toolbox does the `go` command replace?
2. What is `gopls`, and what would you lose without it?

## Notes for the AI
Keep explanations short; this is setup. Do not explain GOPATH history in depth.
