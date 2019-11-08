package go_dfa_sensitive

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
)

type TrieNode struct {
	IsLeaf bool
	Value  string
	SonMap map[rune]*TrieNode
}

func NewTrieNode(value string) *TrieNode {
	return &TrieNode{
		Value:  value,
		SonMap: make(map[rune]*TrieNode),
	}
}

type TrieTree struct {
	Root *TrieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		Root: NewTrieNode(""),
	}
}

func (t *TrieTree) AddOneWord(word string) {
	if 0 == len(word) {
		return
	}
	tWord := []rune(word)
	root := t.Root
	var newRoot *TrieNode
	var exist bool
	for _, c := range tWord {
		newRoot, exist = root.SonMap[c]
		if !exist {
			newRoot = NewTrieNode(string([]rune{c}))
			root.SonMap[c] = newRoot
			root = newRoot
		}
	}
	if root != t.Root {
		root.IsLeaf = true
	}
}

func (t *TrieTree) SetTree(wordArray []string) {
	for _, word := range wordArray {
		t.AddOneWord(word)
	}
}

func (t *TrieTree) SetTreeByFile(filename string) (err error) {
	var file *os.File
	file, err = os.OpenFile(filename, os.O_RDONLY, 0666)
	if nil != err {
		return err
	}
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		t.AddOneWord(line)
		if nil != err {
			break
		}
	}
	if !errors.Is(err, io.EOF) {
		return err
	}
	return
}

func (t *TrieTree) IsLegal(content string) bool {
	var root *TrieNode
	var exist bool
	contentRune := []rune(content)
	for i, c := range contentRune {
		root, exist = t.Root.SonMap[c]
		if !exist {
			continue
		}
		for j := i + 1; j < len(contentRune); j++ {
			n := contentRune[j]
			root, exist = root.SonMap[n]
			if !exist {
				break
			} else if root.IsLeaf {
				return true
			}
		}
	}
	return false
}

func (t *TrieTree) ReplaceChar(content, charReplacer string) string {
	contentRune := []rune(content)
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < len(contentRune); i++ {
		c := contentRune[i]
		root, exist := t.Root.SonMap[c]
		if !exist {
			buffer.WriteString(content[i : i+1])
			continue
		}
		if root.IsLeaf {
			buffer.WriteString(charReplacer)
			continue
		}
		for j := i + 1; j < len(contentRune); j++ {
			nc := contentRune[j]
			root, exist = root.SonMap[nc]
			if !exist {
				break
			}
			if root.IsLeaf {
				for k := 0; k < j+1-i; k++ {
					buffer.WriteString(charReplacer)
				}
				i = j
			}
		}
	}
	return buffer.String()
}
