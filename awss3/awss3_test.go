package awss3_test

import (
	"github.com/aws-contrib/aws-cli/awss3"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ValueSource", func() {
	var source *awss3.ValueSource

	BeforeEach(func() {
		source = awss3.Object("test-bucket", "test-key")
	})

	Context("Initialization", func() {
		It("should set the Bucket and Key correctly", func() {
			Expect(source.Bucket).To(Equal("test-bucket"))
			Expect(source.Key).To(Equal("test-key"))
		})
	})

	Context("String representation", func() {
		It("should return a correctly formatted String", func() {
			Expect(source.String()).To(Equal("s3://test-bucket/test-key"))
		})

		It("should return a correctly formatted GoString", func() {
			Expect(source.GoString()).To(Equal("&ValueSource{Bucket:\"test-bucket\", Key:\"test-key\"}"))
		})
	})

	Context("Objects helper", func() {
		It("should create a ValueSourceChain from S3 URIs", func() {
			chain := awss3.Objects("s3://b1/k1", "s3://b2/k2")
			Expect(chain).NotTo(BeNil())
		})

		It("should skip invalid URIs", func() {
			chain := awss3.Objects("invalid-uri")
			Expect(chain).NotTo(BeNil())
		})
	})
})
