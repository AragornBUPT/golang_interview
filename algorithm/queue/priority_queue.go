package queue
// package main

import "fmt"

type Item struct {
	Priority int
	Value    interface{}
}

type PriorityQueue struct {
	data []*Item
}

func (pq *PriorityQueue) up(index int) {
	for {
		parent := (index - 1) / 2
		if parent == index || pq.data[parent].Priority >= pq.data[index].Priority {
			break
		}

		pq.data[parent], pq.data[index] = pq.data[index], pq.data[parent]
		index = parent
	}
}

func (pq *PriorityQueue) down() {
	index := 0
	length := len(pq.data)
	for {
		parent := index
		left := parent*2 + 1
		right := parent*2 + 2

		if left < length && pq.data[index].Priority < pq.data[left].Priority {
			index = left
		}

		if right < length && pq.data[index].Priority < pq.data[right].Priority {
			index = right
		}

		if parent == index {
			break
		}

		pq.data[parent], pq.data[index] = pq.data[index], pq.data[parent]
	}
}

func NewPriorityQueue() (pq *PriorityQueue) {
	pq = &PriorityQueue{
		data: make([]*Item, 0),
	}

	return
}

func (pq *PriorityQueue) Enqueue(value interface{}, priority int) {
	item := &Item{
		Priority: priority,
		Value:    value,
	}
	pq.data = append(pq.data, item)
	pq.up(len(pq.data) - 1)
}

func (pq *PriorityQueue) Dequeue() (item *Item) {
	lastIndex := len(pq.data) - 1
	pq.data[0], pq.data[lastIndex] = pq.data[lastIndex], pq.data[0]
	item = pq.data[lastIndex]

	pq.data = pq.data[0:lastIndex]
	pq.down()

	return item
}

func (pq *PriorityQueue) Peek() (item *Item) {
	if pq.IsEmpty() {
		return nil
	}
	return pq.data[0]
}

func (pq *PriorityQueue) IsEmpty() (isEmpty bool) {
	if len(pq.data) == 0 {
		return true
	}

	return false
}

// ========== 测试示例 ==========
func main() {
	// 创建优先队列
	pq := NewPriorityQueue()

	// 入队：添加不同优先级的元素
	pq.Enqueue("任务A", 3)
	pq.Enqueue("任务B", 1)
	pq.Enqueue("任务C", 5)
	pq.Enqueue("任务D", 2)

	// 查看队首元素
	fmt.Println("队首元素：", pq.Peek().Value, "（优先级：", pq.Peek().Priority, "）")

	// 依次出队，验证优先级顺序
	fmt.Println("\n出队顺序：")
	for !pq.IsEmpty() {
		item := pq.Dequeue()
		fmt.Printf("值：%s，优先级：%d\n", item.Value, item.Priority)
	}
}
