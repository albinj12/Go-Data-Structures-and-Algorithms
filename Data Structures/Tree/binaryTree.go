package main

import "fmt"

type node struct{
	data int
	left *node
	right *node
}

type binaryTree struct{
	root *node
}

func main(){
	bt := binaryTree{}
	// bt.insertNode(1)
	// bt.insertNode(2)
	// bt.insertNode(3)
	// bt.insertNode(4)
	// bt.insertNode(5)
	// bt.insertNode(6)
	// bt.insertNode(7)
	// bt.insertNode(8)
	// bt.levelOrderTraversal()
	// bt.preorderTraversal(bt.root)
	// bt.inorderTraversal(bt.root)
	// bt.postorderTraversal(bt.root)
	// bt.deleteNode(6)
}

// node is inserted in levelorder
func (bt *binaryTree) insertNode(value int){
	newNode := &node{
		value,
		nil,
		nil,
	}
	if bt.root == nil{
		bt.root = newNode
		return
	}
	var queue []*node
	queue = append(queue, bt.root)
	for len(queue) > 0{
		currentNode := queue[0]
		queue = append(queue[:0], queue[1:]...)
		if currentNode.left == nil{
			currentNode.left = newNode
			return
		}else{
			queue = append(queue, currentNode.left)
		}
		if currentNode.right == nil{
			currentNode.right = newNode
			return
		}else{
			queue = append(queue, currentNode.right)
		}
	}
}

func (bt *binaryTree) levelOrderTraversal(){
	if bt.root == nil{
		return
	}

	var queue []*node
	queue = append(queue, bt.root)
	for len(queue) > 0{
		currentNode := queue[0]
		queue = append(queue[:0], queue[1:]...)
		fmt.Println(currentNode.data)
		if currentNode.left != nil{
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil{
			queue = append(queue, currentNode.right)
		}
	}
}

func (bt *binaryTree) preorderTraversal(btNode *node){
	if btNode != nil{
		fmt.Println(btNode.data)
		bt.preorderTraversal(btNode.left)
		bt.preorderTraversal(btNode.right)
		
	}
}

func (bt *binaryTree) inorderTraversal(btNode *node){
	if btNode != nil{
		bt.inorderTraversal(btNode.left)
		fmt.Println(btNode.data)
		bt.inorderTraversal(btNode.right)
	}
}

func (bt *binaryTree) postorderTraversal(btNode *node){
	if btNode != nil{
		bt.postorderTraversal(btNode.left)
		bt.postorderTraversal(btNode.right)
		fmt.Println(btNode.data)
	}
}

// while deleting, we are replacing the value of deepest node
// with the node we need to delete
func (bt *binaryTree) deleteNode(value int){
	if bt.root == nil{
		return
	}

	var queue []*node
	var deepestNode *node
	var nodeToDelete *node
	queue = append(queue, bt.root)
	// finding deepestNode and node to delete
	for len(queue) > 0{
		deepestNode = queue[0]
		queue = append(queue[:0], queue[1:]...)
		if deepestNode.data == value{
			nodeToDelete = deepestNode
		}
		if deepestNode.left != nil{
			queue = append(queue, deepestNode.left)
		}
		if deepestNode.right != nil{
			queue = append(queue, deepestNode.right)
		}
	}

	nodeToDelete.data = deepestNode.data
	bt.deleteDeepestNode(deepestNode)
}

func (bt *binaryTree) deleteDeepestNode(deepestNode *node){
	var queue []*node
	queue = append(queue, bt.root)
	for len(queue) > 0{
		currentNode := queue[0]
		queue = append(queue[:0], queue[1:]...)
		// if currentNode.left == deepestNode or currentNode.right == deepestNode
		// it means currentNode is the parent of deepestNode
		// then we delete deepestNode and return
		
		if currentNode.left == deepestNode{
			currentNode.left = nil
			return
		}else{
			queue = append(queue, currentNode.left)
		}
		if currentNode.right == deepestNode{
			currentNode.right = nil
			return
		}else{
			queue = append(queue, currentNode.right)
		}
	}
}