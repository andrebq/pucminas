package crdt

type (
	// Catalog contains the list of all valid identities
	Catalog struct {
		entries map[Identity]bool
	}
)

// NewCatalog returns an empty catalog
func NewCatalog() *Catalog {
	return &Catalog{
		entries: make(map[Identity]bool),
	}
}

// Add an identity to the catalog
func (c *Catalog) Add(i Identity) {
	c.entries[i] = true
}

// VerifyIdentity returns true if the given Identity is present in the catalog
func (c *Catalog) VerifyIdentity(i Identity) bool {
	return c.entries[i]
}
