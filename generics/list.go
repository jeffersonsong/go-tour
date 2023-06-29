package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func NewList[T any](val T) *List[T] {
	return &List[T]{nil, val}
}

func (l *List[T]) AddFirst(val T) *List[T] {
	return &List[T]{l, val}
}

func (l *List[T]) Add(val T) *List[T] {
	if l == nil {
		return &List[T]{nil, val}
	}
	curr := l
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = &List[T]{nil, val}
	return l
}

func (l *List[T]) RemoveFirst(filter func(val T) bool) (list *List[T], ok bool) {
	if l == nil {
		return l, false
	} else if filter(l.val) {
		return l.next, true
	}

	prev := l
	for curr := l.next; curr != nil; curr = curr.next {
		if filter(curr.val) {
			prev.next = curr.next
			return l, true
		}
		prev = curr
	}

	return l, false
}

func (l *List[T]) Size() int {
	size := 0
	for curr := l; curr != nil; curr = curr.next {
		size++
	}
	return size
}

func (l *List[T]) Get(index int) (val T, ok bool) {
	if l == nil {
		return val, false
	}

	curr := l
	for i := 0; curr != nil && i < index; i++ {
		curr = curr.next
	}

	if curr != nil {
		return curr.val, true
	} else {
		return val, false
	}
}

func (l *List[T]) Reverse() *List[T] {
	if l == nil {
		return l
	} else if l.next != nil {
		return l.next.Reverse().Add(l.val)
	} else {
		return &List[T]{nil, l.val}
	}
}

func (l *List[T]) String() string {
	switch {
	case l == nil:
		return ""
	case l.next == nil:
		return fmt.Sprintf("%v", l.val)
	default:
		return fmt.Sprintf("%v,%v", l.val, l.next)
	}
}

func (l *List[T]) ForEach(fn func(val T)) {
	for curr := l; curr != nil; curr = curr.next {
		fn(curr.val)
	}
}

func Map[T any, U any](l *List[T], transform func(val T) U) *List[U] {
	if l == nil {
		return nil
	}

	head2 := &List[U]{nil, transform(l.val)}
	tail2 := head2

	for tail := l.next; tail != nil; tail = tail.next {
		tail2.next = &List[U]{nil, transform(tail.val)}
		tail2 = tail2.next
	}

	return head2
}

func main() {
	list := NewList(1).Add(2).Add(3).Add(4)
	fmt.Println("list: ", list)
	fmt.Println("size: ", list.Size())
	val, ok := list.Get(0)
	if ok {
		fmt.Println("list[0] = ", val)
	} else {
		fmt.Println("list[0] not found")
	}

	fn := func(val int) { fmt.Print(val, ",") }
	fmt.Print("ForEach: ")
	list.ForEach(fn)
	fmt.Println()

	val, ok = list.Get(4)
	if ok {
		fmt.Println("list[4] = ", val)
	} else {
		fmt.Println("list[4] not found")
	}

	reversed := list.Reverse()
	fmt.Println("reversed: ", reversed)

	equals3 := func(val int) bool { return val == 3 }

	//Remove(list, 3)
	list, ok = list.RemoveFirst(equals3)
	fmt.Println("list: ", list)

	toString := func(val int) string { return fmt.Sprintf("\"%v\"", val) }
	strList := Map(list, toString)
	fmt.Println("strList: ", strList)
}
