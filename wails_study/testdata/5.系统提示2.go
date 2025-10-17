package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"gopkg.in/toast.v1"
	"log"
)

func ToGBK(s string) string {
	result, _, err := transform.String(simplifiedchinese.GBK.NewEncoder(), s)
	if err != nil {
		return s // 转换失败时返回原字符串
	}
	return result
}

func main() {
	notification := toast.Notification{
		AppID:   "fengfengzhidao",
		Title:   ToGBK("你收到一条消息"),
		Message: ToGBK("有人给你点赞了"),
		Icon:    "E:\\IT\\go_pro\\webview_study\\wails_study\\build\\appicon.png",
		Actions: []toast.Action{
			{
				Type:      "protocol",                                 // 协议类型
				Label:     ToGBK("查看详情"),                              // 按钮文字
				Arguments: "https://www.fengfengzhidao.com?name=xxxx", // 点击后打开该网页
			},
		},
		Audio: toast.IM,
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(err)
}
