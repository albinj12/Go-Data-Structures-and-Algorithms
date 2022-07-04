package main

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

type binarySearchTree struct {
	root *node
}

func main() {
	bst := binarySearchTree{}
}

func (bst *binarySearchTree) insert(value int) {
	if bst.root == nil {
		bst.root = &node{
			data: value,
		}
		return
	}

	insertNode(bst.root, value)

}

func insertNode(bstNode *node, value int) {
	if value == bstNode.data {
		fmt.Println("Duplicate Data")
		return
	}

	if value < bstNode.data {
		if bstNode.left == nil {
			bstNode.left = &node{
				data: value,
			}
		} else {
			insertNode(bstNode.left, value)
		}
	} else {
		if bstNode.right == nil {
			bstNode.right = &node{
				data: value,
			}
		} else {
			insertNode(bstNode.right, value)
		}
	}
}

func (bst *binarySearchTree) search(value int) *node {
	ptr := bst.root
	for ptr != nil {
		if ptr.data == value {
			return ptr
		} else if value < ptr.data {
			ptr = ptr.left
		} else {
			ptr = ptr.right
		}
	}
	return nil
}

func (bst *binarySearchTree) preorderTraversal(bstNode *node) {
	if bstNode != nil {
		fmt.Println(bstNode.data)
		bst.preorderTraversal(bstNode.left)
		bst.preorderTraversal(bstNode.right)
	}
}

func (bst *binarySearchTree) inorderTraversal(bstNode *node) {
	if bstNode != nil {
		bst.inorderTraversal(bstNode.left)
		fmt.Println(bstNode.data)
		bst.inorderTraversal(bstNode.right)
	}
}

func (bst *binarySearchTree) postorderTraversal(bstNode *node) {
	if bstNode != nil {
		bst.postorderTraversal(bstNode.left)
		bst.postorderTraversal(bstNode.right)
		fmt.Println(bstNode.data)
	}
}

func (bst *binarySearchTree) delete(value int) {
	if bst.root == nil {
		fmt.Println("Tree is empty")
		return
	}

	deleteNode(bst.root, value)
}

func deleteNode(bstNode *node, value int) *node {
	if value < bstNode.data {
		bstNode.left = deleteNode(bstNode.left, value)
	} else if value > bstNode.data {
		bstNode.right = deleteNode(bstNode.right, value)
	} else {
		if bstNode.left == nil {
			return bstNode.right
		} else if bstNode.right == nil {
			return bstNode.left
		} else {
			bstNode.data = findInorderSuccessor(bstNode.right)
			bstNode.right = deleteNode(bstNode.right, bstNode.data)
		}
	}
	return bstNode
}

func findInorderSuccessor(bstNode *node) int {
	for bstNode.left != nil {
		bstNode = bstNode.left
	}

	return bstNode.data
}

func (bst *binarySearchTree) levelOrderTraversal(){
	if bst.root == nil{
		return
	}

	var queue []*node
	queue = append(queue, bst.root)
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