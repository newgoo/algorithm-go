package red_black_bst

import (
	"testing"

	"github.com/lunny/log"
)

type RDKey struct {
	Key int64
}

func (r RDKey) CompareTo(key Key) int64 {
	return r.Key - key.GetKey()
}

func (r RDKey) GetKey() int64 {
	return r.Key
}

func TestPut(t *testing.T) {
	root := InitRBRoot(RDKey{Key: 6}, 6)
	log.Info(root)

	root = Put(root, RDKey{Key: 4}, 4)
	root = Put(root, RDKey{Key: 3}, 3)

	log.Info(root.Right)

	log.Info(Get(root, RDKey{Key: 5}))

	log.Info(Max(root), Min(root))

	log.Info(Floor(root, RDKey{Key: 7}))

	log.Info(Ceiling(root, RDKey{Key: 4}))

	log.Info(Select(root, 0))

	log.Info(Rank(root, RDKey{Key: 5}), size(root))

}
