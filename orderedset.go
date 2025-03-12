package oops

type orderedSet[T comparable] struct {
	elements []T
	set      map[T]struct{}
}

func newOrderedSet[T comparable]() *orderedSet[T] {
	return &orderedSet[T]{
		elements: make([]T, 0),
		set:      make(map[T]struct{}),
	}
}

func (o *orderedSet[T]) Add(value T) {
	if _, exists := o.set[value]; !exists {
		o.set[value] = struct{}{}
		o.elements = append(o.elements, value)
	}
}

func (o *orderedSet[T]) Len() int {
	return len(o.elements)
}

func (o orderedSet[T]) GetOrdered() []T {
	return o.elements
}

func (o orderedSet[T]) GetReverseOrdered() []T {
	n := len(o.elements)
	reverse := make([]T, n)
	for i, elem := range o.elements {
		reverse[n-i-1] = elem
	}
	return reverse
}
