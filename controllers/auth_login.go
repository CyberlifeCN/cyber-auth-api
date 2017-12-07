package controllers

import (
	"cyber-auth-api/models"
	"encoding/json"
	"github.com/astaxie/beego"
  "log"
  "fmt"
  "net/rpc"
)


// Operations about Login
type LoginController struct {
	beego.Controller
}

var GlobalRpcClient *rpc.Client
var err error

func init() {
  service := "127.0.0.1:12345"
  client, err := rpc.Dial("tcp", service)
  if err != nil {
      log.Fatal("dialing:", err)
  }
  GlobalRpcClient = client
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
	if (req.Username == "" || req.Password == "") {
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
    Username: req.Username,
    Md5Password: models.GetMd5String(req.Password),
  }
  reply := &models.SessionTicket{}
  err = GlobalRpcClient.Call("Mgo.CreateTicket", args, &reply)
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
