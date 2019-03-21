package goban

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
)

type Position struct {
	Schema   []int
	Size     int
	BlackCap int
	WhiteCap int
	HisNode  Node
	revert   bool
}

// 创建position对象
func NewPosition(size int) Position {
	position := Position{}
	position.Size = size
	position.Schema = position.CreateSchema()
	return position
}

//设置XY的计算公式
func (p *Position) SetRevert(xy bool) {
	p.revert = xy
}

// position坐标规则x*size+y
func (p Position) GetPos(x, y int) int {
	if p.revert {
		return x*p.Size + y
	}
	return x + y*p.Size
}

// 克隆POSITION对象
func (p *Position) Clone() (*Position) {
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(p);
	pos := &Position{}
	gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(pos)
	return pos
}

//创建schema
func (p Position) CreateSchema() []int {
	poss := make([]int, p.Size*p.Size)
	return poss
}

// 设置POSITION坐标对应颜色 Position规则为x*size+y
func (p *Position) SetColor(x, y, c int) {
	if x >= 0 && y >= 0 && x <= p.Size-1 && y <= p.Size-1 {
		p.Schema[p.GetPos(x, y)] = c
	}
}

// 获取position坐标对应的颜色 Position规则为x*size+y
func (p Position) GetColor(x, y int) int {
	if x >= 0 && y >= 0 {
		return p.Schema[p.GetPos(x, y)]
	}
	return 0
}

// 遍历x,y
func (p Position) ForeachXY(cb func(x, y int)) {
	for i := int(0); i < p.Size; i++ {
		for j := int(0); j < p.Size; j++ {
			cb(i, j)
		}
	}
}

//获取对应坐标的四领域,并回调触发
func (p Position) Neighbor4(x, y int, cb func(x, y int)) {
	// up
	if y > 0 {
		cb(x, y-1)
	}
	//left
	if x > 0 {
		cb(x-1, y)
	}
	//down
	if y < p.Size-1 {
		cb(x, y+1)
	}
	//right
	if x < p.Size-1 {
		cb(x+1, y)
	}
}

// 根据坐标和颜色获取是否可提子
func (p *Position) Cap(nodes []Node) {
	for _, v := range nodes {
		p.SetColor(v.X, v.Y, Empty)
	}
}

// 根据坐标和颜色获取是否可提子
func (p *Position) GetDeadByPointColor(x, y, c int) []Node {
	nodes := make([]Node, 0)
	p.Neighbor4(x, y, func(x, y int) {
		if p.GetColor(x, y) == c {
			nodes = append(nodes, p.CalcDeadNode(x, y, c)...)
		}
	})
	return nodes
}

//计算死子但不提子
func (p *Position) CalcDeadNode(x, y, c int) []Node {
	//新建一个计算的POSITION用于判断死子
	calcPos := NewPosition(p.Size)
	isDead := true
	nodes := make([]Node, 0)
	calcPos = p.FindAreaByC(calcPos, x, y, c)
	//判断是否可提子
	p.ForeachXY(func(x, y int) {
		if calcPos.GetColor(x, y) == 3 {
			isDead = false
		}
	})
	//如果可提子进入提子，但是不动原始数据
	if isDead {
		p.ForeachXY(func(i, j int) {
			if calcPos.GetColor(i, j) == c {
				p.SetColor(i, j, Empty)
				nodes = append(nodes, Node{
					X: i,
					Y: j,
					C: c,
				})
			}
		})
	}
	return nodes
}

