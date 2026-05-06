package awssm_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/aws-contrib/aws-cli/awssm"
)

var _ = Describe("ValueSource", func() {
	var source *awssm.ValueSource

	BeforeEach(func() {
		source = awssm.Secret("test-secret")
	})

	Context("Initialization", func() {
		It("should set the SecretId correctly", func() {
			Expect(source.SecretId).To(Equal("test-secret"))
		})
	})

	Context("String representation", func() {
		It("should return a correctly formatted String", func() {
			Expect(source.String()).To(Equal("secret \"test-secret\""))
		})

		It("should return a correctly formatted GoString", func() {
			Expect(source.GoString()).To(Equal("&ValueSource{SecretId:\"test-secret\"}"))
		})
	})

	Context("Secrets helper", func() {
		It("should create a ValueSourceChain with multiple secrets", func() {
			chain := awssm.Secrets("secret1", "secret2")
			Expect(chain).NotTo(BeNil())
			// Chain doesn't expose its length easily without reflection,
			// but we can at least verify it's not nil.
		})
	})
})
