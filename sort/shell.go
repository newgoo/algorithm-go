package main

import (
	"github.com/newgoo/algorithm"
	"github.com/sirupsen/logrus"
)

/*
	希尔排序
		将数组分组，保证每一个组的对应数据都是排好序的数据
	特点:
		1.基于插入排序的快速排序，加快速度简单地改进了插入排序，交换不相邻的元素以对数组的局部进行排序，并最终用插入排序将局部有序的数组排序
		2.希尔排序比插入排序和选择排序要快得多，并且数组越大，优势越大
		3.不需要额外的内存
	思想: 思想是使数组中任意间隔为h 的元素都是有序的
	时间复杂度：平均每个增幅所带来的比较次数约为N^(1/5), 当h为3n+1时，最快
*/
func Shell(list []int64, steps ...int) {
	//h := 3
	for _, h := range steps {
		for i := 0; i < h; i++ {
			for m := h + i; m < len(list); m += h {
				for n := i; n < m; n += h {
					if list[m] < list[n] {
						list[m], list[n] = list[n], list[m]
					}
				}
			}
		}
	}

	logrus.Info(list)
	logrus.Info(algorithm.ValidSorting(list))
}

func main() {
	list := algorithm.Random(1000)
	//3n+1
	Shell(list, 364, 121, 40, 13, 4, 1)
}
