package clock

import "fmt"

type clock struct {
	minute int
}

// Clock is a custom type based on clock struct
type Clock clock

// New creates a new clock
func New(hour, minute int) Clock {
	m := minute + hour*60
	m %= 24 * 60
	if m < 0 {
		m += 24 * 60
	}
	return Clock{m}
}

func (c Clock) String() string {
	hour := c.minute / 60
	minute := c.minute % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

// Add adds minutes to the clock
func (c Clock) Add(minutes int) Clock {
	return New(0, c.minute+minutes)
}

// Subtract subtracts minutes from the clock
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.minute-minutes)
}
