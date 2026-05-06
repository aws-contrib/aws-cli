package awssmp_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/aws-contrib/aws-cli/awssmp"
)

var _ = Describe("ValueSource", func() {
	var source *awssmp.ValueSource

	BeforeEach(func() {
		source = awssmp.Parameter("test-parameter")
	})

	Context("Initialization", func() {
		It("should set the Name correctly", func() {
			Expect(source.Name).To(Equal("test-parameter"))
		})
	})

	Context("String representation", func() {
		It("should return a correctly formatted String", func() {
			Expect(source.String()).To(Equal("name \"test-parameter\""))
		})

		It("should return a correctly formatted GoString", func() {
			Expect(source.GoString()).To(Equal("&ValueSource{Name:\"test-parameter\"}"))
		})
	})

	Context("Parameters helper", func() {
		It("should create a ValueSourceChain with multiple parameters", func() {
			chain := awssmp.Parameters("param1", "param2")
			Expect(chain).NotTo(BeNil())
		})
	})
})
