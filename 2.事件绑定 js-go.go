package main

import (
	"fmt"
	webview "github.com/webview/webview_go"
	"time"
)

func main() {
	w := webview.New(true) // 如果是true 就可以打开f12
	defer w.Destroy()
	w.SetTitle("枫枫知道")
	w.SetSize(1200, 600, webview.HintNone)
	w.SetHtml(`
<h1>hello</h1>
<button onclick="showDate()">show Date 点我</button>
<button onclick="add(1, 2)">add 点我</button>
<button onclick="gu()">getUser 点我</button>
<script>
async function gu(){
	const u = await getUser()
	alert("用户id=" + u)
}
</script>
`)
	w.Bind("showDate", func() {
		fmt.Println(time.Now().Format(time.DateTime))
	})
	w.Bind("add", func(n1, n2 int) {
		fmt.Println("add: ", n1, n2)
	})
	w.Bind("getUser", func() string {
		userID := "xxx001"
		fmt.Println("getUser", userID)
		return userID
	})
	w.Run()
}
