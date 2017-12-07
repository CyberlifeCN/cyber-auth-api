package controllers

import (
	"cyber-auth-api/models"
	// "encoding/json"
	"github.com/astaxie/beego"
  "log"
  "fmt"
)


// Operations about Ticket
type TicketController struct {
	beego.Controller
}


// @Title Get
// @Description retrieve session_ticket by access_token
// @Param	Authorization		header 	string	"Bearer 81063dfdda8911e79b00a45e60efbf2d"		true		"Authorization=Bearer access_token"
// @Param	access_token		path 	string	true		"The key for session_ticket"
// @Success 200 {object} models.RetrieveSessionTicketResp
// @Failure 403 :access_token is empty
// @router /:access_token [get]
func (t *TicketController) Get() {
	uri := t.Ctx.Input.URI()
  beego.Info(uri)

	access_token := t.GetString(":access_token")
  beego.Trace(access_token)

	var args = &models.RetrieveTicketArgs{
    AccessToken: access_token,
  }
  reply := &models.SessionTicket{}
  err = GlobalRpcClient.Call("Auth.RetrieveTicket", args, &reply)
  if err != nil {
    log.Fatal("RetrieveTicket error :", err)
  }
  fmt.Println("RetrieveTicket:", args, reply)

	if (reply != nil) {
		beego.Trace(reply)

		var rs = &models.RetrieveSessionTicketResp{
			Code: 200,
			Msg: "Success",
			Rs: *reply,
		}

		t.Data["json"] = *rs
		t.ServeJSON()
	} else {
		var rs = &models.RetrieveSessionTicketResp{
			Code: 404,
			Msg: "Not Found",
		}

		t.Data["json"] = *rs
		t.ServeJSON()
	}
}