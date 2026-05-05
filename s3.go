package awscli

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/urfave/cli/v3"
)

var _ cli.ValueSource = &S3ObjectValueSource{}

// S3ObjectValueSource retrieves values from an AWS S3 object.
// It implements the cli.ValueSource interface.
type S3ObjectValueSource struct {
	Bucket  string
	Key     string
	Options []LoadOptionsFunc
}

// S3Object creates a new S3ObjectValueSource for the given bucket and key.
// Optional AWS SDK configuration options can be provided.
func S3Object(bucket, key string, opts ...LoadOptionsFunc) *S3ObjectValueSource {
	return &S3ObjectValueSource{
		Bucket:  bucket,
		Key:     key,
		Options: opts,
	}
}

// Lookup retrieves the object content from S3.
// It returns the object content as a string and a boolean indicating whether the retrieval was successful.
func (f *S3ObjectValueSource) Lookup() (string, bool) {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, f.Options...)
	if err != nil {
		return "", false
	}

	output, err := s3.NewFromConfig(cfg).GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(f.Bucket),
		Key:    aws.String(f.Key),
	})
	if err != nil {
		return "", false
	}
	defer output.Body.Close()

	data, err := io.ReadAll(output.Body)
	if err != nil {
		return "", false
	}

	return string(data), true
}

// String returns a string representation of the S3ObjectValueSource.
func (f *S3ObjectValueSource) String() string {
	return fmt.Sprintf("s3://%s/%s", f.Bucket, f.Key)
}

// GoString returns a Go-syntax representation of the S3ObjectValueSource.
func (f *S3ObjectValueSource) GoString() string {
	return fmt.Sprintf("&S3ObjectValueSource{Bucket:%[1]q, Key:%[2]q}", f.Bucket, f.Key)
}
