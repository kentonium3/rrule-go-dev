# Discovery: UNTIL scan bounds

Kent confirmed scope on 2026-10-25, decision 01M5S0X2AW994PNVNMEXVCVENB. Baseline master 886f07adf97d59c5348bd4e64c71ca616227304b includes both prior repairs.

## Evidence

Source: docs/issues/03-until-scan.md and docs/repro/03-until-scan/main.go. DAILY BYSETPOS=2 or 2..366 cannot select a second occurrence from a one-occurrence daily set. The current generate loop compares UNTIL only when emitting a candidate; no candidate means scanning to MAXYEAR. These ordinary accepted combinations are plausible user mistakes and can consume substantial CPU when submitted repeatedly. Historical reported measurements are 128ms and 13.2s; current-baseline measurements remain to be taken.

buildRRule supplies the existing approximately 292-year default UNTIL when absent. The optimization must honor this effective value without changing it. calcDaySet selects an entire year/month for YEARLY/MONTHLY and handles cross-year weeks. A raw iterator date is not necessarily the earliest possible occurrence in its period. Therefore stopping on that raw cursor can remove correct occurrences. DST gaps/folds and timezone conversion also require a conservative bound rather than assuming wall-clock cursor order equals instant order.

## Decisions and open design questions

Confirmed: both short explicit-UNTIL fixtures must finish within 100ms operation time; preserve exact occurrences and inclusive endpoints, API/setters, prior overflow/reachability fixes. No new arbitrary caps or Vikunja changes.

Planning must select a provable lower bound on possible occurrences in the current and later periods, handle DST/date normalization safely, and leave a period eligible whenever it could contain a candidate at or before UNTIL. Benchmark and deterministic evidence should distinguish work proportional to the horizon from machine-dependent timing. Existing candidate checks remain authoritative for inclusive semantics. No algorithm is yet approved.

## Design decision pending

Independent research (gpt-6.1-sol) and review (gpt-6-astra) found that naive local-date cutoff changes existing outputs for real Sao_Paulo and Apia midnight/date gaps. See research/period-bound-design.md, research/period-bound-review.md and the scratch reproduction source. A minimal UTC-only guard is proved but narrows FR-001; broader timezone pruning needs further design and potentially revised constraints. Decision 01M5S1NKNPCYNJNX44929ZRZBQ is open. No production changes or performance claims made.

## Resolution

Kent chose option 1 on 2026-10-25; decision 01M5S1NKNPCYNJNX44929ZRZBQ resolved. Adopt the proved UTC pointer-identity fast path and unchanged fallback. Scope is explicitly narrowed in spec.md. No outstanding design questions.
