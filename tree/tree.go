package tree

import "fmt"

type binaryTree struct {
	key   int
	left  *binaryTree
	right *binaryTree
}

type RootTree struct {
	tree  *binaryTree
	nodes int
}

func (tree *binaryTree) dfs() {
	if tree.left != nil {
		tree.left.dfs()
	}
	fmt.Println(tree.key)
	if tree.right != nil {
		tree.right.dfs()
	}
}

func (root *RootTree) Dfs() {
	root.tree.dfs()
}

func (tree *binaryTree) appendNode(value int) {
	if tree.key <= value {
		if tree.left == nil {
			tree.left = &binaryTree{value, nil, nil}
		} else {
			tree.left.appendNode(value)
		}
	} else {
		if tree.right == nil {
			tree.right = &binaryTree{value, nil, nil}
		} else {
			tree.right.appendNode(value)
		}
	}
}

func (root *RootTree) appendNode(value int) {
	if root.tree == nil {
		root.tree = &binaryTree{value, nil, nil}
	} else {
		appendTree := root.tree
		appendTree.appendNode(value)
	}
	root.nodes++
}

func InitTree() RootTree {
	tree := RootTree{nil, 0}
	for i := 0; tree.nodes != 10; i++ {
		tree.appendNode(i)
	}
	return tree
}
