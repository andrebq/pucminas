package crdt

import "time"

// Less returns true if c is less than o, first the epoch field
// is compared and then the Unix
func (c Clock) Less(o Clock) bool {
	return c.Epoch <= o.Epoch &&
		c.Unix < o.Unix
}

// NewClock returns a clock with the given epoch
func NewClock(epoch uint32) Clock {
	return Clock{
		Epoch: epoch,
		Unix:  time.Now().UnixNano() / int64(time.Millisecond),
	}
}

// Tick returns a new clock that indicates a time after the current clock
// if the OS clock has decreased the current clock time is increased by 1ms
func (c Clock) Tick() Clock {
	nc := NewClock(c.Epoch)
	if !c.Less(nc) {
		nc.Unix = c.Unix + 1
	}
	return nc
}
