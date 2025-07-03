package scalars_test

import (
	"strings"
	"time"

	"github.com/aarondl/null/v8"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	scalars "github.com/nrfta/go-graphql-scalars"
)

var _ = Describe("Marshall/ UnMarshall null time", func() {

	var (
		testDateString    = "2019-01-02"
		testStringBuilder = &strings.Builder{}
		testTime, _       = time.Parse("2006-01-02", testDateString)
	)
	It("Should return of null.Time type if null value passed onto UnmarshalNullDate", func() {
		nullTime, err := scalars.UnmarshalNullDate(nil)
		Expect(err).To(Succeed())
		Expect(nullTime).Should(BeAssignableToTypeOf(null.Time{}))
		Expect(nullTime.Valid).Should(BeEquivalentTo(false))
	})

	It("returns a valid date of null.Time type for valid input", func() {
		notNullDate, _ := scalars.UnmarshalNullDate(testDateString)
		Expect(notNullDate).Should(BeAssignableToTypeOf(null.Time{}))
		Expect(notNullDate.Time).To(Equal(testTime))
	})

	It("returns a function that writes the time from string format into a slice of bytes", func() {
		scalars.MarshalDate(testTime).MarshalGQL(testStringBuilder)
		marshaledTimeString := testStringBuilder.String()
		marshaledTimeString = (marshaledTimeString[1 : len(marshaledTimeString)-1]) // gets rid of extra quotes around Marshalled String
		Expect(marshaledTimeString).To(Equal(testDateString))

		notNullDate, _ := scalars.UnmarshalNullDate(marshaledTimeString)
		Expect(notNullDate).Should(BeAssignableToTypeOf(null.Time{}))
	})
})
