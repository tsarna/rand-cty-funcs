# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.0] - 2026-11-27

### Changed

- **Deprecated: this module has moved to
  [`github.com/tsarna/ctyx`](https://github.com/tsarna/ctyx).** The package is
  now `github.com/tsarna/ctyx/randcty`, merged there with eight sibling
  modules. Its contents and API are unchanged — only the import path moved, so
  migrating is a path rewrite plus a `require` swap.

  No code changed in this release. It exists to carry the `// Deprecated:`
  marker on `go.mod`, which is what `go list -m -u` and pkg.go.dev report. It is
  a minor rather than a patch bump because the module is no longer maintained
  here: this repository is archived, and development continues in ctyx.

## [0.2.0] - 2026-07-15

### Changed

- **BREAKING: the functions are namespaced under `rand::`.** Six of the seven already
  carried a `rand` prefix — a poor-man's namespace — so this makes it a real one: the
  leaf names drop the prefix and sort together. HCL parses `a::b(x)` natively as a single
  flat map key, so this is a naming change, not a structural one. **Existing `.vcl`/`.cty`
  files must be updated.**

  | was | is |
  | --- | --- |
  | `random` | `rand::float` |
  | `randint` | `rand::int` |
  | `randuniform` | `rand::uniform` |
  | `randgauss` | `rand::gauss` |
  | `randchoice` | `rand::choice` |
  | `randsample` | `rand::sample` |
  | `randshuffle` | `rand::shuffle` |

  `random()` becomes `rand::float` rather than `rand::random`: it returns a float in
  `[0.0, 1.0)`, and a leaf name should not repeat its namespace.

- Every function and every parameter now carries a cty `Description`. These functions have
  honest cty signatures — concrete parameter types, no variadics, describable returns — so
  they need no external declaration; the cty metadata is the whole of their documentation.

## [0.1.1] - earlier

## [0.1.0] - earlier

- Initial release: `random`, `randint`, `randuniform`, `randgauss`, `randchoice`,
  `randsample`, `randshuffle`.
