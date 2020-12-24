package bagnode

// BagNode contains information about
// which bag is nested by which
type BagNode struct {
	Bag      string
	BaseBags *[]*BagNode
}

// AppendBaseBag appends pointer to base BagNode
// to list of current BagNode's array of base bags
func (bn *BagNode) AppendBaseBag(baseBagNode *BagNode) {
	*bn.BaseBags = append(*bn.BaseBags, baseBagNode)
}

// FindAllBaseBags searches for all base bags
// in the tree from node BagNode
// Recursively searches for base bags, then for base bags of
// base bags and so on, removing duplicates
func (bn *BagNode) FindAllBaseBags() (baseBags []string) {
	// Append base bags on current bag first
	for _, baseBagNode := range *bn.BaseBags {
		baseBags = append(baseBags, baseBagNode.Bag)
	}
	// Recursively search for base bags of base bag
	for _, baseBagNode := range *bn.BaseBags {
		baseBaseBags := baseBagNode.FindAllBaseBags()
		isBaseBaseBagAlreadyFound := false
		// Removing duplicates in the result
		for _, baseBaseBag := range baseBaseBags {
			for _, alreadyFoundBag := range baseBags {
				if baseBaseBag == alreadyFoundBag {
					isBaseBaseBagAlreadyFound = true
					break
				}
			}
			if !isBaseBaseBagAlreadyFound {
				baseBags = append(baseBags, baseBaseBag)
			}
		}
	}
	return
}
