package estruturadedados

type LinkedList[T any] struct {
	Value T
	Prev  *LinkedList[T]
	Next  *LinkedList[T]
}

func getTail[T any](node *LinkedList[T]) *LinkedList[T] {
	if node.Next == nil {
		return node
	}

	return getTail(node.Next)
}

func (l *LinkedList[T]) Insert(element T) {
	tail := getTail(l)

	newNode := LinkedList[T]{
		Value: element,
		Prev:  nil,
		Next:  nil,
	}

	tail.Next = &newNode
}

func handleSearch[T any](
	node *LinkedList[T], element *any,
) (bool, *LinkedList[T]) {
	if node == nil {
		var emptyNode = LinkedList[T]{}
		return false, &emptyNode
	}

	var value any = node.Value

	if value == *element {
		return true, node
	}

	return handleSearch(node.Next, element)
}

func (l *LinkedList[T]) Search(element T) (bool, *LinkedList[T]) {
	var elem any = element
	return handleSearch(l, &elem)
}

func getParentNode[T any](
	node *LinkedList[T], element *any,
) (bool, *LinkedList[T]) {
	if node.Next == nil {
		var emptyNode = LinkedList[T]{}
		return false, &emptyNode
	}

	var nextValue any = node.Next.Value

	if *element == nextValue {
		return true, node
	}

	return getParentNode(node.Next, element)
}

func (l *LinkedList[T]) Exclude(element T) bool {
	ok, node := l.Search(element)

	if !ok {
		return false
	}

	var elem any = element

	ok, parentNode := getParentNode(l, &elem)

	if !ok {
		return false
	}

	parentNode.Next = node.Next
	return true
}
