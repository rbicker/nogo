# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.0] - 2021-03-07
### Added
* multi os support - all nogo files are saved and received using forward slashes.

## [0.2.1] - 2020-07-06
### Fixed
* removed log statement

## [0.2.0] - 2020-07-06
### Added
* full implementation of Readdir
### Fixed
* bugfix losing information because of decoding pointers instead of vars

## [0.1.6] - 2020-07-06
### Fixed
* fixed exported fields in file info

## [0.1.5] - 2020-07-06
### Fixed
* generator has to return pointer in get method

## [0.1.4] - 2020-07-06
### Fixed
* fixed endless loop while reading

## [0.1.3] - 2020-07-06
### Fixed
* Add method

## [0.1.2] - 2020-07-06
### Fixed
* fix field exports
* code cleanup

## [0.1.1] - 2020-07-06
### Fixed
* fixed read method by using an internal reader for the File struct

## [0.1.0] - 2020-02-04
### Added
* initial version of nogogen to generate nogo.go file from given directories
* basic functionally to use files from nogo.go using "Get" or "Dir"