// Package bst implements a binary search tree of integers.
package bst

// Node stores one value and links to its children. Its fields are private so a
// BST can maintain the ordering invariant through its methods.
type Node struct {
	value int
	left  *Node
	right *Node
}

// BST is a binary search tree that stores unique integer values.
type BST struct {
	root *Node
}

// New returns an empty binary search tree.
func New() *BST { return &BST{} }

// Insert adds value to the tree. It returns false when value is already present.
func (b *BST) Insert(value int) bool {
	if b.root == nil {
		b.root = &Node{value: value}
		return true
	}

	current := b.root
	for {
		if value < current.value {
			if current.left == nil {
				current.left = &Node{value: value}
				return true
			}
			current = current.left
		} else if value > current.value {
			if current.right == nil {
				current.right = &Node{value: value}
				return true
			}
			current = current.right
		} else {
			return false
		}
	}
}

// Remove deletes value from the tree and reports whether it was present.
func (b *BST) Remove(value int) bool {
	var parent *Node
	current := b.root
	for current != nil && current.value != value {
		parent = current
		if value < current.value {
			current = current.left
		} else {
			current = current.right
		}
	}
	if current == nil {
		return false
	}

	replacement := removeNode(current)
	if parent == nil {
		b.root = replacement
	} else if parent.left == current {
		parent.left = replacement
	} else {
		parent.right = replacement
	}
	return true
}

func removeNode(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	if node.right == nil {
		return node.left
	}

	replacement := node.left
	rightmost := replacement
	for rightmost.right != nil {
		rightmost = rightmost.right
	}
	rightmost.right = node.right
	return replacement
}

// InOrderTraversal returns values in ascending order.
func (b *BST) InOrderTraversal() []int {
	values := make([]int, 0)
	var visit func(*Node)
	visit = func(node *Node) {
		if node == nil {
			return
		}
		visit(node.left)
		values = append(values, node.value)
		visit(node.right)
	}
	visit(b.root)
	return values
}

// PreOrderTraversal returns values in root-left-right order.
func (b *BST) PreOrderTraversal() []int {
	values := make([]int, 0)
	var visit func(*Node)
	visit = func(node *Node) {
		if node == nil {
			return
		}
		values = append(values, node.value)
		visit(node.left)
		visit(node.right)
	}
	visit(b.root)
	return values
}

// PostOrderTraversal returns values in left-right-root order.
func (b *BST) PostOrderTraversal() []int {
	values := make([]int, 0)
	var visit func(*Node)
	visit = func(node *Node) {
		if node == nil {
			return
		}
		visit(node.left)
		visit(node.right)
		values = append(values, node.value)
	}
	visit(b.root)
	return values
}

// LevelOrderTraversal returns values one tree level at a time.
func (b *BST) LevelOrderTraversal() []int {
	if b.root == nil {
		return []int{}
	}

	values := make([]int, 0)
	queue := []*Node{b.root}
	for head := 0; head < len(queue); head++ {
		current := queue[head]
		values = append(values, current.value)
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
	return values
}

// Contains reports whether value is present in the tree. It takes O(h) time,
// where h is the height of the tree.
func (b *BST) Contains(value int) bool {
	current := b.root
	for current != nil {
		if value < current.value {
			current = current.left
		} else if value > current.value {
			current = current.right
		} else {
			return true
		}
	}
	return false
}

// Min returns the smallest value in the tree. The second result is false for
// an empty tree.
func (b *BST) Min() (int, bool) {
	if b.root == nil {
		return 0, false
	}

	current := b.root
	for current.left != nil {
		current = current.left
	}
	return current.value, true
}

// Max returns the largest value in the tree. The second result is false for
// an empty tree.
func (b *BST) Max() (int, bool) {
	if b.root == nil {
		return 0, false
	}

	current := b.root
	for current.right != nil {
		current = current.right
	}
	return current.value, true
}

// Height returns the number of nodes on the longest root-to-leaf path. An
// empty tree has height 0 and a single-node tree has height 1.
func (b *BST) Height() int {
	var height func(*Node) int
	height = func(node *Node) int {
		if node == nil {
			return 0
		}

		leftHeight := height(node.left)
		rightHeight := height(node.right)
		if leftHeight > rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	}

	return height(b.root)
}

// IsBalanced reports whether every node's left and right subtrees differ in
// height by at most one. It visits every node at most once.
func (b *BST) IsBalanced() bool {
	var balancedHeight func(*Node) int
	balancedHeight = func(node *Node) int {
		if node == nil {
			return 0
		}

		leftHeight := balancedHeight(node.left)
		if leftHeight == -1 {
			return -1
		}
		rightHeight := balancedHeight(node.right)
		if rightHeight == -1 || leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
			return -1
		}

		if leftHeight > rightHeight {
			return leftHeight + 1
		}
		return rightHeight + 1
	}

	return balancedHeight(b.root) != -1
}

// IsValid reports whether the tree satisfies the strict BST ordering rule:
// every left descendant is smaller and every right descendant is larger.
func (b *BST) IsValid() bool {
	var valid func(node *Node, lower, upper *int) bool
	valid = func(node *Node, lower, upper *int) bool {
		if node == nil {
			return true
		}
		if (lower != nil && node.value <= *lower) || (upper != nil && node.value >= *upper) {
			return false
		}

		return valid(node.left, lower, &node.value) && valid(node.right, &node.value, upper)
	}

	return valid(b.root, nil, nil)
}

// KthSmallest returns the kth smallest value using one iterative in-order
// traversal. k is one-indexed; for example, k == 1 returns the minimum.
// The second result is false when k is outside the number of tree values.
func (b *BST) KthSmallest(k int) (int, bool) {
	if k <= 0 {
		return 0, false
	}

	stack := make([]*Node, 0)
	current := b.root
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return current.value, true
		}
		current = current.right
	}

	return 0, false
}

// LowestCommonAncestor returns the lowest shared ancestor of first and second.
// Both values must be present; otherwise the second result is false.
func (b *BST) LowestCommonAncestor(first, second int) (int, bool) {
	if !b.Contains(first) || !b.Contains(second) {
		return 0, false
	}

	current := b.root
	for current != nil {
		if first < current.value && second < current.value {
			current = current.left
		} else if first > current.value && second > current.value {
			current = current.right
		} else {
			return current.value, true
		}
	}

	return 0, false
}
