package clock

import "fmt"

type clock struct {
	hour   int
	minute int
}

// Clock is a custom type based on clock struct
type Clock clock

// New creates a new clock
func New(hour, minute int) Clock {
	hour += minute / 60
	minute %= 60

	hour += 24
	hour %= 24

	if minute < 0 {
		hour--
		minute += 60
	}

	hour %= 24
	if hour < 0 {
		hour += 24
	}

	return Clock{
		hour:   hour,
		minute: minute,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds minutes to the clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract subtracts minutes from the clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
