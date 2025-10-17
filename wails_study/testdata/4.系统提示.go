package main

import (
	"github.com/martinlindhe/notify"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func toGBK(s string) string {
	result, _, err := transform.String(simplifiedchinese.GBK.NewEncoder(), s)
	if err != nil {
		return s // 转换失败时返回原字符串
	}
	return result
}

func main() {
	notify.Alert("fengfengzhidao", toGBK("收到一条消息"), toGBK("有人给你点赞了"), "E:\\IT\\go_pro\\webview_study\\wails_study\\build\\appicon.png")
	//notify.Alert("APP", "title", "msg", "")
	//notify.Alert("APP", "title", "msg", "E:\\IT\\go_pro\\webview_study\\wails_study\\testdata\\icon.ico")
}
