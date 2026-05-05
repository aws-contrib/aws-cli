# aws-cli

> An extension for `urfave/cli/v3` that provides AWS-backed value sources for CLI flags, allowing you to automatically fetch configuration and secrets from AWS services.

[![CI](https://github.com/aws-contrib/aws-cli/actions/workflows/merge.yml/badge.svg)](https://github.com/aws-contrib/aws-cli/actions/workflows/merge.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## Prerequisites

- [Go](https://go.dev/) 1.21+ (`go`)
- [urfave/cli](https://github.com/urfave/cli) v3

## Installation

```bash
go get github.com/aws-contrib/aws-cli
```

## Usage

You can use the provided value sources inside a `cli.NewValueSourceChain` when defining your application's flags.

```go
import (
    "github.com/aws-contrib/aws-cli"
    "github.com/urfave/cli/v3"
)
```

## Systems Manager (SSM)

Fetch values from Systems Manager Parameter Store. It seamlessly handles standard, secure (automatically decrypted), and list parameters.

```go
&cli.StringFlag{
    Name: "db-password",
    Sources: cli.NewValueSourceChain(
        awscli.Parameter("/prod/db/password"),
    ),
}
```

If you have multiple parameters you want to check as fallbacks, use the `awscli.Parameters` helper:

```go
&cli.StringFlag{
    Name: "db-password",
    Sources: awscli.Parameters("/prod/db/password", "/fallback/db/password"),
}
```

## Secrets Manager

Fetch secrets directly into your CLI flags. It supports both `SecretString` and `SecretBinary` return values.

```go
&cli.StringFlag{
    Name: "api-key",
    Sources: cli.NewValueSourceChain(
        awscli.Secret("my-app-api-key"),
    ),
}
```

If you have multiple secrets you want to check as fallbacks, use the `awscli.Secrets` helper:

```go
&cli.StringFlag{
    Name: "api-key",
    Sources: awscli.Secrets("my-app-api-key", "legacy-app-api-key"),
}
```

## S3

Fetch flag values from the contents of an S3 object.

```go
&cli.StringFlag{
    Name: "config-file",
    Sources: cli.NewValueSourceChain(
        awscli.S3Object("my-bucket", "path/to/config.json"),
    ),
}
```

If you have multiple S3 objects you want to check as fallbacks, use the `awscli.S3Objects` helper using `s3://` URIs:

```go
&cli.StringFlag{
    Name: "config-file",
    Sources: awscli.S3Objects("s3://my-bucket/path/to/config.json", "s3://default-bucket/default.json"),
}
```

## Configuration

### Custom AWS Configuration

You can pass standard AWS SDK functional options to the constructors to configure the underlying AWS client (e.g., setting a specific region or custom endpoint).

```go
awscli.Parameter("/my/param", config.WithRegion("us-west-2"))
```

## License

[MIT](LICENSE) © 2026
