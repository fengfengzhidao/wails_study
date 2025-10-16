package main

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
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

func (a *App) getFileMenu() *menu.Menu {
	m := menu.NewMenu()
	fileMenu := m.AddSubmenu("文件")

	fileMenu.AddText("打开文件", &keys.Accelerator{}, func(data *menu.CallbackData) {
		fmt.Println("打开文件")
		filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
			Title: "枫枫要选择文件了",
			Filters: []runtime.FileFilter{
				{
					DisplayName: "Image Files (*.jpg, *.png)",
					Pattern:     "*.jpg;*.png",
				},
			},
		})
		fmt.Println(filePath, err)
	})
	fileMenu.AddText("保存文件", &keys.Accelerator{}, func(data *menu.CallbackData) {
		fmt.Println("保存文件")
		filePath, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
			Title:            "枫枫要保存文件了",
			DefaultDirectory: "E:\\IT\\go_pro\\webview_study\\wails_study",
			Filters: []runtime.FileFilter{
				{
					DisplayName: "Text Files (*.txt)",
					Pattern:     "*.txt",
				},
			},
		})
		fmt.Println(filePath, err)
		err = os.WriteFile(filePath, []byte("hello"), 0644)
		fmt.Println(err)
	})
	return m
}
