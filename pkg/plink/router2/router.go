package router2

import (
	"strings"
)

type HandlerFun func(ctx any)

type Node struct {
	path     string           // 当前节点的路径
	children map[string]*Node // 子节点集合
	isLeaf   bool             // 是否为叶子节点
	Param    string           // 参数名（如果该节点是变量）
	Handler  HandlerFun       // 绑定的方法
}

func NewNode(path string) *Node {
	return &Node{
		path:     path,
		children: make(map[string]*Node),
		isLeaf:   false,
		Param:    "",
		Handler:  nil,
	}
}

type PrefixTree struct {
	root *Node
}

func NewPrefixTree() *PrefixTree {
	return &PrefixTree{
		root: NewNode(""),
	}
}

func (p *PrefixTree) Insert(path string, handler HandlerFun) {
	node := p.root
	paths := strings.Split(path, "/")

	for _, p := range paths {
		if len(p) == 0 {
			continue
		}

		if strings.HasPrefix(p, ":") {
			child, ok := node.children[":"]
			if !ok {
				child = NewNode(p)
				child.Param = p[1:]
				node.children[":"] = child
			}
			node = child
		} else {
			child, ok := node.children[p]
			if !ok {
				child = NewNode(p)
				node.children[p] = child
			}
			node = child
		}
	}

	node.isLeaf = true
	node.Handler = handler
}

func (p *PrefixTree) Search(path string) (*Node, map[string]string) {
	node := p.root
	params := make(map[string]string)
	paths := strings.Split(path, "/")

	for _, p := range paths {
		if len(p) == 0 {
			continue
		}

		child, ok := node.children[p]
		if !ok {
			child, ok = node.children[":"]
			if !ok {
				return nil, params
			}
			params[child.Param] = p
			node = child
		} else {
			node = child
		}
	}

	if !node.isLeaf {
		return nil, params
	}

	return node, params
}
