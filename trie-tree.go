package go_dfa_sensitive

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type TrieNode struct {
	Value  string
	SonMap map[rune]*TrieNode
}

func NewTrieNode(value string) *TrieNode {
	return &TrieNode{Value: value}
}

func (t *TrieNode) IsLeaf() bool {
	if 0 == len(t.SonMap) {
		return true
	}
	return false
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
	for _, c := range tWord {
		newRoot, exist := root.SonMap[c]
		if !exist {
			newRoot = NewTrieNode(string([]rune{c}))
			root.SonMap[c] = newRoot
		}
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
			} else {
				if root.IsLeaf() {
					return true
				}
			}
		}
	}
	return false
}
