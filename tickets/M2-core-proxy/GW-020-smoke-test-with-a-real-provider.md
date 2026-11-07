---
id: GW-020
title: Smoke test with a real provider
milestone: GW-M2 · Core proxy
labels: api, testing
size: S (about 45 min)
depends_on: GW-019
---

# GW-020 · Smoke test with a real provider

## Goal
The official OpenAI SDK, pointed at your gateway, gets a real answer and real embeddings.

## Why it matters
This proves the gateway is a drop-in replacement: apps only change their `baseURL`.

## Steps
1. In your OpenAI dashboard, set a monthly spending limit (for example $10) **before** you start.
2. Create `config.json` from `config.example.json` (it is git-ignored) and export `OPENAI_API_KEY`.
3. Start the gateway bound to localhost only. There is no authentication yet, so it must not be
   reachable from your network.
4. Create `examples/smoke.ts` (any folder with `npm i openai`):
   ```ts
   import OpenAI from "openai";
   const client = new OpenAI({ baseURL: "http://localhost:8080/v1", apiKey: "not-used-yet" });
   const chat = await client.chat.completions.create({
     model: "fast",
     messages: [{ role: "user", content: "In one sentence: why do dogs wag their tails?" }],
   });
   console.log(chat.choices[0].message.content);
   const emb = await client.embeddings.create({ model: "embed", input: "leash training" });
   console.log(emb.data[0].embedding.length);
   ```
5. Run it and check the gateway's logs.

## Acceptance criteria
- [ ] The script prints an answer and `1536`.
- [ ] The gateway log shows both requests with status 200.
- [ ] A spending limit is set on the provider account.

## Explain it back
1. Why must the gateway listen only on localhost for now?
