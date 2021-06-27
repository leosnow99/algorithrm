package datastruct

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const UpLevelsAbility = 500
const UpLevelsTotal = 1000

type SkipList interface {
	Size() int
	Levels() int
	Get(score int64) interface{}
	Insert(score int64, val interface{})
	Remove(score int64) interface{}
}

type skipListNode struct {
	score int64
	val   interface{}
	next  *skipListNode
	pre   *skipListNode
	up    *skipListNode
	down  *skipListNode
}

type skipList struct {
	head   *skipListNode
	tail   *skipListNode
	size   int
	levels int
}

func NewSkipList() SkipList {
	var sl = new(skipList)
	sl.head = new(skipListNode)
	sl.tail = new(skipListNode)
	sl.head.score = math.MinInt64
	sl.tail.score = math.MaxInt64
	sl.head.next = sl.tail
	sl.tail.pre = sl.head
	sl.size = 0
	sl.levels = 1
	return sl
}

func (sl *skipList) Size() int {
	return sl.size
}

func (sl *skipList) Levels() int {
	return sl.levels
}

func (sl *skipList) Get(score int64) interface{} {
	node := sl.findNode(score)
	if node.score == score {
		return node.val
	}
	return nil
}

func (sl *skipList) Insert(score int64, val interface{}) {
	f := sl.findNode(score)
	if f.score == score {
		f.val = val
		return
	}
	curNode := new(skipListNode)
	curNode.score = score
	curNode.val = val
	sl.insertAfter(f, curNode)

	randier := rand.New(rand.NewSource(time.Now().UnixNano()))
	curLevels := 1
	for randier.Intn(UpLevelsTotal) < UpLevelsAbility {
		curLevels++
		if curLevels > sl.levels {
			sl.newLevels()
		}

		for f.up == nil {
			f = f.pre
		}
		f = f.up
		tmpNode := &skipListNode{score: score}
		curNode.up = tmpNode
		tmpNode.down = curNode
		sl.insertAfter(f, tmpNode)
		curNode = tmpNode
	}
	sl.size++
}

func (sl *skipList) Remove(score int64) interface{} {
	f := sl.findNode(score)
	if f.score != score {
		return nil
	}
	oldValue := f.val
	for f != nil {
		f.pre.next = f.next
		f.next.pre = f.pre
		f = f.up
	}
	return oldValue
}

func (sl *skipList) newLevels() {
	nHead := &skipListNode{score: math.MinInt64}
	nTail := &skipListNode{score: math.MinInt64}
	nHead.next = nTail
	nTail.pre = nHead

	sl.head.up = nHead
	sl.tail.up = nTail
	nHead.down = sl.head
	nTail.down = sl.tail
	sl.head = nTail
	sl.tail = nTail
	sl.levels++
}

func (sl *skipList) findNode(score int64) *skipListNode {
	p := sl.head
	for p != nil {
		if p.score == score {
			if p.down == nil {
				return p
			}
			p = p.down
		} else if p.score < score {
			if p.next.score > score {
				if p.down == nil {
					return p
				}
				p = p.down
			} else {
				p = p.next
			}
		}
	}
	return p
}

func (sl *skipList) insertAfter(pNode, curNode *skipListNode) {
	curNode.next = pNode.next
	curNode.pre = pNode
	pNode.next.pre = curNode
	pNode.next = curNode
}

func (sl *skipList) Print() {
	mapScore := make(map[int64]int)
	p := sl.head
	for p.down != nil {
		p = p.down
	}
	index := 0
	for p != nil {
		mapScore[p.score] = index
		p = p.next
		index++
	}

	p = sl.head
	for i := 0; i < sl.levels; i++ {
		q := p
		preIndex := 0
		for q != nil {
			s := q.score
			if s == math.MinInt64 {
				fmt.Printf("%s", "BEGIN")
				q = q.next
				continue
			}
			index := mapScore[s]
			c := (index - preIndex - 1) * 6
			for m := 0; m < c; m++ {
				fmt.Print("-")
			}
			if s == math.MaxInt64 {
				fmt.Printf("-->%s\n", "END")
			} else {
				fmt.Printf("-->%3d", s)
				preIndex = index
			}
			q = q.next
		}
		p = p.down
	}
}