// FindAreaByC 查找区域连块逻辑
func (p Position) FindAreaByC(pos Position, x, y, c int) Position {
	if pos.GetColor(x, y) != c && p.GetColor(x, y) == c {
		pos.SetColor(x, y, c)
		//上区域连块
		if y > 0 && p.GetColor(x, y-1) == c {
			pos = p.FindAreaByC(pos, x, y-1, c)
		} else if y > 0 && p.GetColor(x, y-1) == Empty {
			pos.SetColor(x, y-1, 3)
			return pos
		}
		//左区域连块
		if x > 0 && p.GetColor(x-1, y) == c {
			pos = p.FindAreaByC(pos, x-1, y, c)
		} else if x > 0 && p.GetColor(x-1, y) == Empty {
			pos.SetColor(x-1, y, 3)
			return pos
		}
		//下区域连块
		if y < p.Size-1 && p.GetColor(x, y+1) == c {
			pos = p.FindAreaByC(pos, x, y+1, c)
		} else if y < p.Size-1 && p.GetColor(x, y+1) == 0 {
			pos.SetColor(x, y+1, 3)
			return pos
		}
		//右区域连块
		if x < p.Size-1 && p.GetColor(x+1, y) == c {
			pos = p.FindAreaByC(pos, x+1, y, c)
		} else if x < p.Size-1 && p.GetColor(x+1, y) == 0 {
			pos.SetColor(x+1, y, 3)
			return pos
		}
	}
	return pos
}
func (p Position) ShowBoard(coors ...bool) string {
	show := true
	if len(coors) > 0 {
		show = coors[0]
	}
	boards := make([]string, p.Size)
	p.ForeachXY(func(x, y int) {
		color := p.GetColor(x, y)
		str := "."
		switch color {
		case B:
			str = "X"
		case W:
			str = "O"
		}
		boards[y] = fmt.Sprintf("%s%+3v", boards[y], str)
	})
	xCoor := ""
	if show {
		for i := 0; i < p.Size; i++ {
			temp := CoorToBoardNode(i, i, p.Size)
			xCoor += fmt.Sprintf("%+3v", string(temp[0]))
			boards[i] = fmt.Sprintf("%+3v%s", temp[1:], boards[i])
		}
		xCoor = fmt.Sprintf("%+3v", " ") + xCoor + "\n"
	}

	return xCoor + strings.Join(boards, "\n")
}

// GetStones 获取各个颜色的棋子列表
func (p Position) GetStones() (blackList []string, whiteList []string) {
	blackList = make([]string, 0)
	whiteList = make([]string, 0)
	p.ForeachXY(func(x, y int) {
		coon := CoorToSgfNode(x, y)
		if p.GetColor(x, y) == B {
			blackList = append(blackList, coon)
		} else if p.GetColor(x, y) == W {
			whiteList = append(whiteList, coon)
		}
	})
	return
}
func (p Position) CalcCap(color int, hisNode Node) (*Node, int) {
	cp := p.Clone()
	deadCount := 0
	result := &Node{}
	p.ForeachXY(func(i, j int) {
		n := cp.GetColor(i, j)
		if n == Empty {
			cp.Neighbor4(i, j, func(x, y int) {
				if cp.GetColor(x, y) != Empty {
					c := p.GetColor(x, y)
					if 0-c == color {
						node, cnt := p.getNextMove(i, j, color, deadCount, hisNode)
						if node != nil {
							deadCount = cnt
							result = node
						}
					}

				}
			})
		}
	})
	if deadCount > 0 {
		return result, deadCount
	}
	return nil, 0
}
func (p Position) getNextMove(x, y, c int, deadCount int, hisNode Node) (*Node, int) {
	p.SetColor(x, y, c)
	nodes := p.GetDeadByPointColor(x, y, -c)
	if len(nodes) > 0 {
		if len(nodes) > deadCount && !(x == hisNode.X && y == hisNode.Y && c == hisNode.C) {
			return &Node{
				X: x,
				Y: y,
				C: c,
			}, len(nodes)
		}
	}
	p.SetColor(x, y, Empty)
	return nil, 0
}

//获取还棋头
func (p Position) GetHeader(data []float64, size int) (black, white int) {
	pos := NewPosition(size)
	pos.SetRevert(true)
	pos.ForeachXY(func(x, y int) {
		if data[pos.GetPos(x, y)] > 0 {
			pos.SetColor(x, y, 1)
		} else if data[pos.GetPos(x, y)] < 0 {
			pos.SetColor(x, y, -1)
		}
	})
	pos.ForeachXY(func(i, j int) {
		color := pos.GetColor(i, j)
		if color != 0 {
			temp := pos.Clone()
			temp.SetRevert(true)
			r := NewPosition(size)
			r.SetRevert(true)
			p.calcHeader(temp, i, j, color, &r)
			if color == B {
				black++
			}
			if color == W {
				white++
			}
			r.ForeachXY(func(x, y int) {
				if r.GetColor(x, y) != Empty {
					pos.SetColor(x, y, 0)
				}
			})
		}
	})
	return
}

func (s Position) calcHeader(pos *Position, i, j, c int, temp *Position) {
	if temp.GetColor(i, j) != c {
		temp.SetColor(i, j, c)
		pos.Neighbor4(i, j, func(x, y int) {
			if pos.GetColor(x, y) != c {
				return
			}
			s.calcHeader(pos, x, y, c, temp)
		})
	}
}
