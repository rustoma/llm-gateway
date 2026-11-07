---
id: GW-021
title: How streaming works - SSE and a streaming fake
milestone: GW-M3 · Streaming and shutdown
labels: go, streaming, learning
size: M (about 1.5 h)
depends_on: GW-020
---

# GW-021 · How streaming works: SSE and a streaming fake

## Goal
The fake upstream can stream an answer token by token, exactly like OpenAI does.

## Why it matters
Users see the answer appear word by word on their phone. Streaming is also where most gateway bugs
hide, so it needs a realistic fake first.

## You will learn
- Server-Sent Events (SSE): lines starting with `data: `, a blank line after each event, and a
  final `data: [DONE]`. You have consumed this in JavaScript with `EventSource` or `fetch` streams.
- Flushing: sending bytes now instead of buffering, with `http.NewResponseController(w).Flush()`.
- Reading line by line with `bufio.Scanner`.

## Steps
1. Ask the AI to show what a real OpenAI streaming response looks like (a few `data:` lines).
2. In `fakeupstream`: when the request has `"stream": true`, send the answer as several
   `chat.completion.chunk` events with a configurable delay between them, then `data: [DONE]`.
3. Test: read the stream with `bufio.Scanner`, check the chunks arrive in order, and check the
   first chunk arrives well before the last one (the response really is streamed).

## Acceptance criteria
- [ ] The test proves chunks arrive separately, not all at the end.
- [ ] The stream ends with `data: [DONE]`.

## Explain it back
1. What would the client see if the fake forgot to flush after each event?
