# REASONS Canvas

## Requirements

Prevent unreachable sub-daily clock filters from tying up service evaluation. FR-001 through FR-003 and NFR-001/NFR-002 in spec.md define acceptance. Done requires timeout-safe red/green evidence, exact reachable controls, independent model review and the charter quality gates.

## Entities

Caller, recurrence rule, starting clock phase, interval, clock filters, occurrence sequence, and fresh iterator after setter changes. See spec.md for canonical boundaries. General calendar nonmatching is outside this mission.

## Approach

Bound the existing sub-daily search by its finite clock-state space; preserve constructor acceptance and exhaust normally. See plan.md for the modular-cycle argument, rejected alternatives and timeout-safe acceptance design.

## Structure

rrule.go owns production changes; rrule_reachability_test.go owns regression fixtures and isolation. No new public interfaces or dependencies. Existing setters rebuild RRule; fresh iterators require no new reachability cache.

## Operations and Safeguards

WP01 owns only rrule.go, rrule_reachability_test.go and issue 02 notes. Red-first tests precede bounds, compatibility checks, full validation and handoff. Orchestrator owns mission evidence and lifecycle; no overlapping source ownership. Independent model review precedes acceptance and consolidation.
