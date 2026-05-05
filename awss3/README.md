# awss3

The `awss3` package provides `urfave/cli/v3` value sources backed by AWS S3. It fetches the contents of an S3 object to populate your CLI flags.

## Usage

```go
import (
    "github.com/aws-contrib/aws-cli/awss3"
    "github.com/urfave/cli/v3"
)

// Fetch a single S3 object
&cli.StringFlag{
    Name: "config-file",
    Sources: cli.NewValueSourceChain(
        awss3.Object("my-bucket", "path/to/config.json"),
    ),
}
```

### Fallbacks (Multiple S3 Objects)

If you have multiple S3 objects you want to check as fallbacks, use the `awss3.Objects` helper using standard `s3://` URIs:

```go
&cli.StringFlag{
    Name: "config-file",
    Sources: awss3.Objects("s3://my-bucket/path/to/config.json", "s3://default-bucket/default.json"),
}
```

### Custom AWS Configuration

You can pass standard AWS SDK functional options to the constructors to configure the underlying AWS client (e.g., setting a specific region or custom endpoint).

```go
import "github.com/aws/aws-sdk-go-v2/config"

awss3.Object("my-bucket", "path/to/config.json", config.WithRegion("us-west-2"))
```
