---
id: GW-026
title: A small, secure Docker image
milestone: GW-M4 · Ship v0.1
labels: docker
size: S (about 1 h)
depends_on: GW-025
---

# GW-026 · A small, secure Docker image

## Goal
A multi-stage Dockerfile builds a static binary into a minimal image that runs as a non-root user.

## Why it matters
This image is what runs locally in compose, in CI, and later on AWS.

## You will learn
- Go compiles to one static binary (`CGO_ENABLED=0`), so the final image needs almost nothing.
- Distroless images: no shell or package manager, smaller attack surface, CA certificates included
  (needed for HTTPS calls to providers).

## Steps
1. `Dockerfile`: build stage `golang:1.27` (`-trimpath`, `-ldflags "-s -w -X ...Version=$VERSION"`),
   final stage `gcr.io/distroless/static-debian12:nonroot` (or the newest `static-debian` tag).
2. `.dockerignore` excluding `.git`, `bin/`, `.env`, `config.local.json`.
3. Build and run: `docker build -t llm-gateway .`, then run with your config mounted and the API
   key passed as an environment variable.
4. Check the image size: `docker images llm-gateway`.
5. Optional: build for ARM too with `docker buildx build --platform linux/amd64,linux/arm64 .`
   (ARM servers are cheaper on AWS).

## Acceptance criteria
- [ ] The image runs, `/healthz` answers, and a chat request works.
- [ ] The container runs as a non-root user.
- [ ] The image is under 30 MB.

## Explain it back
1. Why does the final image not need Go installed?
2. Why are CA certificates needed even though the image has no shell?
