---
name: code-review
description: Review the changes on the current branch before the pull request is opened. Use when the human says "review", or at step 4 of the ticket-workflow skill.
---

# code-review

A review here has two jobs: catch real problems, and leave the human able to explain every line.
The second job is why a "looks good" review is a failed review.

## 1. Look at the actual diff

```
git fetch origin && git diff origin/main...HEAD
```

Read the ticket as well: a change that is correct but outside the ticket is still a finding.

## 2. Check, in this order

**Correctness**
- Is every error checked, wrapped with `%w`, and handled exactly once?
- Is `ctx` passed to everything that blocks? Is a cancelled request actually cancelled upstream?
- Does every `defer cancel()` / `defer Close()` run on every path, and after the error check?
- Can a goroutine outlive its caller, block forever, or write to a closed channel?
- Nil maps, slices sharing an array after `append`, a loop variable captured by a closure.
- Does a partial failure leave a stream half-written or a response double-written?

**The gateway's contract**
- Does `/v1` stay OpenAI-compatible — status codes, error shape, streaming framing, `[DONE]`?
- Unknown fields passed through rather than dropped?
- No API key in a log line, an error message, a test fixture or a commit.

**Tests**
- Does a test fail if the fix is reverted? If not, it is not testing the fix.
- Table-driven, covering the error paths, not just the happy one. No sleeps, no real network.

**Style**
- The `go-style` skill. `gofmt -l .`, `go vet ./...` and the tests must be green — run them and
  paste the output rather than assuming.

**Understanding**
- Any line the human would struggle to explain? Call it out and explain it now, briefly.
- Anything clever that could be plainly obvious instead? Prefer obvious.

## 3. Report

Group findings as:

- **Must fix** — wrong, unsafe, or breaks the OpenAI contract.
- **Should fix** — works, but will hurt later.
- **Consider** — taste; the human decides.

Each finding: `file:line`, what is wrong, why it matters, and the smallest fix. Quote the two or
three lines in question. Say plainly when there is nothing in a category.

Then: the acceptance criteria, ticked or not, with the evidence for each.

## Rules

- No praise padding. No "great work!" — say what is good only when it is worth copying later.
- Do not fix the code during a review of a `hands-on` ticket; report and let the human fix it.
- On a non-`hands-on` ticket, ask before applying fixes, then re-run the checks.
