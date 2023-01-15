package collections

type TrieTreeNode struct {
	Children [26]*TrieTreeNode
	//表示是否是结尾，部分节点可能既是结尾，又不是，如abc、abcd的c节点
	IsFinished bool
}

func (n *TrieTreeNode) GetChildrenNum() int {
	if n == nil {
		return 0
	}
	res := 0
	for _, node := range n.Children {
		if node != nil {
			res++
		}
	}
	return res
}

func newTrieTreeNode() *TrieTreeNode {
	return &TrieTreeNode{}
}

type TrieTree struct {
	Root *TrieTreeNode
}

func newTrieTree() *TrieTree {
	return &TrieTree{
		Root: newTrieTreeNode(),
	}
}

func CreateTrieTree(words []string) *TrieTree {
	trie := newTrieTree()
	for _, word := range words {
		trie.AddWord(word)
	}
	return trie
}

func (t *TrieTree) AddWord(word string) {
	temp := t.Root
	for i, w := range word {
		idx := int(w - 'a')
		if temp.Children[idx] != nil {
			temp = temp.Children[idx]
			if i == len(word)-1 {
				temp.IsFinished = true
			}
			continue
		}
		temp.Children[idx] = newTrieTreeNode()
		temp = temp.Children[idx]
		if i == len(word)-1 {
			temp.IsFinished = true
		}
	}
}

//精确查找
func (t *TrieTree) FindWord(word string) bool {
	temp := t.Root
	for i, w := range word {
		idx := int(w - 'a')
		if temp.Children[idx] == nil {
			return false
		}
		temp = temp.Children[idx]
		if i == len(word)-1 {
			return temp.IsFinished
		}
	}
	return false
}

//前缀查找
func (t *TrieTree) FuzzyFindWord(word string) bool {
	temp := t.Root
	for _, w := range word {
		idx := int(w - 'a')
		if temp.Children[idx] == nil {
			return false
		}
		temp = temp.Children[idx]
	}
	return true
}
