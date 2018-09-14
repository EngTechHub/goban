package goban

import "fmt"

type Node struct {
	X          int
	Y          int
	C          int
	Comment    string
	Parent     *Node               `json:"-"`
	Childrens  []*Node             `json:"-"`
	Steup      []*Node             `json:"-"`
	LastSelect int                 `json:"-"`
	Info       map[string][]string `json:"-"`
}

//创建空对象
func NewNode() *Node {
	return &Node{
		Childrens: make([]*Node, 0),
		Steup:     make([]*Node, 0),
		Parent:    nil,
		Info:      make(map[string][]string),
	}
}

//获取子节点
func (n Node) GetChild(i int) *Node {
	if len(n.Childrens) == 0 {
		return nil
	}
	if len(n.Childrens) <= i {
		i = 0
	}
	n.LastSelect = i
	return n.Childrens[n.LastSelect]
}

//追加子节点
func (n *Node) AppendChild() *Node {
	node := NewNode()
	node.Parent = n
	n.Childrens = append(n.Childrens, node)
	return node
}
func (n *Node) RemoveChild() *Node {
	parent := n.Parent
	if parent!=nil{
		i := parent.LastSelect
		parent.Childrens = append(parent.Childrens[:i], parent.Childrens[i+1:]...)
	}
	return parent
}

// 添加AB/AW标签
func (n *Node) AddSetup(obj *Node) {
	n.Steup = append(n.Steup, obj)
}

// 返回当前节点的颜色字符串
func (n *Node) GetColor() string {
	if n.C == B {
		return "B"
	} else if n.C == W {
		return "W"
	}
	return ""
}

// 返回当前节点的颜色字符串
func (n *Node) IsPass() bool {
	if n.X == -1 && n.Y == -1 {
		return true
	}
	return false
}

func (n *Node) AddInfo(key string,value interface{}) {
	str:=fmt.Sprintf("%v",value)
	switch value.(type) {
	case float64,float32:
		str=fmt.Sprintf("%.1f",value)
	}
	n.Info[key]=[]string{str}
}