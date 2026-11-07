---
id: GW-017
title: Errors in OpenAI format
milestone: GW-M2 · Core proxy
labels: go, api
size: S (about 1 h)
depends_on: GW-016
---

# GW-017 · Errors in OpenAI format

## Goal
Every error the gateway produces itself looks like an OpenAI error, with the right status code.

## Why it matters
OpenAI SDKs (in TypeScript and Python) parse this format. Correct status codes later decide
whether to retry (503) or not (400).

## You will learn
- One small helper used everywhere instead of ad-hoc error writing.
- Mapping Go errors to HTTP status codes with `errors.Is` and `errors.As`.

## Steps
1. Create `writeError(w, status int, code, message string)` producing
   `{"error": {"message": "...", "type": "...", "code": "..."}}`.
2. Use it for:

   | Case | Status | Code |
   |---|---|---|
   | Body is not valid JSON | 400 | `invalid_json` |
   | `model` missing | 400 | `missing_model` |
   | Unknown alias | 404 | `model_not_found` |
   | Body too large | 413 | `request_too_large` |
   | Upstream unreachable | 502 | `upstream_unavailable` |
   | Upstream timed out | 504 | `upstream_timeout` |

3. Detect the too-large body with `errors.As(err, &maxBytesErr)` (type `*http.MaxBytesError`)
   and the timeout with `errors.Is(err, context.DeadlineExceeded)`.
4. Table-driven test covering every row.

## Acceptance criteria
- [ ] Every row of the table has a passing test.
- [ ] The official OpenAI SDK (try it from TypeScript) shows the message for a 404.

## Explain it back
1. When do you use `errors.As` instead of `errors.Is`?
