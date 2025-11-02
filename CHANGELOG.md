# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [6.0.2] - 2025-11-02

### Changed

- Test Helpers

### Documentation

- Correct Typo

### Miscellaneous Tasks

- Upgrade CI Tool Auto Version Release

## [6.0.1] - 2025-04-10

### Added

- Chdir To Test Package

## [6.0.0] - 2025-03-22

### Added

- Utilties to the env Package
- DynamoDB Session Store
- URL To web.request
- Sss/session.go Logging
- SetSessionCookie
- Request To Web Package
- Get Cookies for Web Response To Lambda Response
- Getter For Session Time
- Session Initialization

### Changed

- Updated Depenencies
- Panic On Error Closing File
- Move Error Message
- Logger API
- Lock The Session
- Add Response Writer Assertion
- Revert Session Store
- Session Extend Time
- Session Loading
- Finish Wrapping http.Request
- Finish Wrapping http.Request
- Overwrite Session Data
- Web Response To Implement Response Writer
- Web.Response To support HTTP Headers
- Session Cookie SameSite
- Session ID Getter
- Consolidate Session Restore
- Sessions To Store Bytes
- Web Session Package
- Name of Log Package
- Session ID Cookie

### Fixed

- Session Manager Load From Cookies
- Web Request
- Use Consistent Session Expiration Time
- Web Packge Request Cookies
- Lambda Function Url Response Headers
- Sss Package Session
- Session Manager Remove Feature

### Miscellaneous Tasks

- Add return statements to Test Functions
- Update Contribution Guide
- Cleanup sss Package Messages
- Clean-up Code
- Rearrange Code
- Cleanup Comments
- Cleanup Comments

### Removed

- Map from str Package
- Web and Session To It Own Library
- Object Lock From S3 Session Save

## [5.5.4] - 2024-11-20

### Added

- Package env

## [5.5.3] - 2024-11-14

### Added

- Sha256 Functions

### Miscellaneous Tasks

- Upgrade Kohirens Version Release Orb
- Cleanup Old URL in Comment

## [5.5.2] - 2024-05-02

### Security

- Upgrade Version Release

<a name="unreleased"></a>
## [Unreleased]


<a name="5.5.1"></a>
## [5.5.1] - 2024-03-17
### Changed
- Upgrade Mongo Driver

### Fixed
- Security Vulnerability


<a name="5.5.0"></a>
## [5.5.0] - 2024-02-17
### Added
- Generic Prepend Array Function


<a name="5.4.0"></a>
## [5.4.0] - 2024-02-11
### Added
- Handler To The session Package


<a name="5.3.1"></a>
## [5.3.1] - 2024-02-06
### Changed
- Package session Cookie Expiration


<a name="5.3.0"></a>
## [5.3.0] - 2024-02-06
### Added
- Detect Expired session ID Cookie

### Fixed
- Handling Invalid Session Cookie


<a name="5.2.0"></a>
## [5.2.0] - 2024-02-04
### Added
- Restore Session From Cookie
- MongoDB Read and Update Functions


<a name="5.1.1"></a>
## [5.1.1] - 2024-02-04
### Changed
- Set Session Logger


<a name="5.1.0"></a>
## [5.1.0] - 2024-02-02
### Added
- Web session Package


<a name="5.0.0"></a>
## [5.0.0] - 2024-02-02
### Changed
- Rename Consolidated Packages

### Removed
- ReadCloser Function From io Package


<a name="4.4.0"></a>
## [4.4.0] - 2024-02-02
### Added
- 201 Responder To The web Package
- 405 Response To web Package

### Fixed
- Package web Footer Text

### Removed
- Deprecated Methods From web Package


<a name="4.3.0"></a>
## [4.3.0] - 2024-02-01
### Added
- Web Debug Responder
- Individual Web Response Redirect Methods

### Fixed
- Web 501 Reponse Function Name


<a name="4.2.1"></a>
## [4.2.1] - 2024-01-22
### Fixed
- Log Message


<a name="4.2.0"></a>
## [4.2.0] - 2024-01-21
### Added
- New Constant


<a name="4.1.0"></a>
## [4.1.0] - 2024-01-21
### Added
- Web Library
- Standard Logger To Log Package


<a name="4.0.0"></a>
## [4.0.0] - 2024-01-15
### Added
- Copy A Directory To Another Directory
- New Test Functions
- Test Features
- More Git Functions
- Panf To Log Package

### Changed
- Moved StringMap
- Function To Public
- Added path Package
- Made git Package
- Renamed kstring Package

### Fixed
- Grammer in cli Package Function

### Removed
- Sibling Package Dependency
- Removed Silencer Function
- User Methods
- Flag Types


