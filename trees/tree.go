package trees

import "fmt"

// TreeNode represents a node in the binary tree.
type TreeNode struct {
    Value int
    Left  *TreeNode
    Right *TreeNode
}

// BinaryTree represents a binary tree with a reference to the root node.
type BinaryTree struct {
    Root *TreeNode
}

// NewBinaryTree creates a new instance of a binary tree.
func NewBinaryTree() *BinaryTree {
    return &BinaryTree{}
}

// Insert adds a new element to the binary tree.
func (bt *BinaryTree) Insert(value int) {
    newNode := &TreeNode{Value: value}
    if bt.Root == nil {
        bt.Root = newNode
    } else {
        insertNode(bt.Root, newNode)
    }
}

// insertNode places a new node in the binary tree.
func insertNode(node, newNode *TreeNode) {
    if newNode.Value < node.Value {
        if node.Left == nil {
            node.Left = newNode
        } else {
            insertNode(node.Left, newNode)
        }
    } else {
        if node.Right == nil {
            node.Right = newNode
        } else {
            insertNode(node.Right, newNode)
        }
    }
}

// InOrderTraversal traverses the tree in in-order sequence and prints the values.
func (bt *BinaryTree) InOrderTraversal() {
    inOrderTraversal(bt.Root)
}

// inOrderTraversal is a recursive helper function for InOrderTraversal.
func inOrderTraversal(node *TreeNode) {
    if node != nil {
        inOrderTraversal(node.Left)
        fmt.Print(node.Value, " ")
        inOrderTraversal(node.Right)
    }
}

// PreOrderTraversal traverses the tree in pre-order sequence and prints the values.
func (bt *BinaryTree) PreOrderTraversal() {
    preOrderTraversal(bt.Root)
}

// preOrderTraversal is a recursive helper function for PreOrderTraversal.
func preOrderTraversal(node *TreeNode) {
    if node != nil {
        fmt.Print(node.Value, " ")
        preOrderTraversal(node.Left)
        preOrderTraversal(node.Right)
    }
}

// PostOrderTraversal traverses the tree in post-order sequence and prints the values.
func (bt *BinaryTree) PostOrderTraversal() {
    postOrderTraversal(bt.Root)
}

// postOrderTraversal is a recursive helper function for PostOrderTraversal.
func postOrderTraversal(node *TreeNode) {
    if node != nil {
        postOrderTraversal(node.Left)
        postOrderTraversal(node.Right)
        fmt.Print(node.Value, " ")
    }
}

// PrettyPrint prints the tree in a visually appealing format
func (bt *BinaryTree) PrettyPrint() {
    levels := make([][]int, 0)
    fillLevels(bt.Root, 0, &levels)
    for _, level := range levels {
        for _, val := range level {
            if val == -1 {
                fmt.Print("nil ")
            } else {
                fmt.Printf("%d ", val)
            }
        }
        fmt.Println()
    }
}

// fillLevels fills the slices with the values at each level of the tree
func fillLevels(node *TreeNode, level int, levels *[][]int) {
    if len(*levels) == level {
        *levels = append(*levels, []int{})
    }

    if node == nil {
        (*levels)[level] = append((*levels)[level], -1)
        return
    }

    (*levels)[level] = append((*levels)[level], node.Value)
    fillLevels(node.Left, level+1, levels)
    fillLevels(node.Right, level+1, levels)
}
