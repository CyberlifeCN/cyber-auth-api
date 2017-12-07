package controllers

import (
	"cyber-auth-api/models"
	"encoding/json"
	"github.com/astaxie/beego"
  "log"
  "fmt"
)


// Operations about VerifyCode
type CodeController struct {
	beego.Controller
}


// @Title Code
// @Description retrieve verify_code by uid(username,phone,email) though RPC
// @Param	body		body 	models.VerifyCodeReq	true		"body for retrieve content"
// @Success 200 {object} models.VerifyCodeResp
// @Failure 403 :uid or type is empty
// @router / [post]
func (this *CodeController) Post() {
	uri := this.Ctx.Input.URI()
  beego.Info(uri)
  beego.Info(this.Ctx.Input.RequestBody)

  var req models.VerifyCodeReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	beego.Trace(req)
	if (req.Id == "") {
		var rs = &models.VerifyCodeResp{
			Code: 403,
			Msg: "Bad Request",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
		return
	}

  var args = &models.CreateCodeArgs{
    Id: req.Id,
  }
  reply := &models.CreateCodeReply{}
  err = GlobalRpcClient.Call("Auth.CreateCode", args, &reply)
  if err != nil {
    log.Fatal("CreateCode error :", err)
  }
  fmt.Println("CreateCode:", args, reply)

	if (reply == nil || reply.Code == "") {
    var rs = &models.VerifyCodeResp{
			Code: 404,
			Msg: "Not Found",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
	} else {
		beego.Trace(reply)

		var rs = &models.VerifyCodeResp{
			Code: reply.Status,
			Msg: "Success",
			Rs: reply.Code,
		}
    if (rs.Code == 409) {
      rs.Msg = "Conflict"
    }

		this.Data["json"] = *rs
		this.ServeJSON()
	}
}
