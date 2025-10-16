package main

import (
	_ "embed"
	"fmt"
	"github.com/energye/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) systray() {
	systray.Run(a.onReady, a.onExit)
}

//go:embed testdata/icon.ico
var homeIcon []byte

func (a *App) onReady() {
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
	menu1.Click(func() {
		fmt.Println("打开主界面")
		runtime.Show(a.ctx)
	})

	menu2 := systray.AddMenuItem("隐藏", "")
	menu2.Click(func() {
		fmt.Println("隐藏")
		runtime.Hide(a.ctx)
	})

	systray.AddMenuItem("退出", "").Click(func() {
		a.onExit()
	})
}

func (a *App) onExit() {
	// clean up here
	systray.Quit()
	runtime.Quit(a.ctx)
}
