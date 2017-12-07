# cyber-auth-api

''# yum install go

$ vi ~/.bashrc
export GOROOT=/usr/lib/golang
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin


$ go get github.com/astaxie/beego
$ go get github.com/beego/bee


$ bee api cyber-auth-api
$ cd cyber-auth-api/
$ bee run -gendoc=true -downdoc=true


$ go get -u github.com/go-sql-driver/mysql
$ go get github.com/satori/go.uuid
$ go get gopkg.in/mgo.v2
$ go get github.com/bradfitz/gomemcache/memcache
$ go get github.com/casbin/casbin


# Linux下通过端口查看进程
''# netstat -anp|grep 8086
