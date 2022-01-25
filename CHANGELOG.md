# Changelog
All notable changes to this project will be documented in this file.

## [Unreleased] - 2022-01-25
Swagger api stuff
### Added
- /auth/signUp
### Changed
- /transactions no longer required {id} when we do 
a get request because we can get their id from their Jwt
- /transactions now included post request which will create a new transaction