---
description: Find the next ticket whose dependencies are done
---

Find the next ticket to work on.

1. Read `LEARNING-PATH.md` for the order and `ROADMAP.md` for what is already ticked.
2. Cross-check against GitHub, because a merged pull request closes its issue **without**
   ticking the `ROADMAP.md` checkbox — the two drift:
   ```
   gh issue list --state closed --limit 100 --json number,title
   ```
   A ticket whose issue is closed but whose box is unticked is **done**. Say so, and offer to
   tick it. A ticket whose box is ticked but whose issue is open is worth flagging too.
3. Pick the first ticket that is not done and whose `depends_on` is done.
4. Show its id, title, milestone, size, labels (say clearly if it is `hands-on`) and its Goal.
5. Say in one line what it teaches and what will exist when it is finished.
6. Ask whether to start it. Do not start it or create a branch without an answer.

If the next ticket is in the `pawcite` repo (see `LEARNING-PATH.md`), say so instead of looking
for it here. Some tickets depend across repos — check `depends_on` for an `APP-` id before
assuming a gateway ticket is unblocked.
