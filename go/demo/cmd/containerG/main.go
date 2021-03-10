package main

import (
	"container/heap"
	"fmt"

	"../../containerg"
)

func DemoMinHeap() {
	hp := &containerg.RectHeap{}
	for i := 2; i < 6; i++ {
		*hp = append(*hp, containerg.Rectangle{i, 6 - i})
	}

	fmt.Println("原始slice: ", hp)

	// 堆操作
	heap.Init(hp)
	fmt.Println("init之后,当前slice: ", hp)
	heap.Push(hp, containerg.Rectangle{7, 2})
	fmt.Println("push之后,当前slice: ", hp)
	fmt.Println("top元素：", (*hp)[0])
	fmt.Println("删除并返回最后一个：", heap.Pop(hp)) // 最后 一个元素
	fmt.Println("最终slice: ", hp)
}

func DemoNumHeap() {
	h := &containerg.IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func DemoPriorityQueue() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(containerg.PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &containerg.Item{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &containerg.Item{
		Value:    "orange",
		Priority: 1,
	}
	heap.Push(&pq, item)
	pq.Update(item, item.Value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*containerg.Item)
		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple
}

func main() {
	//DemoMinHeap()
	//DemoNumHeap()
	DemoPriorityQueue()
}
