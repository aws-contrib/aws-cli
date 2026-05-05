# aws-cli

> An extension for `urfave/cli/v3` that provides AWS-backed value sources for CLI flags, allowing you to automatically fetch configuration and secrets from AWS services.

[![CI](https://github.com/aws-contrib/aws-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/aws-contrib/aws-cli/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## Prerequisites

- [Go](https://go.dev/) 1.21+ (`go`)
- [urfave/cli](https://github.com/urfave/cli) v3

## Installation

```bash
go get github.com/aws-contrib/aws-cli
```

## Packages

The project is divided into modular subpackages, allowing you to only import the AWS SDK dependencies you actually need.

| Package      | Description                               | Documentation                 |
| :----------- | :---------------------------------------- | :---------------------------- |
| **`awssmp`** | AWS Systems Manager (SSM) Parameter Store | [View Docs](awssmp/README.md) |
| **`awssm`**  | AWS Secrets Manager                       | [View Docs](awssm/README.md)  |
| **`awss3`**  | AWS S3 Objects                            | [View Docs](awss3/README.md)  |

## Usage Overview

You can use the provided value sources inside a `cli.NewValueSourceChain` when defining your application's flags.

```go
import (
    "github.com/aws-contrib/aws-cli/awssmp"
    "github.com/urfave/cli/v3"
)

func main() {
    cmd := &cli.Command{
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name: "db-password",
                // Automatically fetches the parameter from AWS when parsing flags
                Sources: cli.NewValueSourceChain(
                    awssmp.Parameter("/prod/db/password"),
                ),
            },
        },
    }
}
```

For detailed examples, fallbacks, and configuring the AWS SDK client, please see the individual package documentation linked above.

## Testing

The project uses [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) for testing.

To run the tests, you can use the `go tool` command:

```bash
go tool ginkgo ./...
```

## License

[MIT](LICENSE) © 2026
