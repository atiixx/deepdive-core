package list

import "fmt"

type List[T comparable] struct {
	Head *Node[T]
}

func (l *List[T]) Append(value T) {
	if l.Head == nil {
		l.Head = &Node[T]{Value: value}
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node[T]{Value: value}
}

func (l *List[T]) Prepend(value T) {
	if l.Head == nil {
		l.Head = &Node[T]{Value: value}
		return
	}
	new := &Node[T]{Value: value, Next: l.Head}
	l.Head = new
}

func (l List[T]) String() string {
	result := ""
	for current := l.Head; current != nil; current = current.Next {
		result += fmt.Sprintf("%v", current.Value)
		if current.Next != nil {
			result += "-"
		}
	}
	return result
}

func (l *List[T]) Remove(key T) bool {
	if l.Head == nil {
		return false
	}
	if l.Head.Value == key {
		l.Head = l.Head.Next
		return true
	}
	prev := l.Head
	for curr := l.Head.Next; curr != nil; curr = curr.Next {
		if curr.Value == key {
			prev.Next = curr.Next
			return true
		}
		prev = curr
	}
	return false
}

func (l *List[T]) RemoveAt(indx int) T {
	if l.Head == nil || l.Size() <= indx {
		var zero T
		return zero
	}
	if indx == 0 {
		removedNode := l.Head
		l.Head = l.Head.Next
		return removedNode.Value
	}
	curr := l.Head
	for i := 0; i != indx-1; i++ {
		curr = curr.Next
	}
	removedNode := curr.Next
	curr.Next = curr.Next.Next
	return removedNode.Value
}

func (l List[T]) Size() int {
	size := 0
	for curr := l.Head; curr != nil; curr = curr.Next {
		size += 1
	}
	return size
}

func (l List[T]) Find(key T) int {
	current := l.Head
	for i := 0; current != nil; i++ {
		if current.Value == key {
			return i
		}
		current = current.Next
	}
	return -1
}
