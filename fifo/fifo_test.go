package fifo

import "testing"

func TestFifoCapabilities(t *testing.T) {
	q := NewFifo()

	if !q.IsEmpty() {
		t.Log("FIFO should be empty")
		t.Fail()
	}

	if q.Size() != 0 {
		t.Log("FIFO size should equal 0")
		t.Fail()
	}

	q.Enqueue(1)

	if q.IsEmpty() {
		t.Log("FIFO should not be empty")
		t.Fail()
	}

	if q.Size() != 1 {
		t.Log("FIFO size should equal 1")
		t.Fail()
	}

	if item := q.Dequeue(); item.(int) != 1 {
		t.Log("element should be of type int with the value '1'")
		t.Fail()
	}

	q.Enqueue("two")
	q.Enqueue(true)

	if item := q.Dequeue(); item.(string) != "two" {
		t.Log("element should be of type string with the value 'two'")
		t.Fail()
	}

	if item := q.Dequeue(); !item.(bool) {
		t.Log("element should be of type bool with the value 'true'")
		t.Fail()
	}
}

var q *Fifo = NewFifo()

func BenchmarkEnqueue(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkDequeue(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < q.Size(); i++ {
		q.Dequeue()
	}
}
