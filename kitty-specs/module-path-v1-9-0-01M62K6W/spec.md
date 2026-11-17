# Mission Specification: rrule-go v1.9.0 module path migration

**Mission Branch**: `rrule-go-dev-xfer`  
**Created**: 2026-10-28 (America/New_York)  
**Status**: Specified; implementation and publication pending  
**Input**: Kent's confirmed request to adopt `github.com/kentonium3/rrule-go` for v1.9.0, document migration, prepare release communication and an unposted upstream proposal, and verify maintenance assumptions.

## Intent Summary

Go developers adopting or updating the separately maintained fork should use its own module identity and receive accurate migration and release information. Rename active module/import references from `github.com/teambition/rrule-go` to `github.com/kentonium3/rrule-go` for candidate v1.9.0. Preserve package name `rrule`, exported declarations, recurrence behavior, attribution, and historical release contents. Prepare communications privately without assuming official succession or upstream endorsement.

Kent confirmed scope through decision `01M62K7ADP565BH1469Q0VEJ12`. The branch starts from private `main` at `737ae01be40048d1d2932dbafb623230efa8d361`. Spec Kitty planning/consolidation targets `rrule-go-dev-xfer`; the eventual private PR targets `main`. This request specifies the mission and prepares drafts; it does not authorize a full implementation run, private PR merge, or public actions.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Adopt the fork directly (Priority: P1)

A Go developer can depend on the fork's own declared identity without replacing the original module.

**Why this priority**: Direct adoption is the purpose of the change.

**Independent Test**: A clean consumer exercises parsing and recurrence through the new import against candidate source. The same explicit module-declaration check must fail against pre-migration source and pass against candidate source; a local replacement alone can mask a mismatch. The external consumer build checks local API usability, while real public retrieval remains a later release check. Public-tag retrieval is a separate deferred release check.

**Acceptance Scenarios**:

1. **Given** candidate source with the new identity, **when** a fresh consumer imports it, **then** it compiles and evaluates a representative recurrence without an old-module replacement.
2. **Given** the candidate and pre-migration private-main baseline, **when** the existing golden corpus is evaluated, **then** occurrences and `String()` output match byte-for-byte.
3. **Given** package-qualified calls, **when** imports change, **then** the qualifier remains `rrule` and this mission introduces no exported declarations or signature changes.

### User Story 2 - Migrate an existing application (Priority: P1)

A developer using upstream directly or the transitional replacement can follow complete instructions and validate the resulting dependency graph.

**Why this priority**: A module-identity change does not automatically migrate consumers.

**Independent Test**: Follow the guide on three clean fixtures: direct upstream use, an all-version replacement, and a version-specific replacement.

**Acceptance Scenarios**:

1. **Given** each fixture, **when** the guide is followed, **then** imports use the new path, obsolete replacements are removed correctly, dependency cleanup succeeds, and application tests pass.
2. **Given** a dependency exposing old-path `RRule` or `ROption` values, **when** the guide is read, **then** it explains distinct type identities and coordinated dependency migration. Before removing a protective old-path replacement, inspect module requirements and package imports for remaining old-path consumers. Removal can silently restore original upstream code and lose the fork's panic/hang fixes even when no rrule types cross interfaces. Coordinate those consumers' migration first, or remain on the existing pinned v1.8.3 replacement arrangement. Passing application tests alone do not establish preservation of those fixes.
3. **Given** a consumer staying pinned to an existing release, **when** this work completes privately, **then** no historical tag or published release has been altered.

### User Story 3 - Review accurate release and upstream proposals (Priority: P2)

Kent can review migration communications and an upstream PR proposal without implying an agreed handoff.

**Why this priority**: The availability and relationship claims must be accurate.

**Independent Test**: Inspect the complete draft packet against the maintenance audit, migration fixtures, and publication prerequisites.

**Acceptance Scenarios**:

1. **Given** no verified public own-path release, **when** drafts are reviewed, **then** proposed release copy is clearly gated and does not advertise currently available v1.9.0 installation.
2. **Given** an eventual verified own-path release, **when** the packet is prepared for approval, **then** it includes migration copy, a consumer changelog contribution, release announcement, upstream PR title/body, proposed README notice, and exact advisory deprecation diff.
3. **Given** the upstream proposal, **when** considered by maintainers, **then** it asks their agreement, offers a README-only alternative, preserves their module identity, and leaves deprecation and any final upstream release to them.
4. **Given** notification guidance, **when** reviewed, **then** it distinguishes opt-in release notifications from query-time Go warnings and explains latest-release and unconditional-replacement limits.

### Edge Cases

- Old and new named types differ despite matching declarations; dependencies exposing these types may require coordinated updates.
- Version-specific and all-version replacement directives require different removal syntax.
- Cross-reference the separate v1.9.0 migration from positional `ROption` literals to keyed literals; this identity change does not solve it.
- Historical changelog sections, attribution, evidence, and old-state examples retain truthful old-path references.
- Tags, GitHub Releases, proxy retrieval, and documentation indexing are separate events.
- Inactivity does not establish absence of other authorized maintainers or grant authority over upstream.
- No runtime logging, network calls, notification services, or new dependencies are added.

## Requirements *(mandatory)*

### Functional Requirements

