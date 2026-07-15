# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
