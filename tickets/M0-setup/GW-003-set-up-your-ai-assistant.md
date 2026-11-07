---
id: GW-003
title: Set up your AI coding assistant
milestone: GW-M0 · Setup
labels: setup, ai
size: S (about 30 min)
depends_on: GW-002
---

# GW-003 · Set up your AI coding assistant

## Goal
Claude Code (or Codex) runs in this repo and follows its rules, skills and commands.

## Why it matters
The assistant writes much of the code, so it must follow the project's rules: small changes,
tests, explanations, and leaving `hands-on` tickets to you.

## You will learn
- `AGENTS.md`: the rules every assistant reads. `CLAUDE.md` imports it for Claude Code.
- **Skills** (`.claude/skills/*/SKILL.md`): instructions for a kind of task (teaching, reviewing…).
- **Commands** (`.claude/commands/*.md`): shortcuts like `/explain` that trigger a skill.

## Steps
1. Install Claude Code by following https://docs.claude.com/en/docs/claude-code/overview
   (or Codex: `npm i -g @openai/codex`).
2. Start it in the repo folder and ask: "Summarise the rules of this repo in five bullets."
3. Try the tutor: `/explain what is a Go module` (in Codex: "explain what is a Go module").
4. Read the four skills in `.claude/skills/`. Edit the learner profile in `tutor/SKILL.md` if
   anything about you is wrong.
5. Commit any changes with the message `GW-003: assistant setup`.

## Acceptance criteria
- [ ] The assistant's summary mentions one ticket at a time, `hands-on` tickets, and "Explain it back".
- [ ] The `/explain` answer has a JavaScript comparison, an example from the dog app, and a check question.

## Explain it back
1. What is the difference between `AGENTS.md`, a skill and a command?
2. What should the assistant do when you start a `hands-on` ticket?

## Notes for the AI
This ticket is about the workflow, not Go. If a rule seems unclear to the human, suggest an edit.
