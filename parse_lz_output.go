package goban

import (
	"strconv"
	"strings"
)

// 解析leelazero 数据
func ParseLZOutput(output string, size int, limits ...int) ([]*AIOutput, float64) {
	lim := 10
	if len(limits) > 0 {
		lim = limits[0]
	}
	result := make([]*AIOutput, 0)
	lines := strings.Split(output, "\n")
	rate := 0.0
	isFirst := true
	for _, v := range lines {
		if strings.Contains(v, "->") {

			item := &AIOutput{} //make(map[string]interface{})
			first := strings.Split(v, "->")
			//选点
			item.Select = strings.TrimSpace(first[0])

			second := strings.Split(strings.TrimSpace(first[1]), "(")

			// 模拟次数
			times := strings.TrimSpace(second[0])
			t, _ := strconv.Atoi(times)
			item.Times = t
			// 胜率
			wineRateS := strings.TrimSpace(strings.Replace(strings.Replace(second[1], "V:", "", -1), "%)", "", -1))
			wineRate, _ := strconv.ParseFloat(wineRateS, 64)
			item.WineRate = wineRate
			if isFirst {
				rate = item.WineRate
				isFirst = false
			}
			three := strings.Split(strings.Replace(second[2], "N:", "", -1), "%)")
			// 策略网络概率
			chanceS := strings.TrimSpace(three[0])
			chance, _ := strconv.ParseFloat(chanceS, 64)
			item.Chance = chance
			//变化图
			four := strings.Fields(strings.TrimSpace(three[1]))
			if len(four) > 0 && four[0] == "PV:" {
				diagram := make([]string, 0)
				for j, v := range four[1:] {
					if j >= lim {
						break
					}
					x, y := StoneToXY(v, size)
					diagram = append(diagram, CoorToSgfNode(x, y))
				}
				item.Diagram = diagram
			}
			result = append(result, item)
		}
	}
	return result, rate
}

//解析leelazero heatmap
func ParseLZHeatMap(heatmap string) ([]float64, float64, float64) {
	position := [361]float64{}
	wineRate := 0.0
	pass := 0.0
	for y, v := range strings.Split(heatmap, "\n") {
		lines := strings.Fields(v)
		switch len(lines) {
		case 19:
			for x, p := range lines {
				pp, _ := strconv.ParseFloat(p, 64)
				position[x+y*19] = pp
			}
		case 2:
			if lines[0] == "pass:" {
				pp, _ := strconv.ParseFloat(lines[1], 64)
				pass = pp
			} else if lines[0] == "winrate:" {
				rate, _ := strconv.ParseFloat(lines[1], 64)
				wineRate = rate
			}
		}
	}
	return position[:], pass, wineRate
}

//解析leelazero heatmap
func ParseHeatMap(heatmap string, size int) ([]float64, float64, float64) {
	position := [size * size]float64{}
	wineRate := 0.0
	pass := 0.0
	for y, v := range strings.Split(heatmap, "\n") {
		lines := strings.Fields(v)
		switch len(lines) {
		case size:
			for x, p := range lines {
				pp, _ := strconv.ParseFloat(p, 64)
				position[x+y*size] = pp
			}
		case 2:
			if lines[0] == "pass:" {
				pp, _ := strconv.ParseFloat(lines[1], 64)
				pass = pp
			} else if lines[0] == "winrate:" {
				rate, _ := strconv.ParseFloat(lines[1], 64)
				wineRate = rate
			}
		}
	}
	return position[:], pass, wineRate
}
