package clock

import "fmt"

type clock struct {
	hour   int
	minute int
}

type Clock clock

func New(hour, minute int) Clock {
	var c = Clock{}
	c.Add(minute + hour*60)
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (d Clock) Add(minutes int) Clock {
	c := Clock {
		hour: d.hour,
		minute: d.minute,
	}
	fmt.Println("C from Add2", c, minutes)
	hoursFromMin := minutes / 60
	hoursFromMin += c.hour
	minutes %= 60
	minutes += c.minute

	hoursFromMin += 24
	hoursFromMin %= 24

	if minutes < 0 {
		hoursFromMin += 23
		minutes += c.minute + 60
	}

	// need to mod 60 and add hours and such
	hoursFromMin += minutes / 60
	minutes %= 60
	hoursFromMin %= 24

	c.hour = hoursFromMin
	c.minute = minutes

	return c
}

func (d Clock) Subtract(minutes int) Clock {
	c := Clock {
		hour: d.hour,
		minute: d.minute,
	}
	c.Add(minutes * -1)
	return c
}
