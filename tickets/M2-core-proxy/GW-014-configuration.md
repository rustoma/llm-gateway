---
id: GW-014
title: Configuration - routes and settings
milestone: GW-M2 · Core proxy
labels: go, config
size: M (about 1.5 h)
depends_on: GW-013
---

# GW-014 · Configuration: routes and settings

## Goal
The gateway loads its routes (alias → upstream) from a JSON file, reads API keys from environment
variables, and refuses to start with a helpful message if anything is wrong.

## Why it matters
Clients ask for an **alias** (`chat`, `fast`, `embed`, `judge`); the config decides which provider
and model serve it. Changing models later becomes a config change, not a code change.

## You will learn
- Struct tags: `json:"base_url"` maps JSON fields to Go fields.
- `json.Decoder` with `DisallowUnknownFields()` to catch typos in config.
- Reporting several problems at once with `errors.Join`.

## Steps
1. Create `internal/config/config.go`:
   - `Config` with `Port` and `Routes map[string]Route`.
   - `Route` with `BaseURL`, `Model`, and `APIKeyEnv` (the **name** of the env variable holding the key).
2. `func Load(path string) (*Config, error)`: decode the file, then validate. Every route needs a
   base URL and a model, and its key's env variable must be set.
3. Create `config.example.json` with aliases `chat`, `fast`, `embed` and `judge` pointing at
   OpenAI, with `"api_key_env": "OPENAI_API_KEY"`. Use `text-embedding-3-small` for `embed`
   (1536 dimensions, as in the earlier projects) and current models from the provider's model
   list for the others. Commit it; real configs stay local.
4. `cmd/gateway`: load the file named by `GATEWAY_CONFIG` (default `config.json`); on error, log
   the problems and exit with code 1.
5. Table-driven tests: valid file, unknown field, missing model, unset key variable.

## Acceptance criteria
- [ ] A typo like `"modle"` fails with a message naming the bad field.
- [ ] Two problems in one file are both reported at once.
- [ ] No API key value appears in the config file, the logs or the tests.

## Explain it back
1. Why store the env variable **name** in the config instead of the key itself?
2. Why fail at startup instead of at the first request?
