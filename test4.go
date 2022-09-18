package main

import (
	"time"
)

func Schedule(date string) time.Time {
	ans, _ := time.Parse(time.RFC3339, date)
	return ans
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}

//func main() {
//	fmt.Println(Schedule("2006-01-02T15:04:05Z	"))
//}
