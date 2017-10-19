package clock

import (
	"fmt"
)

const testVersion = 4

type Clock struct {
	hour   int
	minute int
}

func New(hour, minute int) Clock {

	//catch negative values by looping until positive
	for minute < 0 {
		minute += 60
		hour--
	}
	for hour < 0 {
		hour += 24
	}
	//catch hours higher than 23 and minutes higher than 60 with modulus
	return Clock{
		hour:   (hour + minute/60) % 24,
		minute: minute % 60,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {

	if minutes < 0 {
		c.minute += minutes
		for c.minute < 0 {
			c.minute += 60
			c.hour--
		}
		for c.hour < 0 {
			c.hour += 24
		}
	} else {
		rem := (c.minute + minutes) / 60
		c.minute = (c.minute + minutes) % 60
		c.hour = (c.hour + rem) % 24
	}
	return c

}

// Remember to delete all of the stub comments.
// They are just noise, and reviewers will complain.
