package chinaNumber

import (
	"math"
	"strings"
)

var chnNumChar = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
var chnUnitSection = []string{"", "万", "亿", "万亿", "亿亿"}
var chnUnitChar = []string{"", "十", "百", "千"}

var chnNumCharDict = map[string]int{"零": 0, "一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9}
var chnNameValue = map[string]struct {
	value   int
	secUnit bool
}{"十": {value: 10, secUnit: false}, "百": {value: 100, secUnit: false}, "千": {value: 1000, secUnit: false}, "万": {value: 10000, secUnit: true}, "亿": {value: 100000000, secUnit: true}}

func SectionToChinese(section int) string {
	var strIns = ""
	var chnStr = ""
	var unitPos = 0
	var zero = true
	for {
		if section <= 0 {
			break
		}
		var v = int(section) % 10
		if v == 0 {
			if !zero {
				zero = true
				chnStr = chnNumChar[v] + chnStr
			}
		} else {
			zero = false
			strIns = chnNumChar[v]
			strIns += chnUnitChar[unitPos]
			chnStr = strIns + chnStr
		}
		unitPos++
		section = int(math.Floor(float64(section) / 10))
	}
	return chnStr
}

func NumberToChinese(num int) string {
	var unitPos = 0
	var strIns = ""
	var chnStr = ""
	var needZero = false

	if num == 0 {
		return chnNumChar[0]
	}
	for {
		if num <= 0 {
			break
		}
		var section = num % 10000
		if needZero {
			chnStr = chnNumChar[0] + chnStr
		}
		strIns = SectionToChinese(section)
		if section != 0 {
			strIns += chnUnitSection[unitPos]
		} else {
			strIns += chnUnitSection[0]
		}

		chnStr = strIns + chnStr
		needZero = (section < 1000) && (section > 0)
		num = int(math.Floor(float64(num) / 10000))
		unitPos++
	}
	return chnStr
}

func ChineseToNumber(chnStr string) int {
	var rtn = 0
	var section = 0
	var number = 0
	var secUnit = false
	var str = strings.Split(chnStr, "")
	for i := 0; i < len(str); i++ {
		var num = chnNumCharDict[str[i]]
		if num != 0 {
			number = num
			if i == len(str)-1 {
				section += number
			}
		} else {
			var units = chnNameValue[str[i]].value
			secUnit = chnNameValue[str[i]].secUnit
			if secUnit {
				section = (section + number) * units
				rtn += section
				section = 0
			} else {
				section += (number * units)
			}
			number = 0
		}
	}
	return rtn + section
}
