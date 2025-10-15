package main

import (
	_ "embed"
	webview "github.com/webview/webview_go"
	"time"
)

//go:embed index.html
var html string

func main() {
	w := webview.New(true) // 如果是true 就可以打开f12
	defer w.Destroy()
	w.SetTitle("枫枫知道")
	w.SetSize(1200, 600, webview.HintNone)
	w.SetHtml(html)

	go func() {
		time.Sleep(2 * time.Second)
		w.Dispatch(func() {
			w.Eval("showAlert()")
			w.Eval("showText('这是go传递来的数据')")
		})

	}()

	w.Run()
}
