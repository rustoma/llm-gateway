---
id: GW-024
title: Timeouts done right
milestone: GW-M3 · Streaming and shutdown
labels: go, api
size: S (about 1 h)
depends_on: GW-023
---

# GW-024 · Timeouts done right

## Goal
The server and the upstream client have timeouts that protect the gateway without killing long streams.

## Why it matters
Two classic Go traps silently cut off streams after a fixed time. You will avoid both on purpose.

## You will learn
- `http.Client.Timeout` covers reading the **whole** body, so it kills long streams. Use contexts.
- A server-wide `WriteTimeout` cuts off SSE responses. Use `ReadHeaderTimeout` and `IdleTimeout`.
- Transport timeouts: connecting, TLS handshake, and waiting for response headers.

## Steps
1. In `cmd/gateway`, use an `http.Server` with `ReadHeaderTimeout: 10s` and `IdleTimeout: 120s`,
   and no `WriteTimeout`. Add a comment explaining why.
2. Give the upstream `http.Client` a `Transport` with dial, TLS-handshake and response-header
   timeouts from config.
3. Test: a fake that sends headers but then streams slowly for longer than the response-header
   timeout still completes.
4. Test: a fake that never sends headers fails with 504 after the response-header timeout.

## Acceptance criteria
- [ ] Both tests pass.
- [ ] No `Client.Timeout` and no server `WriteTimeout` anywhere, each with a comment saying why.

## Explain it back
1. Which timeout protects you from a provider that accepts the connection but never answers?
