package goban

type AIOutput struct {
	WineRate float64  `json:"wine_rate"` //胜率
	Select   string   `json:"select"`    //选点
	Times    int      `json:"times"`     //模拟次数
	Chance   float64  `json:"chance"`    //神经网络概率
	Diagram  []string `json:"branches"`  //变化图
	LCB float64 `json:"lcb"` //lower confidence bounds 值
}
