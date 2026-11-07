---
id: GW-010
title: Your first Go tests
milestone: GW-M1 · Go basics
labels: go, testing, learning
size: S (about 1 h)
depends_on: GW-009
---

# GW-010 · Your first Go tests

## Goal
A real gateway helper, `ParseBearer`, with a table-driven test.

## Why it matters
API keys arrive as `Authorization: Bearer <key>`. This function is used for real in a later milestone.

## You will learn
- Test files end in `_test.go`; test functions are `func TestXxx(t *testing.T)`.
- Table-driven tests: a slice of cases plus a loop with `t.Run` (≈ Jest's `test.each`).
- `go test ./...`, `-v` for details, `-cover` for coverage, `-run` to pick tests.

## Steps
1. Create `internal/auth/bearer.go` with
   `func ParseBearer(header string) (token string, ok bool)`.
2. Rules: the scheme `Bearer` is case-insensitive, extra spaces around the token are trimmed,
   an empty token or any other scheme (like `Basic`) returns `ok == false`.
3. Write `internal/auth/bearer_test.go` with at least six cases.
4. Run `go test ./... -v -cover`.

## Acceptance criteria
- [ ] All cases pass, including empty header, `Basic ...`, lowercase `bearer`, and extra spaces.
- [ ] Coverage for `internal/auth` is 100%.

## Explain it back
1. Why are table-driven tests popular in Go?
2. What does `t.Run` give you that a plain loop does not?

## Notes for the AI
Let the human write the test cases; you may write the function together.
