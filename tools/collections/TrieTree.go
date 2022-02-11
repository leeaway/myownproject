package collections

var aInt = int('a')

type TrieTreeNode struct {
	Content  int
	Children [26]*TrieTreeNode
}

func (n *TrieTreeNode) IsLeaf() bool {
	if n == nil {
		return false
	}
	for _, node := range n.Children {
		if node != nil {
			return false
		}
	}
	return true
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

func newTrieTreeNode(content int) *TrieTreeNode {
	return &TrieTreeNode{Content: content}
}

type TrieTree struct {
	Root *TrieTreeNode
}

func newTrieTree() *TrieTree {
	return &TrieTree{
		Root: &TrieTreeNode{},
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
	for _, w := range word {
		if temp.Children[int(w)-aInt] != nil {
			temp = temp.Children[int(w)-aInt]
			continue
		}
		curNode := newTrieTreeNode(int(w))
		temp.Children[int(w)-aInt] = curNode
		temp = curNode
	}
}

//精确查找
func (t *TrieTree) FindWord(word string) bool {
	temp := t.Root
	for i, w := range word {
		if temp.Children[int(w)-aInt] == nil {
			return false
		}
		temp = temp.Children[int(w)-aInt]
		if i == len(word)-1 {
			return true
		}
	}
	return false
}

//前缀查找
func (t *TrieTree) FuzzyFindWord(word string) bool {
	temp := t.Root
	for _, w := range word {
		if temp.Children[int(w)-aInt] == nil {
			return false
		}
		temp = temp.Children[int(w)-aInt]
	}
	return true
}
