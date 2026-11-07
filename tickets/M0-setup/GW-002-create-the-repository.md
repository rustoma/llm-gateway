---
id: GW-002
title: Create the repository on GitHub
milestone: GW-M0 · Setup
labels: setup
size: S (about 15 min)
depends_on: GW-001
---

# GW-002 · Create the repository on GitHub

## Goal
This starter kit is a public GitHub repository named `llm-gateway`.

## Why it matters
The repo is your portfolio piece. Every ticket becomes an issue and a pull request in it.

## Steps
1. Unzip the starter kit and open a terminal in the `llm-gateway` folder.
2. Create the first commit:
   ```bash
   git init -b main
   git add .
   git commit -m "GW-002: starter kit"
   ```
3. Create the GitHub repo and push: `gh repo create llm-gateway --public --source . --push`
4. Open the repo in your browser (`gh repo view --web`) and check the README renders.

## Acceptance criteria
- [ ] The repo exists on GitHub and is public.
- [ ] The README, `AGENTS.md`, `tickets/` and `.claude/` folders are visible on GitHub.

## Explain it back
1. Why is the repo public, and what should therefore never be committed to it?

## Notes for the AI
If the human prefers a private repo for now, that is fine; remind them to make it public later.
