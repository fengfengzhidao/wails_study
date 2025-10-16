package main

import (
	"fmt"
	"github.com/MakeNowJust/hotkey"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	hk1 "golang.design/x/hotkey"
	"golang.design/x/hotkey/mainthread"
)

func (a *App) hotKey() {
	hkey := hotkey.New()

	hkey.Register(hotkey.Ctrl, 'J', func() {
		text, err := runtime.ClipboardGetText(a.ctx)
		fmt.Println("全局快捷键触发 剪贴板的数据：", text, err)
	})

	mainthread.Init(func() {
		hk := hk1.New([]hk1.Modifier{hk1.ModCtrl, hk1.ModShift}, hk1.KeyQ)
		err := hk.Register()
		if err != nil {
			fmt.Println(err)
			return
		}

		<-hk.Keydown()
		<-hk.Keyup()
		fmt.Println("全局快捷键触发 退出程序")
		hk.Unregister()
		runtime.Quit(a.ctx)
	})

}
