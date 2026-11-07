# Tickets

Small, numbered steps. Each one teaches something and leaves the project working.

## Where things are

```
tickets/
├── M0-setup/          GW-001 … GW-004
├── M1-go-basics/      GW-005 … GW-012
├── M2-core-proxy/     GW-013 … GW-020
├── M3-streaming/      GW-021 … GW-025
└── M4-ship-v0.1/      GW-026 … GW-028
```

The order across both repos (`llm-gateway` and `pawcite`) is in `../LEARNING-PATH.md`.
Progress is ticked in `../ROADMAP.md`. Tickets for later milestones are written when you get there,
so they can build on what you have learned.

## Ticket format

Every ticket has a header (read by `scripts/create-issues`) and the same sections:

| Section | What it holds |
|---|---|
| Header | `id`, `title`, `milestone`, `labels`, `size`, `depends_on` |
| Goal | The result, in one or two sentences |
| Why it matters | How this fits the gateway and the PawCite app |
| You will learn | The new concepts, with JavaScript comparisons |
| Steps | Small, ordered steps |
| Acceptance criteria | Checkboxes that must all be true |
| Explain it back | Questions to answer before closing the ticket |
| Notes for the AI | Extra instructions for the coding assistant (optional) |

## Sizes and labels

- **S**: about an hour or less. **M**: one to two hours.
- `hands-on`: **you** write the code; the AI explains, hints and reviews.
- Other labels describe the area: `go`, `api`, `streaming`, `testing`, `ci`, `docker`, `config`, `setup`, `ai`, `learning`.

## How to work on a ticket

1. In Claude Code: `/start-ticket GW-005` (or say "start GW-005").
2. Follow the steps with the AI; ask `/explain <topic>` whenever something is unclear.
3. `/review` before opening the pull request.
4. The pull request says `Closes #<issue>`; merging moves the card to Done.
5. `/quiz` for the "Explain it back" questions, then add your notes to `docs/learning-log.md`.

## Definition of done

- All acceptance criteria are ticked, with evidence in the pull request.
- Checks pass locally and in CI (from GW-012 on).
- You can answer the "Explain it back" questions without looking.
