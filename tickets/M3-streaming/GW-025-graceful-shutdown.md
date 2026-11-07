---
id: GW-025
title: Graceful shutdown
milestone: GW-M3 · Streaming and shutdown
labels: go, api
size: M (about 1 h)
depends_on: GW-024
---

# GW-025 · Graceful shutdown

## Goal
On Ctrl+C or SIGTERM, the gateway stops accepting new requests, lets running streams finish
(up to a deadline), and exits cleanly.

## Why it matters
Every deploy stops old containers. Without this, users see answers cut off mid-sentence on every release.

## You will learn
- `signal.NotifyContext`: a context that is cancelled when the process receives a signal.
- `http.Server.Shutdown(ctx)`: closes listeners, then waits for active requests until `ctx` expires.
- Health checks while draining: `/healthz` returns 503 so a load balancer stops sending traffic.

## Steps
1. In `main`: run `ListenAndServe` in a goroutine, wait for the signal context, then call
   `Shutdown` with a 25-second deadline (configurable).
2. Set a "draining" flag on shutdown; `/healthz` returns 503 with `{"status":"draining"}`.
3. Manual test: start a slow stream through the gateway (fake with a 200 ms delay per chunk),
   press Ctrl+C, and watch the stream finish before the process exits.
4. Manual test: with a 1-second deadline, the process exits after one second even with a stream running.

## Acceptance criteria
- [ ] Both manual tests behave as described (record what you saw in the pull request).
- [ ] Logs show "shutting down", the number of seconds waited, and "stopped".

## Explain it back
1. Why does `/healthz` return 503 during shutdown?
2. What happens to requests still running when the deadline passes?
