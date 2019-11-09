package go_dfa_sensitive

import (
	"fmt"
	"testing"
)

func TestTrieTree_ReplaceChar(t *testing.T) {
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

func TestTrieTree_SetTreeByFile(t1 *testing.T) {
	tree := NewTrieTree()
	err := tree.SetTreeByFile("./test-data/words.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("result: %s\n", tree.ReplaceChar("安局办公楼123", "*"))
}
