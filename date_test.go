package scalars_test

import (
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/nrfta/go-graphql-scalars"
)

type sampleTime struct {
	TimeString string
}

var (
	correctTime = "2019-06-24"
	wrongTime   = "2019-06-2435"
)

var _ = Describe("Unmarshal Date Test", func() {
	It("returns a time object when proper time string passed in", func() {
		Expect(scalars.UnmarshalDate(correctTime)).Should(BeAssignableToTypeOf(time.Time{}))
	})
	It("returns an error when improper formatted time passed in", func() {
		_, err := scalars.UnmarshalDate(wrongTime)
		Expect(err).ToNot(BeNil())
	})
})

var (
	testTimeString    = "2019-01-02"
	testTime, _       = time.Parse("2006-01-02", testTimeString)
	testStringBuilder = &strings.Builder{}
)

var _ = Describe("Marshal Date Test", func() {

	It("returns a function that writes the time from string format into a slice of bytes", func() {
		scalars.MarshalDate(testTime).MarshalGQL(testStringBuilder)
		marshaledTimeString := testStringBuilder.String()
		marshaledTimeString = (marshaledTimeString[1 : len(marshaledTimeString)-1]) // gets rid of extra quotes around Marshalled String

		Expect(marshaledTimeString).To(Equal(testTimeString))
		Expect(scalars.UnmarshalDate(correctTime)).Should(BeAssignableToTypeOf(time.Time{}))
	})
})
