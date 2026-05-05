# awssm

The `awssm` package provides `urfave/cli/v3` value sources backed by AWS Secrets Manager. It seamlessly handles both `SecretString` and `SecretBinary` return values.

## Usage

```go
import (
    "github.com/aws-contrib/aws-cli/awssm"
    "github.com/urfave/cli/v3"
)

// Fetch a single secret
&cli.StringFlag{
    Name: "api-key",
    Sources: cli.NewValueSourceChain(
        awssm.Secret("my-app-api-key"),
    ),
}
```

### Fallbacks (Multiple Secrets)

If you have multiple secrets you want to check as fallbacks, use the `awssm.Secrets` helper:

```go
&cli.StringFlag{
    Name: "api-key",
    Sources: awssm.Secrets("my-app-api-key", "legacy-app-api-key"),
}
```

### Custom AWS Configuration

You can pass standard AWS SDK functional options to the constructors to configure the underlying AWS client (e.g., setting a specific region or custom endpoint).

```go
import "github.com/aws/aws-sdk-go-v2/config"

awssm.Secret("my-secret", config.WithRegion("us-west-2"))
```
