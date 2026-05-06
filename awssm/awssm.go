package awssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awssecretsmanager "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/urfave/cli/v3"
)

var _ cli.ValueSource = &ValueSource{}

// ValueSource retrieves values from AWS Secrets Manager.
// It implements the cli.ValueSource interface.
type ValueSource struct {
	SecretId string
	Options  []func(*config.LoadOptions) error
}

// Secret creates a new ValueSource for the given secret ID.
// Optional AWS SDK configuration options can be provided.
func Secret(secretId string, opts ...func(*config.LoadOptions) error) *ValueSource {
	return &ValueSource{
		SecretId: secretId,
		Options:  opts,
	}
}

// Secrets is a helper function to encapsulate a number of ValueSource
// together as a ValueSourceChain.
func Secrets(secretIds ...string) cli.ValueSourceChain {
	sources := make([]cli.ValueSource, len(secretIds))
	for i, secretId := range secretIds {
		sources[i] = Secret(secretId)
	}
	return cli.NewValueSourceChain(sources...)
}

// Lookup retrieves the secret value from AWS Secrets Manager.
// It returns the secret value and a boolean indicating whether the retrieval was successful.
func (f *ValueSource) Lookup() (string, bool) {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, f.Options...)
	if err != nil {
		return "", false
	}

	output, err := awssecretsmanager.NewFromConfig(cfg).GetSecretValue(ctx, &awssecretsmanager.GetSecretValueInput{
		SecretId: aws.String(f.SecretId),
	})
	if err != nil {
		return "", false
	}

	if output.SecretString != nil {
		return aws.ToString(output.SecretString), true
	} else if output.SecretBinary != nil {
		return string(output.SecretBinary), true
	}

	return "", false
}

// String returns a string representation of the ValueSource.
func (f *ValueSource) String() string {
	return fmt.Sprintf("secret %[1]q", f.SecretId)
}

// GoString returns a Go-syntax representation of the ValueSource.
func (f *ValueSource) GoString() string {
	return fmt.Sprintf("&ValueSource{SecretId:%[1]q}", f.SecretId)
}
