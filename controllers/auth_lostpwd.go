package controllers

import (
	"cyber-auth-api/models"
	"encoding/json"
	"github.com/astaxie/beego"
  "log"
  "fmt"
)


// Operations about lost password
type LostpwdController struct {
	beego.Controller
}


// @Title Lost password
// @Description lost password by username,password,verify_code though RPC
// @Param	body		body 	models.LostpwdReq	true		"body for register content"
// @Success 200 {object} models.LostpwdResp
// @Failure 403 :username or password is empty
// @router / [post]
func (this *LostpwdController) Post() {
	uri := this.Ctx.Input.URI()
  beego.Info(uri)
  beego.Info(this.Ctx.Input.RequestBody)

  var req models.LostpwdReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	beego.Trace(req)
	if (req.Id == "" || req.Pwd == "" || req.Code == "") {
		var rs = &models.LostpwdResp{
			Code: 403,
			Msg: "Bad Request",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
		return
	}

  // only for test, unit test can't md5(password) by js
  var args = &models.LostpwdArgs{
    Id: req.Id,
    Md5pwd: models.GetMd5String(req.Pwd),
		Code: req.Code,
  }
  reply := &models.LostpwdReply{}
  err = GlobalRpcClient.Call("Auth.Lostpwd", args, &reply)
  if err != nil {
    log.Fatal("Lostpwd error :", err)
  }
  fmt.Println("Lostpwd:", args, reply)

	if (reply == nil) {
    var rs = &models.LostpwdResp{
			Code: 409,
			Msg: "Conflict",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
	} else {
		beego.Trace(reply)

		var rs = &models.LostpwdResp{
			Code: reply.Status,
			Msg: "Success",
			Rs: reply.Id,
		}
		if (rs.Code == 408) {
      rs.Msg = "Request Timeout"
		} else if (rs.Code == 404) {
      rs.Msg = "Not Found"
    } else if (rs.Code == 409) {
      rs.Msg = "Conflict"
		} else if (rs.Code == 412) {
      rs.Msg = "Precondition Failed"
    }

		this.Data["json"] = *rs
		this.ServeJSON()
	}
}
