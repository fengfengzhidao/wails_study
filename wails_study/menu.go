package main

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
)

func (a *App) getMenu() *menu.Menu {
	m := menu.NewMenu()
	fileMenu := m.AddSubmenu("文件")
	fileMenu.AddText("打开文件", keys.Control("o"), func(data *menu.CallbackData) {
		fmt.Println("打开文件")
	})
	fileMenu.AddText("保存文件", &keys.Accelerator{Key: "s", Modifiers: []keys.Modifier{
		keys.ControlKey,
		keys.ShiftKey,
	}}, func(data *menu.CallbackData) {
		fmt.Println("保存文件")
	})
	fileMenu.AddSeparator()
	fileMenu.AddText("退出", &keys.Accelerator{}, func(data *menu.CallbackData) {
		fmt.Println("退出")
	})
	moreMenu := m.AddSubmenu("更多")
	moreMenu.AddText("关于", &keys.Accelerator{}, func(data *menu.CallbackData) {
		fmt.Println("关于")
	})
	return m
}
