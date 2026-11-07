---
id: GW-019
title: Proxy embeddings
milestone: GW-M2 · Core proxy
labels: go, api
size: S (about 1 h)
depends_on: GW-018
---

# GW-019 · Proxy embeddings

## Goal
`POST /v1/embeddings` works like chat, using the same forwarding code.

## Why it matters
The app turns every paper chunk and every question into a vector through this endpoint.

## You will learn
Refactoring in Go: extract the shared forwarding logic into one function that both handlers call.

## Steps
1. Extract `forward(w, r, upstreamPath string)` from the chat handler.
2. Add the embeddings handler using it.
3. Tests: the embedding alias is rewritten, the fake returns one vector per input, and all
   earlier chat tests still pass.

## Acceptance criteria
- [ ] Chat and embeddings share one forwarding function.
- [ ] All tests pass.

## Explain it back
1. What made this refactor safe to do?
