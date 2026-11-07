---
id: GW-008
title: Methods and interfaces
milestone: GW-M1 · Go basics
labels: go, learning, hands-on
size: M (about 1 h)
depends_on: GW-007
---

# GW-008 · Methods and interfaces

## Goal
Two different "answerers" used through one interface, the same shape as the gateway's providers.

## Why it matters
Later, OpenAI, Anthropic and a fake test server will all be used through one `Provider` interface.

## You will learn
- Methods: functions attached to a type, `func (e EchoAnswerer) Answer(...)`.
- Interfaces are satisfied **implicitly**: no `implements` keyword. This is very close to
  TypeScript's structural typing.
- Pointer vs value receivers: `func (c *Counter) Add()` can change the value; a value receiver cannot.

## Steps
1. Create `playground/answer/main.go`.
2. Define `type Answerer interface { Answer(question string) (string, error) }`.
3. Implement `EchoAnswerer` (returns "You asked: ...") and `GrumpyAnswerer` (always returns an error).
4. Write `func ask(a Answerer, q string)` that prints the answer or the error.
5. Call `ask` with both answerers.
6. Add a `CountingAnswerer` with a pointer receiver that counts how many questions it answered.

## Acceptance criteria
- [ ] `ask` works with all three answerers without knowing their concrete types.
- [ ] The counter shows the right number after three questions.

## Explain it back
1. How does Go know `EchoAnswerer` satisfies `Answerer`?
2. Why does `CountingAnswerer` need a pointer receiver?

## Notes for the AI
Hands-on. Use the TypeScript structural-typing comparison; mention "accept interfaces, return structs".
