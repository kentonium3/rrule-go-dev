# Work Packages: UTC UNTIL scan bound

Planning/base/consolidation branch: fix/until-scan. Later private PR target: master.

## WP01 — UTC period guard and compatibility evidence

P1; no dependencies. IC-01 and IC-02 are one atomic root-cause fix. Independent test: both UTC repros finish below 100ms operation time; exact baseline sequences remain unchanged everywhere.

- [ ] T001 Capture baseline and isolated public-path red tests (WP01)
- [ ] T002 Add the proved UTC-only period bound (WP01)
- [ ] T003 Verify exact compatibility and setters/default horizons (WP01)
- [ ] T004 Run quality gates and defect-sensitive controls (WP01)
- [ ] T005 Document measured results and fallback limitation (WP01)

Prompt: tasks/WP01-utc-period-bound.md. One production owner; orchestrator owns mission lifecycle/evidence and private PR. Different-model review required.
