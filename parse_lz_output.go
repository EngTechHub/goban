package goban

import (
	"strings"
	"strconv"
)

// 解析leelazero 数据
func ParseLZOutput(output string,size int,move string) ([]*AIOutput,float64) {
	result := make([]*AIOutput, 0)
	lines := strings.Split(output, "\n")
	rate:=0.0
	//result := make([]map[string]interface{}, 0)
	for _, v := range lines {
		if strings.Contains(v, "->") {
			item := &AIOutput{} //make(map[string]interface{})
			first := strings.Split(v, "->")
			//选点
			item.Select = strings.TrimSpace(first[0])

			second := strings.Split(strings.TrimSpace(first[1]), "(")

			// 模拟次数
			times := strings.TrimSpace(second[0])
			t, err := strconv.Atoi(times)
			if err != nil {
				item.Times = 0
			} else {
				item.Times = t
			}
			// 胜率
			wineRateS := strings.TrimSpace(strings.Replace(strings.Replace(second[1], "V:", "", -1), "%)", "", -1))
			wineRate, err := strconv.ParseFloat(wineRateS, 64)
			if err != nil {
				item.WineRate = 0
			} else {
				item.WineRate = wineRate
			}
			if item.Select==move{
				rate=item.WineRate
			}
			three := strings.Split(strings.Replace(second[2], "N:", "", -1), "%)")
			// 策略网络概率
			chanceS := strings.TrimSpace(three[0])
			chance, err := strconv.ParseFloat(chanceS, 64)
			if err != nil {
				item.Chance = 0
			} else {
				item.Chance = chance
			}
			//变化图
			four := strings.Fields(strings.TrimSpace(three[1]))
			if len(four) > 0 && four[0] == "PV:" {
				diagram:=make([]string,0)
				for _,v:=range four[1:]{
					x,y:=StoneToXY(v,size)
					diagram=append(diagram, CoorToSgfNode(x,y))
				}
				item.Diagram = diagram
			}
			result = append(result, item)
		}
	}
	return result,rate
}

//解析leelazero heatmap
func ParseLZHeatMap(heatmap string) ([362]float64, float64) {
	position := [362]float64{}
	wineRate := 0.0
	for x, v := range strings.Split(heatmap, "\n") {
		lines := strings.Fields(v)
		switch len(lines) {
		case 19:
			for y, p := range lines {
				pp, _ := strconv.ParseFloat(p, 64)
				position[x+y*19] = pp
			}
		case 2:
			if lines[0] == "pass:" {
				pp, _ := strconv.ParseFloat(lines[1], 64)
				position[361] = pp
			} else if lines[0] == "winrate:" {
				rate, _ := strconv.ParseFloat(lines[1], 64)
				wineRate = rate
			}
		}
	}
	return position, wineRate
}
