# Work Packages: Unreachable clock-filter termination

Planning/base/consolidation branch: fix/unreachable-byset-hang. Later private PR target: master.

## WP01 — Bounded clock search and regression evidence

Priority P1. Dependencies: none. Covers IC-01 and IC-02 together because code, tests and issue disposition form one small reviewable fix. Independent test: both constructors and public queries terminate promptly for unreachable fixtures while exact reachable sequences remain identical.

T001 Capture baseline and add isolated red-first regressions for three frequency families (WP01)
T002 Bound the three clock searches while preserving overflow checks and final-candidate matches (WP01)
T003 Verify setter transitions, combined/mixed filters and exact reachable controls (WP01)
T004 Run full validation and frequency-by-frequency negative controls (WP01)
T005 Update issue notes and hand off reproducible evidence for independent review (WP01)

Prompt: tasks/WP01-bounded-clock-search.md. One production owner; no parallel source edits. Orchestrator persists /tmp evidence under the mission, owns review lifecycle and creates the private PR after acceptance/consolidation. Different implementation/review models are mandatory.
