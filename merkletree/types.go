package merkletree

// Branches are made by combining Leaf/branch nodes
type Branch struct {
	Hash   []byte
	LChild *Branch
	RChild *Branch
}
