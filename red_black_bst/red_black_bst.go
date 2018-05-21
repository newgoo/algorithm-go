package red_black_bst

type Key interface {
	CompareTo(key Key) int64
	GetKey() int64
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

func InitRBRoot(key Key, value interface{}) *Node {
	return &Node{
		Key:   key,
		Value: value,
		N:     1,
		Color: Black,
	}
}

//小于key的最大节点
func Floor(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}

	cmp := key.CompareTo(node.Key)
	if cmp == 0 {
		return node
	}
	if cmp < 0 {
		return Floor(node.Left, key)
	}
	if t := Floor(node.Right, key); t != nil {
		return t
	}
	return node
}

//大于key的最大节点
func Ceiling(node *Node, key Key) *Node {
	if node == nil {
		return nil
	}

	cmp := key.CompareTo(node.Key)
	if cmp == 0 {
		return node
	}
	if cmp > 0 {
		return Ceiling(node.Right, key)
	}
	if t := Ceiling(node.Left, key); t != nil {
		return t
	}
	return node
}

//计算某排名的节点(前序遍历--数组是有序的)
func Select(node *Node, t int64) *Node {
	if node == nil {
		return nil
	}
	if t < size(node.Left) {
		return Select(node.Left, t)
	} else if t > size(node.Left) {
		return Select(node.Right, t-size(node.Left)-1) //1为根节点
	} else {
		return node
	}
}

//计算某节点的排名(前序遍历--数组是有序的)
func Rank(node *Node, key Key) int64 {
	if node == nil {
		return -1
	}

	cmp := key.CompareTo(node.Key)
	if cmp < 0 {
		r := Rank(node.Left, key)
		if r == -1 {
			return -1
		} else {
			return r
		}

	} else if cmp > 0 {
		r := Rank(node.Right, key)
		if r == -1 {
			return -1
		} else {
			return size(node.Left) + 1 + r //1为根节点
		}
	} else {
		return size(node.Left)
	}
}

//最大的节点
func Max(node *Node) *Node {
	if node.Right == nil {
		return node
	} else {
		return Max(node.Right)
	}
}

//最小的节点
func Min(node *Node) *Node {
	if node.Left == nil {
		return node
	} else {
		return Min(node.Left)
	}
}

func Get(root *Node, key Key) interface{} {
	if root == nil {
		return nil
	}
	cmp := key.CompareTo(root.Key)
	if cmp > 0 {
		return Get(root.Right, key)
	} else if cmp < 0 {
		return Get(root.Left, key)
	} else {
		return root.Value
	}
}

func Put(root *Node, key Key, value interface{}) *Node {
	root = put(root, key, value)
	root.Color = Black
	return root
}

func put(node *Node, key Key, value interface{}) *Node {
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
		node.Left = put(node.Left, key, value)
	} else if cmp > 0 {
		node.Right = put(node.Right, key, value)
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
	node.Right = x.Left
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
