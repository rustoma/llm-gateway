# llm-gateway

An OpenAI-compatible LLM gateway written in Go: one endpoint in front of every model call, with
failover, rate limits, budgets, caching, metrics and tracing.

> **Status:** early. Built ticket by ticket as a learning project. See [ROADMAP.md](ROADMAP.md).

## Why

Apps talk to the gateway instead of to providers directly. They only change their `baseURL`:

```ts
const client = new OpenAI({ baseURL: "http://localhost:8080/v1", apiKey: process.env.GATEWAY_KEY });
await client.chat.completions.create({ model: "chat", messages });  // "chat" is an alias
```

The gateway decides which provider and model serve each alias, and handles the hard parts:
retries, failover between providers, streaming, timeouts, per-key limits and budgets.

## Architecture

```
 PawCite API (Python) ─┐
 Evals, workers ───────┼──▶  llm-gateway (Go)  ──▶  OpenAI · Anthropic · vLLM · Ollama
 Any OpenAI SDK ───────┘     /v1, OpenAI-compatible
                              │
                              ├── PostgreSQL: keys, request log, cache
                              └── Prometheus metrics · OpenTelemetry traces
```

## Project context

This gateway serves **PawCite**, a dog-research assistant that answers questions from published
studies (repo: `pawcite`). It is ported from a TypeScript gateway, and a later milestone proves
both behave the same and benchmarks them.

## Working on this repo

- Start with [LEARNING-PATH.md](LEARNING-PATH.md) and the first ticket in [tickets/](tickets/).
- AI assistants follow [AGENTS.md](AGENTS.md); skills and commands are in [.claude/](.claude/).
- What I learned along the way: [docs/learning-log.md](docs/learning-log.md).
