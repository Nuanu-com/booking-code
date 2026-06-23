package booking_code

import (
	"fmt"
	"strings"
	"time"
)

func Unhash(s string) int {
	result := 0
	for _, c := range s {
		for i, ch := range CharArray {
			if ch == string(c) {
				result = result*NumberBase + i
				break
			}
		}
	}
	return result
}

// ParseBookingCode reverses BookingCode, returning the funnel code, order date, and sequence.
// The year millennium is inferred from the current year.
func ParseBookingCode(code string) (funnelCode string, orderDate time.Time, orderSequence int, err error) {
	parts := strings.SplitN(code, "-", 2)
	if len(parts) != 2 || len(parts[1]) != 8 {
		return "", time.Time{}, 0, fmt.Errorf("invalid booking code: %q", code)
	}

	funnelCode = parts[0]
	suffix := parts[1]
	pad := CharArray[0]

	day := Unhash(strings.TrimLeft(suffix[0:2], pad))
	month := Unhash(suffix[2:3])
	yearDigit := Unhash(strings.TrimLeft(suffix[3:5], pad))
	orderSequence = Unhash(strings.TrimLeft(suffix[5:8], pad))

	fullYear := (time.Now().Year()/1000)*1000 + yearDigit
	orderDate = time.Date(fullYear, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return
}
