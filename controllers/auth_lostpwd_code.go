package controllers

import (
	"cyber-auth-api/models"
	"encoding/json"
	"github.com/astaxie/beego"
  "log"
  "fmt"
)


// Operations about lost password verify_code
type LostpwdCodeController struct {
	beego.Controller
}


// @Title Code
// @Description retrieve lostpwd_verify_code by uid(username,phone,email) though RPC
// @Param	body		body 	models.LostpwdVerifyCodeReq	true		"body for retrieve content"
// @Success 200 {object} models.LostpwdVerifyCodeResp
// @Failure 403 :uid or type is empty
// @router / [post]
func (this *LostpwdCodeController) Post() {
	uri := this.Ctx.Input.URI()
  beego.Info(uri)
  beego.Info(this.Ctx.Input.RequestBody)

  var req models.LostpwdVerifyCodeReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	beego.Trace(req)
	if (req.Id == "") {
		var rs = &models.LostpwdVerifyCodeResp{
			Code: 403,
			Msg: "Bad Request",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
		return
	}

  var args = &models.CreateLostpwdCodeArgs{
    Id: req.Id,
  }
  reply := &models.CreateLostpwdCodeReply{}
  err = GlobalRpcClient.Call("Auth.CreateLostpwdCode", args, &reply)
  if err != nil {
    log.Fatal("CreateLostpwdCode error :", err)
  }
  fmt.Println("CreateLostpwdCode:", args, reply)

	if (reply == nil || reply.Code == "") {
    var rs = &models.LostpwdVerifyCodeResp{
			Code: 404,
			Msg: "Not Found",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
	} else {
		beego.Trace(reply)

		var rs = &models.LostpwdVerifyCodeResp{
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
