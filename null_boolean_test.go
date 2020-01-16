package scalars_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/volatiletech/null"

	"github.com/nrfta/go-graphql-scalars"
)

var _ = Describe("UnmMarshall Null String", func() {

	var testBool = true
	It("Should return a null.Bool type if null input passed in ", func() {
		nullBool, _ := scalars.UnmarshalNullBoolean(nil)
		Expect(nullBool).Should(BeAssignableToTypeOf(null.Bool{}))
		Expect(nullBool.Valid).Should(BeEquivalentTo(false))
	})

	It("Should return a bool value for valid input  ", func() {
		notNullBool, _ := scalars.UnmarshalNullBoolean(testBool)
		Expect(notNullBool).Should(BeAssignableToTypeOf(null.Bool{}))
		Expect(notNullBool.Bool).Should(BeEquivalentTo(testBool))
	})
})
