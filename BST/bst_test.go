package bst

import (
	"reflect"
	"testing"
)

func TestInsertAndTraversals(t *testing.T) {
	tree := New()
	for _, value := range []int{50, 30, 70, 20, 40, 60, 80, 75} {
		if !tree.Insert(value) {
			t.Fatalf("Insert(%d) returned false", value)
		}
	}
	if tree.Insert(50) {
		t.Fatal("duplicate insert returned true")
	}

	tests := []struct {
		name string
		got  []int
		want []int
	}{
		{"in-order", tree.InOrderTraversal(), []int{20, 30, 40, 50, 60, 70, 75, 80}},
		{"pre-order", tree.PreOrderTraversal(), []int{50, 30, 20, 40, 70, 60, 80, 75}},
		{"post-order", tree.PostOrderTraversal(), []int{20, 40, 30, 60, 75, 80, 70, 50}},
		{"level-order", tree.LevelOrderTraversal(), []int{50, 30, 70, 20, 40, 60, 80, 75}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !reflect.DeepEqual(test.got, test.want) {
				t.Errorf("got %v; want %v", test.got, test.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tree := New()
	for _, value := range []int{50, 30, 70, 20, 40, 60, 80, 75} {
		tree.Insert(value)
	}
	for _, value := range []int{20, 80, 70, 50} {
		if !tree.Remove(value) {
			t.Fatalf("Remove(%d) returned false", value)
		}
	}
	if tree.Remove(999) {
		t.Fatal("Remove of absent value returned true")
	}

	if got, want := tree.InOrderTraversal(), []int{30, 40, 60, 75}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestEmptyTree(t *testing.T) {
	tree := New()
	if tree.Remove(1) {
		t.Fatal("Remove from empty tree returned true")
	}
	if got := tree.InOrderTraversal(); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("got %v; want empty slice", got)
	}
}

func TestInterviewMethods(t *testing.T) {
	tree := New()
	for _, value := range []int{50, 30, 70, 20, 40, 60, 80, 75} {
		tree.Insert(value)
	}

	if !tree.Contains(75) || tree.Contains(99) {
		t.Error("Contains returned an unexpected result")
	}
	if value, ok := tree.Min(); !ok || value != 20 {
		t.Errorf("Min() = (%d, %t), want (20, true)", value, ok)
	}
	if value, ok := tree.Max(); !ok || value != 80 {
		t.Errorf("Max() = (%d, %t), want (80, true)", value, ok)
	}
	if got, want := tree.Height(), 4; got != want {
		t.Errorf("Height() = %d, want %d", got, want)
	}
	if !tree.IsBalanced() {
		t.Error("IsBalanced() = false, want true")
	}
	if !tree.IsValid() {
		t.Error("IsValid() = false, want true")
	}
	if value, ok := tree.KthSmallest(5); !ok || value != 60 {
		t.Errorf("KthSmallest(5) = (%d, %t), want (60, true)", value, ok)
	}
	if _, ok := tree.KthSmallest(0); ok {
		t.Error("KthSmallest(0) reported a value")
	}
	if value, ok := tree.LowestCommonAncestor(20, 40); !ok || value != 30 {
		t.Errorf("LowestCommonAncestor(20, 40) = (%d, %t), want (30, true)", value, ok)
	}
	if _, ok := tree.LowestCommonAncestor(20, 99); ok {
		t.Error("LowestCommonAncestor returned a value for an absent node")
	}
}

func TestEmptyTreeInterviewMethods(t *testing.T) {
	tree := New()
	if _, ok := tree.Min(); ok {
		t.Error("Min() returned a value for an empty tree")
	}
	if _, ok := tree.Max(); ok {
		t.Error("Max() returned a value for an empty tree")
	}
	if got := tree.Height(); got != 0 {
		t.Errorf("Height() = %d, want 0", got)
	}
	if !tree.IsBalanced() || !tree.IsValid() {
		t.Error("an empty tree should be balanced and valid")
	}
}

func TestUnbalancedAndInvalidTree(t *testing.T) {
	tree := &BST{
		root: &Node{
			value: 10,
			left: &Node{
				value: 5,
				left:  &Node{value: 2, left: &Node{value: 1}},
			},
		},
	}
	if tree.IsBalanced() {
		t.Error("IsBalanced() = true, want false for a left-skewed tree")
	}

	// A value in the left subtree larger than the root violates the BST rule.
	tree.root.left.right = &Node{value: 15}
	if tree.IsValid() {
		t.Error("IsValid() = true, want false for an invalid tree")
	}
}
