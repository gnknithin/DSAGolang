package fibonacci

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFibonacci(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fibonacci Suite")
}

var _ = Describe("Fibonacci Tests", func() {

	Context("Fibonacci", func() {

		It("Calculate(0) should equal 0", func() {
			Expect(Calculate(0)).Should(Equal(0))
		})

		It("Calculate(5) should equal 5", func() {
			Expect(Calculate(5)).Should(Equal(5))
		})

		It("Calculate(14) should equal 377", func() {
			Expect(Calculate(14)).Should(Equal(377))
		})

		It("Calculate(15) should equal 610", func() {
			Expect(Calculate(15)).Should(Equal(610))
		})

		It("Calculate(18) should equal 2584", func() {
			Expect(Calculate(18)).Should(Equal(2584))
		})

		It("Calculate(19) should equal 4181", func() {
			Expect(Calculate(19)).Should(Equal(4181))
		})

		It("Calculate(20) should equal 6765", func() {
			Expect(Calculate(20)).Should(Equal(6765))
		})

		It("Calculate(25) should equal 75025", func() {
			Expect(Calculate(25)).Should(Equal(75025))
		})
	})
})
