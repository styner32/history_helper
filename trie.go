package history_helper

type Node struct {
	Value  rune
	Depth  int
	Parent *Node
	Next   *Node
	Prev   *Node
	Child  *Node
}

func CreateTrie() *Node {
	return &Node{Value: 0, Parent: nil, Next: nil, Child: nil}
}

func (node *Node) AddWord(word string) {
	currentNode := node.FindCommonParentOf(word)

	for _, value := range word[currentNode.Depth:len(word)] {
		if currentNode.Child == nil {
			currentNode = currentNode.CreateChild(value)
		} else {
			currentNode = currentNode.Child.CreateSibling(value)
		}
	}
}

func (node *Node) CreateChild(value rune) *Node {
	child := &Node{Value: value, Depth: node.Depth + 1, Parent: node, Prev: nil, Next: nil, Child: nil}
	node.Child = child
	return child
}

func (node *Node) CreateSibling(value rune) *Node {
	sibling := &Node{Value: value, Depth: node.Depth, Parent: node.Parent, Prev: node, Next: nil, Child: nil}
	node.Next = sibling
	return sibling
}

func (node *Node) FindCommonParentOf(word string) *Node {
	currentNode := node

	for _, value := range word {
		for child := currentNode.Child; child != nil; child = child.Next {
			if child.Value == value {
				currentNode = child
				break
			}
		}
		if currentNode.Value != value {
			return currentNode
		}
	}
	return currentNode
}
