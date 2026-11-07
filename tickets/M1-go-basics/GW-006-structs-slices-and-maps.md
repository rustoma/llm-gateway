---
id: GW-006
title: Structs, slices and maps
milestone: GW-M1 · Go basics
labels: go, learning, hands-on
size: M (about 1.5 h)
depends_on: GW-005
---

# GW-006 · Structs, slices and maps

## Goal
A small program that stores dogs in a slice, counts breeds in a map and finds the oldest dog.

## Why it matters
The gateway is full of these: a request is a struct, routes are a map, upstreams are a slice.

## You will learn
- `struct`: a typed group of fields (≈ a TypeScript `type` describing an object).
- Slice `[]Dog` (≈ an array) and `append`.
- Map `map[string]int` (≈ `Map` or `Record`). Map order is random: sort keys to print in order.
- Zero values: an `int` starts at `0`, a `string` at `""`. Nothing is ever `undefined`.
- Returning two values: `(Dog, bool)` means "result, and whether one was found".

## Steps
1. Create `playground/dogs/main.go`.
2. Define `type Dog struct { Name string; Breed string; AgeYears int }`.
3. Make a slice with at least five dogs, two of the same breed.
4. Count dogs per breed in a `map[string]int` and print the counts sorted by breed name
   (use `slices.Sorted(maps.Keys(counts))`).
5. Write `func oldest(dogs []Dog) (Dog, bool)`; it returns `false` for an empty slice.
6. Print the oldest dog, and also handle the empty case.

## Acceptance criteria
- [ ] Breed counts print in alphabetical order.
- [ ] `oldest` returns `false` for an empty slice and the program handles it.
- [ ] `gofmt -l playground/` prints nothing.

## Explain it back
1. Why does `oldest` return a `bool` instead of `nil` for "no dog"?
2. Why must you sort the map keys before printing?

## Notes for the AI
Hands-on. Relate every concept to TypeScript; point out where the analogy breaks (zero values).
