# Migration identities and invariants

No runtime data model changes. Existing recurrence objects, parser metadata and serialization stay unchanged.

| Entity | Fields / relationships | Invariant |
|---|---|---|
| Module identity | Declared path, package imports, source commit | Candidate declares github.com/kentonium3/rrule-go; package name remains rrule. |
| Historical identity | Original path, immutable published version | Historical releases/evidence retain github.com/teambition/rrule-go where truthful. |
| Consumer fixture | Initial import/require/replace form, local source path, resulting graph | New import selects candidate; obsolete old replacement removed; old-path consumers must be identified before protective replacement removal; coordinate migration or remain on existing arrangement. |
| Golden run | Source SHA/directory, corpus digest, ordered rows, outputs | Same corpus/row order before and after; zero migration differences; expected failures classified by existing corpus. |
| Communication packet | Audience, draft text, verification evidence, approval state | Private draft is not a released fact or authorization to publish. |
| Upstream proposal | Original repository/module, proposed README/deprecation, maintainer decision | Maintainers retain authority; original module declaration stays intact. |
| Vikunja handoff | Candidate commit, import/replace steps, requested integration scenarios | Prepared later; no cross-session contact or source changes without readiness and authorization. |

Consumer transitions: original dependency or transitional replacement → documented import/directive edits → local candidate validation → later separately approved real tagged-release retrieval. Old/new named Go types remain distinct; no alias bridge is introduced.

Release transitions are documentary gates, not new runtime state: private preparation → exact-copy authorization → separately executed public port → release verification. This plan executes only private preparation. No schema, database, cache or network protocol is added.
