package multiplies_test

import (
	. "testing/multiplies"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MultipliesGinkgo", func() {
	var (
		strategies []string
	)

	BeforeEach(func() {
		strategies = []string{
			"normal",
			"egyptian",
			"long",
			"russian",
		}
	})

	Describe("Multiplies should work", func() {
		It("should work for each strategy", func() {
			x := 67
			y := 89

			for _, sgy := range strategies {
				result := Multiply(x, y, sgy)
				Expect(result).To(Equal(5963))
			}
		})
	})
})
