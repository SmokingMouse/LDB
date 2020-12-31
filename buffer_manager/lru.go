package buffer_manager

type LruKey interface{}
type LruValue interface{}

type Lru struct {
	m  map[LruKey]*ListNode
	li *List
	c  int
}

func NewLru(c int) *Lru {
	return &Lru{
		m:  make(map[LruKey]*ListNode),
		li: NewList(c),
		c:  c,
	}
}

func (lru *Lru) Get(key LruKey) LruValue {
	if _, ok := lru.m[key]; !ok {
		return nil
	}
	n := lru.m[key]
	lru.li.Delete(n)
	lru.li.Push(n)
	return n.lruVal
}

func (lru *Lru) Set(key LruKey, value LruValue) {
	if _, ok := lru.m[key]; ok {
		lru.m[key].lruVal = value
		lru.Get(key)
	} else {
		if lru.li.lenth == lru.c {
			toDelete := lru.li.head.prv
			delete(lru.m, toDelete.lruKey)
			lru.li.Delete(toDelete)
		}
		lru.m[key] = lru.li.InsertValue(lru.li.head, key, value)
	}
}

type ListNode struct {
	prv    *ListNode
	nxt    *ListNode
	lruVal LruValue
	lruKey LruKey
}

type List struct {
	head  *ListNode
	lenth int
}

func NewList(lenth int) *List {
	he := &ListNode{}
	he.prv = he
	he.nxt = he
	return &List{
		head: he,
	}
}

func (l *List) InsertValue(p *ListNode, k LruKey, v LruValue) *ListNode {
	ln := &ListNode{
		prv:    p,
		nxt:    p.nxt,
		lruKey: k,
		lruVal: v,
	}
	p.nxt.prv = ln
	p.nxt = ln
	l.lenth++
	return ln
}

func (l *List) Insert(p *ListNode, q *ListNode) {
	q.prv = p
	q.nxt = p.nxt
	p.nxt.prv = q
	p.nxt = q
	l.lenth++
}

func (l *List) Delete(p *ListNode) {
	if p == l.head {
		return
	}
	p.prv.nxt = p.nxt
	p.nxt.prv = p.prv
	p.nxt = nil
	p.prv = nil
	l.lenth--
}

func (l *List) Push(p *ListNode) {
	l.Insert(l.head, p)
}
