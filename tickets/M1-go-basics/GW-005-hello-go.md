---
id: GW-005
title: Hello, Go
milestone: GW-M1 · Go basics
labels: go, learning, hands-on
size: S (about 45 min)
depends_on: GW-004
---

# GW-005 · Hello, Go

## Goal
Your first Go program prints a sentence about a dog, built by a function you wrote.

## Why it matters
Variables, functions and types are the base of every later ticket.

## You will learn
- `package main` and `func main()`: the entry point (like the file Node runs first).
- `:=` (declare and assign, type inferred) vs `var` (explicit declaration).
- Static types: `string`, `int`, `bool`. A `string` can never become `undefined`.
- `fmt.Println` (≈ `console.log`) and `fmt.Sprintf` (≈ template literals).
- Capitalised names are **exported** (public); lowercase names are private to the package.

## Steps
1. Create `playground/hello/main.go`.
2. Declare `name := "Luna"` and `age := 3`.
3. Write `func describe(name string, age int) string` returning, for example,
   `Luna is 3 years old and loves research papers.`
4. Print the result from `main`.
5. Run it: `go run ./playground/hello/main.go`
6. Format it: `gofmt -w playground/` (Go has exactly one formatting style).

## Acceptance criteria
- [ ] The program prints the sentence built by `describe`.
- [ ] `gofmt -l playground/` prints nothing.

## Explain it back
1. When do you use `:=` and when `var`?
2. Why would `describe` be invisible to other packages, and how would you make it visible?

## Notes for the AI
Hands-on: the human writes the code. Explain concepts first, then hint; no full solution unless asked.
