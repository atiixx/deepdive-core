package list

import "fmt"

type List[T comparable] struct {
	head *Node[T]
}

func (l *List[T]) Append(value T) {
	if l.head == nil {
		l.head = &Node[T]{Value: value}
		return
	}
	current := l.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node[T]{Value: value}
}

func (l *List[T]) Prepend(value T) {
	if l.head == nil {
		l.head = &Node[T]{Value: value}
		return
	}
	new := &Node[T]{Value: value, Next: l.head}
	l.head = new
}

func (l *List[T]) Print() {
	for current := l.head; current != nil; current = current.Next {
		fmt.Print(current.Value)
		if current.Next != nil {
			fmt.Print("-")
		}
	}
	fmt.Println()
}

func (l *List[T]) Remove(key T) bool {
	if l.head == nil {
		return false
	}
	if l.head.Value == key {
		l.head = l.head.Next
		return true
	}
	prev := l.head
	for curr := l.head.Next; curr != nil; curr = curr.Next {
		if curr.Value == key {
			prev.Next = curr.Next
			return true
		}
		prev = curr
	}
	return false
}

func (l List[T]) Find(key T) int {
	current := l.head
	for i := 0; current != nil; i++ {
		if current.Value == key {
			return i
		}
		current = current.Next
	}
	return -1
}
