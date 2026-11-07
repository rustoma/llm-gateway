---
id: GW-022
title: Stream responses through the gateway
milestone: GW-M3 · Streaming and shutdown
labels: go, streaming
size: M (about 2 h)
depends_on: GW-021
---

# GW-022 · Stream responses through the gateway

## Goal
Requests with `"stream": true` are streamed from the upstream to the client event by event.

## Why it matters
This loop is where later features plug in: failover before the first token, first-token
timeouts, and counting tokens for budgets.

## You will learn
- **Goroutines**: a function running concurrently, started with `go` (cheap: thousands are fine).
- **Channels**: typed pipes between goroutines (think of a conveyor belt of tokens).
- **`select`**: wait on several channels at once, for example "next event" or "client gone".
- Rule: the **sender** closes the channel.

## Steps
1. In `forward`, detect `"stream": true`.
2. Start a goroutine that reads the upstream body line by line and sends each event on a
   channel; it closes the channel when the body ends.
3. In the handler: set `Content-Type: text/event-stream` and `Cache-Control: no-cache`, then loop
   with `select` over the events channel and `r.Context().Done()`. Write and flush each event.
4. Tests with the streaming fake: all chunks and `[DONE]` arrive in order, and the first chunk
   reaches the client before the fake has finished.

## Acceptance criteria
- [ ] Streaming tests pass with `-race`.
- [ ] The smoke script from GW-020, changed to `stream: true`, prints tokens as they arrive.

## Explain it back
1. Why use a goroutine and a channel here instead of a simple copy loop? (Hint: what will the
   loop need to wait on later besides the next event?)
2. Who closes the channel, and why not the other side?

## Notes for the AI
Explain goroutines, channels and select with the tutor skill before writing code. Use the
conveyor-belt example from the skill.
