package tree

type Node struct {
	right *Node
	left  *Node
	value int
}

type Tree struct {
	root  *Node
	level int
}

func CreateTree(root int) *Node {
	return &Node{
		right: nil,
		left:  nil,
		value: root,
	}
}

func (doubleTree *Tree) BuildTree(root, left, right int) *Node {

	doubleTree.level++
	parentNode := FindParentNode(doubleTree.root, root)
	parentNode.left = &Node{
		right: nil,
		left:  nil,
		value: left,
	}
	parentNode.right = &Node{
		right: nil,
		left:  nil,
		value: right,
	}
	return parentNode
}

func FindParentNode(node *Node, value int) *Node {

	//for node != nil && node.value != value {
	//	if node.value < value {
	//		node = node.left
	//	} else {
	//		node = node.right
	//	}
	//}
	//return node

	if node == nil {
		return nil
	}
	if node.value == value {
		return node
	}
	if node.value > value {
		node := node.left
		FindParentNode(node, value)
	} else {
		node := node.right
		FindParentNode(node, value)
	}
	return nil
}
