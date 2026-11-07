---
id: GW-004
title: Turn tickets into GitHub issues and a project board
milestone: GW-M0 · Setup
labels: setup, learning
size: S (about 45 min)
depends_on: GW-003
---

# GW-004 · Turn tickets into GitHub issues and a project board

## Goal
Every ticket is a GitHub issue with labels and a milestone, shown on a project board.

## Why it matters
Pull requests close their issue automatically ("Closes #12"), so the board shows your progress
without extra work. Employers looking at the repo see how you work.

## You will learn
How to read your first Go program: `scripts/create-issues/main.go`.

## Steps
1. Preview: `go run ./scripts/create-issues/main.go --dry-run`
2. Create everything: `go run ./scripts/create-issues/main.go`
3. Create a board: GitHub → your profile → **Projects** → **New project** → **Board**.
   Name it `PawCite learning`. Add the `llm-gateway` issues (the `pawcite` issues join later).
4. Read the script with the AI next to you. Find and explain:
   - `package main` and `func main()`
   - a `struct` and a `map`
   - `:=`
   - `if err != nil`
   - `for _, t := range tickets`
   - `fmt.Errorf("...: %w", err)`

## Acceptance criteria
- [ ] All tickets appear as issues with labels and milestones.
- [ ] The board shows them in a "Todo" column.
- [ ] Running the script again creates nothing new ("already exists").
- [ ] You can explain the six Go constructs listed in step 4.

## Explain it back
1. Why does the `gh` function return two values?
2. What does `_` mean in `for _, t := range tickets`?

## Notes for the AI
Explain the constructs one at a time with the `tutor` skill. Do not change the script.
