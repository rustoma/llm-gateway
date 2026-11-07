---
description: Start a ticket: read it, branch, explain the new concepts
argument-hint: <ticket id, e.g. GW-007>
---

Start ticket **$1** following the `ticket-workflow` skill.

1. Find and read the whole ticket file under `tickets/` for id `$1`.
2. Check `depends_on` against `ROADMAP.md`; if a dependency is not done, say so and stop.
3. Create the branch from an up-to-date `main`.
4. Say back: the goal, the acceptance criteria, whether this ticket is `hands-on`, and which Go
   concepts are new to this repo.
5. Explain each new concept with the `tutor` skill before any code is written.

If the ticket is `hands-on`, the human writes the code. Explain, hint and review only — no
solution unless they say "show me the solution".
