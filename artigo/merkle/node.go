package merkle

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
)

type (
	// Hash holds a SHA-1 hash
	Hash [20]byte

	// Node contains a list of children nodes
	Node struct {
		children map[byte]*Node
		block    []byte
		hash     Hash
	}

	// NodeExport contains an exported version of the node
	// all pointer references are replaced by the hash of the node
	NodeExport struct {
		Self     Hash
		Children map[byte]Hash
		Value    []byte
	}

	// Storage allows the items to be indexed by their Hash
	Storage struct {
		items map[Hash]Node
	}
)

// NewNode returns a new tree node
func NewNode() *Node {
	n := &Node{}
	n.prehash()
	return n
}

// Add includes the given k/v pair in the tree
// a new *Node is returned since this is a immutable
// structure.
//
// Only the affected path is changed, everything else is shared
// between other instances
func (n *Node) Add(newkey, newvalue []byte) *Node {
	if len(newkey) == 0 {
		ret := &Node{block: newvalue}
		ret.prehash()
		return ret
	}
	copy := &Node{}
	copy.children = make(map[byte]*Node)
	k0 := newkey[0]
	for k, v := range n.children {
		if k != k0 {
			copy.children[k] = v
		} else {
			copy.children[k] = v.Add(newkey[1:], newvalue)
		}
	}
	if _, ok := copy.children[k0]; ok {
		// already copied
		return copy
	}

	// missing node
	newchildren := &Node{}
	copy.children[k0] = newchildren.Add(newkey[1:], newvalue)
	newchildren.prehash()
	copy.prehash()
	return copy
}

// Get returns the value associated with the given key
func (n *Node) Get(key []byte) ([]byte, bool) {
	if len(key) == 0 {
		return append(n.block[:0], n.block...), true
	}
	k0 := key[0]
	if len(n.children) == 0 {
		return nil, false
	}
	n = n.children[k0]
	if n == nil {
		return nil, false
	}
	return n.Get(key[1:])
}

// Hash returns the hash of this node
func (n *Node) Hash() Hash {
	return n.hash
}

func (n *Node) prehash() *Node {
	h := sha1.New()
	h.Write(n.block)

	ints := make(sort.IntSlice, len(n.children))
	var i int
	for k := range n.children {
		ints[i] = int(k)
		i++
	}
	sort.Sort(ints)

	for _, v := range ints {
		c := n.children[byte(v)]
		h.Write(c.hash[:])
	}
	buf := h.Sum(nil)
	copy(n.hash[:], buf)
	return n
}

// Export returns the export version of this node
func (n *Node) Export() NodeExport {
	ne := NodeExport{}
	ne.Self = n.Hash()
	ne.Value = append(ne.Value, n.block...)
	ne.Children = make(map[byte]Hash)
	for k, v := range n.children {
		ne.Children[k] = v.Hash()
	}
	return ne
}

// String implements Stringer interface
func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

// String implements Stringer interface
func (ne NodeExport) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "{%v, [", ne.Self.String())
	for k, v := range ne.Children {
		fmt.Fprintf(buf, " (%v, %v)", k, v.String())
	}
	fmt.Fprintf(buf, " ] %v}", hex.EncodeToString(ne.Value))
	return buf.String()
}
