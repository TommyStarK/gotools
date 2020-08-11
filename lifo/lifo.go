package lifo

import "sync"

// const defaultLifoCapacity = 1 << 6

// // Lifo represents the LIFO
// type Lifo struct {
// 	items []interface{}

// 	mutex sync.RWMutex
// }

// // NewLifo creates a new empty LIFO
// func NewLifo() *Lifo {
// 	return &Lifo{
// 		items: make([]interface{}, 0, defaultLifoCapacity),
// 		mutex: sync.RWMutex{},
// 	}
// }

// // Dequeue returns the last element of the FIFO or nil
// func (l *Lifo) Dequeue() interface{} {
// 	l.mutex.Lock()
// 	defer l.mutex.Unlock()

// 	if len(l.items) == 0 {
// 		return nil
// 	}

// 	data := l.items[len(l.items)-1]
// 	l.items = l.items[:len(l.items)-1]
// 	return data
// }

// // Enqueue allows to enqueue some data
// func (l *Lifo) Enqueue(data interface{}) {
// 	l.mutex.Lock()
// 	l.items = append(l.items, data)
// 	l.mutex.Unlock()
// 	return
// }

// // IsEmpty returns whether the FIFO is empty or not
// func (l *Lifo) IsEmpty() (empty bool) {
// 	l.mutex.RLock()
// 	empty = len(l.items) == 0
// 	l.mutex.RUnlock()
// 	return
// }

// // Size returns the actual FIFO size
// func (l *Lifo) Size() (size int) {
// 	l.mutex.RLock()
// 	size = len(l.items)
// 	l.mutex.RUnlock()
// 	return
// }

type element struct {
	data interface{}
	prev *element
	next *element
}

// Lifo represents the LIFO
type Lifo struct {
	size int
	head *element
	tail *element

	mutex sync.RWMutex
}

// NewLifo creates a new empty LIFO
func NewLifo() *Lifo {
	return &Lifo{
		size:  0,
		head:  nil,
		tail:  nil,
		mutex: sync.RWMutex{},
	}
}

// Dequeue returns the last element of the FIFO or nil
func (l *Lifo) Dequeue() interface{} {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.tail == nil {
		return nil
	}

	l.size--
	tail := l.tail
	if tail.prev == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = tail.prev
	}

	return tail.data
}

// Enqueue allows to enqueue some data
func (l *Lifo) Enqueue(data interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.size++
	var element = &element{
		data: data,
	}

	if l.head == nil && l.tail == nil {
		l.head = element
		l.tail = element
		return
	}

	last := l.tail
	last.next = element
	element.prev = last
	l.tail = element
	return
}

// IsEmpty returns whether the LIFO is empty or not
func (l *Lifo) IsEmpty() (empty bool) {
	l.mutex.RLock()
	empty = l.size == 0
	l.mutex.RUnlock()
	return
}

// Size returns the actual LIFO size
func (l *Lifo) Size() (size int) {
	l.mutex.RLock()
	size = l.size
	l.mutex.RUnlock()
	return
}
