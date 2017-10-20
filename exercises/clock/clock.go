package clock

import (
	"fmt"
)

const testVersion = 4

type Clock struct {
	hour, minute int
}

func New(hour, minute int) Clock {
	return Clock{hour, minute}.formatTime()
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	c.minute += minutes
	return c.formatTime()
}

func (c Clock) formatTime() Clock {
	//catch negative values by looping until positive
	if c.minute < 0 || c.hour < 0 {
		for c.minute < 0 {
			c.minute += 60
			c.hour--
		}
		for c.hour < 0 {
			c.hour += 24
		}
	} else {
		rem := (c.minute) / 60
		c.minute = (c.minute) % 60
		c.hour = (c.hour + rem) % 24
	}
	return c
}
