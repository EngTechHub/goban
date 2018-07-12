package main

import (
	"gitee.com/larry_dev/goban"
	"io/ioutil"
)

func main() {
	v, _ := ioutil.ReadFile("./all_1280184.sgf")
	kifu := goban.ParseSgf(string(v))
	kifu.Last()
	println(kifu.ToSgf())
}