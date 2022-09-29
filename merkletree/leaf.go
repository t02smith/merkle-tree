package merkletree

func NewLeaf(hash []byte) *Branch {
	return &Branch{
		Hash:   hash,
		LChild: nil,
		RChild: nil,
	}
}
