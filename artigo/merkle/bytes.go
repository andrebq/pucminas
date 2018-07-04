package merkle

// Bytes returns a copy of the path data appended to out
func (p Bytes) Bytes(out []byte) []byte {
	out = append(out, p.buf...)
	return out
}

// Concat returns a copy of this byte array with new items added to the end
func (p Bytes) Concat(items ...byte) Bytes {
	out := Bytes{
		buf: make([]byte, 0, len(p.buf)+len(items)),
	}
	out.buf = append(out.buf, p.buf...)
	out.buf = append(out.buf, items...)
	return out
}

// String returns the underlying bytes as a string (usefull when the keys are utf8 sequences of bytes)
func (p Bytes) String() string {
	return string(p.buf)
}

// Len returns the size of this bytes
func (p Bytes) Len() int {
	return len(p.buf)
}
