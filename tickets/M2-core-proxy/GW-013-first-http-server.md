---
id: GW-013
title: Your first HTTP server
milestone: GW-M2 · Core proxy
labels: go, api
size: M (about 1.5 h)
depends_on: GW-012
---

# GW-013 · Your first HTTP server

## Goal
`GET /healthz` returns `{"status":"ok","version":"dev"}`, and every request is logged as JSON.

## Why it matters
Load balancers and Docker health checks will call `/healthz`. The logging middleware is the first
of several layers (auth, rate limits, metrics) that wrap every request.

## You will learn
- A handler: `func(w http.ResponseWriter, r *http.Request)` (≈ Express's `(req, res)`, reversed).
- `http.NewServeMux()` with method patterns: `mux.HandleFunc("GET /healthz", ...)`.
- Writing JSON with `json.NewEncoder(w).Encode(...)`.
- Middleware as a function: `func(next http.Handler) http.Handler` (≈ Express middleware).
- Structured logging with `log/slog`.

## Steps
1. Create `internal/server/server.go` with `func New(logger *slog.Logger) http.Handler`.
2. Add the `/healthz` handler.
3. Add a logging middleware that logs method, path, status and duration. (Hint: wrap the
   `ResponseWriter` to remember the status code.)
4. In `cmd/gateway/main.go`: create a JSON `slog` logger and listen on `:8080`, or on the `PORT`
   environment variable if set.
5. Test with `httptest.NewRecorder` in `internal/server/server_test.go`.
6. Try it: `go run ./cmd/gateway`, then `curl -i localhost:8080/healthz`.

## Acceptance criteria
- [ ] `curl localhost:8080/healthz` returns 200 with the JSON body.
- [ ] `POST /healthz` returns 405 (the mux does this for you).
- [ ] Each request produces one JSON log line with method, path, status and duration.
- [ ] The handler test passes.

## Explain it back
1. Why does the middleware need to wrap the `ResponseWriter` to know the status code?
2. What does `"GET /healthz"` give you that a plain `"/healthz"` pattern does not?
