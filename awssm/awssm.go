package awssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awssecretsmanager "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/urfave/cli/v3"
)

var _ cli.ValueSource = &SecretValueSource{}

// SecretValueSource retrieves values from AWS Secrets Manager.
// It implements the cli.ValueSource interface.
type SecretValueSource struct {
	SecretId string
	Options  []func(*config.LoadOptions) error
}

// Secret creates a new SecretValueSource for the given secret ID.
// Optional AWS SDK configuration options can be provided.
func Secret(secretId string, opts ...func(*config.LoadOptions) error) *SecretValueSource {
	return &SecretValueSource{
		SecretId: secretId,
		Options:  opts,
	}
}

// Secrets is a helper function to encapsulate a number of SecretValueSource
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
func (f *SecretValueSource) Lookup() (string, bool) {
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

// String returns a string representation of the SecretValueSource.
func (f *SecretValueSource) String() string {
	return fmt.Sprintf("secret %[1]q", f.SecretId)
}

// GoString returns a Go-syntax representation of the SecretValueSource.
func (f *SecretValueSource) GoString() string {
	return fmt.Sprintf("&SecretValueSource{SecretId:%[1]q}", f.SecretId)
}
