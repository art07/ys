package main

import "strings"

type Tree struct {
	line      string
	leftTree  *Tree
	rightTree *Tree
}

func createTree(value string) (tree *Tree) {
	tree = &Tree{
		line:      value,
		leftTree:  nil,
		rightTree: nil,
	}
	return
}

func (t *Tree) insert(nodeLine string, column int) *Tree {
	if strings.Split(nodeLine, ",")[column-1] < strings.Split(t.line, ",")[column-1] {
		if t.leftTree == nil {
			t.leftTree = createTree(nodeLine)
		} else {
			t.leftTree.insert(nodeLine, column)
		}
	} else {
		if t.rightTree == nil {
			t.rightTree = createTree(nodeLine)
		} else {
			t.rightTree.insert(nodeLine, column)
		}
	}
	return t
}

func (t *Tree) forEach(fn func(string)) {
	if t.leftTree != nil {
		t.leftTree.forEach(fn)
	}
	fn(t.line)
	if t.rightTree != nil {
		t.rightTree.forEach(fn)
	}
}

func binaryTreeSort(linesArr []string, column int) *Tree {
	mainTree := createTree(linesArr[0])
	for _, nodeLine := range linesArr[1:] {
		mainTree.insert(nodeLine, column)
	}
	return mainTree
}
