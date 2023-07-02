package booking

import "time"
import "strconv"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	formatTime, _ := time.Parse(layout, date)
	return formatTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t := time.Now()
	parsedTime, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.After(parsedTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return parsedTime.Hour() >= 12 && parsedTime.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	return Schedule(date).Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	parsedTime, _ := time.Parse("2006-01-02", strconv.Itoa(time.Now().Year())+"-09-15")
	return parsedTime
}
