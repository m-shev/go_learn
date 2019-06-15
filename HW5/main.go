package main

import "fmt"

//List is doubly linked list
type List struct {
	length uint
	first  *Item
	last   *Item
}

//PushFront adds an item to the end of the list
func (list *List) PushFront(value interface{}) {
	itemToSet := &Item{value, nil, nil, list}

	switch list.length {
	case 0:
		list.first = itemToSet
		list.last = itemToSet
	case 1:
		itemToSet.prev = list.first
		list.first.next = itemToSet
		list.last = itemToSet
	default:
		temp := list.last
		list.last = &Item{value, temp, nil, list}
		temp.next = list.last
	}

	list.length++
}

//PushBack adds an item to the start of the list
func (list *List) PushBack(value interface{}) {
	itemToSet := &Item{value, nil, nil, list}

	switch list.length {
	case 0:
		list.first = itemToSet
		list.last = list.first

	case 1:
		itemToSet.next = list.first
		list.last = list.first
		list.first = itemToSet

	default:
		itemToSet.next = list.first
		itemToSet.next.prev = itemToSet
		list.first = itemToSet
	}

	list.length++
}

//Len returns size of doubly linked list
func (list *List) Len() uint {
	return list.length
}

//First returns first Item
func (list *List) First() *Item {
	return list.first
}

//Last returns last Item
func (list *List) Last() *Item {
	return list.last
}

func main() {
	list := List{}

	//list.PushBack(1)
	//list.PushBack(2)
	//list.PushBack(3)
	//list.PushBack(4)

	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	list.PushFront(4)

	last := list.last
	fmt.Println("remove", last.prev.Value())
	fmt.Println("remove", last.next)
	last.Remove()

	fmt.Println(list.Len())
	current := list.First()

	for current != nil {
		fmt.Println(current.Value())
		current = current.next
	}
}
