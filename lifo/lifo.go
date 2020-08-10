package lifo

import "sync"

const defaultLifoCapacity = 1 << 6

// Lifo represents the LIFO
type Lifo struct {
	items []interface{}

	mutex sync.RWMutex
}

// NewLifo creates a new empty LIFO
func NewLifo() *Lifo {
	return &Lifo{
		items: make([]interface{}, 0, defaultLifoCapacity),
		mutex: sync.RWMutex{},
	}
}

// Dequeue returns the last element of the FIFO or nil
func (l *Lifo) Dequeue() interface{} {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if len(l.items) == 0 {
		return nil
	}

	data := l.items[len(l.items)-1]
	l.items = l.items[:len(l.items)-1]
	return data
}

// Enqueue allows to enqueue some data
func (l *Lifo) Enqueue(data interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.items = append(l.items, data)
}

// IsEmpty returns whether the FIFO is empty or not
func (l *Lifo) IsEmpty() bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return len(l.items) == 0
}

// Size returns the actual FIFO size
func (l *Lifo) Size() int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return len(l.items)
}
