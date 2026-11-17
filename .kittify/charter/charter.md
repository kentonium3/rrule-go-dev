# rrule-go project charter

RFC 5545 is normative in all cases. Preserve exported API signatures; change
behavior only as required by the RFC within issue acceptance scope, or to prevent
accepted input panics or hangs. Prove other occurrences and String() output
byte-identical to v1.8.3 plus inherited contributor fix76c646d with the golden comparison. Prefer small root-cause
fixes, no new dependencies, and state findings' conditions and likelihood.
Go 1.16 remains the minimum. An exported addition requires plan justification
and a minor release. Record RFC corrections and changes in spec and release notes.

The current mission covers public kentonium3/rrule-go issues #5 (DST) and #6
(floating UNTIL), two work packages, one per issue. Contributor issue7 fix76c646d
is inherited baseline; do not assess or extend it. Kent explicitly approved
private ROption parse metadata and targetv1.9.0 on2026-10-25: signatures/keyed
callers remain, but external positional ROption literals break; document the
source break and migration to keyed literals.
Follow docs/runbooks/teambition-backlog.md through specify, plan, independent
post-plan review, tasks, analyze, implement/review, accept, consolidate and
integrated final review. Vikunja adoption is separate.

Regression tests must fail before fixes. Validate with tests, race tests, vet,
gofmt, 60 seconds of FuzzIterate and the v1.8.3 golden comparison. Each work
package requires different implementer and reviewer models; record both and
the date. Same-model checks are degraded.

The maintained library is kentonium3/rrule-go, which alone carries release tags,
issues and contributions. This development repository kentonium3/rrule-go-dev is
public since 2026-11-17 (private history is kept in a private archive) and holds
Spec Kitty records and development docs only; never commit content from private
systems or other private repositories here. Never push upstream. Public posting,
pushes to the library repo, PRs, tags and releases require Kent's approval of
exact copy. Messages require explicit authorization.
Preserve the MIT license and README modified-copy notice.

Use the autonomous D-3 fault protocol: preserve evidence before mutation,
diagnose, validate and record safe recovery. Stop on unsafe failures or genuine
scope decisions. Use scoped Conventional Commits, no issue-closing keywords,
no force-push or stash. Develop on scoped branches against main. Kent authorizes
merging the separate charter PR first, then delivery of the implementation as a
PR to main that must not be merged.

The runtime authority is charter.yaml. Supporting rules are AGENTS.md,
docs/CONTEXT.md and relevant issue briefs. Kent authorizes policy changes and
exceptions; keep this companion synchronized with the structured charter.
