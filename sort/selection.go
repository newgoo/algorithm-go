package main

import (
	"github.com/newgoo/algorithm"
	"github.com/sirupsen/logrus"
)

/*
	选择排序是一种很容易理解和实现的简单排序算法
	特点: a.运行时间和输入无关; b.数据移动是最少的
	时间复杂度为N^2/2

	首先，找到数组中最小的那个元素，其次，将它和数组的第一个元素交换位置（如果第一个元素就是最小元素那么它就和自己交换）。
	再次，在剩下的元素中找到最小的元素，将它与数组的第二个元素交换位置。如此往复，直到将整个数组排序。
	这种方法叫做选择排序，因为它在不断地选择剩余元素之中的最小者。
*/
func Selection(list []int64) {
	logrus.Info(list)
	var min = 0
	for i := 0; i < len(list); i++ {
		min = i
		for j := i + 1; j < len(list); j++ {
			if list[min] > list[j] {
				min = j
			}
		}
		list[min], list[i] = list[i], list[min]
	}

	logrus.Info(list)
	logrus.Info(algorithm.ValidSorting(list))

}

func main() {
	list := algorithm.Random(100)
	Selection(list)
}
