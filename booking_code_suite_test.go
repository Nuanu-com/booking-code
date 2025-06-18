package booking_code_test

import (
	"testing"
	"time"

	booking_code "github.com/Nuanu-com/booking-code"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBookingCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BookingCode Suite")
}

var _ = Describe("QuotientAndReminder", func() {
	It("returns the reminder and the quotient from the base number", func() {
		q, r := booking_code.QuotientAndReminder(26)

		Expect(q).To(Equal(1))
		Expect(r).To(Equal(0))

		q, r = booking_code.QuotientAndReminder(25)

		Expect(q).To(Equal(0))
		Expect(r).To(Equal(25))

		q, r = booking_code.QuotientAndReminder(28)

		Expect(q).To(Equal(1))
		Expect(r).To(Equal(2))
	})
})

var _ = Describe("Hash", func() {
	It("Hash the number", func() {
		Expect(booking_code.Hash(0)).To(Equal("M"))
		Expect(booking_code.Hash(1)).To(Equal("N"))
		Expect(booking_code.Hash(25)).To(Equal("Q"))
		Expect(booking_code.Hash(26)).To(Equal("NM"))
		Expect(booking_code.Hash(27)).To(Equal("NN"))
		Expect(booking_code.Hash(28)).To(Equal("NB"))
		Expect(booking_code.Hash(29)).To(Equal("NV"))
		Expect(booking_code.Hash(30)).To(Equal("NC"))
		Expect(booking_code.Hash(31)).To(Equal("NX"))

		Expect(booking_code.Hash(2025)).To(Equal("BQE"))
		Expect(booking_code.Hash(9999)).To(Equal("SYA"))
	})
})

var _ = Describe("LeftPad", func() {
	It("add text padding", func() {
		Expect(booking_code.LeftPad("A", 2, "B")).To(Equal("AB"))
		Expect(booking_code.LeftPad("A", 4, "B")).To(Equal("AAAB"))
		Expect(booking_code.LeftPad("A", 4, "BBCD")).To(Equal("BBCD"))
	})
})

var _ = Describe("BookingCode", func() {
	It("returns hashed result", func() {
		result := booking_code.BookingCode("110", time.Date(2025, 06, 27, 0, 0, 0, 0, time.Local), 3034)

		Expect(result).To(Equal("110-NNZAQCFI"))

		result = booking_code.BookingCode("110", time.Date(2025, 06, 27, 0, 0, 0, 0, time.Local), 4)
		Expect(result).To(Equal("110-NNZAQAAC"))
	})
})
