package noun

import "sync"

type noun struct {
	sync.RWMutex
	size, cap  int
	head, tail *item
	values     *segMap
}

type item struct {
	key        string
	value      interface{}
	prev, next *item
}

func NewNoun(size int) *noun {
	return &noun{
		cap:    size,
		values: NewSegMap(size),
	}
}

func (n *noun) Set(key string, value interface{}) {
	n.Lock()
	defer n.Unlock()
	iteM, ok := n.values.Get(key)
	if ok {
		iteM.value = value
		n.moveToFirst(iteM)
		return
	} else {
		if n.size == n.cap {
			n.values.Delete(n.tail.key)
			n.removeLast()
		} else {
			n.size++
		}
		item := &item{
			key:   key,
			value: value,
		}
		n.values.Put(item)
		n.insertToFirst(item)
	}
}

func (n *noun) Get(key string) interface{} {
	n.Lock()
	defer n.Unlock()
	if item, ok := n.values.Get(key); ok {
		n.moveToFirst(item)
		return item.value
	}
	return nil
}

func (n *noun) moveToFirst(node *item) {
	switch node {
	case n.head:
		return
	case n.tail:
		n.removeLast()
	default:
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	n.insertToFirst(node)
}

func (n *noun) insertToFirst(node *item) {
	if n.tail == nil {
		n.tail = node
	} else {
		node.next = n.head
		n.head.prev = node
	}
	n.head = node
}

func (n *noun) removeLast() {
	if n.tail.prev != nil {
		n.tail.prev.next = nil
	} else {
		n.head = nil
	}
	n.tail = n.tail.prev
}
