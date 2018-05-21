package binary_search

import (
	"testing"

	"github.com/lunny/log"
)

func TestBinaryTree_Get(t *testing.T) {
	root := InitRoot(3, 3)

	root = Put(root, 2, 2)
	root = Put(root, 3, 3)
	root = Put(root, 5, 5)
	root = Put(root, 4, 4)
	root = Put(root, 6, 6)
	log.Info(root, root.Left, root.right)

	log.Info("通过key=1002获取：", Get(root, 1002), root)

	log.Info("最大节点：", Max(root), root)

	log.Info("最小节点：", Min(root), root)

	log.Info("小于key的最大节点(floor)", Floor(root, 1006), root)

	log.Info("大于key的最大节点(ceiling)", Ceiling(root, 1002), root)

	//数据为[2,3,4,5,6]从0开始索引
	log.Info("排名为n的节点(select)", Select(root, 3), root)

	log.Info("key为n的节点的排名(rank)", Rank(root, 7), root)

	//前序遍历
	log.Info("keys链表中指定范围的key，包含low和high", Keys(root, []int64{}, 0, 6))

	//root = DeleteMin(root)
	//log.Info("删除最小节点(deletemin)", Keys(root, []int64{}, 0, 6), root)
	root = DeleteMin(root)
	log.Info("删除最小节点(deletemin)", Keys(root, []int64{}, 0, 6), root)

	root = DeleteMax(root)
	log.Info("删除最小节点(deletemax)", Keys(root, []int64{}, 0, 6), root)

	root = Put(root, 2, 2)
	root = Put(root, 6, 6)

	log.Info(Keys(root, []int64{}, 0, 6))
	root = Delete(root, 5)
	log.Info("删除节点(delete)", Keys(root, []int64{}, 0, 6), root)
}
