package algorithm

import (
	"math/rand"
	"time"
)

func Random(len int64) []int64 {
	list := make([]int64, len)
	for index := range list {
		list[index] = rand.Int63n(time.Now().Unix())
	}
	return list
}

func ValidSorting(list []int64) bool {
	for i := 0; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			if list[i] > list[j] {
				return false
			}
		}
	}
	return true
}
