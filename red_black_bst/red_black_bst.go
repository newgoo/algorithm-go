package red_black_bst

type Key interface {
	CompareTo(key Key) int
}

const (
	Red   = true
	Black = false
)

type Node struct {
	Key         Key
	Value       interface{}
	Left, Right *Node
	N           int64
	Color       bool
}

func Put(node *Node, key Key, value interface{}) *Node {
	if node == nil {
		return &Node{
			Key:   key,
			Value: value,
			N:     1,
			Color: Red,
		}
	}

	cmp := key.CompareTo(node.Key)
	if cmp < 0 {
		node.Left = Put(node.Left, key, value)
	} else if cmp > 0 {
		node.Right = Put(node.Right, key, value)
	} else {
		node.Value = value
	}

	if isRed(node.Left) && isRed(node.Left.Left) {
		node = retateRight(node)
	}
	if isRed(node.Right) && !isRed(node.Left) {
		node = retateLeft(node)
	}
	if isRed(node.Right) && isRed(node.Left) {
		flipColors(node)
	}
	node.N = size(node.Left) + size(node.Right) + 1
	return node
}

func retateLeft(node *Node) *Node {
	x := node.Right
	x.Left = node
	x.Color = node.Color
	node.Color = Red
	x.N = node.N
	node.N = 1 + size(node.Left) + size(node.Right)
	return x
}

func retateRight(node *Node) *Node {
	x := node.Left
	node.Left = x.Right
	x.Right = node
	x.Color = node.Color
	node.Color = Red
	x.N = node.N
	node.N = 1 + size(node.Left) + size(node.Right)
	return x
}

func flipColors(node *Node) {
	node.Color = Red
	node.Left.Color = Black
	node.Left.Color = Black
}

func isRed(node *Node) bool {
	if node == nil {
		return false
	}
	return node.Color == Red
}

func size(node *Node) int64 {
	if node == nil {
		return 0
	}
	return node.N
}
