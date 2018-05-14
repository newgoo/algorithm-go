package main

import (
	"github.com/newgoo/algorithm"
	"github.com/sirupsen/logrus"
)

/*
	快速排序(切分)
		将数组切分成两个部分，其中左边部分比这个数小，右边部分比这个数大，经过递归，两边数据分别排好序，数组就自动排好序了
		特点: 使用最广泛的排序方式(单个算法最快排序算法)
		改进：
			1. 切换到插入排序(通常是5-15)
			2. 三取样切分：(待排序数组总随机取其中三个元素，取值为中位数的元素作为切分元素)；代价为需要每次计算中位数
			3. 嫡最优排序：假设某一个数组里面的数据全部相同，快排还是会一次排下去; 因此，将数组元素拆分成三个部分，low，middle，high三个部分，其中，low和high排序
		时间复杂度: NlogN

*/
func portion(arr []int64, low, high int) (l int, h int) {
	temp := arr[low]
	for low < high {
		//从右往左，找到一个比temp小的元素
		for low < high && temp <= arr[high] {
			high--
		}
		arr[low] = arr[high]

		//从左往右，找到一个比temp大的元素
		for low < high && temp >= arr[low] {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = temp
	//改进三
	for ; low >= 0 && arr[low] == temp; low-- {
		l = low
	}
	for ; high < len(arr)-1 && arr[high] == temp; high++ {
		h = high
	}
	return l, h
}

func QuickSort(arr []int64, low, high int) {
	if low < high {
		l, h := portion(arr, low, high)
		QuickSort(arr, low, l-1)
		QuickSort(arr, h+1, high)
	}
	return
}

func main() {
	//list := algorithm.Random(102)
	list := []int64{3, 1, 1, 1, 2}
	logrus.Info(list)
	QuickSort(list, 0, len(list)-1)
	logrus.Info(algorithm.ValidSorting(list))
	logrus.Info(list)
}
