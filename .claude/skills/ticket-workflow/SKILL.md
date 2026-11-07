---
name: ticket-workflow
description: Start, work on and finish a ticket from tickets/. Use when the human says "start GW-007", "next", or asks to continue or close the current ticket.
---

# ticket-workflow

One ticket = one branch = one pull request.

## 1. Start

1. Read the **whole** ticket file: `tickets/<milestone>/<ID>-<name>.md`.
2. Check `depends_on`. If a dependency is not ticked in `ROADMAP.md`, say so and stop.
3. Branch from an up-to-date `main`, named from the id and title, lowercase and hyphenated:
   ```
   git switch main && git pull && git switch -c gw-007-errors-and-defer
   ```
4. Say back, in three or four lines: the goal, the acceptance criteria, whether the ticket is
   `hands-on`, and which concepts are new to this repo.
5. New concept → use the `tutor` skill **before** any code is written.

## 2. Work

- `hands-on` label → **the human writes the code.** You explain, hint, point at the line that is
  wrong, and review. Do not write the solution, not even "to show the shape of it", unless the
  human says "show me the solution". A hint names the thing to look up; it does not paste the
  answer.
- Not `hands-on` → you write the code, following the `go-style` skill, in small steps the human
  can follow. Explain each new construct as it appears.
- Follow the ticket's Steps in order. If a step is wrong or impossible, say so and propose the fix
  before doing something different.
- Stay inside the ticket. Something else that needs doing becomes a note, not a surprise commit.

## 3. Check

Run the checks and **show the output**. Never claim green without it.

- From GW-011 on: `make check`
- Before that: `gofmt -l .` (prints nothing), `go vet ./...`, `go test ./...`

Then walk the acceptance criteria one by one and show the evidence for each. A criterion with no
evidence is not met.

## 4. Review and pull request

1. Use the `code-review` skill on the branch.
2. Commit in small, readable commits; the subject is imperative and under 72 characters.
3. Find the issue number for this ticket — it is not the ticket number:
   ```
   gh issue list --search "GW-007" --state all --json number,title
   ```
4. Open the pull request with the ticket id in the title and `Closes #<issue>` in the body, plus
   the acceptance criteria as a checklist with evidence.
   ```
   gh pr create --title "GW-007 · Errors and defer" --body "Closes #11 ..."
   ```
5. CI must be green (from GW-012 on).

## 5. Finish — "Explain it back"

Do not skip this. It is why the project exists.

1. Ask the ticket's **Explain it back** questions **one at a time**. Wait for each answer.
2. Correct gently: what is right, in one sentence, and why.
3. Add 2–3 bullets to `docs/learning-log.md`, in the human's own words from their answers —
   not your summary. Head them with the ticket id.
4. Tick the ticket in `ROADMAP.md`.
5. Only now is the ticket done.

## Definition of done

Every acceptance criterion met with evidence · checks green · "Explain it back" answered ·
learning log updated · `ROADMAP.md` ticked.
