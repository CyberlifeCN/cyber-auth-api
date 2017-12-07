package controllers

import (
	"cyber-auth-api/models"
	"encoding/json"
	"github.com/astaxie/beego"
  "log"
  "fmt"
)


// Operations about Register
type RegisterController struct {
	beego.Controller
}


// @Title Register
// @Description register by username & password though RPC
// @Param	body		body 	models.RegisterReq	true		"body for register content"
// @Success 200 {object} models.RegisterResp
// @Failure 403 :username or password is empty
// @router / [post]
func (this *RegisterController) Post() {
	uri := this.Ctx.Input.URI()
  beego.Info(uri)
  beego.Info(this.Ctx.Input.RequestBody)

  var req models.RegisterReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	beego.Trace(req)
	if (req.Username == "" || req.Password == "") {
		var rs = &models.RegisterResp{
			Code: 403,
			Msg: "Bad Request",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
		return
	}

  // only for test, unit test can't md5(password) by js
  var args = &models.CreateLoginArgs{
    Username: req.Username,
    Md5Password: models.GetMd5String(req.Password),
  }
  reply := &models.CreateLoginReply{}
  err = GlobalRpcClient.Call("Auth.CreateLogin", args, &reply)
  if err != nil {
    log.Fatal("CreateLogin error :", err)
  }
  fmt.Println("CreateLogin:", args, reply)

	if (reply == nil || reply.Id == "") {
    var rs = &models.RegisterResp{
			Code: 409,
			Msg: "Conflict",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
	} else {
		beego.Trace(reply)

		var rs = &models.RegisterResp{
			Code: 200,
			Msg: "Success",
			Rs: reply.Id,
		}

		this.Data["json"] = *rs
		this.ServeJSON()
	}
}
