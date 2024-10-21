package merkletree_model 

type MerkleTree struct {
	RootNode *MerkleNode
}

func NewMerkleTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	if len(data)%2 ==0 {
		data = append(data, data[len(data)-1])
	}

	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}
	
	for {
		var newLevel []MerkleNode
		if len(nodes)%2 == 1 {
			nodes = append(nodes, nodes[len(nodes)-1])
		}
		for j:=0; j<len(nodes); j+=2 {
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil) 
			newLevel = append(newLevel, *node)
		}
		nodes = newLevel

		if len(nodes) == 1 {
			break
		}
	}

	mTree := MerkleTree{&nodes[0]}
	return &mTree
}