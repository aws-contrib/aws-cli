package awss3

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/urfave/cli/v3"
)

var _ cli.ValueSource = &ValueSource{}

// ValueSource retrieves values from an AWS S3 object.
// It implements the cli.ValueSource interface.
type ValueSource struct {
	Bucket  string
	Key     string
	Options []func(*config.LoadOptions) error
}

// Object creates a new ValueSource for the given bucket and key.
// Optional AWS SDK configuration options can be provided.
func Object(bucket, key string, opts ...func(*config.LoadOptions) error) *ValueSource {
	return &ValueSource{
		Bucket:  bucket,
		Key:     key,
		Options: opts,
	}
}

// Objects is a helper function to encapsulate a number of ValueSource
// together as a ValueSourceChain. It expects S3 URIs in the format s3://bucket/key.
func Objects(uris ...string) cli.ValueSourceChain {
	sources := make([]cli.ValueSource, 0, len(uris))
	for _, uri := range uris {
		u, err := url.Parse(uri)
		if err != nil || u.Scheme != "s3" {
			continue
		}
		bucket := u.Host
		key := strings.TrimPrefix(u.Path, "/")
		if bucket != "" && key != "" {
			sources = append(sources, Object(bucket, key))
		}
	}
	return cli.NewValueSourceChain(sources...)
}

// Lookup retrieves the object content from S3.
// It returns the object content as a string and a boolean indicating whether the retrieval was successful.
func (f *ValueSource) Lookup() (string, bool) {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, f.Options...)
	if err != nil {
		return "", false
	}

	output, err := awss3.NewFromConfig(cfg).GetObject(ctx, &awss3.GetObjectInput{
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

// String returns a string representation of the ValueSource.
func (f *ValueSource) String() string {
	return fmt.Sprintf("s3://%s/%s", f.Bucket, f.Key)
}

// GoString returns a Go-syntax representation of the ValueSource.
func (f *ValueSource) GoString() string {
	return fmt.Sprintf("&ValueSource{Bucket:%[1]q, Key:%[2]q}", f.Bucket, f.Key)
}
