package red_black_bst

type RDKey struct {
	Key int64
}

func (r RDKey) CompareTo(key RDKey) int64 {
	return r.Key - key.Key
}
