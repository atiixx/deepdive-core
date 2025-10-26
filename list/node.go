package list

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}
