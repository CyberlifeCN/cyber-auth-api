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
// @Description register by uid,password,verify_code though RPC
// @Param	body		body 	models.RegisterReq	true		"body for register content"
// @Success 200 {object} models.RegisterResp
// @Failure 403 :uid, password or code is empty
// @Failure 409 :uid has register
// @Failure 412 :code not match
// @Failure 408 :code was timeout
// @router / [post]
func (this *RegisterController) Post() {
	uri := this.Ctx.Input.URI()
  beego.Info(uri)
  beego.Info(this.Ctx.Input.RequestBody)

  var req models.RegisterReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	beego.Trace(req)
	if (req.Id == "" || req.Pwd == "" || req.Code == "") {
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
    Id: req.Id,
    Md5pwd: models.GetMd5String(req.Pwd),
		Code: req.Code,
  }
  reply := &models.CreateLoginReply{}
  err = GlobalRpcClient.Call("Auth.CreateLogin", args, &reply)
  if err != nil {
    log.Fatal("CreateLogin error :", err)
  }
  fmt.Println("CreateLogin:", args, reply)

	if (reply == nil) {
    var rs = &models.RegisterResp{
			Code: 409,
			Msg: "Conflict",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
	} else {
		beego.Trace(reply)

		var rs = &models.RegisterResp{
			Code: reply.Status,
			Msg: "Success",
			Rs: reply.Id,
		}
		if (rs.Code == 408) {
      rs.Msg = "Request Timeout"
    } else if (rs.Code == 409) {
      rs.Msg = "Conflict"
		} else if (rs.Code == 412) {
      rs.Msg = "Precondition Failed"
    }

		this.Data["json"] = *rs
		this.ServeJSON()
	}
}


// @Title Code
// @Description retrieve register verify_code by uid(username,phone,email) though RPC
// @Param	body		body 	models.RegisterVerifyCodeReq	true		"body for retrieve content"
// @Success 200 {object} models.RegisterVerifyCodeResp
// @Failure 403 :uid is empty
// @Failure 412 :uid has register
// @Failure 409 :uid has already apply an verify_code for register, that code is expires at 5mins
// @router /code [post]
func (this *RegisterController) Code() {
	uri := this.Ctx.Input.URI()
  beego.Info(uri)
  beego.Info(this.Ctx.Input.RequestBody)

  var req models.RegisterVerifyCodeReq
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	beego.Trace(req)
	if (req.Id == "") {
		var rs = &models.RegisterVerifyCodeResp{
			Code: 403,
			Msg: "Bad Request",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
		return
	}

  var args = &models.CreateRegisterCodeArgs{
    Id: req.Id,
  }
  reply := &models.CreateRegisterCodeReply{}
  err = GlobalRpcClient.Call("Auth.CreateRegisterCode", args, &reply)
  if err != nil {
    log.Fatal("CreateCode error :", err)
  }
  fmt.Println("CreateCode:", args, reply)

	if (reply == nil || reply.Code == "") {
    var rs = &models.RegisterVerifyCodeResp{
			Code: 404,
			Msg: "Not Found",
		}

		this.Data["json"] = *rs
		this.ServeJSON()
	} else {
		beego.Trace(reply)

		var rs = &models.RegisterVerifyCodeResp{
			Code: reply.Status,
			Msg: "Success",
			Rs: reply.Code,
		}
    if (rs.Code == 409) {
      rs.Msg = "Conflict"
    } else if (rs.Code == 412) {
      rs.Msg = "Precondition Failed"
		}

		this.Data["json"] = *rs
		this.ServeJSON()
	}
}
