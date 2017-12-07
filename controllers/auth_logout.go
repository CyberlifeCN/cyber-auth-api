package controllers

import (
	"cyber-auth-api/models"
	// "encoding/json"
	"github.com/astaxie/beego"
  "log"
  "fmt"
	"strings"
)


// Operations about Logout
type LogoutController struct {
	beego.Controller
}


// @Title Login
// @Description logout by access_token though RPC
// @Param	Authorization		header 	string	"Bearer 81063dfdda8911e79b00a45e60efbf2d"		true		"Authorization=Bearer access_token"
// @Success 200
// @Failure 403 :access_token is empty
// @Failure 404 :access_token not found
// @router / [delete]
func (this *LogoutController) Delete() {
	uri := this.Ctx.Input.URI()
  beego.Info(uri)

	authorString := this.Ctx.Request.Header.Get("Authorization")
	fmt.Println("Authorization:", authorString)
	s := strings.SplitN(authorString, " ", 2)
	if len(s) != 2 || s[0] != "Bearer" {
		var rs = &models.LoginResp{
			Code: 403,
			Msg: "Bad Request",
		}
		this.Data["json"] = *rs
		this.ServeJSON()
	}

  var args = &models.LogoutArgs{
    AccessToken: s[1],
  }
  reply := &models.LogoutReply{}
  err = GlobalRpcClient.Call("Auth.Logout", args, &reply)
  if err != nil {
    log.Fatal("Logout error :", err)
  }
  fmt.Println("Logout:", args, reply)

	var rs = &models.LogoutResp{
		Code: reply.Status,
		Msg: "Success",
	}
	if (rs.Code == 404) {
    rs.Msg = "Not Found"
  }

	this.Data["json"] = *rs
	this.ServeJSON()
}