<a name="3.2.1"></a>
## [3.2.1] - 2023-09-22
### Changed
- Usage Interface Feature For cli Package


<a name="3.2.0"></a>
## [3.2.0] - 2023-09-22
### Added
- Interface Features


<a name="3.1.1"></a>
## [3.1.1] - 2023-09-20
### Changed
- Displaying CLI Options


<a name="3.1.0"></a>
## [3.1.0] - 2023-09-18
### Added
- OS Exit To Fatal


<a name="3.0.0"></a>
## [3.0.0] - 2023-09-17
### Removed
- All Lowercase Variables From cli Usage


<a name="2.5.2"></a>
## [2.5.2] - 2023-09-17
### Fixed
- Duplicate Option Header


<a name="2.5.1"></a>
## [2.5.1] - 2023-09-16
### Changed
- Subcommand Template Optional
- Package cli Usage Output


<a name="2.5.0"></a>
## [2.5.0] - 2023-07-25
### Added
- UsageTmplVar Variable cli Package

### Removed
- Cleanup Default Usage Output


<a name="2.4.0"></a>
## [2.4.0] - 2023-07-24
### Added
- NewReadCloser To io Package


<a name="2.3.0"></a>
## [2.3.0] - 2023-07-18
### Added
- GitHub API


<a name="2.2.0"></a>
## [2.2.0] - 2023-07-17
### Added
- Package io


<a name="2.1.1"></a>
## [2.1.1] - 2023-07-12
### Changed
- Return Style For CLI RunCommand


<a name="2.1.0"></a>
## [2.1.0] - 2023-07-06
### Added
- Usage Function


<a name="2.0.0"></a>
## [2.0.0] - 2023-06-13
### Added
- Function RunCommandWithInputAndEnv

### Changed
- RunCommand In CLI Package


<a name="1.9.0"></a>
## [1.9.0] - 2023-06-08
### Added
- Run A Command With Input


<a name="1.8.0"></a>
## [1.8.0] - 2023-06-01
### Added
- Package cli


<a name="1.7.0"></a>
## [1.7.0] - 2023-05-28
### Added
- ResetDir Function To Test Package


<a name="1.6.0"></a>
## [1.6.0] - 2023-05-28
### Added
- AbsPath To Test Package


<a name="1.5.0"></a>
## [1.5.0] - 2023-05-28
### Added
- More Testing Tools


<a name="1.4.0"></a>
## [1.4.0] - 2023-05-28
### Added
- Method To Setup A Tmp Git Repository


<a name="1.3.0"></a>
## [1.3.0] - 2023-04-30
### Added
- Print Line After Error Message


<a name="1.2.0"></a>
## [1.2.0] - 2023-03-24
### Added
- Test Package


<a name="1.1.0"></a>
## [1.1.0] - 2023-02-19
### Added
- Normalize Path Function


<a name="1.0.1"></a>
## [1.0.1] - 2023-02-19
### Fixed
- File Extension Checker File Handling


<a name="1.0.0"></a>
## [1.0.0] - 2023-02-18
### Added
- kstring Package

### Changed
- Removed Obsolete Log Functions
- Moved StrToCamel Into kstring Package


<a name="0.5.0"></a>
## [0.5.0] - 2023-02-08
### Added
- CopyToDir Function


<a name="0.4.0"></a>
## [0.4.0] - 2023-02-05
### Added
- Slice Prepend Functions


<a name="0.3.0"></a>
## [0.3.0] - 2023-02-04
### Added
- Logging Methods


<a name="0.2.0"></a>
## [0.2.0] - 2023-02-03
### Added
- File Checker Extensions
- Logging Functions
- Strings To Camael Case


<a name="0.1.2"></a>
## [0.1.2] - 2022-02-12
### Changed
- Upgrade auto-relase Orb to 0.7.3.


<a name="0.1.1"></a>
## [0.1.1] - 2022-02-12
### Changed
- API documentation.
- PathExist returns false for all errors in path.


<a name="0.1.0"></a>
## [0.1.0] - 2021-06-09
### Changed
- Changed IsTextFile to new type FileExtCheck.
- Use LOCALAPPDATA on Windows.

### Fixed
- Logic in FileExtCheck.InValid.


<a name="0.0.1"></a>
## 0.0.1 - 2021-05-22
### Added
- DirExist function.
- VS Code configs.
- Readme.
- Go deps.

### Changed
- Update development environment config.

### Fixed
- Made back into a library.
- Package name.
- Development environment


