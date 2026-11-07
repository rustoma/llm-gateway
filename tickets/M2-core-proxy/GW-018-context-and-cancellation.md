---
id: GW-018
title: Context - timeouts and cancellation
milestone: GW-M2 · Core proxy
labels: go, api, learning
size: M (about 1.5 h)
depends_on: GW-017
---

# GW-018 · Context: timeouts and cancellation

## Goal
Upstream calls have a timeout, and a client that disconnects cancels its upstream call.

## Why it matters
If someone closes the app while a model is still writing, we stop paying for that answer.

## You will learn
- `context.Context`: carries cancellation and deadlines through every call (Node has
  `AbortSignal`, but Go uses context everywhere, by convention as the first parameter).
- `context.WithTimeout` and why you always `defer cancel()`.
- Your first `go` statement and `select` (in the test).
- A real-world gotcha: Go's HTTP/1.1 server only notices a client disconnect **after** it has read
  the request body. A fake server must read the body first, like real model servers do.

## Steps
1. Add `upstream_timeout_seconds` to the config (default 120).
2. Derive the upstream request's context from `r.Context()` with that timeout.
3. Test 1: the fake upstream is slower than a 50 ms timeout → the client gets 504.
4. Test 2: start the request in a goroutine, cancel the client's context, and check (with
   `select` and a 2-second `time.After`) that the fake upstream saw its request cancelled.
   The fake must read the body before waiting on `r.Context().Done()`.

## Acceptance criteria
- [ ] Both tests pass with `-race`.
- [ ] No test uses `time.Sleep` longer than 100 ms.

## Explain it back
1. What happens to the upstream call if you build its request with `context.Background()`?
2. Why must `cancel` always be called, even when the request succeeds?
