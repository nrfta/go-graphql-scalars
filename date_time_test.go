package scalars_test

import (
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nrfta/go-graphql-scalars"
)

var _ = Describe("Marshal/ Unmarshal DateTime Test", func() {
	var (
		correctDateTime    = "2006-01-02T15:04:05Z"
		wrongDateTime      = "2019-06-2435"
		testDateTime, _    = time.Parse(time.RFC3339, correctDateTime)
		testStringBuilders = &strings.Builder{}
	)

	var _ = Context("Unmarshal DateTime test", func() {
		It("returns a time object when valid datetime string passed in to UnmarshalDateTime", func() {
			dateTime, err := scalars.UnmarshalDateTime(correctDateTime)
			Expect(err).To(Succeed())
			Expect(dateTime).Should(BeAssignableToTypeOf(time.Time{}))
			Expect(dateTime).To(Equal(testDateTime))
		})

		It("returns an error when invalid datetime format passed in to UnmarshalDateTime", func() {
			_, err := scalars.UnmarshalDateTime(wrongDateTime)
			Expect(err).ToNot(BeNil())
		})
	})

	var _ = Context("Marshal DateTime Test", func() {

		It("returns a function that writes the time from string format into a slice of bytes", func() {
			scalars.MarshalDateTime(testDateTime).MarshalGQL(testStringBuilders)
			marshaledDateTimeString := testStringBuilders.String()
			marshaledDateTimeString = (marshaledDateTimeString[1 : len(marshaledDateTimeString)-1]) // gets rid of extra quotes around Marshalled String
			Expect(marshaledDateTimeString).To(Equal(correctDateTime))
			Expect(scalars.UnmarshalDateTime(marshaledDateTimeString)).Should(BeAssignableToTypeOf(time.Time{}))
		})
	})
})
