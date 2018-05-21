package binary_search

/*
	二叉查找树，每个节点又有一个可比较的key
		读取和写入的效率都是~2lnN 1.39logN,效率比较高
*/
type Node struct {
	Key         int64
	Value       interface{}
	N           int64 //子节点总数
	Left, right *Node
}

func InitRoot(key int64, value interface{}) *Node {
	return &Node{
		Key:   key,
		Value: value,
		N:     1,
	}
}

//获取key对应的值
func Get(node *Node, key int64) interface{} {
	if node == nil {
		return nil
	}
	if key < node.Key {
		return Get(node.Left, key)
	} else if key > node.Key {
		return Get(node.right, key)
	} else {
		return node.Value
	}
}

//将键值对存入对应值中
func Put(node *Node, key int64, value interface{}) *Node {
	if node == nil {
		return &Node{
			Key:   key,
			Value: value,
			N:     1,
		}
	}
	if key < node.Key {
		node.Left = Put(node.Left, key, value)
	} else if key > node.Key {
		node.right = Put(node.right, key, value)
	} else {
		node.Value = value
	}
	node.N = size(node.Left) + size(node.right) + 1
	return node
}

//最大的节点
func Max(node *Node) *Node {
	if node.right == nil {
		return node
	} else {
		return Max(node.right)
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

//小于key的最大节点
func Floor(node *Node, key int64) *Node {
	if node == nil {
		return nil
	}
	if key == node.Key {
		return node
	}
	if key < node.Key {
		return Floor(node.Left, key)
	}
	if t := Floor(node.right, key); t != nil {
		return t
	}
	return node
}

//大于key的最大节点
func Ceiling(node *Node, key int64) *Node {
	if node == nil {
		return nil
	}
	if key == node.Key {
		return node
	}
	if key > node.Key {
		return Ceiling(node.right, key)
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
		return Select(node.right, t-size(node.Left)-1) //1为根节点
	} else {
		return node
	}
}

//计算某节点的排名(前序遍历--数组是有序的)
func Rank(node *Node, key int64) int64 {
	if node == nil {
		return 0
	}
	if key < node.Key {
		return Rank(node.Left, key)
	} else if key > node.Key {
		return size(node.Left) + 1 + Rank(node.right, key) //1为根节点
	} else {
		return size(node.Left)
	}
}

func Delete(node *Node, key int64) *Node {
	if node == nil {
		return nil
	}
	if key < node.Key {
		node.Left = Delete(node.Left, key)
	} else if key > node.Key {
		node.right = Delete(node.right, key)
	} else {
		if node.Left == nil {
			return node.right
		}
		if node.right == nil {
			return node.Left
		}
		t := node
		node = Min(t.right)
		node.right = DeleteMin(t.right)
		node.Left = t.Left
	}
	node.N = size(node.Left) + size(node.right) + 1
	return node
}

func DeleteMin(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node.right
	} else {
		node.Left = DeleteMin(node.Left)
	}
	node.N = size(node.Left) + size(node.right) + 1
	return node
}

func DeleteMax(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node.Left
	} else {
		node.right = DeleteMax(node.right)
	}
	node.N = size(node.Left) + size(node.right) + 1
	return node
}

//从key到key包含两边的所有key
func Keys(node *Node, keys []int64, low, high int64) []int64 {
	if node == nil {
		return keys
	}
	if low < node.Key {
		keys = Keys(node.Left, keys, low, high)
	}

	if low <= node.Key && high >= node.Key {
		keys = append(keys, node.Key)
	}

	if high > node.Key {
		keys = Keys(node.right, keys, low, high)
	}
	return keys
}

func size(node *Node) int64 {
	if node == nil {
		return 0
	} else {
		return node.N
	}
}
