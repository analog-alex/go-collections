package extra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var words = []string{"hello", "help", "world", "computer", "science", "GoLang", "fun"}

func TestTrie_Add(t *testing.T) {
	trie := MakeTrie()

	for _, word := range words {
		assert.NotPanics(t, func() { trie.Add(word) })
	}
}

func TestTrie_Contains(t *testing.T) {
	trie := MakeTrie()

	for _, word := range words {
		trie.Add(word)
	}

	contains := trie.Contains("hello")
	assert.True(t, contains)
}

func TestTrie_DoesNotContains(t *testing.T) {
	trie := MakeTrie()

	for _, word := range words {
		trie.Add(word)
	}

	contains := trie.Contains("something")
	assert.False(t, contains)
}

func TestTrie_DoesNotContainsEvenThoughItIsAPrefix(t *testing.T) {
	trie := MakeTrie()

	for _, word := range words {
		trie.Add(word)
	}

	contains := trie.Contains("hel")
	assert.False(t, contains)
}

func TestTrie_Starts(t *testing.T) {
	trie := MakeTrie()

	for _, word := range words {
		trie.Add(word)
	}

	contains := trie.Starts("hel")
	assert.True(t, contains)
}

func TestTrie_DoesNotStart(t *testing.T) {
	trie := MakeTrie()

	for _, word := range words {
		trie.Add(word)
	}

	contains := trie.Starts("some")
	assert.False(t, contains)
}

func TestTrie_Suggestions(t *testing.T) {
	trie := MakeTrie()

	for _, word := range words {
		trie.Add(word)
	}

	suggestions := trie.Suggestions("he")
	assert.Equal(t, []string{"hello", "help"}, suggestions)
}
