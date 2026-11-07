---
id: GW-007
title: Errors and defer
milestone: GW-M1 · Go basics
labels: go, learning, hands-on
size: M (about 1 h)
depends_on: GW-006
---

# GW-007 · Errors and defer

## Goal
Parse dog ages from text with proper Go error handling, and clean up a file with `defer`.

## Why it matters
The gateway's reliability (retries, failover, clear error messages) is built on error handling.

## You will learn
- Errors are **values**: functions return `(result, error)`; there is no `throw`.
- `if err != nil { return ... }`: handle errors where they happen.
- Wrapping: `fmt.Errorf("parse age %q: %w", s, err)` adds context and keeps the cause.
- Sentinel errors: `var ErrNegativeAge = errors.New("age cannot be negative")`, checked with
  `errors.Is(err, ErrNegativeAge)` (never by comparing strings).
- `defer` runs when the function returns (≈ `finally`, but written next to the resource).

## Steps
1. Create `playground/ages/main.go`.
2. Write `func parseAge(s string) (int, error)` using `strconv.Atoi`. Wrap its error with `%w`.
   Return `ErrNegativeAge` for negative numbers.
3. Parse `"3"`, `"-1"` and `"three"`; print a friendly message for each error, using
   `errors.Is` to recognise the negative case.
4. Create `playground/ages/dogs.txt` with one age per line. Open it with `os.Open`,
   `defer f.Close()`, and read lines with `bufio.Scanner`.

## Acceptance criteria
- [ ] Each of the three inputs prints a different, clear result.
- [ ] The negative case is detected with `errors.Is`, not string matching.
- [ ] The file is closed with `defer` right after opening it.

## Explain it back
1. What does `%w` do that `%v` does not?
2. When exactly does a deferred call run?

## Notes for the AI
Hands-on. Compare with try/catch/finally in JavaScript, and say why Go chose explicit errors.
