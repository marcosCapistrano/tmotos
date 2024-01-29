package memory

import (
	"fmt"
	"strings"
)

type ProductNode struct {
	ID       int
	Name     string
	Parent   *ProductNode
	Children []*ProductNode
	URL      string
}

type ProductTree struct {
	root *ProductNode
}

func NewProductTree() ProductTree {
	root := &ProductNode{
		Name:     "Todos",
		Parent:   nil,
		Children: make([]*ProductNode, 0),
		URL:      "/",
	}

	return ProductTree{
		root: root,
	}
}

func (node *ProductNode) Display(spacing string) {
	fmt.Println(spacing + node.Name)

	spacing = spacing + "\t"
	for _, child := range node.Children {
		child.Display(spacing)
	}
}

func (tree ProductTree) Display() {
	tree.root.Display("")
}

func (node *ProductNode) Insert(name, url string, isProduct bool) *ProductNode {
	var children []*ProductNode

	if !isProduct {
		children = make([]*ProductNode, 0)
	}

	newNode := &ProductNode{
		Name:     name,
		Parent:   node,
		Children: children,
		URL:      url,
	}

	node.Children = append(node.Children, newNode)

	return newNode
}

func (tree ProductTree) Insert(name, url string, isProduct bool) *ProductNode {
	return tree.root.Insert(name, url, isProduct)
}

func (node *ProductNode) FindSubtree(url string) *ProductNode {
	if url == "" {
		return node
	}

	splitURL := strings.Split(url, "/")
	nextCategory := splitURL[0]

	for _, child := range node.Children {
		if strings.Compare(child.URL, nextCategory) == 0 {
			return child.FindSubtree(strings.Join(splitURL[1:], "/"))
		}
	}

	return node
}

func (tree ProductTree) FindSubtree(url string) *ProductNode {
	root := tree.root

	return root.FindSubtree(url)
}

func (node *ProductNode) CollectLeaves() []*ProductNode {
	if node == nil {
		return nil
	}

	var leaves []*ProductNode
	stack := []*ProductNode{node}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(node.Children) == 0 {
			// It's a leaf node
			leaves = append(leaves, node)
		} else {
			// Push children to the stack
			stack = append(stack, node.Children...)
		}
	}

	return leaves
}
