# Learning log

Two or three bullets per ticket: what I learned, in my own words.

## GW-001 Install Go and your tools

- The `go` command is node + npm + tsc + jest + prettier in one binary. No bundler, because
  `go build` emits one static binary with nothing to bundle.
- `go vet` is much narrower than eslint — real bugs only, no style. golangci-lint fills the gap
  in GW-011. There is no `scripts` section, which is what the Makefile is for.
- `gopls` is the language server (Go's tsserver). Without it I would be hand-editing imports
  constantly, because an unused import is a compile error in Go, not a warning.

## GW-002 Create the repository on GitHub

- Public because the repo is the portfolio piece — the tickets, PRs and learning log show how
  I work.
- So: never commit sensitive credentials, or anything else hackers could use. Keys come from
  environment variables only.
- Git history is permanent. Deleting a key in a later commit does nothing — the only fix is to
  rotate it at the provider.

## GW-003 Set up your AI coding assistant

- `AGENTS.md` is always loaded; a skill is loaded by the assistant when the task matches its
  `description`; a command is loaded by me typing `/name`. Commands usually just point at a skill.
- On a `hands-on` ticket I write the code. The assistant explains, hints and reviews, and only
  writes the solution if I say "show me the solution". GW-005 to GW-008 are hands-on.

## GW-004 Turn tickets into GitHub issues and a project board

- `gh` returns `(string, error)` because Go has no exceptions — a function that can fail says so
  in its type. It always returns both; the error is nil when nothing went wrong.
- `_` is the blank identifier, discarding the index from `range`. Needed because an unused
  variable is a compile error. Careful: one variable gives the index, not the element.
