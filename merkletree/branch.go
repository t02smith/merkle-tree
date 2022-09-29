package merkletree

import "crypto/sha256"

func (b *Branch) generateHash() error {
	sha256 := sha256.New()

	// Left child node

	lHash, err := b.LChild.GetHash()
	if err != nil {
		return err
	}

	_, err = sha256.Write(lHash)
	if err != nil {
		return err
	}

	// right child node

	rHash, err := b.RChild.GetHash()
	if err != nil {
		return err
	}

	_, err = sha256.Write(rHash)
	if err != nil {
		return err
	}

	// final hash

	b.Hash = sha256.Sum([]byte{})
	return nil
}

func (b *Branch) GetHash() ([]byte, error) {
	if b.Hash == nil {
		err := b.generateHash()
		if err != nil {
			return nil, err
		}
	}

	return b.Hash, nil
}
