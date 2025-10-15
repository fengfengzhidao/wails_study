package main

import (
	webview "github.com/webview/webview_go"
)

func main() {
	w := webview.New(true) // 如果是true 就可以打开f12
	defer w.Destroy()
	w.SetTitle("枫枫知道")
	w.SetSize(1200, 600, webview.HintNone)
	w.Navigate("https://www.fengfengzhidao.com")
	w.Run()
}
