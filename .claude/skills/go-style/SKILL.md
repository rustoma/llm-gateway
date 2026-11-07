---
name: go-style
description: House rules for writing or changing Go in this repo. Read before any Go code is written or edited.
---

# go-style

Go 1.27, **standard library first**. A dependency is added only when a ticket introduces it
(later: pgx, the Prometheus client, OpenTelemetry, a Redis client). Ask before adding anything
else.

## Errors

- Errors are **values**. Return them; do not panic. No `panic` in a request path, ever.
- Wrap with context and `%w` so callers can still `errors.Is` / `errors.As`:
  ```go
  if err != nil {
      return fmt.Errorf("calling upstream %s: %w", name, err)
  }
  ```
- The message says what failed, lowercase, no trailing punctuation, no "error:" prefix — the
  caller adds its own layer.
- Sentinel errors are `var ErrNoSuchAlias = errors.New("no such alias")`; compare with
  `errors.Is`, never by string.
- Handle an error once: either wrap and return it, or log it. Not both.

## Context

- `ctx context.Context` is the **first parameter** of anything that blocks (I/O, a request, a
  wait). Never store it in a struct.
- Pass the request's context down to the upstream call so a cancelled client cancels the call.
- Timeouts belong to the caller: `ctx, cancel := context.WithTimeout(ctx, d)` then
  `defer cancel()` — always, even on the success path.

## defer

- `defer` runs when the **function** returns, not at the end of the block. Not inside a loop
  unless you mean it; pull the body into its own function instead.
- Check the error *before* deferring the close:
  ```go
  resp, err := client.Do(req)
  if err != nil {
      return fmt.Errorf("posting to upstream: %w", err)
  }
  defer resp.Body.Close()
  ```

## Interfaces and types

- Small interfaces, **defined where they are used** (the consumer), not next to the
  implementation. One or two methods is normal.
- Accept interfaces, return concrete types.
- No global mutable state. Dependencies go in a struct and are passed in; that is what makes the
  tests easy.
- Exported identifiers (capitalised) need a doc comment starting with the name.

## Concurrency

- Every goroutine has a clear owner and a clear way to stop: a context, a closed channel, or a
  `sync.WaitGroup` the caller waits on. A goroutine with no exit is a leak.
- Do not leak goroutines on an early return — this is the main streaming bug (GW-023).
- `select` with `case <-ctx.Done():` is how a blocked goroutine gives up.
- Prefer a channel for handing work over; prefer a `sync.Mutex` for guarding a field.

## Tests

- **Table-driven**, with `t.Run(tc.name, ...)` subtests:
  ```go
  tests := []struct {
      name    string
      header  string
      want    string
      wantErr error
  }{ ... }
  ```
- `t.Parallel()` where the test allows it. `t.Cleanup` over manual teardown.
- Black-box tests that exercise the HTTP surface go in `test/`; unit tests sit beside the code.
- `httptest.NewServer` for a fake upstream. No network in tests, no sleeping to synchronise.

## Layout

```
cmd/        programs (cmd/gateway)
internal/   packages, importable only by this module
test/       black-box tests
scripts/    helper programs
```

## Naming and shape

- Short names in short scopes (`r`, `w`, `ctx`, `err`); full words for package-level names.
- Package names are short, lowercase, no underscores, and never `util` or `common`.
- Handle the error and return early; keep the happy path at the left margin.
- `gofmt` decides formatting. There is no debate and no second style.

## Secrets

API keys come from environment variables only. Never commit one, never log one, never put one in
an error message.

## Before you say it is done

`gofmt -l .` prints nothing · `go vet ./...` passes · `go test ./...` passes
(from GW-011: `make check`). Show the output.
