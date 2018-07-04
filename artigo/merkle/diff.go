package merkle

type (
	// DiffFn is called with the leaf key and the values from the two different nodes
	//
	// If false is returned the diff process stops, if an error is returned the diff process
	// stops and the error is exposed to the client
	DiffFn func(k Bytes, a, b Bytes) (bool, error)
)

// Diff calls fn with the different leaf
func Diff(a, b *Node, fn DiffFn) error {
	aLeafs := make(map[string]*Node)
	bLeafs := make(map[string]*Node)

	addToNode := func(dest map[string]*Node) WalkFn {
		return func(k Bytes, n *Node) (bool, error) {
			if !n.Leaf() {
				return true, nil
			}
			dest[k.String()] = n
			return true, nil
		}
	}

	// this is the works option, but works as a reference implementation
	a.Walk(addToNode(aLeafs))
	b.Walk(addToNode(bLeafs))

	type pair struct {
		a *Node
		b *Node
	}

	diffs := make(map[string]pair)

	for k, v := range aLeafs {
		b, ok := bLeafs[k]
		if !ok {
			// key is present on a but not on b
			diffs[k] = pair{aLeafs[k], nil}
			continue
		}

		if v.Hash() != b.Hash() {
			// key is present on both
			diffs[k] = pair{v, b}
		}
	}

	for k, v := range bLeafs {
		if _, ok := diffs[k]; ok {
			// key already added to diff, skip
			continue
		}
		_, ok := aLeafs[k]
		if ok {
			// item exists on a, but not on diff
			// means it was scanned and the hash is the same
			continue
		}
		// key is present on b but not on a
		diffs[k] = pair{nil, v}
	}

	for k, p := range diffs {
		ok, err := fn(Bytes{[]byte(k)}, p.a.Value(), p.b.Value())
		if err != nil {
			return err
		}
		if !ok {
			return nil
		}
	}

	return nil
}
