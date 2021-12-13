package bbolt

import "crypto/sha256"

func UpdatingHash(index int, n *node) {
	/// updating current nodes' hash after the index-th child of node has changed
	// todo: in case of children of n contains another bucket, need further processing.
	hashingBytes := make([]byte, 0)
	// for updating
	for i, v := range n.inodes {
		if i != index {
			hashingBytes = append(hashingBytes, v.hash...)
		} else {
			hashing := sha256.Sum256(v.value)
			hashingBytes = append(hashingBytes, hashing[:]...)
		}
	}
	t := sha256.Sum256(hashingBytes)
	n.hash = t[:]
}

func CalculatingHashing(n *node, t bool) {
	// t: true for leaf node, false for non leaf node
	hashingBytes := make([]byte, 0)
	for _, v := range n.inodes {
		if t {
			hash := sha256.Sum256(v.value)
			hashingBytes = append(hashingBytes, hash[:]...)
		} else {
			hashingBytes = append(hashingBytes, v.hash...)
		}
	}
	hash := sha256.Sum256(hashingBytes)
	n.hash = hash[:]
}
