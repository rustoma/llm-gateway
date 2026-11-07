---
id: GW-028
title: Release v0.1.0
milestone: GW-M4 · Ship v0.1
labels: ci, docker
size: M (about 1.5 h)
depends_on: GW-027
---

# GW-028 · Release v0.1.0

## Goal
Version 0.1.0 is tagged, documented, and its Docker image is published automatically to GitHub's
container registry.

## Why it matters
The PawCite app will use a specific gateway version (`ghcr.io/<you>/llm-gateway:0.1.0`), exactly
like platform teams publish services for other teams.

## Steps
1. README: what works today, how to run it (compose and real provider), config format, and the
   endpoints, with `curl` examples.
2. `CHANGELOG.md` with a 0.1.0 entry.
3. `.github/workflows/release.yml`: on tags `v*`, build the image (with the version injected) and
   push it to `ghcr.io` (needs `permissions: packages: write`).
4. Tag and push: `git tag v0.1.0 && git push origin v0.1.0`. Check the GitHub release and the package.
5. Add a "Milestone summary" section to `docs/learning-log.md`: the five most important things you
   learned in Go so far.

## Acceptance criteria
- [ ] `docker pull ghcr.io/<you>/llm-gateway:0.1.0` works.
- [ ] The README lets a stranger run the gateway in under five minutes.

## Explain it back
1. Why publish versioned images instead of always using `latest`?
