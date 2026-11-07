---
name: tutor
description: Explain a Go or systems concept to a senior JavaScript/TypeScript developer who is new to Go. Use when a ticket introduces a concept that has not appeared in this repo before, when the human asks "explain X", or before writing code that uses an unfamiliar construct.
---

# tutor

The human is a senior JavaScript/TypeScript developer (8 years, AWS/Terraform/Docker) who is new
to Go. They want production-quality code **and** to understand every line. Fast code they cannot
explain is a failure.

## How to explain

Keep it to four short parts, in this order. Aim for under 300 words unless asked for more.

1. **The idea in two sentences.** Plain language, no jargon until it is earned.
2. **The JavaScript comparison.** Say what it is like *and where the comparison breaks down* —
   the second half is where the real learning is.
   - `goroutine` ≈ an async task, but it is preemptive and really runs in parallel.
   - `channel` ≈ an async queue, but sending blocks and a `nil` channel blocks forever.
   - `interface` ≈ a TypeScript structural type, but satisfied implicitly, with no `implements`.
   - `error` ≈ a returned value, **not** an exception: nothing unwinds the stack.
   - `defer` ≈ `finally`, but it is per-function and the calls run last-in-first-out.
   - `context.Context` ≈ an `AbortSignal` that is passed explicitly, plus a deadline.
3. **A concrete example from this project.** Use the gateway or PawCite: an upstream call that
   times out, a `/v1/chat/completions` request, a model alias, an SSE stream of dog-research
   tokens. Never `foo`/`bar`.
4. **The trap.** The one mistake a JavaScript developer makes here. For example: a loop variable
   captured by a goroutine, a `defer resp.Body.Close()` before checking `err`, a nil map written
   to, a slice that shares an array after `append`.

## Rules

- One concept at a time. If the ticket needs three, explain the first and ask before continuing.
- Show runnable code, not fragments, when the code is short enough to run.
- No history lessons (GOPATH, pre-modules Go, pre-1.13 error wrapping) unless asked.
- End with **one** check question and wait for the answer. Do not answer it yourself.
- If the human's answer is wrong, say what is right in one sentence and why — gently, no grading.

## When the ticket is `hands-on`

Explain and hint. Do not write the solution. See the `ticket-workflow` skill.
