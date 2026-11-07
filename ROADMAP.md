# Roadmap

Tick a ticket when its pull request is merged. Later milestones list planned tickets; their
detailed files are written when you reach them.

## GW-M0 · Setup
- [ ] GW-001 Install Go and your tools
- [x] GW-002 Create the repository on GitHub
- [x] GW-003 Set up your AI coding assistant
- [ ] GW-004 Turn tickets into GitHub issues and a project board

## GW-M1 · Go basics
- [ ] GW-005 Hello, Go *(hands-on)*
- [ ] GW-006 Structs, slices and maps *(hands-on)*
- [ ] GW-007 Errors and defer *(hands-on)*
- [ ] GW-008 Methods and interfaces *(hands-on)*
- [ ] GW-009 Create the Go module and project layout
- [ ] GW-010 Your first Go tests
- [ ] GW-011 Format, vet, lint, and a Makefile
- [ ] GW-012 CI with GitHub Actions

## GW-M2 · Core proxy
- [ ] GW-013 Your first HTTP server
- [ ] GW-014 Configuration: routes and settings
- [ ] GW-015 A fake upstream for tests
- [ ] GW-016 Proxy chat completions
- [ ] GW-017 Errors in OpenAI format
- [ ] GW-018 Context: timeouts and cancellation
- [ ] GW-019 Proxy embeddings
- [ ] GW-020 Smoke test with a real provider

## GW-M3 · Streaming and shutdown
- [ ] GW-021 How streaming works: SSE and a streaming fake
- [ ] GW-022 Stream responses through the gateway
- [ ] GW-023 Streaming edge cases
- [ ] GW-024 Timeouts done right
- [ ] GW-025 Graceful shutdown

## GW-M4 · Ship v0.1
- [ ] GW-026 A small, secure Docker image
- [ ] GW-027 Local development with docker compose and the fake upstream
- [ ] GW-028 Release v0.1.0

---

## Planned (tickets written later)

### GW-M5 · Reliability
Error classification (retry? fail over?) · retries with backoff and jitter · fallback across
upstreams with alias rules (judge: one upstream only; embed: same model only) · circuit breaker
with a fake clock · streaming commit point and first-token timeout · chaos testing ·
a `dev` alias routed to local models with Ollama

### GW-M6 · Keys, limits and budgets
PostgreSQL with pgx and migrations · API keys (hashed) · per-key token-bucket rate limits ·
prices and cost per request · daily budgets · asynchronous request log · judge response cache ·
per-text embedding cache · `gatewayctl` admin CLI

### GW-M7 · Observability
Prometheus metrics (requests, latency, time to first token, attempts, breaker state, tokens,
cost) · OpenTelemetry tracing with one span per upstream attempt · Grafana and Jaeger in compose ·
dashboards in the README

### GW-M8 · Conformance
The TypeScript gateway as a frozen reference image · black-box conformance suite · run it
against both gateways in CI · switch PawCite to the Go gateway

### GW-M9 · AWS
Terraform module (ECR, ECS Fargate, ALB, RDS, Secrets Manager, CloudWatch) · GitHub Actions
with OIDC · deploy-and-destroy workflow · Redis-backed rate limits for several instances

### GW-M10 · Native Anthropic adapter
Messages API request translation · SSE event translation · structured outputs · conformance tests

### GW-M11 · Benchmark
Fixed-latency fake upstreams · load generator · Go vs TypeScript: added latency, time to first
token, throughput, memory, CPU · `BENCHMARKS.md`

### GW-M12 · Kubernetes operator (stretch)
Config hot-reload · CRDs for routes and API keys · controller with kubebuilder · local cluster with kind
