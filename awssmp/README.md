# awssmp

The `awssmp` package provides `urfave/cli/v3` value sources backed by AWS Systems Manager (SSM) Parameter Store. It seamlessly handles standard, secure (automatically decrypted), and list parameters.

## Usage

```go
import (
    "github.com/aws-contrib/aws-cli/awssmp"
    "github.com/urfave/cli/v3"
)

// Fetch a single parameter
&cli.StringFlag{
    Name: "db-password",
    Sources: cli.NewValueSourceChain(
        awssmp.Parameter("/prod/db/password"),
    ),
}
```

### Fallbacks (Multiple Parameters)

If you have multiple parameters you want to check as fallbacks, use the `awssmp.Parameters` helper:

```go
&cli.StringFlag{
    Name: "db-password",
    Sources: awssmp.Parameters("/prod/db/password", "/fallback/db/password"),
}
```

### Custom AWS Configuration

You can pass standard AWS SDK functional options to the constructors to configure the underlying AWS client (e.g., setting a specific region or custom endpoint).

```go
import "github.com/aws/aws-sdk-go-v2/config"

awssmp.Parameter("/my/param", config.WithRegion("us-west-2"))
```
