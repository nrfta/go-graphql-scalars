package scalars_test

import (
	"github.com/aarondl/null/v8"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	scalars "github.com/nrfta/go-graphql-scalars"
)

var _ = Describe("UnmMarshall Null String", func() {
	var (
		testInt = 231
	)
	It("Should return a Null String if null passed in ", func() {
		nullInt, _ := scalars.UnmarshalNullInt(nil)
		Expect(nullInt).Should(BeAssignableToTypeOf(null.Int{}))
		Expect(nullInt.Valid).Should(BeEquivalentTo(false))
	})

	It("Should return a String value for valid input", func() {
		notNullInt, _ := scalars.UnmarshalNullInt(testInt)

		Expect(notNullInt).Should(BeAssignableToTypeOf(null.Int{}))
		Expect(notNullInt.Int).Should(BeEquivalentTo(testInt))
	})
})
