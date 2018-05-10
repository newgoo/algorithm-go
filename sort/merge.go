package main

import (
	"github.com/newgoo/algorithm"
	"github.com/sirupsen/logrus"
)

/*
	自顶向下的归并排序
	将数组分成若干个数组，分组排序后再合并(二叉树结构)
	特点：a.速度比较快; b.占用内存比较大
	时间复杂度: 1/2NlgN到NlgN
	改进：
		a. 小规模的子数组采用插入排序；
		b. 在判断中如果left的最大值，已经小于right的最大值，直接合并
		c. 组合使用插入排序等排序，避免产生新的数组

*/

func merge(left, right []int64) []int64 {
	result := make([]int64, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

func MergeSort(arr []int64) []int64 {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2

	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	//改进
	if len(left) > 0 && len(right) > 0 && left[len(left)-1] < right[0] {
		return append(left, right...)
	}

	return merge(left, right)
}

func main() {
	list := algorithm.Random(100000)
	logrus.Info(list)
	ls := MergeSort(list)
	logrus.Info(algorithm.ValidSorting(ls))
	logrus.Info(ls)
}

//1576664
//1636271
