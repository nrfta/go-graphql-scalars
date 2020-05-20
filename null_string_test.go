package scalars_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/volatiletech/null/v8"

	scalars "github.com/nrfta/go-graphql-scalars"
)

var _ = Describe("UnmMarshall Null String", func() {
	var (
		testString = "notNullString"
	)
	It("Should return a Null String if null passed in ", func() {

		nullString, _ := scalars.UnmarshalNullString(nil)

		Expect(nullString).Should(BeAssignableToTypeOf(null.String{}))
		Expect(nullString.Valid).Should(BeEquivalentTo(false))
	})

	It("Should return a String value for valid input", func() {
		notNullString, _ := scalars.UnmarshalNullString(testString)

		Expect(notNullString).Should(BeAssignableToTypeOf(null.String{}))
		Expect(notNullString.String).Should(BeEquivalentTo(testString))
	})
})
