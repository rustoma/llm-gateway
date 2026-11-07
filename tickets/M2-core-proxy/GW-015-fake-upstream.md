---
id: GW-015
title: A fake upstream for tests
milestone: GW-M2 · Core proxy
labels: go, testing
size: M (about 1.5 h)
depends_on: GW-014
---

# GW-015 · A fake upstream for tests

## Goal
A test server that pretends to be OpenAI: it answers chat and embedding requests, can be told to
fail or be slow, and records what it received.

## Why it matters
Tests must never call real providers: that would be slow, flaky and cost money. This fake will be
used by almost every later test, and later by the benchmarks.

## You will learn
- `httptest.NewServer`: a real HTTP server on a random local port, made for tests.
- Each request runs in its **own goroutine**, so shared data (the recorded requests) needs a
  `sync.Mutex`. This is your first concurrency primitive.

## Steps
1. Create `internal/fakeupstream/fakeupstream.go` with `func New() *Server`, where `Server`
   wraps `*httptest.Server` and exposes `URL()` and `Close()`.
2. `POST /chat/completions` returns a fixed OpenAI-style response whose content includes the
   requested model name. `POST /embeddings` returns one 1536-number vector per input.
3. Record every request body; expose `Requests() []map[string]any` (copy under the lock).
4. Add behaviours for tests: `FailNext(n int, status int)` and `SetDelay(d time.Duration)`.
5. Test the fake itself: a normal call, a failing call, and a recorded request.

## Acceptance criteria
- [ ] Chat and embedding responses parse as valid OpenAI-style JSON.
- [ ] `FailNext(2, 503)` makes exactly the next two calls fail.
- [ ] `go test -race ./internal/fakeupstream` passes.

## Explain it back
1. Why does recording requests need a mutex even though your test looks single-threaded?
2. Why does `Requests()` return a copy instead of the internal slice?
