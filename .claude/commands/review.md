---
description: Review the current branch before opening the pull request
---

Review the changes on the current branch using the `code-review` skill.

Diff against `origin/main`, read the ticket this branch belongs to, and report findings grouped
as **Must fix** / **Should fix** / **Consider**, each with `file:line` and the smallest fix.

Run `gofmt -l .`, `go vet ./...` and `go test ./...` (or `make check` once it exists) and paste
the output. Finish with the acceptance criteria, ticked or not, and the evidence for each.
