package merkletree

func NewLeaf(hash []byte) *Branch {
	return &Branch{
		hash:   hash,
		lChild: nil,
		rChild: nil,
	}
}
