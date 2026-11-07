---
id: GW-011
title: Format, vet, lint, and a Makefile
milestone: GW-M1 · Go basics
labels: go, ci
size: S (about 45 min)
depends_on: GW-010
---

# GW-011 · Format, vet, lint, and a Makefile

## Goal
One command, `make check`, runs every quality check.

## Why it matters
The same command will run in CI, before every commit, and in the AI's workflow.

## You will learn
- `gofmt` (one official style), `go vet` (finds suspicious code), and golangci-lint (runs many
  linters at once; ≈ ESLint with a big preset).
- Makefiles are the conventional task runner in Go repos (≈ `npm run` scripts).

## Steps
1. Install golangci-lint v2 (instructions at https://golangci-lint.run).
2. Create `.golangci.yml`:
   ```yaml
   version: "2"
   linters:
     default: standard
   formatters:
     enable:
       - gofmt
   ```
3. Create a `Makefile` with targets `fmt`, `vet`, `lint`, `test` (`go test -race ./...`),
   `check` (all four), `build` and `run`.
4. Run `make check` and fix anything it reports.
5. Windows: `make` works in Git Bash or WSL. The race detector (`-race`) needs a C compiler;
   if it is missing locally, run tests without `-race` and let CI run it.

## Acceptance criteria
- [ ] `make check` passes.
- [ ] Adding an unused variable makes `make check` fail; removing it makes it pass again.

## Explain it back
1. What kind of problem does `go vet` catch that the compiler does not?
2. What does `-race` do?
