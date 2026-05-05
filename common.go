package awscli

import (
	"github.com/aws/aws-sdk-go-v2/config"
)

// LoadOptionsFunc is a functional option for configuring the AWS SDK.
type LoadOptionsFunc = func(*config.LoadOptions) error
