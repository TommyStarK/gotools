package fifo

import (
	"sync"
)

type element struct {
	data interface{}
	next *element
}

// Fifo represents the FIFO
type Fifo struct {
	size int
	head *element
	tail *element

	mutex sync.RWMutex
}

// NewFifo creates a new empty FIFO
func NewFifo() *Fifo {
	return &Fifo{
		size:  0,
		head:  nil,
		tail:  nil,
		mutex: sync.RWMutex{},
	}
}

// Dequeue returns the first element of the FIFO or nil
func (f *Fifo) Dequeue() interface{} {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if f.head == nil {
		return nil
	}

	f.size--
	head := f.head
	if head.next == nil {
		f.head = nil
		f.tail = nil
	} else {
		f.head = head.next
	}

	return head.data
}

// Enqueue allows to enqueue some data
func (f *Fifo) Enqueue(data interface{}) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	f.size++
	var element = &element{
		data: data,
	}

	if f.head == nil && f.tail == nil {
		f.head = element
		f.tail = element
		return
	}

	last := f.tail
	last.next = element
	f.tail = element
	return
}

// IsEmpty returns whether the FIFO is empty or not
func (f *Fifo) IsEmpty() (empty bool) {
	f.mutex.RLock()
	empty = f.size == 0
	f.mutex.RUnlock()
	return
}

// Size returns the actual FIFO size
func (f *Fifo) Size() (size int) {
	f.mutex.RLock()
	size = f.size
	f.mutex.RUnlock()
	return
}

// const defaulFifoCapacity = 1 << 6

// // Fifo represents the FIFO
// type Fifo struct {
// 	items []interface{}

// 	mutex sync.RWMutex
// }

// // NewFifo creates a new empty FIFO
// func NewFifo() *Fifo {
// 	return &Fifo{
// 		items: make([]interface{}, 0, defaulFifoCapacity),
// 		mutex: sync.RWMutex{},
// 	}
// }

// // Dequeue returns the first element of the FIFO or nil
// func (f *Fifo) Dequeue() interface{} {
// 	f.mutex.Lock()
// 	defer f.mutex.Unlock()

// 	if len(f.items) == 0 {
// 		return nil
// 	}

// 	data := f.items[0]
// 	f.items = f.items[1:]
// 	return data
// }

// // Enqueue allows to enqueue some data
// func (f *Fifo) Enqueue(data interface{}) {
// 	f.mutex.Lock()
// 	f.items = append(f.items, data)
// 	f.mutex.Unlock()
// 	return
// }

// // IsEmpty returns whether the FIFO is empty or not
// func (f *Fifo) IsEmpty() (empty bool) {
// 	f.mutex.RLock()
// 	empty = len(f.items) == 0
// 	f.mutex.RUnlock()
// 	return
// }

// // Size returns the actual FIFO size
// func (f *Fifo) Size() (size int) {
// 	f.mutex.RLock()
// 	size = len(f.items)
// 	f.mutex.RUnlock()
// 	return
// }