[Unreleased]: https://github.com/kohirens/stdlib.git/compare/5.5.1...HEAD
[5.5.1]: https://github.com/kohirens/stdlib.git/compare/5.5.0...5.5.1
[5.5.0]: https://github.com/kohirens/stdlib.git/compare/5.4.0...5.5.0
[5.4.0]: https://github.com/kohirens/stdlib.git/compare/5.3.1...5.4.0
[5.3.1]: https://github.com/kohirens/stdlib.git/compare/5.3.0...5.3.1
[5.3.0]: https://github.com/kohirens/stdlib.git/compare/5.2.0...5.3.0
[5.2.0]: https://github.com/kohirens/stdlib.git/compare/5.1.1...5.2.0
[5.1.1]: https://github.com/kohirens/stdlib.git/compare/5.1.0...5.1.1
[5.1.0]: https://github.com/kohirens/stdlib.git/compare/5.0.0...5.1.0
[5.0.0]: https://github.com/kohirens/stdlib.git/compare/4.4.0...5.0.0
[4.4.0]: https://github.com/kohirens/stdlib.git/compare/4.3.0...4.4.0
[4.3.0]: https://github.com/kohirens/stdlib.git/compare/4.2.1...4.3.0
[4.2.1]: https://github.com/kohirens/stdlib.git/compare/4.2.0...4.2.1
[4.2.0]: https://github.com/kohirens/stdlib.git/compare/4.1.0...4.2.0
[4.1.0]: https://github.com/kohirens/stdlib.git/compare/4.0.0...4.1.0
[4.0.0]: https://github.com/kohirens/stdlib.git/compare/3.2.1...4.0.0
[3.2.1]: https://github.com/kohirens/stdlib.git/compare/3.2.0...3.2.1
[3.2.0]: https://github.com/kohirens/stdlib.git/compare/3.1.1...3.2.0
[3.1.1]: https://github.com/kohirens/stdlib.git/compare/3.1.0...3.1.1
[3.1.0]: https://github.com/kohirens/stdlib.git/compare/3.0.0...3.1.0
[3.0.0]: https://github.com/kohirens/stdlib.git/compare/2.5.2...3.0.0
[2.5.2]: https://github.com/kohirens/stdlib.git/compare/2.5.1...2.5.2
[2.5.1]: https://github.com/kohirens/stdlib.git/compare/2.5.0...2.5.1
[2.5.0]: https://github.com/kohirens/stdlib.git/compare/2.4.0...2.5.0
[2.4.0]: https://github.com/kohirens/stdlib.git/compare/2.3.0...2.4.0
[2.3.0]: https://github.com/kohirens/stdlib.git/compare/2.2.0...2.3.0
[2.2.0]: https://github.com/kohirens/stdlib.git/compare/2.1.1...2.2.0
[2.1.1]: https://github.com/kohirens/stdlib.git/compare/2.1.0...2.1.1
[2.1.0]: https://github.com/kohirens/stdlib.git/compare/2.0.0...2.1.0
[2.0.0]: https://github.com/kohirens/stdlib.git/compare/1.9.0...2.0.0
[1.9.0]: https://github.com/kohirens/stdlib.git/compare/1.8.0...1.9.0
[1.8.0]: https://github.com/kohirens/stdlib.git/compare/1.7.0...1.8.0
[1.7.0]: https://github.com/kohirens/stdlib.git/compare/1.6.0...1.7.0
[1.6.0]: https://github.com/kohirens/stdlib.git/compare/1.5.0...1.6.0
[1.5.0]: https://github.com/kohirens/stdlib.git/compare/1.4.0...1.5.0
[1.4.0]: https://github.com/kohirens/stdlib.git/compare/1.3.0...1.4.0
[1.3.0]: https://github.com/kohirens/stdlib.git/compare/1.2.0...1.3.0
[1.2.0]: https://github.com/kohirens/stdlib.git/compare/1.1.0...1.2.0
[1.1.0]: https://github.com/kohirens/stdlib.git/compare/1.0.1...1.1.0
[1.0.1]: https://github.com/kohirens/stdlib.git/compare/1.0.0...1.0.1
[1.0.0]: https://github.com/kohirens/stdlib.git/compare/0.5.0...1.0.0
[0.5.0]: https://github.com/kohirens/stdlib.git/compare/0.4.0...0.5.0
[0.4.0]: https://github.com/kohirens/stdlib.git/compare/0.3.0...0.4.0
[0.3.0]: https://github.com/kohirens/stdlib.git/compare/0.2.0...0.3.0
[0.2.0]: https://github.com/kohirens/stdlib.git/compare/0.1.2...0.2.0
[0.1.2]: https://github.com/kohirens/stdlib.git/compare/0.1.1...0.1.2
[0.1.1]: https://github.com/kohirens/stdlib.git/compare/0.1.0...0.1.1
[0.1.0]: https://github.com/kohirens/stdlib.git/compare/0.0.1...0.1.0
