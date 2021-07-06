package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	List
	// Place your code here.
	items []ListItem
}

func (l *list) Len() int {
	return len(l.items)
}

func (l *list) Front() *ListItem {
	return &l.items[0]
}

func (l *list) Back() *ListItem {
	last := len(l.items) - 1
	return &l.items[last]
}

func NewList() List {
	return new(list)
}
