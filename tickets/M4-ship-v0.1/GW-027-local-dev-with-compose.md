---
id: GW-027
title: Local development with docker compose and the fake upstream
milestone: GW-M4 · Ship v0.1
labels: docker, testing
size: M (about 1.5 h)
depends_on: GW-026
---

# GW-027 · Local development with docker compose and the fake upstream

## Goal
`docker compose up` starts the gateway and a standalone fake upstream, so you can develop and
demo without spending money.

## Why it matters
Free, offline development; the same fake will drive the benchmarks later.

## Steps
1. Create `cmd/fakeupstream/main.go`: a tiny program serving the `fakeupstream` handlers on a port
   (configurable chunk delay and failure rate through env variables).
2. Create `config.compose.json` routing all aliases to `http://fakeupstream:9000`.
3. Create `compose.yaml` with services `gateway` (built from the Dockerfile) and `fakeupstream`.
   Add a health check for the gateway using `/healthz`.
4. Add `make up` and `make down` targets.
5. Run the smoke script from GW-020 against the compose setup.

## Acceptance criteria
- [ ] `docker compose up` starts both services and the gateway becomes healthy.
- [ ] The smoke script works with streaming and without, with no API key and no cost.

## Explain it back
1. How does the gateway container find the fake upstream by name?
