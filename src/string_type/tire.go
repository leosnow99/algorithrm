package string_type

/**
字典数=树

字典树又称为前缀树或Trie 树，是处理字符串常见的数据结构。假设组成所有单词的字符
仅是“a”~“z”，请实现字典树结构，并包含以下四个主要功能:
	 void insert(String word)：添加word，可重复添加。
	 void delete(String word)：删除word，如果word 添加过多次，仅删除一个。
	 boolean search(String word)：查询word 是否在字典树中。
	 int prefixNumber(String pre)：返回以字符串pre 为前缀的单词数量。
*/

type Tire interface {
	Insert(string)
	Delete(string)
	Search(string) bool
	PrefixNumber(string) int
}

type tireNode struct {
	path int         // 表示有多少个单次共用这个节点
	end  int         // 表示有多少个单次以这个节点为结尾
	maps []*tireNode // key 代表该节点的一条字符路径，value 表示字符路径指向的节点
}

func newTireNode() *tireNode {
	return &tireNode{
		path: 0,
		end:  0,
		maps: make([]*tireNode, 26),
	}
}

type tireImpl struct {
	root *tireNode
}

func (t tireImpl) Insert(word string) {
	if len(word) == 0 {
		return
	}
	node := t.root
	node.path++
	index := 0
	chs := []byte(word)
	for i := 0; i < len(chs); i++ {
		index = int(chs[i] - 'a')
		if node.maps[index] == nil {
			node.maps[index] = newTireNode()
		}
		node = node.maps[index]
		node.path++
	}
	node.end++
}

func (t tireImpl) Delete(word string) {
	if !t.Search(word) {
		return
	}
	chs := []byte(word)
	node := t.root
	index := 0
	for i := 0; i < len(chs); i++ {
		index = int(chs[i] - 'a')
		if node.maps[index].path == 1 {
			node.maps[index] = nil
			return
		}
		node.maps[index].path--
		node = node.maps[index]
	}
	node.end--
}

func (t tireImpl) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	chs := []byte(word)
	node := t.root
	index := 0
	for i := 0; i < len(chs); i++ {
		index = int(chs[i] - 'a')
		node = node.maps[index]
		if node == nil {
			return false
		}
	}
	return node.end != 0
}

func (t tireImpl) PrefixNumber(pre string) int {
	if len(pre) == 0 {
		return 0
	}
	chs := []byte(pre)
	node := t.root
	index := 0
	for i := 0; i < len(chs); i++ {
		index = int(chs[i] - 'a')
		if node.maps[index] == nil {
			return 0
		}
		node = node.maps[index]
	}
	return node.path
}

func NewTire() Tire {
	return tireImpl{
		root: newTireNode(),
	}
}