package queue

type Queue []interface{}

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	h := *q
	var el interface{}
	if len(h) > 0 {
		el = h[0]
		*q = h[1:]
	}
	return el
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func New() *Queue {
	return &Queue{}
}
