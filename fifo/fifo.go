package fifo

import "sync"

type element struct {
	data interface{}
	next *element
}

// Fifo represents the FIFO
type Fifo struct {
	len  int
	head *element
	tail *element

	mutex sync.RWMutex
}

// NewFifo creates a new empty FIFO
func NewFifo() *Fifo {
	return &Fifo{
		len:   0,
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

	head := f.head
	if head.next == nil {
		f.head = nil
		f.tail = nil
	} else {
		f.head = head.next
	}

	f.len--
	return head.data
}

// Enqueue allows to enqueue some data
func (f *Fifo) Enqueue(data interface{}) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	element := &element{
		data: data,
		next: nil,
	}

	if f.head == nil && f.tail == nil {
		f.head = element
		f.tail = element
	} else {
		last := f.tail
		last.next = element
		f.tail = element
	}

	f.len++
}

// IsEmpty returns whether the FIFO is empty or not
func (f *Fifo) IsEmpty() bool {
	f.mutex.RLock()
	defer f.mutex.RUnlock()
	return f.len == 0 && f.head == nil && f.tail == nil
}

// Size returns the actual FIFO size
func (f *Fifo) Size() int {
	f.mutex.RLock()
	f.mutex.RUnlock()
	return f.len
}
