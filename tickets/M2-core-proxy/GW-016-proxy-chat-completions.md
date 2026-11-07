---
id: GW-016
title: Proxy chat completions
milestone: GW-M2 · Core proxy
labels: go, api
size: M (about 2 h)
depends_on: GW-015
---

# GW-016 · Proxy chat completions

## Goal
`POST /v1/chat/completions` forwards the request to the upstream for its alias, replacing only the
`model` field, and returns the upstream's response unchanged.

## Why it matters
This is the heart of the gateway. Every later feature (failover, budgets, caching) wraps this path.

## You will learn
- `json.RawMessage`: keep JSON fields as raw bytes so unknown fields (`tools`,
  `response_format`, new OpenAI options) pass through untouched.
- `http.NewRequestWithContext`, one shared `http.Client` (it pools connections).
- `defer resp.Body.Close()` and `io.Copy` to stream the response back.

## Steps
1. Create `internal/proxy/proxy.go` with a `Proxy` type that holds the routes and an `*http.Client`.
2. Handler: read the body (`http.MaxBytesReader`, 10 MB), decode into
   `map[string]json.RawMessage`, look up the alias from `model`, replace `model` with the
   upstream's model, re-encode, and forward to `<base_url>/chat/completions` with
   `Authorization: Bearer <key>`.
3. Copy the upstream status code, `Content-Type` and body back to the client.
4. Register the route in the server.
5. Tests with `fakeupstream`:
   - the model is rewritten and other fields (`temperature`, `tools`) arrive unchanged,
   - the upstream's auth header is the configured key,
   - an upstream 400 reaches the client as a 400 with the same body.

## Acceptance criteria
- [ ] All three tests pass with `-race`.
- [ ] A request with fields the gateway does not know about still works.

## Explain it back
1. What would break if you decoded into a full struct instead of `map[string]json.RawMessage`?
2. Why share one `http.Client` instead of creating one per request?

## Notes for the AI
If the human shares the Project 4 TypeScript gateway (`src/gateway/app.ts`), use it as the reference
design. Keep this ticket to non-streaming requests; streaming comes in GW-022.
