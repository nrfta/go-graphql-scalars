package scalars_test

import (
	"strings"
	"time"

	"github.com/aarondl/null/v8"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	scalars "github.com/nrfta/go-graphql-scalars"
)

var _ = Describe("Marshal/ Unmarshal null DateTime Test", func() {
	var (
		correctDateTime   = "2006-01-02T15:04:05Z"
		testDateTime, _   = time.Parse(time.RFC3339, correctDateTime)
		testStringBuilder = &strings.Builder{}
	)

	It("Should return value of null.Time type if null passed onto UnmarshalNullDate", func() {
		nullTime, _ := scalars.UnmarshalNullDateTime(nil)
		Expect(nullTime).Should(BeAssignableToTypeOf(null.Time{}))
		Expect(nullTime.Valid).Should(BeEquivalentTo(false))
	})

	It("returns datetime of null.Time type for valid input", func() {
		notNullDate, err := scalars.UnmarshalNullDateTime(correctDateTime)
		Expect(err).To(Succeed())
		Expect(notNullDate).Should(BeAssignableToTypeOf(null.Time{}))
		Expect(notNullDate.Time).To(Equal(testDateTime))
	})

	It("returns a function that writes the time from string format into a slice of bytes", func() {
		scalars.MarshalNullDateTime(null.TimeFrom(testDateTime)).MarshalGQL(testStringBuilder)
		marshaledDateTimeString := testStringBuilder.String()
		marshaledDateTimeString = (marshaledDateTimeString[1 : len(marshaledDateTimeString)-1]) // gets rid of extra quotes around Marshalled String
		Expect(marshaledDateTimeString).To(Equal(correctDateTime))

		notNullDate, _ := scalars.UnmarshalNullDateTime(marshaledDateTimeString)
		Expect(notNullDate).Should(BeAssignableToTypeOf(null.Time{}))
	})
})
