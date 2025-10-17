package main

import (
	"fmt"
	"github.com/martinlindhe/notify"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"time"
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

func (a *App) getClipboardMenu() *menu.Menu {
	m := menu.NewMenu()
	clipboardMenu := m.AddSubmenu("剪贴板")
	clipboardMenu.AddText("复制", keys.Control("c"), func(data *menu.CallbackData) {
		runtime.ClipboardSetText(a.ctx, "设置当前时间："+time.Now().Format(time.DateTime))
	})
	clipboardMenu.AddText("粘贴", keys.Control("c"), func(data *menu.CallbackData) {
		text, err := runtime.ClipboardGetText(a.ctx)
		fmt.Println(text, err)
	})
	return m
}

func (a *App) getScreenMenu() *menu.Menu {
	m := menu.NewMenu()
	screenMenu := m.AddSubmenu("屏幕")
	screenMenu.AddCheckbox("全屏", false, keys.Key("f11"), func(data *menu.CallbackData) {
		if runtime.WindowIsFullscreen(a.ctx) {
			runtime.WindowUnfullscreen(a.ctx)
			screenMenu.Items[0].SetChecked(false)
		} else {
			runtime.WindowFullscreen(a.ctx)
			screenMenu.Items[0].SetChecked(true)
		}
	})
	return m
}

func (a *App) getWindowMenu() *menu.Menu {
	m := menu.NewMenu()
	windowMenu := m.AddSubmenu("窗口")
	windowMenu.AddText("刷新js", keys.Key("f5"), func(data *menu.CallbackData) {
		fmt.Println("刷新")
		runtime.WindowReload(a.ctx)
	})
	windowMenu.AddText("刷新应用", keys.Control("f5"), func(data *menu.CallbackData) {
		fmt.Println("刷新应用")
		runtime.WindowReloadApp(a.ctx)
	})
	windowMenu.AddText("隐藏", &keys.Accelerator{}, func(data *menu.CallbackData) {
		fmt.Println("隐藏")
		runtime.WindowHide(a.ctx)
		go func() {
			time.Sleep(2 * time.Second)
			runtime.WindowShow(a.ctx)
		}()
	})
	var isTop bool
	windowMenu.AddCheckbox("置顶", false, keys.Key("f10"), func(data *menu.CallbackData) {
		isTop = !isTop
		runtime.WindowSetAlwaysOnTop(a.ctx, isTop)
		windowMenu.Items[3].SetChecked(isTop)
	})

	windowMenu.AddText("执行js", &keys.Accelerator{}, func(data *menu.CallbackData) {
		fmt.Println("执行js")
		runtime.WindowExecJS(a.ctx, "alert('hello')")
	})
	return m
}

func (a *App) getBrowserMenu() *menu.Menu {
	m := menu.NewMenu()
	windowMenu := m.AddSubmenu("浏览器")
	windowMenu.AddText("关于枫枫", &keys.Accelerator{}, func(data *menu.CallbackData) {
		runtime.BrowserOpenURL(a.ctx, "https://www.fengfengzhidao.com")
	})
	return m
}

func (a *App) getRouterMenu() *menu.Menu {
	m := menu.NewMenu()
	routerMenu := m.AddSubmenu("路由")
	routerMenu.AddText("设置", &keys.Accelerator{}, func(data *menu.CallbackData) {
		runtime.EventsEmit(a.ctx, "router", "settings")
	})
	routerMenu.AddText("系统消息", &keys.Accelerator{}, func(data *menu.CallbackData) {
		notify.Alert("APP", "rev a msg", "have people digg", "E:\\IT\\go_pro\\webview_study\\wails_study\\build\\appicon.png")
	})
	return m
}
