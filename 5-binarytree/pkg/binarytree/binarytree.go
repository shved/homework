// Package binarytree implements a binary tree data structure.
package binarytree

import "websearch/pkg/resource"

// Tree - binary search tree
type Tree struct {
	root      *Node
	currentID uint
}

// Node - a tree node
type Node struct {
	left, right *Node
	Value       *resource.Document
}

// Insert a web document into the tree. Returns ID assigned to this document.
func (t *Tree) Insert(doc *resource.Document) uint {
	doc.ID = t.currentID
	t.currentID++

	n := &Node{Value: doc}
	if t.root == nil {
		t.root = n
		return doc.ID
	}

	return t.insert(t.root, n)
}

func (t *Tree) insert(node, new *Node) uint {
	if new.Value.ID < node.Value.ID {
		if node.left == nil {
			node.left = new
			return new.Value.ID
		}

		return t.insert(node.left, new)
	}

	if node.right == nil {
		node.right = new
		return new.Value.ID
	}

	return t.insert(node.right, new)
}

// Find - find a web document in a tree by its ID
func (t *Tree) Find(id uint) *resource.Document {
	return search(t.root, id)
}

func search(n *Node, id uint) *resource.Document {
	if n == nil {
		return nil
	}

	if n.Value.ID == id {
		return n.Value
	}

	if n.Value.ID < id {
		return search(n.right, id)
	}

	return search(n.left, id)
}
