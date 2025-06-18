package goban

import (
	"fmt"
	"strings"
	"strconv"
)

// ToNum 根据字符串转为ASCII码
func StrToASCII(value string, index int) int {
	if value == "" {
		return -1
	}else if value=="tt"{
		return -1
	}
	temp := string(value[index])
	xr := []rune(temp)
	return int(xr[0]) - 97
}

//坐标转为sgf中的坐标字符串
func CoorToSgfNode(x, y int) string {
	xChar := fmt.Sprintf("%s", string(rune('a'+x)))
	yChar := fmt.Sprintf("%s", string(rune('a'+y)))
	if x == -1 {
		xChar = "t"
	}

	if y == -1 {
		yChar = "t"
	}
	return fmt.Sprintf("%s%s", xChar, yChar)
}

//坐标转为棋盘中的坐标字符串
func CoorToBoardNode(x, y, size int) string {
	if x == -1 && y == -1 {
		return "pass"
	}
	if x >= 8 {
		x++
	}
	return strings.ToUpper(fmt.Sprintf("%s%d", string(rune('a'+x)), size-y))
}

// 棋盘坐标转为X,Y坐标
func StoneToXY(move string, size int) (int, int) {
	move = strings.TrimSpace(move)
	if move == "" {
		return -1, -1
	} else if "pass" == strings.ToLower(move) {
		return -1, -1
	}
	temp := strings.ToLower(string(move[0]))
	xr := []rune(temp)
	xInt := xr[0]
	if xInt > 105 {
		xInt = xInt - 1
	}
	xInt = xInt - 97
	y, err := strconv.Atoi(string(move[1:]))
	//utils.CheckError(err)
	if err != nil {
		return -1, -1
	}
	yInt := size - y
	return int(xInt), yInt
}

