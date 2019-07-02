package main

import (
	"fmt"
	"gitee.com/larry_dev/goban"
)

func main() {
	kifu := goban.ParseSgf("(;SZ[19]AP[WGo.js:2]FF[4]GM[1]CA[UTF-8];B[dd];W[de];B[ce];W[ee];B[ed];W[da];B[df];W[eb];B[fe];W[ef];B[eg];W[ff];B[gf];W[fg];B[fh];W[gg];B[hg];W[gh];B[gi];W[hh];B[ih];W[hi];B[hj];W[ii];B[ji];W[ij];B[ik];W[jj];B[kj];W[jk];B[jl];W[kk];B[lk];W[kl];B[km];W[ll];B[ml];W[lm];B[ln])")
	eyeList := make(map[int]bool)
	//kifu.GoTo(10)
	//if eyeList[poss.GetPos(kifu.CurNode.X,kifu.CurNode.Y)]{
	//	fmt.Println("第一次")
	//}
	total := 0
	kifu.EachNode(func(n *goban.Node, move int) bool {
		kifu.GoTo(move)
		if move>=9{
			fmt.Println(1)
		}
		pos := *kifu.CurPos.Clone()
		if len(eyeList) > 0 {
			if eyeList[pos.GetPos(kifu.CurNode.X, kifu.CurNode.Y)] {
				total++
			}
		}
		eyeList=make(map[int]bool)
		pos.ForeachXY(func(x, y int) {
			if pos.GetColor(x, y) != goban.Empty {
				hands := find(pos, x, y, pos.GetColor(x, y))
				if len(hands) == 1 {
					hand := hands[0]
					eyeList[pos.GetPos(hand.X, hand.Y)] = true
				}
			}
		})
		return false
	})
	fmt.Println(total)
}

func find(pos goban.Position, x, y, c int) []goban.Node {
	arr := make([]goban.Node, 0)
	pos.Neighbor4(x, y, func(i, j int) {
		if pos.GetColor(i, j) == goban.Empty {
			arr = append(arr, goban.Node{
				X: i,
				Y: j,
			})
		}
	})
	return arr
}
