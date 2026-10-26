package main

import (
	"fmt"

	bst "myproject/BST"
)

func main() {
	tree := bst.New()
	for _, value := range []int{50, 30, 70, 20, 40, 60, 80, 75} {
		tree.Insert(value)
	}

	fmt.Println("Inorder:", tree.InOrderTraversal())
	fmt.Println("Preorder:", tree.PreOrderTraversal())
	fmt.Println("Postorder:", tree.PostOrderTraversal())
	fmt.Println("Level order:", tree.LevelOrderTraversal())
	fmt.Println("Height:", tree.Height())
	fmt.Println("Balanced:", tree.IsBalanced())
	fmt.Println("Valid BST:", tree.IsValid())

	if value, ok := tree.KthSmallest(3); ok {
		fmt.Println("3rd smallest:", value)
	}
	if value, ok := tree.LowestCommonAncestor(20, 40); ok {
		fmt.Println("LCA of 20 and 40:", value)
	}

	tree.Remove(70)
	fmt.Println("Inorder after removing 70:", tree.InOrderTraversal())
}
