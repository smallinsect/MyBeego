# MyBeego

帮助文档

https://beego.me/quickstart

创建项目

```
bee new LibrarySystem
```

## bee工具：管理beego项目

- 安装bee工具
  - go get -u github.com/beego/bee 
- 使用bee工具
  - 运行bee可以看到bee参数，比如new，api，run

## 安装beego

- go get -u github.com/astaxie/beego
- 升级
  - go get -u github.com/astaxie/beego
- 源码下载升级
  - https://github.com/astaxie/beego

## 配置GOROOT、GOPATH路径

GOPATH E:\smallinsect\GoWorkSpace 工作路径

GOROOT D:\Go go的安装路径

%GOPATH%\bin;%GOROOT%\bin;

## 创建项目

```
cd E:\smallinsect\GoWorkSpace\src
bee new beego_project
```

## 启动一个项目

```
进入到创建好的项目目录
cd E:\smallinsect\GoWorkSpace\src\MyBeego\LibrarySystem
运行项目
bee run
```

## 函数对应请求

```
type MainController struct {
  beego.Controller
}

func (c *MainController) Get() {
  c.Data["Website"] = "beego.me"
  c.Data["Email"] = "astaxie@gmail.com"
  c.TplName = "index.tpl"
}
```

Get对应GET请求

Post函数对应POST请求

## 修改项目页面路径

// beego.SetViewsPath("front")// 将项目前端路径设为front目录下

## 结构体渲染

```
type User struct {
	Name string
}

user := User{Name:"小昆虫"}
c.Data["user"] = user

{{.user.Name}}
```