| ID | Title | User Story | Priority | Status | Delivery | No-op passable? |
|----|-------|------------|----------|--------|----------|-----------------|
| FR-001 | Own module identity | New consumers can use candidate module declaration and active imports `github.com/kentonium3/rrule-go`, retaining package name `rrule`. | High | Open | [build] | no |
| FR-002 | Classified references | Active imports, external tests, examples, applicable reproduction programs, installation links and badges use the new identity under an approved occurrence map; explicit historical/attribution exceptions retain the old identity. | High | Open | [build] | no |
| FR-003 | Migration guide | Developers receive verified steps for upstream use, both replacement forms, import changes, dependency cleanup, type-identity boundaries, retained-old-dependency safety before replacement removal, and the relevant keyed-literal migration. | High | Open | [build] | no |
| FR-004 | Release packet | Kent receives a proposed consumer Unreleased entry, release announcement excerpt, installation/migration copy, publication prerequisites and release-notification guidance consistent with selected release contents. | High | Open | [build] | no |
| FR-005 | Upstream proposal | Kent receives unposted PR title/body, README notice, exact advisory deprecation diff retaining the upstream module declaration, README-only alternative and separate upstream release prerequisites. | High | Open | [build] | no |
| FR-006 | Relationship evidence | Dated history/discussion/governance evidence identifies historical maintainer-like contributors and present-authority uncertainty; copy describes a separately maintained fork without asserting official succession. | High | Open | [build] | no |
| FR-007 | Honest notification guidance | Explain opt-in GitHub Release notifications and Go deprecation lookup, latest-release and replacement caveats, with no universal warning or automatic migration promise. | Medium | Open | [build] | no |
| FR-008 | Acceptance evidence | Record the same explicit module-declaration check failing before the change and passing after, a successful local external-consumer build, and successful candidate migration walkthroughs for all three FR-003 fixture forms. | High | Open | [build] | no |

### Non-Functional Requirements

| ID | Title | Requirement | Category | Priority | Status |
|----|-------|-------------|----------|----------|--------|
| NFR-001 | Behavior equivalence | Match the pre-migration private-main baseline on 100% of the existing golden corpus for occurrences and `String()` output; zero recurrence algorithm changes belong to this migration. | Compatibility | High | Open |
| NFR-002 | Verification quality | Pass `go test ./...`, `go test -race ./...`, `go vet ./...`, formatting with zero reported files, and the changelog check. Record identity-sensitive red/green consumer evidence. | Quality | High | Open |
| NFR-003 | Public packet integrity | Explicitly allowlist every proposed public file; zero private mission IDs, internal evidence paths or workflow records appear in candidate public content. | Release integrity | High | Open |

### Constraints

| ID | Title | Constraint | Category | Priority | Status |
|----|-------|------------|----------|----------|--------|
| C-001 | Private preparation | No public pushes, tags, releases, PRs, comments, messages or publication requests. Do not contact upstream. Future public actions require separate authorization and exact-copy approval. | Authorization | High | Open |
| C-002 | Historical releases | Do not rewrite existing releases/tags or retract working versions to advertise migration. Preserve truthful historical statements. | Compatibility | High | Open |
| C-003 | Library contract | Preserve exported declarations/signatures, package name, MIT license/attribution and Go 1.16 floor; add no dependencies. Only module/import identity changes are authorized here. Separately approved v1.9.0 changes are baseline context. | Technical | High | Open |
| C-004 | Branch/handoff | Consolidate onto `rrule-go-dev-xfer`; eventual private PR targets `main`. Follow `docs/runbooks/changelog-flow.md`; later public port uses an explicit file allowlist. Do not publish this private repository or merge the private PR here. | Workflow | High | Open |
| C-005 | Upstream authority | Proposed upstream declaration remains `github.com/teambition/rrule-go`. Original maintainers decide README/deprecation/releases. Silence and inactivity do not imply endorsement or transfer. | Relationship | High | Open |
| C-006 | Review arc | Full planning, approved occurrence classification, independent post-plan Codex review, tasks/analyze, different-model WP implementation/review and integrated post-consolidation review remain required. Workflow commands own state/worktrees; faults follow stop-and-capture rules. | Workflow | High | Open |

### Domain Language

- **Original module**: `github.com/teambition/rrule-go`, controlled by upstream.
- **Maintained fork / new module**: `github.com/kentonium3/rrule-go`, separately maintained, not an asserted official successor.
- **Development repository**: private `kentonium3/rrule-go-dev`; its name is not the consumer import path.
- **Candidate v1.9.0**: planned release content, not evidence of public availability.
- **Transition proposal**: unposted request for consideration, not a maintainer decision.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: One fresh consumer adopts the new identity and all three migration fixtures pass acceptance checks. — [build] · no-op passable: no
- **SC-002**: Zero occurrence/output differences across the existing golden corpus, paired with SC-001's successful new-identity adoption on the same candidate. — [ratchet] · no-op passable: yes; positive control SC-001
- **SC-003**: Kent receives one complete release packet and one upstream proposal packet covering all migration scenarios and publication prerequisites. — [build] · no-op passable: no
- **SC-004**: A dated maintenance audit cites history, discussions and governance clues, names historical maintainer-like contributors and records present-authority limits. — [build] · no-op passable: no

## Dependencies and Assumptions

- The baseline includes already merged DST/time-form work. Do not revert or claim to reimplement it. Behavior equivalence here is against this mission's pre-migration baseline, not the pre-DST release alone.
- v1.9.0 is Kent's chosen candidate target. Confirm actual public tag/release state before a future port; if already shipped, stop for a version/scope decision rather than rewrite it.
- Public proxy/documentation/release validation is deferred until separately authorized publication. Local validation must not request public indexing of private candidate source.
- Planning produces an eight-category `occurrence_map.yaml`, explicit exceptions and consumer validation strategy before changes. Active references are renamed; historical references and attribution are classified individually.
- Maintenance evidence and proposed copy are under `evidence/` and `drafts/`. The audit did not inspect the private author conversation or enumerate current permissions.
- Kent's confirmed new scope supersedes the charter's earlier DST/UNTIL-specific scope directive for this mission. General compatibility, review and privacy rules remain in force.
