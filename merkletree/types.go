package merkletree

// Branches are made by combining Leaf/branch nodes
type Branch struct {
	hash   []byte
	lChild *Branch
	rChild *Branch
}
