# FurbotServer-Go

[![Go Report Card](https://goreportcard.com/badge/github.com/Rehtt/FurbotServer-Go)](https://goreportcard.com/report/github.com/Rehtt/FurbotServer-Go)

根据[Furbot-MiraiGo](https://github.com/Rehtt/Furbot-MiraiGo) 反推出的服务端，兼容[Furbot-Mirai](https://github.com/furleywolf/Furbot-Mirai)

## 安装

```shell
go install github.com/Rehtt/FurbotServer-Go@latest
FurbotServer-Go
```

OR

```shell
git clone github.com/Rehtt/FurbotServer-Go
cd FurbotServer-Go
go install
FurbotServer-Go
```

OR

```shell
git clone github.com/Rehtt/FurbotServer-Go
cd FurbotServer-Go
go build .
./FurbotServer-Go
```

## 使用
修改配置文件，启动程序`FurbotServer-Go`

可以使用参数指定配置文件路径：`FurbotServer-Go -c ./config.yaml`

图片暂时保存在本地，以后会添加其他存储方式 ~~咕咕咕~~

配置文件内容：
```yaml
# 服务监听
server:
  addr: 0.0.0.0
  port: 80

# 管理员密钥
authKey: Xz0aCixrjL8hgIAG

imagePath: "./image"    # 图片存储路径（暂时存在本地）

db:
  use: mysql            # 选择数据库，mysql与sqlite 二选一
  mysql:                # mysql 配置
    addr: 127.0.0.1
    username: root
    password: root
    port: 3306
    database: test
  sqlite:               # sqlite配置
    path: ./database.db

```
