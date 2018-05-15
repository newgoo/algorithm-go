package main

import (
	"github.com/lunny/log"
)

/*
	优先队列排序(二叉堆排序)
		删除最大(最小元素),和插入元素
		通常我们在处理有优先级蓄力的任务时，我们需要在处理完一个任务，或者添加一个任务，都需要去寻找下一个需要处理的任务，这是使用优先队列排序算法实现
	特点：1.使用的是完全二叉树的结构；
	时间复杂度：插入logN, 删除logN

	应用场景：
		1，对有有优先级处理的逻辑，的增加或者删除，不需要对数据进行完全排序
		2. 在数据量很大的情况，但又不考虑大部分数据的情况(如需要在一个很大数据中取出最大十个，我们将这十个数据存储下来后，在检测后面数据，查看是否需要更改里面的内容)
*/

//自底向上堆有序化(上浮) ---插入元素
func swim(arr []int64, k int) {
	for k > 1 && arr[k/2] < arr[k] {
		arr[k/2], arr[k] = arr[k], arr[k/2]
		k /= 2
	}
}

//至上向下堆有序化(下沉) ---删除元素
func sink(arr []int64, k int) {
	var j = 0
	for 2*k < len(arr) {
		j = 2 * k
		if j < len(arr)-1 && arr[j] < arr[j+1] {
			j++
		}
		if arr[j] <= arr[k] {
			break
		}
		arr[j], arr[k] = arr[k], arr[j]
		k = j
	}
}

func Del(arr []int64) []int64 {
	//max := arr[1]
	arr[1] = arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	sink(arr, 1)
	return arr
}

func Insert(arr []int64, v int64) []int64 {
	arr = append(arr, v) //数组地址发生了变化
	swim(arr, len(arr)-1)
	return arr
}

func main() {
	arr := []int64{0, 9, 8, 4, 4, 3, 2, 1}
	log.Info("原数组：", arr)
	arr = Insert(arr, 5)
	log.Info(arr)
	arr = Del(arr) //数组地址发生了变化
	arr = Del(arr)
	log.Info(arr)
}
