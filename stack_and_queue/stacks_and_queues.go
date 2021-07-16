package stackandqueue

// Stack represents stack that holds a slice
type Stack struct {
	items []int
}

// Push will add a value to the end
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end
// and RETURNs the removed value OR
// RETURNS -1 if the slice is empty
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	lastItem := len(s.items) - 1
	toRemove := s.items[lastItem]
	s.items = s.items[:lastItem]
	return toRemove
}

// Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

// Enqueue adds a value to the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue will remove a value at the front
// and RETURNs the removed value OR
// RETURNS -1 if the slice is empty
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}
