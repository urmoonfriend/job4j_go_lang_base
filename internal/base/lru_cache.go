package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	size int
	items map[string]*Node
	Head *Node
	Tail *Node
}

func NewLruCache(size int) *LruCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return &LruCache{
		size: size,
		Head: head,
		Tail: tail,
		items: make(map[string]*Node),
	}
}

func (l *LruCache) Put(key string, value string) {
	if node, ok := l.items[key]; ok {
		node.Value = value
		l.moveToHead(node)
		return
	}

	newNode := &Node{
		Key: key,
		Value: value,
	}
	l.items[key] = newNode
	l.addToHead(newNode)

	if len(l.items) > l.size {
		tail := l.removeTail()
		delete(l.items, tail.Key)
	}
}

func (l *LruCache) Get(key string) *string {
	if node, ok := l.items[key]; ok {
		l.moveToHead(node)
		return &node.Value
	}
	return nil
}

func (l *LruCache) moveToHead(node *Node) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LruCache) addToHead(node *Node) {
	node.Next = l.Head.Next
	node.Prev = l.Head
	l.Head.Next.Prev = node
	l.Head.Next = node
}

func (l *LruCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (l *LruCache) removeTail() *Node {
	tail := l.Tail.Prev
	l.removeNode(tail)
	return tail
}

