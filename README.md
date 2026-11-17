> **This is a development workspace, not the library's home.** To use rrule-go, report a
> bug, ask a question or contribute, go to
> **[kentonium3/rrule-go](https://github.com/kentonium3/rrule-go)**. Issues are turned off
> here, pull requests here are not accepted, and releases and tags are published only
> there. Report security problems privately through its
> [security advisories](https://github.com/kentonium3/rrule-go/security/advisories/new).
>
> This repository is a [Spec Kitty](https://github.com/spec-kitty/spec-kitty) ("kittified")
> development environment. Alongside the same library code it carries mission records
> (`kitty-specs/`, `.kittify/`, `kitty-ops/`), development docs and agent instructions
> that library users don't need. It is Kent Gale's modified copy of
> [teambition/rrule-go](https://github.com/teambition/rrule-go) (MIT, © Teambition),
> based on `master` at `e74d163`. Agents: read [AGENTS.md](AGENTS.md) first.

# rrule-go

Go library for working with recurrence rules for calendar dates. The maintained
fork follows [RFC 5545](https://www.rfc-editor.org/rfc/rfc5545); it began as a partial
port of [python-dateutil](https://github.com/dateutil/dateutil).

[![CI](https://github.com/kentonium3/rrule-go-dev/actions/workflows/ci.yml/badge.svg)](https://github.com/kentonium3/rrule-go-dev/actions/workflows/ci.yml)
[![CodeQL](https://github.com/kentonium3/rrule-go-dev/actions/workflows/codeql.yml/badge.svg)](https://github.com/kentonium3/rrule-go-dev/actions/workflows/codeql.yml)
[![License](https://img.shields.io/badge/license-mit-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/kentonium3/rrule-go.svg)](https://pkg.go.dev/github.com/kentonium3/rrule-go)

## Installation and migration

Version v1.9.0 is [released](https://github.com/kentonium3/rrule-go/releases/tag/v1.9.0)
under module path `github.com/kentonium3/rrule-go`; the package name remains `rrule`.

```sh
go get github.com/kentonium3/rrule-go@v1.9.0
```

Existing consumers: read [MIGRATION.md](MIGRATION.md) before changing imports or
removing an old-path replacement. Coordinate remaining old-path consumers or
keep the pinned v1.8.3 arrangement until ready.

## Documentation

- [Documentation index](docs/README.md): guides by task.
- [Usage and examples](docs/usage.md): rule construction, parsing, sets and queries.
- [Time zones](docs/timezones.md): DTSTART/UNTIL forms, gaps, folds and serialization.
- [Design](docs/design.md) and [limitations](docs/limitations.md): contributor map and supported boundaries.
- [Migration](MIGRATION.md) and [changelog](CHANGELOG.md): upgrades and release history.

## Development material

Use the [development index](docs/development.md) for context, issue briefs,
runbooks, release research, reproductions and validation. The
[changelog handoff guide](docs/runbooks/changelog-flow.md) describes how changes
move from here to the library repository.

The root [CHANGELOG](CHANGELOG.md) records actual released and unreleased changes;
dated investigation records are historical evidence. Records dated before
2026-11-17 cite issue and PR numbers from this repository's earlier private
history, which is kept in a private archive; those links don't resolve publicly.
Development records, configuration and this README are not copied into the
library repository.

## Demo

The full demonstrations moved to the [usage guide](docs/usage.md#demonstrations).
Existing section links remain below for compatibility.

### rrule.RRule

[Rule construction and queries](docs/usage.md#rrulerrule).

### rrule.Set

[Rule sets and explicit dates](docs/usage.md#rruleset).

### rrule.StrToRRule

[Parsing a rule](docs/usage.md#rrulestrtorrule).

### rrule.StrToRRuleSet

[Parsing a rule set](docs/usage.md#rrulestrtorruleset).

## License

rrule-go is licensed under the [MIT](LICENSE) license.
Copyright © 2017–2023 Teambition. Original attribution is preserved.
