package graph

import "container/heap"

type Item struct {
	Vertex   string
	Distance int
	Index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func NewPriorityQueue() *PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}

func (pq *PriorityQueue) Insert(vertex string, distance int) {
	item := &Item{
		Vertex:   vertex,
		Distance: distance,
	}
	heap.Push(pq, item)
}

func (pq *PriorityQueue) ExtractMin() *Item {
	if pq.Len() == 0 {
		return nil
	}
	return heap.Pop(pq).(*Item)
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}
