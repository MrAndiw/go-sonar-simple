[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/MrAndiw/go-sonar-simple/sonarcloud.yml?branch=master)](https://github.com/MrAndiw/go-sonar-simple/actions)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=MrAndiw_go-sonar-simple&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=MrAndiw_go-sonar-simple)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=MrAndiw_go-sonar-simple&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=MrAndiw_go-sonar-simple)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=MrAndiw_go-sonar-simple&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=MrAndiw_go-sonar-simple)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=MrAndiw_go-sonar-simple&metric=coverage)](https://sonarcloud.io/summary/new_code?id=MrAndiw_go-sonar-simple)
<br>
[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-black.svg)](https://sonarcloud.io/summary/new_code?id=MrAndiw_go-sonar-simple)

# Go Sonar Simple

This is a simple Go project that includes unit tests and static code analysis using SonarQube or SonarCloud via GitHub Actions.

## Features

- Simple Go function (`GetUser`, `GetProfile`)
- Unit tests for the function
- Continuous Integration with GitHub Actions
- Code analysis and coverage using SonarScanner

## Requirements

- [Go](https://golang.org/doc/install) (1.21 or later)
- [SonarQube](https://www.sonarqube.org/) or [SonarCloud](https://sonarcloud.io/) for code analysis
- [GitHub Actions](https://docs.github.com/en/actions) for CI/CD

## Project Structure

```plaintext
.
├── .github
│   └── workflows
│       └── sonarcloud.yml          # GitHub Action for SonarQube scanning
├── pkg
│   └── user_test.go                # pkg user
│   └── user.go                     # Unit test for user.go
├── sonar-project.properties        # Configuration for SonarScanner
├── go.mod                          # Go module file
└── go.sum                          # Go dependencies
```
