# REASONS Canvas

## Requirements

Bound nonmatching scans by effective UNTIL while preserving all correct occurrences. FR-001/002/003 and NFR-001/002 define acceptance in spec.md. Done includes exact compatibility and defect-sensitive performance evidence, independent model review and charter quality gates.

## Entities

Effective UNTIL, recurrence period, candidate occurrence, current rule state and fresh iterator after setters. A period can contain candidates earlier than the raw cursor; safe termination must account for that distinction.

## Approach

Use UTC period lower bound with unchanged non-UTC fallback, selected by Kent; see plan.md and independent research. Reject timezone padding/name heuristics as unproved.

## Structure

One production guard, public regression tests and issue notes; no new state, APIs or dependencies. See plan.md concern map.

## Operations and Safeguards

WP01: baseline/red tests → UTC guard → exact compatibility/default/setters → quality gates → issue notes. Only rrule.go, rrule_until_test.go and docs/issues/03-until-scan.md are implementer-owned. Orchestrator owns mission artifacts and lifecycle. See tasks.md.
