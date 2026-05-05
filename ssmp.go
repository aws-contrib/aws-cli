package awscli

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/urfave/cli/v3"
)

var _ cli.ValueSource = &ParameterValueSource{}

// ParameterValueSource retrieves values from AWS Systems Manager (SSM) Parameter Store.
// It implements the cli.ValueSource interface.
type ParameterValueSource struct {
	Name    string
	Options []LoadOptionsFunc
}

// Parameter creates a new ParameterValueSource for the given parameter name.
// Optional AWS SDK configuration options can be provided.
func Parameter(name string, opts ...LoadOptionsFunc) *ParameterValueSource {
	return &ParameterValueSource{
		Name:    name,
		Options: opts,
	}
}

// Lookup retrieves the parameter value from the SSM Parameter Store.
// It returns the parameter value and a boolean indicating whether the retrieval was successful.
func (f *ParameterValueSource) Lookup() (string, bool) {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, f.Options...)
	if err != nil {
		return "", false
	}

	output, err := ssm.NewFromConfig(cfg).GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(f.Name),
		WithDecryption: aws.Bool(true),
	})

	return aws.ToString(output.Parameter.Value), err == nil
}

// String returns a string representation of the ParameterValueSource.
func (f *ParameterValueSource) String() string {
	return fmt.Sprintf("name %[1]q", f.Name)
}

// GoString returns a Go-syntax representation of the ParameterValueSource.
func (f *ParameterValueSource) GoString() string {
	return fmt.Sprintf("&ParameterValueSource{Name:%[1]q}", f.Name)
}
