# CLI 命令行实用程序开发实战 - Agenda

秦晨曦 17343096

﻿### 1、概述

  命令行实用程序并不是都象 cat、more、grep 是简单命令。[go](https://go-zh.org/cmd/go/)项目管理程序，类似 java 项目管理 maven、Nodejs 项目管理程序 npm、git 命令行客户端、 docker 与 kubernetes 容器管理工具等等都是采用了较复杂的命令行。即一个实用程序同时支持多个子命令，每个子命令有各自独立的参数，命令之间可能存在共享的代码或逻辑，同时随着产品的发展，这些命令可能发生功能变化、添加新命令等。因此，符合 [OCP](https://en.wikipedia.org/wiki/Open/closed_principle)原则 的设计是至关重要的编程需求。

### 任务目标

1.熟悉 go 命令行工具管理项目
2.综合使用 go 的函数、数据结构与接口，编写一个简单命令行应用 agenda
3.使用面向对象的思想设计程序，使得程序具有良好的结构命令，并能方便修改、扩展新的命令,不会影响其他命令的代码
4.项目部署在 Github 上，合适多人协作，特别是代码归并
5.支持日志（原则上不使用debug调试程序）

## tip2: 安装 cobra

使用命令 `go get -v github.com/spf13/cobra/cobra` 下载过程中，会出提示如下错误
>Fetching https://golang.org/x/sys/unix?go-get=1
>
>https fetch failed: Get https://golang.org/x/sys/unix?go-get=1: dial tcp 216.239.37.1:443: i/o timeout

这是熟悉的错误，请在 `$GOPATH/src/golang.org/x` 目录下用 `git clone` 下载 `sys` 和 `text` 项目，然后使用 `go install github.com/spf13/cobra/cobra`, 安装后在 `$GOBIN` 下出现了 `cobra` 可执行程序。

## Cobra 的简单使用

创建一个处理命令 `agenda register -uTestUser` 或 `agenda register --user=TestUser` 的小程序。

简要步骤如下：
>cobra init
>
>cobra add register

需要的文件就产生了。 你需要阅读 `main.go` 的 `main()` ; `root.go` 的 `Execute()`; 最后修改 `register.go`, `init()` 添加：

>registerCmd.Flags().StringP("user", "u", "Anonymous", "Help message for username")

Run 匿名回调函数中添加：

>username, _ := cmd.Flags().GetString("user")
>
>fmt.Println("register called by " + username)

测试命令：
>$ go run main.go register --user=TestUser
>
>register called by TestUser

---

## 测试

> $ go get -u github.com/Tomy-Lee/Agenda-golang

> $GOPATH/bin/agenda -h
>
>Agenda is a meeting manager based on CLI using cobra library.
>It supports different operation on meetings including register, create meeting, query and so on.
>It's a cooperation homework assignment for service computing.
>
>Usage:
>  Agenda [command]
>
>
```
Available Commands:
  cancel              Cancel your own meeting by specifying title name.
  changeParticipators Change your own meetings' participators.
  clear               Clear all meetings you attended or created.
  createMeetings      Create meetings.
  delete              A brief description of your command
  help                Help about any command
  list
  listMeetingsCmd     List all of your own meetings during a time interval.
  login               Login
  logout              Logout
  quit                Quit meetings.
  register            Register user.
```

>Flags:
  -d, --debug   display log message
  -h, --help    help for Agenda
>
>Use "Agenda [command] --help" for more information about a command.


## 测试用例



#### 1.注册

>./Agenda-golang register -u Tomy -p 060505 -m www.Tomy.com -t 12345678987

#### 2.再次注册已存在账户


>there's another user with username Tomy


#### 3.登录


>./Agenda-golang login -u Tomy -p 060505

#### 4.密码错误

>Authentication Fail

#### 5.未退出再次登录

>Action login requires an logout state

#### 6.创建会议

>./Agenda-golang createMeetings -t Party -p

创建后的几种情况输出：
>meeting hosted

>there's another meeting with title: ABC_Meeting

>there are time conflict of some participants

>meeting should end later than start


#### 7.更改参与者

>./Agenda-golang changeParticipators -t Party -p AAA

更改后的几种情况输出：

>user 'AAA' is already a participant of meeting 'Party'

>meeting doesn't exist: PPP


#### 8.取消会议

>./Agenda-golang cancel -t Party


会议不存在：
>meeting doesn't exist: ZZZ

#### 9.退出会议

>./Agenda-golang quit -t Party


不是参与者无法退出会议：

>user 'Li' is not a participant of meeting 'Party'


#### 10.列出会议
列出所有会议：
>./Agenda-golang list

```
Username Email Phone
'Tomy' 'www.Tomy.com' '12345678987'
'AAA' 'www.AAA.com' '13712345678'
'BBB' 'www.BBB.com' '13898765432'
'CCC' 'www.CCC.com' '13978945632'
```
根据时间列出会议

>./Agenda-golang listMeetings -s 2019-10-20 -e 2019-12-01

```
title: TTT
  host: AAA
  time: 2019-10-22 to 2019-12-01
  participants: ZZZ
```
列出会议时情况输出


>meeting should end later than start

>invalid time format: 2019-10-1


#### 11.清除会议


>./Agenda-golang clear

#### 12.删除会议

>./agenda delete


### 使用cobra

#### 获取帮助和日志记录：
>-h, --help
>-d, --debug

#### 注册

```
Flags:
  -m, --mail string       email.
  -p, --password string   Help message for username
  -t, --phone string      Phone
  -u, --user string       Username
```

#### 登录

```
Flags:
  -p, --password string   Input password
  -u, --user string       Input username
```

#### 创建会议

```
Flags:
  -e, --end string             Input end time as the format of (yyyy-mm-dd).
  -p, --participators string   Input participator name.
  -s, --start string           Input start time as the format of (yyyy-mm-dd).
  -t, --title string           Input title name.
```

