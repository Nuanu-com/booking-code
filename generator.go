package booking_code

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"time"
)

var CharArray []string

const NumberBase = 26

//go:embed .config.json
var configByte []byte

func init() {
	if err := json.Unmarshal(configByte, &CharArray); err != nil {
		panic(err)
	}
}

func QuotientAndReminder(number int) (int, int) {
	quotient := number / NumberBase
	reminder := number % NumberBase

	return quotient, reminder
}

func Hash(number int) string {
	hashResult := ""

	q, r := QuotientAndReminder(number)

	if l := CharArray[r]; l != "" {
		hashResult = fmt.Sprintf("%s%s", hashResult, l)
	}

	if q == 0 {
		return hashResult
	} else {
		return fmt.Sprintf("%s%s", Hash(q), hashResult)
	}
}

func LeftPad(pad string, width int, hash string) string {
	if len(hash) >= width {
		return hash
	}

	rem := width - len(hash)

	padResult := ""

	for range rem {
		padResult = fmt.Sprintf("%s%s", pad, padResult)
	}

	return fmt.Sprintf("%s%s", padResult, hash)
}

// Expected format {XXX}-{DDMYYXXX}
// Where {XXX} = is the funnel code
// {DDMYYXXX} => the value is the hashed version of
func BookingCode(funnelCode string, orderDate time.Time, orderSequence int) string {
	day := orderDate.Day()
	month := orderDate.Month()
	year := orderDate.Year()

	yearDigit := year % 1000 // To get the last 2 digit

	dayF := LeftPad("A", 2, Hash(day))
	monthF := Hash(int(month))
	yearF := LeftPad("A", 2, Hash(yearDigit))
	sequenceF := LeftPad("X", 3, Hash(orderSequence))

	return fmt.Sprintf("%s-%s%s%s%s", funnelCode, dayF, monthF, yearF, sequenceF)
}
