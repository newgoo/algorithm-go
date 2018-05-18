package main

import (
	"github.com/newgoo/algorithm"
	"github.com/sirupsen/logrus"
)

/*
	插入排序
	特点: a.能最快发现是已经排好序的数组；b.对于实际应用中常见的某些类型的非随机数组很有效
	时间复杂度: N^2/4(平均情况),最坏N^2/2 (通常情况比选择排序快一倍)

	通常人们整理桥牌的方法是一张一张的来，将每一张牌插入到其他已经有序的牌中的适当位置。
	在计算机的实现中，为了给要插入的元素腾出空间，我们需要将其余所有元素在插入之前都向右移动一

	常见使用情况:
	a. 数组中每个元素距离它的最终位置都不远
	b.一个有序的大数组接一个小数组
	c.数组中只有几个元素的位置不正确
	d.插入排序对于部分有序的数组十分高效，也很适合小规模数组

	提高插入排序速度:只需要在内循环中将较大的元素都向右移动而不总是交换两个元素（这样访问数组的次数就能减半）

	命题：
		长度为N的数组，平均需要~N^2/4次比较和~N^2/4次交换，最坏情况~N^2/2次比较和~N^2/2次交换；最好情况为N-1此比较和0次交换
*/
func Insertion(list []int64) {
	logrus.Info(list)
	t := 0
	for i := 1; i < len(list); i++ {
		for j := 0; j < i; j++ {
			if list[i] < list[j] {
				list[i], list[j] = list[j], list[i]
				t++
			}
		}
	}
	logrus.Info(list)
	logrus.Info(t)
	logrus.Info(algorithm.ValidSorting(list))
}

func main() {
	list := algorithm.Random(100)
	Insertion(list)
}
