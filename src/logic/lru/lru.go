package lru

import (
	"sync"
	"unsafe"
)

type Node struct {
	value int
	key   string
	pre   *Node
	next  *Node
}

func NewNode(key string, value int) *Node {
	return &Node{
		value: value,
		key:   key,
	}
}

type LinkList struct {
	count    int64
	capacity int64
	head     *Node
	tail     *Node
}

func (list *LinkList) moveToHead(node *Node) {

	//node 是尾巴
	if node.pre != nil && node.next == nil && unsafe.Pointer(node) == unsafe.Pointer(list.tail) {
		list.tail = node.pre
	}
	//node 本来就是head
	if node.pre == nil {
		return
	}
	//断开原来的关系链
	node.pre.next = node.next
	node.next.pre = node.pre
	node.pre = nil

	//插入到头部
	head := list.head
	head.pre = node
	node.next = head

	list.head = node
}

func (list *LinkList) add(node *Node) {
	head := list.head
	if head == nil {
		list.head = node
		list.tail = node
		list.count++
		return
	}
	node.next = head
	head.pre = node
	list.head = node
	list.count++
	if list.count > list.capacity {
		tail := list.tail
		tail.pre.next = nil
		list.tail = tail.pre
		tail.pre = nil //gc
		list.count--
	}

}
func (list *LinkList) remove(node *Node) {
	ht := false
	if node.pre == nil {
		list.head = nil
		ht = true
	}
	if node.next == nil {
		list.tail = nil
		ht = true
	}
	if ht {
		list.count--
		return
	}
	node.pre.next = node.next
	node.next.pre = node.pre
	list.count--
}

type Cache struct {
	sync.Mutex
	keys map[string]*Node
	list LinkList
}

func NewLRUCache(capacity int64) *Cache {

	return &Cache{
		Mutex: sync.Mutex{},
		keys:  map[string]*Node{},
		list:  LinkList{capacity: capacity},
	}
}

func (cache *Cache) Get(key string) any {
	cache.Lock()
	defer cache.Unlock()
	node, ok := cache.keys[key]
	if !ok {
		return nil
	}
	cache.list.moveToHead(node)
	return node.value
}
func (cache *Cache) Add(key string, value int) {

	node, ok := cache.keys[key]
	if ok {
		node.value = value
		cache.list.moveToHead(node)
		return
	}
	node = NewNode(key, value)
	cache.keys[node.key] = node
	cache.list.add(node)
}
func (cache *Cache) Remove(key string) {
	node, ok := cache.keys[key]
	if !ok {
		return
	}
	cache.list.remove(node)
}
