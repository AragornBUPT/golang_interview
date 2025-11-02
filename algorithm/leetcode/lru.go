package leetcode

import "fmt"

type LinkedNode struct {
	key  int
	val  int
	prev *LinkedNode
	next *LinkedNode
}

type LRUCache struct {
	capacity int
	cache    map[int]*LinkedNode
	head     *LinkedNode
	tail     *LinkedNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*LinkedNode),
		head:     nil,
		tail:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if nodePtr, ok := this.cache[key]; ok {
		this.moveHead(nodePtr)
		return nodePtr.val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if nodePtr, ok := this.cache[key]; ok {
		nodePtr.val = value
		this.moveHead(nodePtr)
	} else {
		// 加入新kv，并将其放在链表的头部
		nodePtr := &LinkedNode{
			key:  key,
			val:  value,
			prev: nil,
			next: nil,
		}
		this.cache[key] = nodePtr
		this.addHead(nodePtr)

		// 判断是否超出容量
		if len(this.cache) > this.capacity {
			this.removeTail()
		}
	}
}

// 新增节点
func (this *LRUCache) addHead(nodePtr *LinkedNode) {
	if this.head == nil { //如果链表没有数据，则将该节点作为首尾节点
		this.head = nodePtr
		this.tail = nodePtr
	} else { //如果链表已经有了至少一个数据，则将该节点作为头节点
		nodePtr.next = this.head
		this.head.prev = nodePtr
		this.head = nodePtr
	}
	this.print()
}

func (this *LRUCache) removeTail() {
	key := this.tail.key
	nodePtr, _ := this.cache[key]
	delete(this.cache, key)

	this.tail = this.tail.prev
	nodePtr.prev = nil
	this.tail.next = nil

	this.print()
}

func (this *LRUCache) moveHead(nodePtr *LinkedNode) {
	if nodePtr == this.head {
		return
	}

	// 断开前后的链
	if  nodePtr == this.tail {
		this.tail = nodePtr.prev
		this.tail.next = nil	//断开 ->node
	} else {
		nodePtr.next.prev = nodePtr.prev	//断开 node<-
		nodePtr.prev.next = nodePtr.next	//断开 ->node
	}

	nodePtr.prev = nil	//断开<-node
	nodePtr.next = this.head	//连接node->
	this.head.prev = nodePtr	//连接node<-head
	this.head = nodePtr

	this.print()
}

func (this *LRUCache) print() {
	this.printHead()
	fmt.Println()
	this.printTail()
	fmt.Println()
	fmt.Println()
}

func (this *LRUCache) printHead() {
	ptr := this.head
	for ptr != nil {
		fmt.Printf("%d, %d -> ", ptr.key, ptr.val)
		ptr = ptr.next
	}
}

func (this *LRUCache) printTail() {
	ptr := this.tail
	for ptr != nil {
		fmt.Printf("%d, %d -> ", ptr.key, ptr.val)
		ptr = ptr.prev
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
