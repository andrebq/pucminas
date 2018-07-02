package crdt

// Less returns true if c message is less than o
// first it checks the clock, if they are the same,
// then toggle 0 wins
func (c ControlMessage) Less(o ControlMessage) bool {
	return c.Clock.Less(o.Clock)
}
