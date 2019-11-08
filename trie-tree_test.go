package go_dfa_sensitive

import (
	"fmt"
	"testing"
)

func TestNewTrieTree(t *testing.T) {
	tree := NewTrieTree()
	tree.AddOneWord("fuck")
	result := tree.ReplaceChar("fuck you", "*")
	fmt.Printf("result: %s\n", result)
}
