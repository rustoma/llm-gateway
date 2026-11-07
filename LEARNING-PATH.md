# Learning path

The order to work through both repos. The same file is in `llm-gateway` and `pawcite`.
Work top to bottom; each part ends with something that works.

| Part | Tickets | Repo | You end up with |
|---|---|---|---|
| 1. Setup and Go basics | GW-001 … GW-012 | llm-gateway | Go installed, AI assistant set up, Go fundamentals, CI |
| 2. Core proxy | GW-013 … GW-020 | llm-gateway | A working gateway answering real requests |
| 3. Setup and Python basics | APP-001 … APP-012 | pawcite | Python fundamentals, workspace, checks, CI, database |
| 4. First API | APP-013 … APP-018 | pawcite | A FastAPI service answering through your gateway |
| 5. Streaming and shutdown | GW-021 … GW-025 | llm-gateway | Streaming answers, timeouts, graceful shutdown |
| 6. First mobile app | APP-019 … APP-026 | pawcite | Questions asked from your phone, end to end |
| 7. Ship gateway v0.1 | GW-026 … GW-028 | llm-gateway | Docker image, compose with a free fake model, v0.1.0 |

## After part 7

The next tickets are written when you get there (see each repo's `ROADMAP.md`). The planned order:

1. **Gateway reliability** (GW-M5): retries, failover, circuit breaker, chaos tests
2. **Papers in** (APP-M5) and **answers with citations** (APP-M6)
3. **Gateway keys, limits and budgets** (GW-M6)
4. **Paper cards and evidence badges** (APP-M7)
5. **Observability** (GW-M7), then **evals in Python** (APP-M8)
6. **Accounts and limits** (APP-M9), **conformance** (GW-M8)
7. **Research agent and digests** (APP-M10)
8. **AWS** (GW-M9, APP-M11), **Anthropic adapter** (GW-M10), **benchmark** (GW-M11)
9. **Safety** (APP-M12), **distillation** (APP-M13), **operator** (GW-M12, stretch)
