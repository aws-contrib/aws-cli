package awssmp_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSMP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AWS SMP Suite")
}
