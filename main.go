package main

import (
	_ "cyber-auth-api/routers"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/plugins/auth"
	// "cyber-api-auth/auth"
	// "github.com/casbin/beego-authz/authz"
	// "cyber-auth-api/authz"
	// "github.com/casbin/casbin"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// authenticate every request.
	// beego.InsertFilter("*", beego.BeforeRouter, auth.Basic("alice", "123"))
	// authorize every request.
	// beego.InsertFilter("*", beego.BeforeRouter, authz.NewAuthorizer(casbin.NewEnforcer("./conf/authz_model.conf", "./conf/authz_policy.csv")))

	beego.Run()
}
