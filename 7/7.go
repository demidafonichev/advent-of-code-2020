package main

import (
	bn "7/bagnode"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// GetBagNode returns BagNode for the bag considering:
// BagNode is already in the tree - simply return pointer
// BagNode is not in the tree - create one and return pointer
func GetBagNode(tree *[]*bn.BagNode, bag string) (bagNode *bn.BagNode) {
	isBagInTree := false
	// If BagNode is in the tree
	for _, node := range *tree {
		if node.Bag == bag {
			isBagInTree = true
			bagNode = node
			return
		}
	}
	// If BagNode is not in the tree
	if !isBagInTree {
		bagNode = &bn.BagNode{Bag: bag, BaseBags: &[]*bn.BagNode{}}
		*tree = append(*tree, bagNode)
	}
	return
}

func main() {
	input, _ := os.Open("input.txt")
	r := bufio.NewReader(input)
	fileEnded := false

	bagsTree := []*bn.BagNode{}

	for !fileEnded {
		rule, err := r.ReadString('\n')
		if err == io.EOF {
			fileEnded = true
			break
		}
		// Would be [light, red, bags, contain, 1, bright, white, bag, 2, muted, yellow, bags.]
		ruleParts := strings.Split(rule, " ")

		baseBag := ruleParts[0] + ruleParts[1]
		baseBagNode := GetBagNode(&bagsTree, baseBag)

		for i := 4; i < len(ruleParts); {
			// ruleParts[i] would be "no" for rules where
			// bag contains no other bags
			if ruleParts[i] != "no" {
				nestedBag := ruleParts[i+1] + ruleParts[i+2]
				nestedBagNode := GetBagNode(&bagsTree, nestedBag)
				nestedBagNode.AppendBaseBag(baseBagNode)
				i += 4
			} else {
				break
			}
		}
	}

	for _, bagNode := range bagsTree {
		if bagNode.Bag == "shinygold" {
			baseBags := bagNode.FindAllBaseBags()
			fmt.Println(len(baseBags))
		}
	}
}
