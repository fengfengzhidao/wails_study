## webview

go桌面端开发的基石，了解webview的使用，对后续wails框架的学习有很大的帮助

在桌面端开发中，WebView 是一种特殊的组件或控件，它允许在桌面应用程序中嵌入和显示网页内容（HTML、CSS、JavaScript 等），相当于在桌面应用里内置了一个轻量级的浏览器引擎



```Go
go get github.com/webview/webview_go
```



### 基本示例

```Go
package main

import (
  webview "github.com/webview/webview_go"
)

func main() {
  w := webview.New(false) // 如果是true 就可以打开f12
  defer w.Destroy()
  w.SetTitle("枫枫知道")
  w.SetSize(1200, 600, webview.HintNone)
  w.Navigate("https://www.fengfengzhidao.com")
  w.Run()
}

```



刚开始运行可能会报错

github.com/webview/webview_go: build constraints exclude all Go files in

```Go
// 开启cgo就好
go env -w CGO_ENABLED=1
```



cgo: C compiler “gcc” not found: exec: “gcc”: executable file not found in %PATH%

win下用Go语言的cgo时需要用到GCC编译器，windows下需要安装MinGW

下载链接 [https://github.com/niXman/mingw-builds-binaries/releases](https://github.com/niXman/mingw-builds-binaries/releases)

然后将mingw64的bin目录加到环境变量里面去，然后重启项目



应该就可以正常运行了



### 事件绑定

上面我们是使用w.Navigate直接打开一个目标地址

也可以使用SetHtml，这样就可以在html里面调后端方法了

#### js → go

在js中执行go的函数

```Go
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

```



#### go → js

在go中执行js的函数

js部分

```HTML
<h1>hello</h1>
<div id="result"></div>
<script>
    function showAlert(){
        alert("hello")
    }

    function showText(text){
        document.getElementById("result").innerText = text
    }
</script>
```



go部分

如果执行的操作在协程中，一定要把函数放到Dispatch里面去

然后可以使用embed嵌入的方式，把html代码打入go程序中

```Go
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

```



> 一定要注意一个误区，很多人把js那一部分叫前端，把go那部分叫后端，在桌面端开发中这样说不太严谨，严谨的说法，只要是在桌面环境中的代码，比如上面的js部分，go代码部分，都叫桌面端，或者叫前端，到现在还没有引入后端



## wails

webview了解个大概就ok了，真要做一点像样的桌面端产品

还是需要使用wails这样的桌面端开发框架

Wails是一个开源项目，旨在让开发者能够使用Go和Web技术（如React、Vue等）来构建桌面应用。它提供了一种轻量级且高效的解决方案，相比传统的Electron框架，Wails构建的应用具有更小的体积和更快的启动速度。



wails官网 [https://wails.io/zh-Hans/docs/gettingstarted/installation/](https://wails.io/zh-Hans/docs/gettingstarted/installation/)



### 安装

```Bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```



创建项目，以vue3+ts为例

```Markdown
wails init -n wails_study -t vue-ts
```



项目结构

```Go
.
├── build/
│   ├── appicon.png           // 应用程序图标
│   ├── darwin/
│   └── windows/
├── frontend/                //  前端项目文件
├── go.mod
├── go.sum
├── main.go                  // 主应用
└── wails.json               // 项目配置
```



使用wails dev就可以运行项目了

前端和go部分代码修改之后，会自动重启运行



### js到go方法调用

go新增一个方法

```Go
func (a *App) Add(n1, n2 int) int {
  return n1 + n2
}

```



新增之后，wails会自动变更

js部分使用

```JavaScript
import {Greet, Add} from '../../wailsjs/go/main/App'
async function add(){
  const res = await Add(1, 2)
  data.addResult = `结果是 ${res}`

}
<button class="btn" @click="add">add</button>
```



还可以通过Event事件实现js到go的方法调用

```Go
func (a *App) startup(ctx context.Context) {
  a.ctx = ctx
  runtime.EventsOn(a.ctx, "user-click", func(events ...any) {
    fmt.Println("收到点击事件，数据：", events)
  })
}
```



```TypeScript
import {EventsEmit} from "../../wailsjs/runtime";

EventsEmit("user-click", "xxxx", "xxx")
```



### go到js的方法调用

通过事件注册机制

```Go
func (a *App) CallJavaScript() {
  // 发送事件到前端，附带参数
  runtime.EventsEmit(a.ctx, "fromGo", "hello from go")
}

```



```TypeScript
import {EventsOn} from "../../wailsjs/runtime";

EventsOn("fromGo", (data)=>{
  console.log("go消息", data)
})
```





### 方法绑定和事件注册怎么选

1. 如果是同步调用，选择绑定方法
2. 如果是异步调用，选择事件注册





### 菜单

```Go
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
```



### 菜单综合练习

1. 文件操作
2. 剪贴板操作
3. 全屏操作
4. 刷新
5. 窗口大小操作
6. 置顶
7. 跳链接
8. 跳路由



#### 文件操作

```Go
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
```



#### 剪贴板操作

```Go
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
```



#### 全屏操作

```Go
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
```



#### 窗口操作

1. 刷新
2. 显示、隐藏
3. 置顶
4. 执行js

```Go
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
```



#### 应用跳转

跳转地址

```Go
runtime.BrowserOpenURL(a.ctx, "https://www.fengfengzhidao.com")
```

跳转其他应用

```JavaScript
// 跳转qq聊天
runtime.BrowserOpenURL(a.ctx, "tencent://message/?uin=qq号&Site=qq&Menu=yes")
// 跳转微信主界面
runtime.BrowserOpenURL(a.ctx, "weixin://")

```



### 路由

先下载vue-router

```JavaScript
npm i vue-router
```



在src下创建router/index.ts

```TypeScript
import { createRouter, createWebHashHistory } from "vue-router";

export const router = createRouter({
    history: createWebHashHistory(),
    routes: [
        {
            path: "/",
            name: "index",
            component: ()=>import("../views/index.vue"),
        },
        {
            path: "/home",
            name: "home",
            component: ()=>import("../views/home.vue"),
        },
        {
            path: "/about",
            name: "about",
            component: ()=>import("../views/about.vue"),
        }
    ],
});
```



然后创建对应的视图

然后在App.vue中修改

```TypeScript
<script lang="ts" setup>

</script>

<template>
  <div class="nav">
    <router-link to="/">首页</router-link>
    <router-link to="/home">home</router-link>
    <router-link to="/about">关于</router-link>
  </div>
  <div class="view">
    <router-view/>
  </div>

</template>
```



在main.ts中引入router

```JavaScript
import {createApp} from 'vue'
import App from './App.vue'
import {router} from "./router";

createApp(App).use(router).mount('#app')

```



#### 在菜单中跳转路由

因为wails的v2版本不能实现新开窗口，新开标签页的功能

所以我们需要曲线救国，把一些配置项通过web页面写出来，然后通过菜单跳转路由的方式

```Go
func (a *App) getRouterMenu() *menu.Menu {
  m := menu.NewMenu()
  routerMenu := m.AddSubmenu("路由")
  routerMenu.AddText("设置", &keys.Accelerator{}, func(data *menu.CallbackData) {
    runtime.EventsEmit(a.ctx, "router", "settings")
  })
  return m
}

```



在前端里面监听事件，然后跳路由即可

```TypeScript
import {EventsOn} from "../wailsjs/runtime";
import {router} from "./router";

EventsOn("router", function (name) {
  router.push({name})
})

```



### 快捷键

程序激活的时候，菜单上的快捷键就可以正常使用

但是要想实现程序未激活的时候使用快捷键，就得借助全局快捷键了

[https://github.com/makenowjust/hotkey](https://github.com/makenowjust/hotkey)

```Go
package main

import (
  "fmt"

  "github.com/MakeNowJust/hotkey"
)

func main() {
  hkey := hotkey.New()

  quit := make(chan bool)

  hkey.Register(hotkey.Ctrl, 'Q', func() {
    fmt.Println("Quit")
    quit <- true
  })

  fmt.Println("Start hotkey's loop")
  fmt.Println("Push Ctrl-Q to escape and quit")
  <-quit
}
```

上面那个库只能实现两个按键的快捷键，不过用起来很简单



如果要实现多个按键组合按键，就得使用golang.design/x/hotkey这个库了

[https://github.com/golang-design/hotkey](https://github.com/golang-design/hotkey)

```Go
package main

import (
  "log"

  "golang.design/x/hotkey"
  "golang.design/x/hotkey/mainthread"
)

func main() {
  mainthread.Init(fn)
}

func fn() {
  hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyS)
  err := hk.Register()
  if err != nil {
    log.Fatalf("hotkey: failed to register hotkey: %v", err)
    return
  }

  log.Printf("hotkey: %v is registered\n", hk)
  <-hk.Keydown()
  log.Printf("hotkey: %v is down\n", hk)
  <-hk.Keyup()
  log.Printf("hotkey: %v is up\n", hk)
  hk.Unregister()
  log.Printf("hotkey: %v is unregistered\n", hk)
}

```





#### 快捷键与wails结合

```Go
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

```



在app的startup方法中调用

```Go
func (a *App) startup(ctx context.Context) {
  a.ctx = ctx
  a.hotKey()
}
```



### 系统托盘

[https://github.com/energye/systray](https://github.com/energye/systray)



![](http://image.fengfengzhidao.com/fengfeng_110920251016225102.png?key=fengfengbuzhidao)



```Go
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


```



#### 系统托盘和wails结合

```Go
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

```





### 系统提示

![](https://image.fengfengzhidao.com/rj_102520251016181519.png)



可以使用notify这个库做简单通知

```Go
package main

import (
  "github.com/martinlindhe/notify"
)

func main() {
  notify.Alert("APP", "title", "msg", "D:\IT\fengfeng\test\wails_study\build\appicon.png")
}
```



如果有中文的话，显示会乱码，因为在windows中使用的编码是gbk

```Go
package main

import (
  "github.com/martinlindhe/notify"
  "golang.org/x/text/encoding/simplifiedchinese"
  "golang.org/x/text/transform"
)

// 将 UTF-8 字符串转换为 GBK 编码（适用于 Windows）
func toGBK(s string) string {
  result, _, err := transform.String(simplifiedchinese.GBK.NewEncoder(), s)
  if err != nil {
    return s // 转换失败时返回原字符串
  }
  return result
}

func main() {
  // 对中文内容进行编码转换
  title := toGBK("收到一条消息")
  message := toGBK("有人给你发消息了")
  appName := toGBK("APP")

  notify.Alert(appName, title, message, "")
}
```



如果是要做一些复杂的通知，比如可以点击通知消息，那可以使用toast.v1这个库

```Go
package main

import (
  "fmt"
  "gopkg.in/toast.v1"
  "log"
)

func main() {
  notification := toast.Notification{
    AppID:   "fengfengzhidao",
    Title:   "你收到一条消息",
    Message: "有人给你点赞了",
    Icon:    "D:\\IT\\fengfeng\\test\\wails_study\\build\\appicon.png",
    Actions: []toast.Action{
      {
        Type:      "protocol",                       // 协议类型
        Label:     "查看详情",                           // 按钮文字
        Arguments: "https://www.fengfengzhidao.com", // 点击后打开该网页
      },
    },
  }
  err := notification.Push()
  if err != nil {
    log.Fatalln(err)
  }
  fmt.Println(err)
}

```



> 有时候会不通知，看看是不是图标路径不对、字符串编码不对



### 无边框

有时候系统自带的菜单和操作逻辑很难做的很酷炫

比如菜单上还有输入框，显示用户头像这些，原生系统菜单几乎是做不出来的

所以可以直接把标题栏和菜单去掉，直接通过web的形式实现边框

那么就需要实现几个逻辑

1. 窗口拖动
2. 窗口最小化、窗口关闭



![](https://image.fengfengzhidao.com/rj_102520251017103713.png)



项目开启无边框

```Go
err := wails.Run(&options.App{
    Frameless: true,
})
```



前端部分设置

```TypeScript
<script lang="ts" setup>
import {WindowMinimise, WindowMaximise, Quit} from "../wailsjs/runtime";

</script>

<template>
  <div class="nav" style="--wails-draggable:drag" >
    导航栏

    <div class="action">
      <span @click="WindowMinimise">-</span>
      <span @click="WindowMaximise">□</span>
      <span @click="Quit">x</span>
    </div>
  </div>
  <div class="view">
    <router-view></router-view>
  </div>

</template>

<style>
body{
  margin: 0;
}
.nav{
  background-color: #333333;
  display: flex;
  justify-content: center;
  align-items: center;
  height: 40px;
  color: white;
  position: relative;

  .action{
    position: absolute;
    right: 20px;
    font-size: 14px;
    span{
      margin-left: 15px;
      cursor: pointer;
    }
  }
}
</style>

```



### 模板

我们之前的创建项目的命令，前端部分没有路由，没有pinia，没有UI库，没有菜单

写复杂的项目的时候，每次都需要我们自己去把这些东西配好，很麻烦

```Markdown
wails init -n wails_study -t vue-ts
```



可以用github上wails的第三方模板

[https://wails.io/zh-Hans/docs/community/templates](https://wails.io/zh-Hans/docs/community/templates)

可以在这上面选择你喜欢的模板

```Bash
wails init -n "Your Project Name" -t https://github.com/misitebao/wails-template-vue
```



也可以自己按照项目库的结构，做一个自己项目需要的模板

比如我可能会需要 一个UI组件库，pinia，axios，并且配置好代理

```Bash
wails init -n "Project Name" -t https://github.com/fengfengzhidao/wails-template-arcodesign-ts
```



如果拉取失败的话，可以设置一下代理

```Batch
set http_proxy=http://127.0.0.1:7890
set https_proxy=http://127.0.0.1:7890

set http_proxy=
set https_proxy=

```



## wails v3版本

最大的两个升级就是

1. 支持多窗口了
2. 集成了系统托盘

截至目前还是alpha版本



文档： [https://v3alpha.wails.io/whats-new/#multiple-windows](https://v3alpha.wails.io/whats-new/#multiple-windows)





## 桌面端项目开发

开发一个备忘录的桌面端软件吧

1. 笔记基本操作
    - 新建笔记（标题 + 内容）、编辑、删除、复制
    - 自动保存（每 3 秒 / 输入暂停时）
    - 快捷键支持（新建`Ctrl+N`、保存`Ctrl+S`、删除`Ctrl+D`）
2. 分类管理
    - 创建 / 删除笔记本（如「工作」「生活」）
    - 笔记归属到指定笔记本（新建时选择 / 后续修改）
    - 左侧导航栏展示笔记本列表，点击切换
3. 本地存储与备份
    - 所有数据默认存储在本地 SQLite 数据库（用户目录下，如`~/.memo/data.db`）
    - 手动备份：支持导出选中笔记为 Markdown 文件（单文件 / 批量打包）



## 参考文档

go运行webview  [https://blog.csdn.net/qq_43660595/article/details/139641147](https://blog.csdn.net/qq_43660595/article/details/139641147)

go_webview [https://bbs.itying.com/topic/6876c4174715aa0088487c64](https://bbs.itying.com/topic/6876c4174715aa0088487c64)

