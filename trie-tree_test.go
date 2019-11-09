package go_dfa_sensitive

import (
	"fmt"
	"testing"
)

func TestNewTrieTree(t *testing.T) {
	tree := NewTrieTree()
	tree.AddOneWord("fuck")
	tree.AddOneWord("fuck you")
	tree.AddOneWord("fuck your")
	tree.AddOneWord("草")
	tree.AddOneWord("草泥")
	fmt.Printf("result: %s\n", tree.ReplaceChar("fuck yourfuck you fucyfuck", "*"))
	fmt.Printf("result: %s\n", tree.ReplaceChar("fuc", "*"))
	fmt.Printf("result: %s\n", tree.ReplaceChar("草泥马", "*"))
	fmt.Printf("result: %s\n", tree.ReplaceChar("马草泥", "*"))
	fmt.Printf("result: %s\n", tree.ReplaceChar("草草马泥", "*"))
}
