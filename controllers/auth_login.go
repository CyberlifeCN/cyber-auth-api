package controllers

import (
	"cyber-auth-api/models"
	"encoding/json"
	"github.com/astaxie/beego"
  "log"
  "fmt"
)


// Operations about Login
type LoginController struct {
	beego.Controller
}


// @Title Login
// @Description login by username & password though RPC
// @Param	body		body 	models.LoginReq	true		"body for login content"
// @Success 200 {object} models.SessionTicket
// @Failure 403 :username or password is empty
// @router / [post]
func (this *LoginController) Post() {
	uri := this.Ctx.Input.URI()
  beego.Info(uri)
  beego.Info(this.Ctx.Input.RequestBody)

  var req models.LoginReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	beego.Trace(req)
	if (req.Id == "" || req.Pwd == "") {
		var rs = &models.LoginResp{
			Code: 403,
			Msg: "Bad Request",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
		return
	}

  // only for test, unit test can't md5(password) by js
  var args = &models.CreateTicketArgs{
    Id: req.Id,
    Md5pwd: models.GetMd5String(req.Pwd),
  }
  reply := &models.SessionTicket{}
  err = GlobalRpcClient.Call("Auth.CreateTicket", args, &reply)
  if err != nil {
    log.Fatal("CreateTicket error :", err)
  }
  fmt.Println("CreateTicket:", args, reply)

	if (reply == nil || reply.Id == "") {
    var rs = &models.LoginResp{
			Code: 404,
			Msg: "Not Found",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
	} else {
		beego.Trace(reply)

		var rs = &models.LoginResp{
			Code: 200,
			Msg: "Success",
			Rs: *reply,
		}

		this.Data["json"] = *rs
		this.ServeJSON()
	}
}
