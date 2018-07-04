package merkle

type (

	// Tree implements the merkle-tree
	Tree struct {
		root *Node
	}

	// TreeExport contains an exported version of the tree
	TreeExport struct {
		Root  Hash
		Nodes []NodeExport
	}
)

// NewTree creates an empty merkle tree
func NewTree() *Tree {
	return &Tree{
		root: NewNode(),
	}
}

// Add includes a new value into the tree
func (t *Tree) Add(key, value []byte) *Tree {
	newnode := t.root.Add(key, value)
	newtree := *t
	newtree.root = newnode
	return &newtree
}

// Get returns the value of the given key
func (t *Tree) Get(key []byte) (Bytes, bool) {
	return t.root.Get(key)
}

// Hash returns the hash of this tree
func (t *Tree) Hash() Hash {
	return t.root.Hash()
}

// Export returns an export version of the tree to be sent over the wire/disk
func (t *Tree) Export() TreeExport {
	te := TreeExport{
		Root: t.root.Hash(),
	}
	t.root.Walk(func(p Bytes, n *Node) (bool, error) {
		te.Nodes = append(te.Nodes, n.Export())
		return true, nil
	})
	return te
}
