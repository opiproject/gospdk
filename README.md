# gospdk

[![Linters](https://github.com/opiproject/gospdk/actions/workflows/linters.yml/badge.svg)](https://github.com/opiproject/gospdk/actions/workflows/linters.yml)
[![CodeQL](https://github.com/opiproject/gospdk/actions/workflows/codeql.yml/badge.svg)](https://github.com/opiproject/gospdk/actions/workflows/codeql.yml)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/opiproject/gospdk/badge)](https://securityscorecards.dev/viewer/?platform=github.com&org=opiproject&repo=gospdk)
[![Go](https://github.com/opiproject/gospdk/actions/workflows/go.yml/badge.svg)](https://github.com/opiproject/gospdk/actions/workflows/go.yml)
[![Docker](https://github.com/opiproject/gospdk/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/opiproject/gospdk/actions/workflows/docker-publish.yml)
[![License](https://img.shields.io/github/license/opiproject/gospdk?style=flat-square&color=blue&label=License)](https://github.com/opiproject/gospdk/blob/master/LICENSE)
[![codecov](https://codecov.io/gh/opiproject/gospdk/branch/main/graph/badge.svg)](https://codecov.io/gh/opiproject/gospdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/opiproject/gospdk)](https://goreportcard.com/report/github.com/opiproject/gospdk)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/opiproject/gospdk)
[![Pulls](https://img.shields.io/docker/pulls/opiproject/gospdk.svg?logo=docker&style=flat&label=Pulls)](https://hub.docker.com/r/opiproject/gospdk)
[![Last Release](https://img.shields.io/github/v/release/opiproject/gospdk?label=Latest&style=flat-square&logo=go)](https://github.com/opiproject/gospdk/releases)
[![GitHub stars](https://img.shields.io/github/stars/opiproject/gospdk.svg?style=flat-square&label=github%20stars)](https://github.com/opiproject/gospdk)
[![GitHub Contributors](https://img.shields.io/github/contributors/opiproject/gospdk.svg?style=flat-square)](https://github.com/opiproject/gospdk/graphs/contributors)

The Storage Performance Development Kit (SPDK) provides a set of tools and libraries for writing high performance, scalable, user-mode storage applications. It achieves high performance by moving all of the necessary drivers into userspace and operating in a polled mode instead of relying on interrupts, which avoids kernel context switches and eliminates interrupt handling overhead.

This project contains Go implementation of the SPDK json-rpc protocol <https://spdk.io/doc/jsonrpc.html>

## I Want To Contribute

This project welcomes contributions and suggestions.  We are happy to have the
Community involved via submission of **Issues and Pull Requests** (with
substantive content  or even just fixes). We are hoping for the documents,
test framework, etc. to become a community process with active engagement.
PRs can be reviewed by any number of people, and a maintainer may accept.

See [CONTRIBUTING](https://github.com/opiproject/opi/blob/main/CONTRIBUTING.md)
and [GitHub Basic Process](https://github.com/opiproject/opi/blob/main/doc-github-rules.md)
for more details.

## Installation

There are several ways of running this CLI.

### Docker

```sh
docker pull opiproject/gospdk:<version>
```

You can specify a version like `0.1.1` or use `latest` to get the most up-to-date version.

Run latest version of the CLI in a container:

```sh
docker run --rm opiproject/gospdk:latest --help
```

Replace `--help` with any `gospdk` command, without `gospdk` itself.

### Golang

```sh
go install -v github.com/opiproject/gospdk@latest
```

or import

```go
import (
        "github.com/opiproject/gospdk/spdk"
)
```
