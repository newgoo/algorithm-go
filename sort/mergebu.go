package main

//import (
//	"github.com/newgoo/algorithm"
//	"github.com/sirupsen/logrus"
//)
//
///*
//
// */
//func mg(left, right []int64) []int64 {
//	result := make([]int64, 0, len(left)+len(right))
//	for len(left) > 0 || len(right) > 0 {
//		if len(left) == 0 {
//			return append(result, right...)
//		}
//		if len(right) == 0 {
//			return append(result, left...)
//		}
//		if left[0] <= right[0] {
//			result = append(result, left[0])
//			left = left[1:]
//		} else {
//			result = append(result, right[0])
//			right = right[1:]
//		}
//	}
//	return result
//}
//
//func MergeBUSort(arr []int64) []int64 {
//	for i := 1; i < len(arr); i += i {
//		ls := make([]int64, 0)
//		for j := 0; j < len(arr)-i; j += 2 * i {
//			logrus.Info(i, j, arr[j:j+i], arr[j+i:j+2*i])
//			if j >= len(arr)/i {
//				ls = append(ls, mg(arr[j:j+i], arr[j+i:])...)
//			} else {
//				ls = append(ls, mg(arr[j:j+i], arr[j+i:j+2*i])...)
//			}
//		}
//		arr = ls
//		logrus.Info(ls)
//	}
//
//	return arr
//}
//
//func main() {
//	list := algorithm.Random(5)
//	logrus.Info(list)
//	ls := MergeBUSort(list)
//	logrus.Info(algorithm.ValidSorting(ls))
//	logrus.Info(ls)
//}
