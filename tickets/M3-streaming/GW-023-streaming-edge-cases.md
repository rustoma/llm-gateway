---
id: GW-023
title: Streaming edge cases
milestone: GW-M3 · Streaming and shutdown
labels: go, streaming, testing
size: M (about 1.5 h)
depends_on: GW-022
---

# GW-023 · Streaming edge cases

## Goal
Streaming survives long lines, client disconnects and upstreams that end badly, with no
goroutines left running.

## Why it matters
These are the bugs that appear in production, not in demos.

## You will learn
- `bufio.Scanner` rejects lines over 64 KB by default; `scanner.Buffer` raises the limit.
- **Goroutine leaks**: a goroutine blocked forever on a send nobody will receive.
  Fix: the sender also `select`s on `ctx.Done()`.

## Steps
1. Test: a fake event larger than 64 KB (say 200 KB) passes through. Raise the scanner buffer to 1 MB.
2. Test: the client disconnects mid-stream → the upstream request is cancelled and the reader
   goroutine exits.
3. Test: the upstream stops without `[DONE]` → the gateway ends the stream cleanly.
4. Make the reader goroutine stop when the context is cancelled, even while it is trying to send.

## Acceptance criteria
- [ ] All three tests pass with `-race`.
- [ ] You can show (with the AI's help) where the goroutine would have leaked before the fix.

## Explain it back
1. What exactly is a goroutine leak, and why doesn't Go clean it up for you?
