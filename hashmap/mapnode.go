package hashmap

type MapNode[V any] struct {
	Key    string
	Value  V
	active bool
}

func NewMapNode[V any]() MapNode[V] {
	return MapNode[V]{active: true}
}
