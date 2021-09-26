package date_time

import (
	"fmt"
	"os"
	"time"
)

/*
The king of working with times and dates in Go is the time.Time data type,
which represents an instant in time with nanosecond precision.

UNIX epoch time, which is the number of seconds that have elapsed since 00:00:00 UTC, January 1, 1970.

The time.Since() function calculates the time that has passed since a given time and returns a time.Duration variable


*/
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}

	// this code illustrates how to work with epoch time in Go and showcases the parsing process
	start := time.Now()

	dateString := os.Args[1]

	d, err := time.Parse("02 January 2006", dateString)

	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date -me-l:", d.Day(), d.Month(), d.Year())
	}

	// Is this a date + time?
	d, err = time.Parse("02 January 2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full: -fille-", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// Is this a date + time with month represented as a number?
	d, err = time.Parse("02-01-2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full -hit-:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// Is it time only?
	d, err = time.Parse("15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch time:", t)
	// Convert Epoch time to time.Time
	d = time.Unix(t, 0)
	fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())

	duration := time.Since(start)
	fmt.Println("Execution time:", duration)

	// fmt.Println(start)
}
