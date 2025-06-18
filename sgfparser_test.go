package goban

import (
	"os"
	"fmt"
	"github.com/magiconair/properties/assert"
	"strings"
	"testing"
)

func TestParseSgfFile(t *testing.T) {
	// 读取文件内容
	sgf, err := os.ReadFile("../sgf_import_test_20250610151827.sgf")
	if err != nil {
		fmt.Printf("读取文件出错: %v", err)
		return
	}
	kifu := ParseSgf(string(sgf))
	// assert.Equal(t, kifu.ToSgf(), sgf)
	fmt.Printf("rs:%v\n", kifu.ToSgf())
}

func TestParseHeatMap(t *testing.T) {
	log := ` 0   0   0   0   0   0   0   0   0   0   0   0   0 
  0   0   0   0   1   1   1   1   1   1   0   0   0 
  0   0   8  35   1   6   3 241  53   9   0   0   0 
  0   0   1  11   1   1   2  15  34   1   0   0   0 
  0   0   0   0   0   1   1   1   1   1   0   0   0 
  0   0   0   1   1   1   1   0   1   1   1   0   0 
  0   0   0   0   1   1   1   1   1   0   0   0   0 
  0   0   0   0   1   1   0   1   0   1   1   0   0 
  0   0   0   0   0   1   0   0   1   4   1   0   0 
  0   0   1   0   0   0   1   1   5   5   0   0   0 
  0   0   1   2   1   6   3  17   1 393   4   0   0 
  0   0   0   0   2   1   1   0   0   1   9   0   0 
  0   0   0   0   0   0   0   0   0   0   0   0   0 
pass: 0
winrate: 0.584828
`
	ss, f, i := ParseHeatMap(log, 13)
	t.Log(ss,f,i)
}
func TestParseSgf(t *testing.T) {
	//正常解析测试

	sgf := "(;SZ[19]KM[6.5]HA[1]AB[ab]AW[bb]C[123132]GM[1];B[ba])"
	kifu := ParseSgf(sgf)
	kifu.Last()
	assert.Equal(t, kifu.Size, 19)
	assert.Equal(t, kifu.Komi, 6.5)
	assert.Equal(t, kifu.Handicap, 1)
	assert.Equal(t, len(kifu.Root.Steup), 2)
	assert.Equal(t, kifu.NodeCount, 1)
	assert.Equal(t, kifu.Play(0, 0, -1), false)
	assert.Equal(t, kifu.ToSgf(), "(;SZ[19]KM[6.5]HA[1]AB[ab]AW[bb]C[123132]GM[1];B[ba])")
	assert.Equal(t, kifu.CurColor, W)

	sgf = "(;SZ[19]AB[ii]AW[jj];B[aa];W[ab];B[ac];W[ba];B[bb];W[ca];B[aa])"
	kifu = ParseSgf(sgf)
	kifu.Last()
	assert.Equal(t, kifu.NodeCount, 7)
	assert.Equal(t, kifu.Play(0, 1, -1), false)
	assert.Equal(t, kifu.Play(0, 0, -1), false)
	assert.Equal(t, kifu.Play(-1, -1, -1), true)
	assert.Equal(t, kifu.Play(10, 10, 1), true)
	assert.Equal(t, kifu.Play(0, 1, -1), true)
	assert.Equal(t, kifu.NodeCount, 10)
	assert.Equal(t, kifu.ToSgf(), "(;SZ[19]KM[7.5]HA[0]AB[ii]AW[jj];B[aa];W[ab];B[ac];W[ba];B[bb];W[ca];B[aa];W[tt];B[kk];W[ab])")
	kifu.GoTo(0)

	assert.Equal(t, kifu.Play(18, 18, 1), true)
	assert.Equal(t, kifu.Play(15, 15, -1), true)
	assert.Equal(t, kifu.NodeCount, 12)
	assert.Equal(t, kifu.ToSgf(), "(;SZ[19]KM[7.5]HA[0]AB[ii]AW[jj](;B[aa];W[ab];B[ac];W[ba];B[bb];W[ca];B[aa];W[tt];B[kk];W[ab])(;B[ss];W[pp]))")
	assert.Equal(t, kifu.ToCurSgf(), "(;SZ[19]KM[7.5]HA[0]AB[ii]AW[jj];B[ss];W[pp])")
	kifu.GoTo(0)
	kifu.CurNode.LastSelect = 2
	kifu.GoTo(1)
	assert.Equal(t, kifu.ToCurSgf(), "(;SZ[19]KM[7.5]HA[0]AB[ii]AW[jj];B[aa])")

	sgf = "(;SZ[a9]KM[7aa]HA[aa]AB[tt](;B[pd];W[dd];B[qp];W[dq];B[oq];W[qf];B[nc];W[rd];B[fc];W[cf];B[qh];W[qc];B[qe];W[re];B[pf];W[pg];B[qg];W[rf];B[og];W[hc];B[pj];W[jq];B[cn];W[cp];B[dj];W[en];B[cl];W[ci];B[cj];W[mq];B[db];W[cc];B[ed];W[kc];B[de];W[ce];B[dc];W[cd];B[he];W[ke];B[ib];W[ic];B[hb];W[jb];B[gc];W[ia];B[ga];W[ie];B[no];W[qo];B[rp];W[ol];B[ml];W[ql];B[qk];W[oj];B[pl];W[pm];B[pk];W[ok];B[om];W[pi];B[qi];W[rk];B[qj];W[rm];B[pn];W[qm];B[ro];W[rj];B[ri];W[sj];B[sl];W[si];B[sh];W[sg];B[ph];W[rh];B[hq];W[io];B[ep];W[eq];B[fq];W[go];B[fo];W[gp];B[fp];W[fn];B[do];W[fr];B[gr];W[er];B[gq];W[ip];B[bi];W[cb];B[bg];W[hf];B[ge];W[mj];B[nm];W[nh];B[mf];W[lg];B[mg];W[mh];B[lf];W[kg];B[pc];W[qb];B[bo];W[bp];B[kr];W[kq];B[lr];W[mr];B[jr];W[lq];B[gn];W[gm];B[hn];W[il];B[fm];W[hm];B[em];W[in];B[dn];W[ho];B[lk];W[kj];B[hh];W[gf];B[ff];W[fg];B[ef];W[gh];B[gi];W[fi];B[gj];W[fj];B[gk];W[fk];B[jk];W[kk];B[jl];W[kl];B[jm];W[km];B[jj];W[ig];B[ki];W[lj];B[jn];W[kn];B[jo];W[ko];B[eh];W[fh];B[ae];W[ad];B[af];W[be];B[ab];W[bf];B[ag];W[cg];B[ch];W[dh];B[di];W[eg];B[dg];W[df];B[dh];W[bc];B[ac];W[bb];B[bd];W[ee];B[fe];W[ad];B[kf];W[jf];B[bd];W[eb];B[ec];W[ad];B[jh];W[jg];B[bd];W[md];B[nd];W[ad];B[hg];W[gg];B[bd];W[oe];B[ne];W[ad];B[fl];W[gl];B[bd];W[ng];B[nf];W[ad];B[ap];W[aq];B[bd];W[oh];B[pb];W[ad];B[ao];W[br];B[bd];W[mb];B[nb];W[ad];B[bq];W[cq];B[bd];W[rg];B[oi];W[ad];B[ni];W[mi];B[bd];W[of];B[pg];W[ad];B[ar];W[as];B[bd];W[qd];B[pe];W[ad];B[iq];W[jp];B[bd];W[oa];B[pa];W[ad];B[qa];W[aa];B[rb];W[sm];B[nr];W[ms];B[nq];W[qn];B[po];W[ek];B[el];W[hd];B[gd];W[da];B[ea];W[fa];B[fb];W[mc];B[ma];W[la];B[na];W[dk];B[ck];W[dl];B[dm];W[ls];B[ir];W[mn];B[mo];W[mm];B[nl];W[mk];B[ll];W[nn];B[on];W[mp];B[np];W[lo];B[lm];W[ln];B[le];W[ld];B[rc];W[sc];B[ns];W[sb];B[ra];W[co];B[bn];W[ks];B[js];W[gs];B[hs];W[fs];B[sn];W[rl];B[rn];W[ca];B[ea];W[ja];B[tt])(;B[aa]))"
	kifu = ParseSgf(sgf)
	kifu.Last()

	assert.Equal(t, kifu.CurNode.GetColor(), "B")
	assert.Equal(t, kifu.CurNode.Parent.GetColor(), "W")
	assert.Equal(t, kifu.Root.GetColor(), "")
	assert.Equal(t, len(kifu.ToSgfList()), 2)
	assert.Equal(t, kifu.ToSgf(), "(;SZ[19]KM[7.5]HA[0]AB[tt](;B[pd];W[dd];B[qp];W[dq];B[oq];W[qf];B[nc];W[rd];B[fc];W[cf];B[qh];W[qc];B[qe];W[re];B[pf];W[pg];B[qg];W[rf];B[og];W[hc];B[pj];W[jq];B[cn];W[cp];B[dj];W[en];B[cl];W[ci];B[cj];W[mq];B[db];W[cc];B[ed];W[kc];B[de];W[ce];B[dc];W[cd];B[he];W[ke];B[ib];W[ic];B[hb];W[jb];B[gc];W[ia];B[ga];W[ie];B[no];W[qo];B[rp];W[ol];B[ml];W[ql];B[qk];W[oj];B[pl];W[pm];B[pk];W[ok];B[om];W[pi];B[qi];W[rk];B[qj];W[rm];B[pn];W[qm];B[ro];W[rj];B[ri];W[sj];B[sl];W[si];B[sh];W[sg];B[ph];W[rh];B[hq];W[io];B[ep];W[eq];B[fq];W[go];B[fo];W[gp];B[fp];W[fn];B[do];W[fr];B[gr];W[er];B[gq];W[ip];B[bi];W[cb];B[bg];W[hf];B[ge];W[mj];B[nm];W[nh];B[mf];W[lg];B[mg];W[mh];B[lf];W[kg];B[pc];W[qb];B[bo];W[bp];B[kr];W[kq];B[lr];W[mr];B[jr];W[lq];B[gn];W[gm];B[hn];W[il];B[fm];W[hm];B[em];W[in];B[dn];W[ho];B[lk];W[kj];B[hh];W[gf];B[ff];W[fg];B[ef];W[gh];B[gi];W[fi];B[gj];W[fj];B[gk];W[fk];B[jk];W[kk];B[jl];W[kl];B[jm];W[km];B[jj];W[ig];B[ki];W[lj];B[jn];W[kn];B[jo];W[ko];B[eh];W[fh];B[ae];W[ad];B[af];W[be];B[ab];W[bf];B[ag];W[cg];B[ch];W[dh];B[di];W[eg];B[dg];W[df];B[dh];W[bc];B[ac];W[bb];B[bd];W[ee];B[fe];W[ad];B[kf];W[jf];B[bd];W[eb];B[ec];W[ad];B[jh];W[jg];B[bd];W[md];B[nd];W[ad];B[hg];W[gg];B[bd];W[oe];B[ne];W[ad];B[fl];W[gl];B[bd];W[ng];B[nf];W[ad];B[ap];W[aq];B[bd];W[oh];B[pb];W[ad];B[ao];W[br];B[bd];W[mb];B[nb];W[ad];B[bq];W[cq];B[bd];W[rg];B[oi];W[ad];B[ni];W[mi];B[bd];W[of];B[pg];W[ad];B[ar];W[as];B[bd];W[qd];B[pe];W[ad];B[iq];W[jp];B[bd];W[oa];B[pa];W[ad];B[qa];W[aa];B[rb];W[sm];B[nr];W[ms];B[nq];W[qn];B[po];W[ek];B[el];W[hd];B[gd];W[da];B[ea];W[fa];B[fb];W[mc];B[ma];W[la];B[na];W[dk];B[ck];W[dl];B[dm];W[ls];B[ir];W[mn];B[mo];W[mm];B[nl];W[mk];B[ll];W[nn];B[on];W[mp];B[np];W[lo];B[lm];W[ln];B[le];W[ld];B[rc];W[sc];B[ns];W[sb];B[ra];W[co];B[bn];W[ks];B[js];W[gs];B[hs];W[fs];B[sn];W[rl];B[rn];W[ca];B[ea];W[ja];B[tt])(;B[aa]))")
	assert.Equal(t, kifu.ToCurSgf(), "(;SZ[19]KM[7.5]HA[0]AB[tt];B[pd];W[dd];B[qp];W[dq];B[oq];W[qf];B[nc];W[rd];B[fc];W[cf];B[qh];W[qc];B[qe];W[re];B[pf];W[pg];B[qg];W[rf];B[og];W[hc];B[pj];W[jq];B[cn];W[cp];B[dj];W[en];B[cl];W[ci];B[cj];W[mq];B[db];W[cc];B[ed];W[kc];B[de];W[ce];B[dc];W[cd];B[he];W[ke];B[ib];W[ic];B[hb];W[jb];B[gc];W[ia];B[ga];W[ie];B[no];W[qo];B[rp];W[ol];B[ml];W[ql];B[qk];W[oj];B[pl];W[pm];B[pk];W[ok];B[om];W[pi];B[qi];W[rk];B[qj];W[rm];B[pn];W[qm];B[ro];W[rj];B[ri];W[sj];B[sl];W[si];B[sh];W[sg];B[ph];W[rh];B[hq];W[io];B[ep];W[eq];B[fq];W[go];B[fo];W[gp];B[fp];W[fn];B[do];W[fr];B[gr];W[er];B[gq];W[ip];B[bi];W[cb];B[bg];W[hf];B[ge];W[mj];B[nm];W[nh];B[mf];W[lg];B[mg];W[mh];B[lf];W[kg];B[pc];W[qb];B[bo];W[bp];B[kr];W[kq];B[lr];W[mr];B[jr];W[lq];B[gn];W[gm];B[hn];W[il];B[fm];W[hm];B[em];W[in];B[dn];W[ho];B[lk];W[kj];B[hh];W[gf];B[ff];W[fg];B[ef];W[gh];B[gi];W[fi];B[gj];W[fj];B[gk];W[fk];B[jk];W[kk];B[jl];W[kl];B[jm];W[km];B[jj];W[ig];B[ki];W[lj];B[jn];W[kn];B[jo];W[ko];B[eh];W[fh];B[ae];W[ad];B[af];W[be];B[ab];W[bf];B[ag];W[cg];B[ch];W[dh];B[di];W[eg];B[dg];W[df];B[dh];W[bc];B[ac];W[bb];B[bd];W[ee];B[fe];W[ad];B[kf];W[jf];B[bd];W[eb];B[ec];W[ad];B[jh];W[jg];B[bd];W[md];B[nd];W[ad];B[hg];W[gg];B[bd];W[oe];B[ne];W[ad];B[fl];W[gl];B[bd];W[ng];B[nf];W[ad];B[ap];W[aq];B[bd];W[oh];B[pb];W[ad];B[ao];W[br];B[bd];W[mb];B[nb];W[ad];B[bq];W[cq];B[bd];W[rg];B[oi];W[ad];B[ni];W[mi];B[bd];W[of];B[pg];W[ad];B[ar];W[as];B[bd];W[qd];B[pe];W[ad];B[iq];W[jp];B[bd];W[oa];B[pa];W[ad];B[qa];W[aa];B[rb];W[sm];B[nr];W[ms];B[nq];W[qn];B[po];W[ek];B[el];W[hd];B[gd];W[da];B[ea];W[fa];B[fb];W[mc];B[ma];W[la];B[na];W[dk];B[ck];W[dl];B[dm];W[ls];B[ir];W[mn];B[mo];W[mm];B[nl];W[mk];B[ll];W[nn];B[on];W[mp];B[np];W[lo];B[lm];W[ln];B[le];W[ld];B[rc];W[sc];B[ns];W[sb];B[ra];W[co];B[bn];W[ks];B[js];W[gs];B[hs];W[fs];B[sn];W[rl];B[rn];W[ca];B[ea];W[ja];B[tt])")
	assert.Equal(t, CoorToBoardNode(8, 0, 19), "J19")
	assert.Equal(t, CoorToBoardNode(-1, -1, 19), "pass")
	log := `Thinking at most 36.0 seconds...
NN eval=0.468024

 Q17 ->       2 (V: 50.69%) (N:  9.74%) PV: Q17 R4
 C16 ->       0 (V:  0.00%) (N:  9.66%) PV: C16 
2.0 average depth, 3 max depth
2 non leaf nodes, 1.00 average children
3 visits, 1083 nodes, 2 playouts, 13 n/s
`
	result, wr := ParseLZOutput(log, kifu.Size)
	assert.Equal(t, len(result), 2)
	assert.Equal(t,wr,50.69)
	assert.Equal(t, result[0].Times, 2)
	result, _ = ParseLZOutput(log, kifu.Size,1)
	assert.Equal(t, len(result[0].Diagram),1)

	log = `  0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0 999   0   0   0   0
 0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0   0
pass: 0
winrate: 1.000000
`
	heatMap, pass, wr := ParseLZHeatMap(log)
	assert.Equal(t, pass, 0.0)
	assert.Equal(t, len(heatMap), 361)
	assert.Equal(t, heatMap[337], 999.0)
	assert.Equal(t, wr, 1.0)
	x, y := StoneToXY("J19", 19)
	assert.Equal(t, x, 8)
	assert.Equal(t, y, 0)
	x, y = StoneToXY("PASS", 19)
	assert.Equal(t, x, -1)
	assert.Equal(t, y, -1)
	x, y = StoneToXY("", 19)
	assert.Equal(t, x, -1)
	assert.Equal(t, y, -1)
	x, y = StoneToXY("JJ", 19)
	assert.Equal(t, x, -1)
	assert.Equal(t, y, -1)
	x, y = StoneToXY("D16", 19)

	sgf = "(;SZ[19]KM[7.5]HA[0]AB[aa];B[pd];W[cc])"
	kifu = ParseSgf(sgf)
	assert.Equal(t, kifu.NodeCount, 2)
	assert.Equal(t,kifu.ToCleanSgf(),"(;SZ[19]AB[aa];B[pd];W[cc])")
	kifu.Last()
	assert.Equal(t,kifu.ToCurSgf(),"(;SZ[19]KM[7.5]HA[0]AB[aa];B[pd];W[cc])")
	assert.Equal(t,kifu.ToSgfByNode(kifu.CurNode),"(;SZ[19]KM[7.5]HA[0]AB[aa];B[pd];W[cc])")
	tv:=0
	kifu.EachNode(func(n *Node, move int)bool {
		tv++
		return false
	})
	assert.Equal(t,tv,kifu.NodeCount+1)
	tv=0
	kifu.EachNode(func(n *Node, move int)bool {
		tv++
		return true
	})
	assert.Equal(t,tv,1)
	kifu.GoTo(2)
	kifu.Remove()
	assert.Equal(t,kifu.ToSgfByNode(kifu.CurNode),"(;SZ[19]KM[7.5]HA[0]AB[aa];B[pd])")

	sgf="(;SZ[19];B[tt];W[aa])"
	kifu=ParseSgf(sgf)
	kifu.GoTo(1)
	assert.Equal(t,kifu.CurNode.IsPass(),true)
	kifu.Last()
	assert.Equal(t,kifu.CurNode.IsPass(),false)
	kifu.CurNode.AddInfo("C","this is comment")
	kifu.CurNode.AddInfo("WR",9.111)
	assert.Equal(t,kifu.CurNode.GetInfo("C"),"this is comment")
	assert.Equal(t,kifu.CurNode.GetInfo("WR"),"9.1")
	assert.Equal(t,kifu.CurNode.GetInfo("BR"),"")
	assert.Equal(t,kifu.CurPos.GetColor(0,0),-1)

	sgf="(;SZ[19];B[])"
	kifu=ParseSgf(sgf)
	kifu.Last()
	assert.Equal(t,kifu.CurNode.X,-1)

	position:=NewPosition(19)
	position.SetRevert(true)
	position.SetColor(1,1,1)
	position.SetColor(2,2,-1)
	assert.Equal(t,position.GetColor(1,1),1)
	assert.Equal(t,position.ShowBoard(false),`  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  X  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  O  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .`)
	assert.Equal(t,position.ShowBoard(true),`     A  B  C  D  E  F  G  H  J  K  L  M  N  O  P  Q  R  S  T
 19  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 18  .  X  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 17  .  .  O  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 16  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 15  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 14  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 13  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 12  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 11  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
 10  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  9  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  8  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  7  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  6  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  5  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  4  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  3  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  2  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .
  1  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .  .`)
	black,white:=position.GetStones()
	assert.Equal(t,strings.Join(black,""),"bb")
	assert.Equal(t,strings.Join(white,""),"cc")

	kifu=ParseSgf("(;SZ[19];B[tt];W[tt])")
	kifu.Last()
	assert.Equal(t,true,kifu.DoublePass)
	TestLastCheck(t)
	TestCapGame(t)
}

