# cyber-auth-api

## 安装golang
    # yum install go

## 编辑环境变量
    $ vi ~/.bashrc
    export GOROOT=/usr/lib/golang
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

## 安装beego
    $ go get github.com/astaxie/beego
    $ go get github.com/beego/bee

## 安装依赖包
    $ go get -u github.com/go-sql-driver/mysql
    $ go get github.com/satori/go.uuid
    $ go get gopkg.in/mgo.v2
    $ go get github.com/bradfitz/gomemcache/memcache
    $ go get github.com/casbin/casbin

## 创建工程
    $ bee api cyber-auth-api
    $ cd cyber-auth-api/

## 启动
    $ cd rpc
    $ go run rpc_server.go
    $ bee run -gendoc=true -downdoc=true

## Demo
    [API文档]: http://auth.domicake.com/swagger/  "可以直接作为单元测试工具使用"

## Linux下通过端口查看进程
    # netstat -anp|grep 8086
