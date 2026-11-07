# AGENTS.md: rules for AI coding assistants

This file is for AI assistants (Claude Code, Codex and others). Humans are welcome to read it too.

## About this project

`llm-gateway` is an OpenAI-compatible LLM gateway written in Go. Every model call from the
**PawCite** app (a dog-research assistant, repo `pawcite`) goes through it: chat, embeddings,
judges and agents.

It is also a **learning project**. The human:
- is a senior JavaScript/TypeScript developer (8 years) who knows AWS, Terraform and Docker,
- is **new to Go**,
- wants production-quality code **and** to understand every line of it.

Both goals matter equally. Fast code the human cannot explain is a failure.

## How we work

1. **One ticket at a time.** Tickets live in `tickets/<milestone>/<ID>-<name>.md`. Read the whole
   ticket before doing anything. The order is in `LEARNING-PATH.md`; check `depends_on`.
2. **Explain before you write.** When a ticket uses a Go concept that has not appeared in this repo
   before, explain it first using the `tutor` skill. Keep it short and concrete.
3. **Respect `hands-on` tickets.** For tickets labelled `hands-on`, the human writes the code. You
   explain, give hints and review. Do not write the solution unless the human explicitly asks
   "show me the solution".
4. **Small changes.** One ticket = one branch = one pull request. Branch names: `gw-007-errors-and-defer`.
5. **Done means done.** Never mark a ticket done until every acceptance criterion is met and the
   checks pass. Show the evidence (test output, command output).
6. **Finish with "Explain it back".** Ask the ticket's questions one at a time, wait for answers,
   and correct gently. Then add 2–3 bullets to `docs/learning-log.md`.

## Technical rules

- **Go 1.27, standard library first.** Dependencies are added only when a ticket introduces them
  (later: pgx, Prometheus client, OpenTelemetry, a Redis client). Ask before adding anything else.
- **Idiomatic Go:** errors are values and are wrapped with `%w`; `context.Context` is the first
  parameter of anything that blocks; no `panic` in request paths; small interfaces defined where
  they are used; table-driven tests; no global mutable state. Details: `go-style` skill.
- **Layout:** `cmd/` (programs), `internal/` (packages), `test/` (black-box tests), `scripts/`.
- **Checks:** `make check` once the Makefile exists (ticket GW-011). Before that:
  `gofmt -l .`, `go vet ./...`, `go test ./...`.
- **Secrets:** API keys come from environment variables only. Never commit them, never print them.
- **Compatibility:** the HTTP API under `/v1` must stay OpenAI-compatible.

## Skills

Skills live in `.claude/skills/<name>/SKILL.md`. When a task matches a skill, read that file first.

| Skill | Use it when |
|---|---|
| `tutor` | A new concept appears, or the human asks you to explain something |
| `ticket-workflow` | Starting, working on or finishing a ticket |
| `code-review` | Reviewing changes before a pull request is merged |
| `go-style` | Writing or changing Go code |

## Phrases the human may use

| The human says | You do |
|---|---|
| "start GW-007" | Follow `ticket-workflow` for ticket GW-007 |
| "explain channels" | Follow `tutor` |
| "review" | Follow `code-review` for the current branch |
| "quiz me" | Ask the current ticket's "Explain it back" questions |
| "next" | Find the next ticket in `LEARNING-PATH.md` whose dependencies are done |

Claude Code also has these as slash commands in `.claude/commands/`.
