package crdt

// Less returns true if c is less than o, first the epoch field
// is compared and then the Unix
func (c Clock) Less(o Clock) bool {
	return c.Epoch <= o.Epoch &&
		c.Unix < o.Unix
}
