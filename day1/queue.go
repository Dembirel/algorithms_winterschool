
package main 

type simpleQueue []int

func (q *simpleQueue) EnQueue(x int) {
	*q = append(*q, x)
}

func (q *simpleQueue) DeQueue() (int, bool) {
	if len(*q) == 0 {
		return 0, false
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, true
}