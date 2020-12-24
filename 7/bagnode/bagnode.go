package bagnode

// BagNode contains information about
// which bag is nested by which
type BagNode struct {
	Bag        string
	NestedBags map[*BagNode]int
}

// AppendNestedBag appends pointer to nested BagNode
// to list of current BagNode's array of nested bags
func (bn *BagNode) AppendNestedBag(nestedBagNode *BagNode, count int) {
	bn.NestedBags[nestedBagNode] = count
}

// FindCountAllNestedBags count  for all nested bags in the tree from node BagNode
// Recursively searches for nested bags, then for nested bags of nested bags and so on
// Bags with no bested bags return 1
// All other bags return 1 + result of call on nested bag * count
func (bn *BagNode) FindCountAllNestedBags() (countNestedBags int) {
	countNestedBags = 1
	for nestedBagNode, count := range bn.NestedBags {
		countNestedBags += nestedBagNode.FindCountAllNestedBags() * count
	}
	return
}
