package deadline

import "time"

// Deadline allows for commands to expire when they run too long
type Deadline struct {
	unixNano int64
	hit      bool
}

// New returns a new deadline object
func New(deadline time.Time) *Deadline {
	return &Deadline{unixNano: deadline.UnixNano()}
}

// Empty deadline does nothing, just a place holder for future updates
func Empty() *Deadline {
	return &Deadline{}
}

// Update the deadline from a given time object
func (deadline *Deadline) Update(newDeadline time.Time) {
	deadline.unixNano = newDeadline.UnixNano()
}

// Check the deadline and panic when reached
//go:noinline
func (deadline *Deadline) Check() {
	if deadline == nil || deadline.unixNano == 0 {
		return
	}
	if !deadline.hit && time.Now().UnixNano() > deadline.unixNano {
		deadline.hit = true
		panic("deadline")
	}
}

// Hit returns true if the deadline has been hit
func (deadline *Deadline) Hit() bool {
	return deadline.hit
}

// GetDeadlineTime returns the time object for the deadline, and an "empty" boolean
func (deadline *Deadline) GetDeadlineTime() (time.Time, bool) {
	return time.Unix(0, deadline.unixNano), deadline.unixNano == 0
}