func TestLastCheck(t *testing.T)  {
	sgf:="(;SZ[19];B[aa];W[dd])"
	kifu:=ParseSgf(sgf)
	assert.Equal(t,kifu.LastAndCheck(),nil)
	assert.Equal(t,kifu.CurColor,1)
	sgf="(;SZ[19]AP[WGo.js:2]FF[4]GM[1]CA[UTF-8];B[ab];W[db];B[ba];W[aa])"
	kifu=ParseSgf(sgf)
	assert.Equal(t,kifu.LastAndCheck()!=nil,true)
	assert.Equal(t,kifu.CurColor,-1)
}
func TestCapGame(t *testing.T)  {
	sgf:="(;SZ[19]AP[WGo.js:2]FF[4]GM[1]CA[UTF-8];B[ba];W[aa])"
	kifu:=ParseSgf(sgf)
	kifu.Last()
	println(kifu.CurPos.ShowBoard())
	n,i:=kifu.CurPos.CalcCap(kifu.CurColor,kifu.GetLiberty())
	fmt.Println(n,i)
}
func TestParseLZOutputV17(t *testing.T) {
	text:=`Thinking at most 36.0 seconds...
NN eval=0.428765

 R16 ->       5 (V: 42.93%) (LCB: 34.93%) (N:  3.14%) PV: R16 D16 Q4 D4
 Q16 ->       4 (V: 43.28%) (LCB: 29.06%) (N: 16.28%) PV: Q16 D4 Q4 D16
  D4 ->       4 (V: 43.34%) (LCB: 29.01%) (N: 16.68%) PV: D4 Q16 Q4 D16
 D16 ->       4 (V: 43.60%) (LCB: 26.91%) (N: 16.02%) PV: D16 Q4 Q16
  C4 ->       4 (V: 43.12%) (LCB: 26.47%) (N:  3.23%) PV: C4 Q4 D16 Q16
  Q3 ->       4 (V: 43.23%) (LCB: 25.80%) (N:  3.16%) PV: Q3 Q16 D4 D16
  Q4 ->       3 (V: 43.59%) (LCB:  0.00%) (N: 16.50%) PV: Q4 D16 Q16
  D3 ->       3 (V: 43.41%) (LCB:  0.00%) (N:  3.24%) PV: D3 D16 Q4
  R4 ->       3 (V: 43.14%) (LCB:  0.00%) (N:  3.16%) PV: R4 D4 Q16
 Q17 ->       3 (V: 43.12%) (LCB:  0.00%) (N:  3.21%) PV: Q17 Q4 D16
3.2 average depth, 5 max depth
26 non leaf nodes, 1.42 average children
38 visits, 13634 nodes, 37 playouts, 617 n/s
`
	outputs, f := ParseLZOutputV17(text, 19, 3)
	fmt.Println(outputs,f)
}