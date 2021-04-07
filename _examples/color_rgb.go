package main

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

// go run ./_examples/color_rgb.go
// FORCE_COLOR=on go run ./_examples/color_rgb.go
func main() {
	color.RGB(30, 144, 255).Println("message. use RGB number")
	color.HEX("#1976D2").Println("blue-darken")
	color.RGBStyleFromString("213,0,0").Println("red-accent. use RGB number")
	// foreground: eee, background: D50000
	color.HEXStyle("eee", "D50000").Println("deep-purple color")

	color.Infoln("\n==== Chinese traditional colors ====\n")
	index := 1
	for _, txt := range chinaColors {
		nodes := strings.Split(txt, " ")
		color.HEXStyle("ccc", nodes[1]).Printf("%-12s", nodes[1])
		if index%9 == 0 {
			fmt.Println()
		}

		index++
	}
	fmt.Println()
}

// from http://tools.jb51.net/color/chinacolor
var chinaColors = []string{
	"蔚蓝 #70f3ff",
	"蓝 #44cef6",
	"碧蓝 #3eede7",
	"石青 #1685a9",
	"靛青 #177cb0",
	"靛蓝 #065279",
	"花青 #003472",
	"宝蓝 #4b5cc4",
	"蓝灰色 #a1afc9",
	"藏青 #2e4e7e",
	"藏蓝 #3b2e7e",
	"黛 #4a4266",
	"黛绿 #426666",
	"黛蓝 #425066",
	"黛紫 #574266",
	"紫色 #8d4bbb",
	"紫酱 #815463",
	"酱紫 #815476",
	"紫檀 #4c221b",
	"绀青 #003371",
	"紫棠 #56004f",
	"青莲 #801dae",
	"群青 #4c8dae",
	"雪青 #b0a4e3",
	"丁香色 #cca4e3",
	"藕色 #edd1d8",
	"藕荷色 #e4c6d0",
	"朱砂 #ff461f",
	"火红 #ff2d51",
	"朱膘 #f36838",
	"妃色 #ed5736",
	"洋红 #ff4777",
	"品红 #f00056",
	"粉红 #ffb3a7",
	"桃红 #f47983",
	"海棠红 #db5a6b",
	"樱桃色 #c93756",
	"酡颜 #f9906f",
	"银红 #f05654",
	"大红 #ff2121",
	"石榴红 #f20c00",
	"绛紫 #8c4356",
	"绯红 #c83c23",
	"胭脂 #9d2933",
	"朱红 #ff4c00",
	"丹 #ff4e20",
	"彤 #f35336",
	"酡红 #dc3023",
	"炎 #ff3300",
	"茜色 #cb3a56",
	"绾 #a98175",
	"檀 #b36d61",
	"嫣红 #ef7a82",
	"洋红 #ff0097",
	"枣红 #c32136",
	"殷红 #be002f",
	"赫赤 #c91f37",
	"银朱 #bf242a",
	"赤 #c3272b",
	"胭脂 #9d2933",
	"栗色 #60281e",
	"玄色 #622a1d",
	"松花色 #bce672",
	"柳黄 #c9dd22",
	"嫩绿 #bddd22",
	"柳绿 #afdd22",
	"葱黄 #a3d900",
	"葱绿 #9ed900",
	"豆绿 #9ed048",
	"豆青 #96ce54",
	"油绿 #00bc12",
	"葱倩 #0eb83a",
	"葱青 #0eb83a",
	"青葱 #0aa344",
	"石绿 #16a951",
	"松柏绿 #21a675",
	"松花绿 #057748",
	"绿沈 #0c8918",
	"绿色 #00e500",
	"草绿 #40de5a",
	"青翠 #00e079",
	"青色 #00e09e",
	"翡翠色 #3de1ad",
	"碧绿 #2add9c",
	"玉色 #2edfa3",
	"缥 #7fecad",
	"艾绿 #a4e2c6",
	"石青 #7bcfa6",
	"碧色 #1bd1a5",
	"青碧 #48c0a3",
	"铜绿 #549688",
	"竹青 #789262",
	"墨灰 #758a99",
	"墨色 #50616d",
	"鸦青 #424c50",
	"黯 #41555d",
	"樱草色 #eaff56",
	"鹅黄 #fff143",
	"鸭黄 #faff72",
	"杏黄 #ffa631",
	"橙黄 #ffa400",
	"橙色 #fa8c35",
	"杏红 #ff8c31",
	"橘黄 #ff8936",
	"橘红 #ff7500",
	"藤黄 #ffb61e",
	"姜黄 #ffc773",
	"雌黄 #ffc64b",
	"赤金 #f2be45",
	"缃色 #f0c239",
	"雄黄 #e9bb1d",
	"秋香色 #d9b611",
	"金色 #eacd76",
	"牙色 #eedeb0",
	"枯黄 #d3b17d",
	"黄栌 #e29c45",
	"乌金 #a78e44",
	"昏黄 #c89b40",
	"棕黄 #ae7000",
	"琥珀 #ca6924",
	"棕色 #b25d25",
	"茶色 #b35c44",
	"棕红 #9b4400",
	"赭 #9c5333",
	"驼色 #a88462",
	"秋色 #896c39",
	"棕绿 #827100",
	"褐色 #6e511e",
	"棕黑 #7c4b00",
	"赭色 #955539",
	"赭石 #845a33",
}
