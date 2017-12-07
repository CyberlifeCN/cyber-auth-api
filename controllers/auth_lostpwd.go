package controllers

import (
	"cyber-auth-api/models"
	"encoding/json"
	"github.com/astaxie/beego"
  "log"
  "fmt"
)


// Operations about reset password
type LostpwdController struct {
	beego.Controller
}


// @Title Lost password
// @Description reset password by uid,password,verify_code though RPC
// @Param	body		body 	models.LostpwdReq	true		"body for reset password content"
// @Success 200 {object} models.LostpwdResp
// @Failure 403 :uid, password or code is empty
// @Failure 404 :uid is not register
// @Failure 412 :code not match
// @Failure 408 :code was timeout
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
		} else if (rs.Code == 412) {
      rs.Msg = "Precondition Failed"
    }

		this.Data["json"] = *rs
		this.ServeJSON()
	}
}


// @Title Code
// @Description retrieve reset password verify_code by uid(username,phone,email) though RPC
// @Param	body		body 	models.LostpwdVerifyCodeReq	true		"body for retrieve content"
// @Success 200 {object} models.LostpwdVerifyCodeResp
// @Failure 403 :uid is empty
// @Failure 404 :uid is not register
// @Failure 409 :uid has already apply an verify_code for lost password, that code is expires at 5mins
// @router /code [post]
func (this *LostpwdController) Code() {
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
