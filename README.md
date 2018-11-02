# Agenda Go！

> A CLI made by golang

## 安装运行

```shell
$ go get github.com/Zophyr/AgendaGo
```

## 程序设计

本次作业的「AgendaGo」是我们初级实训中的「Agenda」会议管理系统的变种。因此我们将模仿在初级实训中「Agenda」的程序设计。

### 构架设计

#### 三层构架

在初级实训中，采取了三层结构设计。因此，我们将仿照其设计方式来实现AgendaGo。

##### 表示层 `cmd`

- 负责与用户的交互操作。例如，进行命令的操作、执行某一命令。
- 负责接收用户的输入。接收用户输入的指令与参数，以及相关命令的数据。并将数据传输给业务逻辑层。
- 负责进行信息的输出。输出程序执行情况，与交互语句。

##### 业务逻辑层 `service`

- 业务逻辑的执行，调取`实体层`提供的相关API进行操作。
- 判断表示层传输进来数据、命令的合法性。

##### 实体层 `entity`

- 暴露相关数据操作接口。
- 直接对数据进行操作。
- 文件的读取与存储。

### 命令设计

#### POSIX/GNU-风格参数处理

```shell
Available Commands:
  add           Add participators to a meeting
  clearMeeting  Clear all the meeting
  createMeeting Create a meeting
  delete        Delete the meeting
  delete        Delete a participator from meeting
  deleteUser    Delete one account of Agenda and log out
  help          Help about any command
  login         User login
  logout        User logout
  queryMeeting  Query the meeting by its title
  queryUser     Show all registered users
  quitMeeting   Help the current user to quit the correctsponding meeting by its title
  register      Used to register account
```

[AgendaGo 命令详细介绍](https://github.com/Zophyr/AgendaGo/blob/master/cmd-design.md)