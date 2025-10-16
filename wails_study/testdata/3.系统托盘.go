package main

import (
	_ "embed"
	"fmt"
	"github.com/energye/systray"
)

func main() {
	fmt.Println("运行中")
	systray.Run(onReady, onExit)
}

//go:embed icon.ico
var homeIcon []byte

//go:embed favicon.ico
var i1 []byte

//go:embed app.ico
var i2 []byte

func onReady() {
	systray.SetIcon(homeIcon)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	systray.SetOnClick(func(menu systray.IMenu) {
		fmt.Println("单击")
	})
	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})

	menu1 := systray.AddMenuItem("打开主界面", "")
	menu1.SetIcon(i1)
	menu1.Click(func() {
		fmt.Println("打开主界面")
	})

	menu2 := systray.AddMenuItem("基本设置", "")
	menu2.SetIcon(i2)
	menu2.Click(func() {
		fmt.Println("基本设置")
	})

	systray.AddMenuItem("退出", "").Click(func() {
		onExit()
	})
}

func onExit() {
	// clean up here
	systray.Quit()
}
