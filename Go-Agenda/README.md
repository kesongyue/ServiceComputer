# Go-Agenda
服务计算课程作业——CLI 命令行实用程序开发实战 - Agenda

## 2018.10.26

已经实现的功能：

- help ：列出命令说明（这个可以不要了，cobra会自动帮我们生成）
- 注册用户：

```
go run main.go register -u qiuxy -p 1233 -e 1233 -t 1233
```

* 登录：

```
go run main.go login -u qiuxy -p 1233
```

- 注销

```
go run main.go logout
```

* 列出所有的用户

```
go run main.go queryUser
```

* 删除用户（这里缺少删除会议中的用户的相关操作，有心情再做）：

```
go run main.go delete
```

基本文件：

```
Go-Agenda
├── cmd（对应的命令文件）
│   |── root.go：
|   └── delete.go，login.go，logout.go，queryUser.go，register.go
├── entity
│   ├── Data
│   │   ├── MJson: text格式，保存所有会议信息
│   │   ├── UJson: text格式，保存所有用户信息
|   |   └── ULoginJson: 如果登录了，这个文件里会有一行用户信息
│   ├── Date.go: 将日期的年月日作为string封装到Data.go，用于Meeting.go中会议的日期
型
│   ├── JsonIOForUM.go: 实现对User和Meeting类型对象的json文件读写
│   ├── Meeting.go: 将Meeting的实体内容及相关方法封装
|	├── Util.go: 一些基本的功能函数
│   └── User.go: 将User的实体内容及相关方法封装
└── main.go
```

其他的命令操作看一下已经建好的命令基本就可以了，重复的操作而已：

1. 先在命令行输入cobra add 命令的名字：

```
F:\Users\HP\GOPATH\src\Go-Agenda>cobra add login
```

2. 然后在自动生成的login.go文件的`init`函数中配置参数
3. 最后在`Run: func(cmd *cobra.Command, args []string)` 中添加基本的函数逻辑

一些之后可能用到的函数，要使用的时候直接调用即可

用于判断是否登录的函数：entity.IsLogin()，返回值是布尔值

保存用户和会议信息到文件：entity.WriteJson()

获取用户的数组：entity.Users

获取会议的数组：entity.Meetings

