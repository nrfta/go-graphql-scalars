package scalars_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/volatiletech/null"

	"github.com/nrfta/go-graphql-scalars"
)

var _ = Describe("UnmMarshall Null String", func() {
	var (
		testFloat = 231.02
	)
	It("Should return a Null String if null passed in ", func() {
		nullFloat, _ := scalars.UnmarshalNullFloat64(nil)
		Expect(nullFloat).Should(BeAssignableToTypeOf(null.Float64{}))
		Expect(nullFloat.Valid).Should(BeEquivalentTo(false))
	})

	It("Should return a String value for valid input", func() {
		notNullFloat, _ := scalars.UnmarshalNullFloat64(testFloat)

		Expect(notNullFloat).Should(BeAssignableToTypeOf(null.Float64{}))
		Expect(notNullFloat.Float64).Should(BeEquivalentTo(testFloat))
	})
})
