# AgendaGo 命令设计

## 注册 register

### 语法

```shell
$ AgendaGo register [flags]
```

### 命令参数 flags
- -u, --username    string  用户名
- -p, --password    string  密码
- -m, --email       string  邮箱
- -t, --phone       string  电话
- -h, --help                帮助

### 实例

```shell
$ AgendaGo register -u test -p testpass -m test@test.com -n 12345678909
```

## 登陆 login

### 语法

```shell
$ AgendaGo login [flags]
```

### 命令参数 flags
- -u, --username    string  用户名
- -p, --password    string  密码
- -h, --help                帮助

### 实例

```shell
$ AgendaGo login -u test -p testpass
```

## 登出 logout

### 语法

```shell
$ AgendaGo logout [flags]
```

### 命令参数 flags
- -h, --help                帮助

### 实例

```shell
$ AgendaGo logout
```

## 用户查询 queryUser

### 语法

```shell
$ AgendaGo queryUser [flags]
```

### 命令参数 flags
- -h, --help                帮助

### 实例

```shell
$ AgendaGo queryUser
```

## 用户删除 deleteUser

### 语法

```shell
$ AgendaGo deleteUser [flags]
```

### 命令参数 flags
- -h, --help                帮助

### 实例

```shell
$ AgendaGo deleteUser
```

## 创建会议 createMeeting

### 语法

```shell
$ AgendaGo createMeeting [flags]
```

### 命令参数 flags
- -t, --title           string      会议主题
- -p, --participator    string      参加者（用多个-p 添加多个参加者）
- -s, --starttime       string      开始时间（格式XXXX-XXX-XX/XX:XX:XX, 24小时制）
- -e, --endtime         string      结束时间（格式XXXX-XXX-XX/XX:XX:XX, 24小时制）
- -h, --help                        帮助

### 实例

```shell
$ AgendaGo createMeeting -t testMeeting01 -p test01 -s 2018-02-01/12:00:00 -e 2018-02-01/13:00:00
```

## 增加会议参与者 add

### 语法

```shell
$ AgendaGo add [flags]
```

### 命令参数 flags
- -t, --title           string      会议主题
- -p, --participator    string      参加者（用多个-p 添加多个参加者）
- -h, --help                        帮助

### 实例

```shell
$ AgendaGo add -t testMeeting01 -p test01
```

## 删除会议参与者 remove

### 语法

```shell
$ AgendaGo remove [flags]
```

### 命令参数 flags
- -t, --title           string      会议主题
- -p, --participator    string      参加者（用多个-p 添加多个参加者）
- -h, --help                        帮助

### 实例

```shell
$ AgendaGo remove -t testMeeting01 -p test01
```

## 查询会议 queryMeeting

### 语法

```shell
$ AgendaGo queryMeeting [flags]
```

### 命令参数 flags
- -t, --title           string      会议主题
- -h, --help                        帮助

### 实例

```shell
$ AgendaGo queryMeeting -t testMeeting01
```

## 取消会议 deleteMeeting

**相关参数**

### 语法

```shell
$ AgendaGo deleteMeeting [flags]
```

### 命令参数 flags
- -t, --title           string      会议主题
- -h, --help                        帮助

### 实例

```shell
$ AgendaGo deleteMeeting -t testMeeting01
```

## 退出会议 quitMeeting

### 语法

```shell
$ AgendaGo quitMeeting [flags]
```

### 命令参数 flags
- -t, --title           string      会议主题
- -h, --help                        帮助

### 实例

```shell
$ AgendaGo quitMeeting -t testMeeting01
```

## 清空会议 clearMeeting

### 语法

```shell
$ AgendaGo clearMeeting [flags]
```

### 命令参数 flags
- -h, --help                        帮助

### 实例

```shell
$ AgendaGo clearMeeting
```



