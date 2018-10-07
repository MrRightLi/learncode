package main

import (
	"project/google_code/tree"
)

type MyTreeNode struct {
	node *tree.Node
}


// 中续遍历
func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.Left}
	right := MyTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	root.Print()
	root.SetValue(100)

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	// nil指针也可以调用方法，不会报错
	var qRoot *tree.Node
	qRoot.SetValue(2000)
	qRoot = &root
	qRoot.SetValue(300)
	qRoot.Print()

	root.Traverse()
	myRoot := MyTreeNode{&root}
	myRoot.postOrder()
}
