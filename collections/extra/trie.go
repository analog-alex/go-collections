package extra

type trieNode struct {
	val      rune
	isWord   bool
	children map[rune]*trieNode
}

type Trie struct {
	root *trieNode
}

func MakeTrie() *Trie {
	return &Trie{root: &trieNode{val: '*', isWord: false, children: make(map[rune]*trieNode)}}
}

// Add adds a new word to the trie
// Idempotent is the word already exists
func (t *Trie) Add(word string) {
	node := t.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			// we are adding a new rune, however until we are done with the insertion, this is not a word
			node.children[c] = &trieNode{val: c, isWord: false, children: make(map[rune]*trieNode)}
		}
		node = node.children[c]
	}

	// we got to the end of the insertion so this is a word
	node.isWord = true
}

// Contains checks if the word exists in the trie
// This means full word, not just a prefix i.e. 'hello' is a word but 'hel' is not (assuming English, etc.)
func (t *Trie) Contains(word string) bool {
	node := t.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}
	return node.isWord
}

// Starts checks if the trie contains any word that starts with the given prefix
// E.g 'hel' should return true if 'hello' was added to the trie
func (t *Trie) Starts(prefix string) bool {
	node := t.root
	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return false
		}
		node = node.children[c]
	}
	return true
}

// Remove removes the word from trie i.e. the sequence of letters remain but the word is no longer considered
func (t *Trie) Remove(word string) {
	node := t.root
	for _, c := range word {
		if _, ok := node.children[c]; !ok {
			return
		}
		node = node.children[c]
	}
	node.isWord = false
}

// Suggestions returns all possible words that start with the given prefix
// E.g. 'hel' should return ['hello', 'help'] if both words were added to the trie
func (t *Trie) Suggestions(prefix string) []string {
	node := t.root
	for _, c := range prefix {
		if _, ok := node.children[c]; !ok {
			return []string{}
		}
		node = node.children[c]
	}

	// starting from the final letter of the prefix, get all full words by transversing the trie
	// pick out all possible path, by doing an exhaustive search
	// we can, for performance reason, limit the number of suggestions returned
	// in the future we can do a depth first search and stop when we have enough suggestions
	return getAllFullWordsStartingFromNode(node, prefix)
}

func getAllFullWordsStartingFromNode(node *trieNode, runeSequence string) []string {
	var suggestions []string
	if node.isWord {
		suggestions = append(suggestions, runeSequence)
	}
	for _, child := range node.children {
		suggestions = append(suggestions, getAllFullWordsStartingFromNode(child, runeSequence+string(child.val))...)
	}
	return suggestions
}
