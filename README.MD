# How to use
```sh
go get github.com/Nuanu-com/booking-code@v0.1.4
```

```go
import (
  booking_code "github.com/Nuanu-com/booking-code"
  "time"
)

func main() {
  funnelID := "010"
  orderDate := time.Date(2025, 06, 27, 0, 0, 0, 0, time.Local)
  orderSequence := 3034

  result := booking_code.BookingCode(funnelID, orderDate, orderSequence)

  fmt.Printf("%s", result) // => 110-NNZAQCFI
}
```

Read [Here](https://nuanu-team.monday.com/docs/2028089117) For more detailed information regarding the logic