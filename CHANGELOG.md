# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).



## [Unreleased]

## [4.0.0] - 2021-11-29

### Changed

- Drop `apiextensions` dependency.

## [3.0.2] - 2021-08-26

### Fixed

- Fix app-admission-controller webhook name in validation error matchers.

## [3.0.1] - 2021-08-19

## [3.0.0] - 2021-08-18

### Changed

- Update `giantswarm/operatorkit` from v4 to v5.

### Changed

- Update `giantswarm/app` to v4.2.0.

## [2.3.0] - 2020-12-10

### Added

- Update app library to v4.

## [2.2.0] - 2020-12-03

### Added

- Allow annotations from current app CR to remain if config has these.

## [2.1.0] - 2020-11-27

### Added

- Handle validation errors from app-admission-controller if configmap or
kubeconfig are not created yet.

## [2.0.1] - 2020-11-23

### Changed

- Update `apiextensions`, `operatorkit` dependencies.

## [2.0.0] - 2020-08-11

### Changed

- Update Kubernetes dependencies to v1.18.5.

## [0.2.0] 2020-03-26

### Changed

- Switch from dep to Go modules.
- Use architect orb.



## [0.1.0] 2020-03-19

### Added

- First release.



[Unreleased]: https://github.com/giantswarm/resource/compare/v4.0.0...HEAD
[4.0.0]: https://github.com/giantswarm/resource/compare/v3.0.2...v4.0.0
[3.0.2]: https://github.com/giantswarm/resource/compare/v3.0.1...v3.0.2
[3.0.1]: https://github.com/giantswarm/resource/compare/v3.0.0...v3.0.1
[3.0.0]: https://github.com/giantswarm/resource/compare/v2.3.0...v3.0.0
[2.3.0]: https://github.com/giantswarm/resource/compare/v2.2.0...v2.3.0
[2.2.0]: https://github.com/giantswarm/resource/compare/v2.1.0...v2.2.0
[2.1.0]: https://github.com/giantswarm/resource/compare/v2.0.1...v2.1.0
[2.0.1]: https://github.com/giantswarm/resource/compare/v2.0.0...v2.0.1
[2.0.0]: https://github.com/giantswarm/resource/compare/v0.2.0...v2.0.0
[0.2.0]: https://github.com/giantswarm/resource/compare/v0.1.0...v0.2.0

[0.1.0]: https://github.com/giantswarm/resource/releases/tag/v0.1.0
