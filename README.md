[![image](http://b2oks-cover.b0.upaiyun.com/default/cyberlife-logo.jpg)](http://cyber-life.cn)
# cyber-auth-api

任何一个信息软件系统，用户认证都是最基础的模块。创建一个项目时，首先总要完成这一部分。把多年项目中的实践归纳、整理出来这个最精简的模块，供大家参考，HTTP Restful API。它包含8个API：
* 登录
* 登出
* 获取注册验证码
* 注册
* 获取重置密码验证码
* 忘记密码
* 通过存储在cookie中access_token重新获取session_ticket
* 通过存储在cookie中refresh_token重新获取session_ticket

### 安装golang
    # yum install go

### 编辑环境变量
    $ vi ~/.bashrc
    export GOROOT=/usr/lib/golang
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

### 安装beego
    $ go get github.com/astaxie/beego
    $ go get github.com/beego/bee

### 安装依赖包
    $ go get -u github.com/go-sql-driver/mysql
    $ go get github.com/satori/go.uuid
    $ go get gopkg.in/mgo.v2
    $ go get github.com/bradfitz/gomemcache/memcache
    $ go get github.com/casbin/casbin

### 创建工程
    $ bee api cyber-auth-api
    $ cd cyber-auth-api/

### 启动
    $ cd rpc
    $ go run rpc_server.go
    $ bee run -gendoc=true -downdoc=true

### Demo
[API文档](http://auth.domicake.com/swagger/ "可以直接作为单元测试工具使用")

### Linux下通过端口查看进程
    # netstat -anp|grep 8086
