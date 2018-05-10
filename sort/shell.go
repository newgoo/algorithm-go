package main

import (
	"github.com/newgoo/algorithm"
	"github.com/sirupsen/logrus"
)

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
	Shell(list, 16, 8, 2, 1)
}
