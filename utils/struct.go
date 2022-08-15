package utils

type Queue struct {
	items []string
}

func NewQueue() *Queue {
	list := make([]string, 0)
	return &Queue{items: list}
}
func (q *Queue) Push(data string) {
	q.items = append(q.items, data)
}

func (q *Queue) Pop() string {
	if len(q.items) == 0 {
		return ""
	}
	res := q.items[0]
	q.items = q.items[1:]
	return res
}

type Stack struct {
	items []string
}

func NewStack() *Stack {
	list := make([]string, 0)
	return &Stack{items: list}
}

func (s *Stack) Push(t string) {
	s.items = append(s.items, t)
}

func (s *Stack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}
