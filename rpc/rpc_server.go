package main

import (
    "fmt"
    "net"
    "net/rpc"
    "os"
    // "log"

    "cyber-auth-api/models"
)

type Auth int


func (t *Auth) CreateLogin(args *models.CreateLoginArgs, reply *models.CreateLoginReply) error {
  fmt.Println("Auth.CreateLogin: %d", args)

  login := models.FindAuthLogin(args.Username)
  if login != nil {
    var rs = &models.CreateLoginReply{}
    rs.Status = 409
    *reply = *rs
    return nil
  } else {
    timestamp := models.GetTimestamp()

    // check verify_code
    code := models.FindVerifyCode(args.Username)
    if (code == nil) {
      var rs = &models.CreateLoginReply{}
      rs.Status = 412 // 服务器未满足请求者在请求中设置的其中一个前提条件。
      *reply = *rs
      return nil
    } else {
      if (code.Code != args.Code) {
        var rs = &models.CreateLoginReply{}
        rs.Status = 412 // 服务器未满足请求者在请求中设置的其中一个前提条件。
        *reply = *rs
        return nil
      }
    }

    if (code.ExpiresAt < timestamp) {
      var rs = &models.CreateLoginReply{}
      rs.Status = 408
      *reply = *rs
      return nil
    }

    // create account_id for login
    var register = &models.AuthLogin{}
    register.Id = args.Username
    register.AccountId = models.GetUuidString()
    register.Ctime = timestamp
    register.Salt = models.RandomString(10, "a0")
    md5_salt := models.GetMd5String(register.Salt)
    hash_pwd := models.GetMd5String(args.Md5Password + md5_salt)
    register.HashPwd = hash_pwd

    models.AddAuthLogin(*register)
    var rs = &models.CreateLoginReply{}
    rs.Status = 200
    rs.Id = register.AccountId
    *reply = *rs
    return nil
  }
}


func (t *Auth) CreateTicket(args *models.CreateTicketArgs, reply *models.SessionTicket) error {
  fmt.Println("Auth.CreateTicket: %d", args)

  login := models.FindAuthLogin(args.Username)
  if login != nil {
    fmt.Println("got login: %d", login)
    md5_salt := models.GetMd5String(login.Salt)
    hash_pwd := models.GetMd5String(args.Md5Password + md5_salt)
    fmt.Println("got hash_pwd: %d", hash_pwd)
    if (hash_pwd == login.HashPwd) {
      fmt.Println("got loginname & password are correct")

      timestamp := models.GetTimestamp()

      // create access_token for session_ticket
      var ticket = &models.SessionTicket{}
      ticket.AccountId = login.AccountId
      ticket.ExpiresAt = timestamp + 7200000 // 2hours
      ticket.RefreshToken = models.GetUuidString()
      ticket.TokenType = "Bearer"
      ticket.Scope = "all"
      ticket.Id = models.GetUuidString()
      models.AddSessionTicket(*ticket)

      // create refresh_token for session_ticket
      var refresh_ticket = &models.RefreshTicket{}
      refresh_ticket.AccountId = login.AccountId
      refresh_ticket.AccessToken = ticket.Id
      refresh_ticket.ExpiresAt = timestamp + 108000000 // 30days
      refresh_ticket.Id = ticket.RefreshToken
      refresh_ticket.TokenType = "Bearer"
      refresh_ticket.Scope = "ticket"
      models.AddRefreshTicket(*refresh_ticket)

      *reply = *ticket
      return nil
    }
  }

  reply = nil
  return nil
}


func (t *Auth) RetrieveTicket(args *models.RetrieveTicketArgs, reply *models.SessionTicket) error {
  fmt.Println("Auth.RetrieveTicket: %d", args)

  ticket := models.FindSessionTicket(args.AccessToken)
  if ticket != nil {
    if (ticket.ExpiresAt < models.GetTimestamp()) {
      reply = nil
      return nil
    } else {
      *reply = *ticket
      return nil
    }
  } else {
    reply = nil
    return nil
  }
}


func (t *Auth) RefreshTicket(args *models.RefreshTicketArgs, reply *models.RefreshTicket) error {
  fmt.Println("Auth.RefreshTicket: %d", args)

  timestamp := models.GetTimestamp()
  refresh_ticket := models.FindRefreshTicket(args.RefreshToken)
  if refresh_ticket != nil {
    if (refresh_ticket.ExpiresAt < timestamp) {
      reply = nil
      return nil
    } else {
      // create access_token for session_ticket
      var session_ticket = &models.SessionTicket{}
      session_ticket.AccountId = refresh_ticket.AccountId
      session_ticket.ExpiresAt = timestamp + 7200000 // 2hours
      session_ticket.RefreshToken = refresh_ticket.Id
      session_ticket.TokenType = "Bearer"
      session_ticket.Scope = "ticket"
      session_ticket.Id = models.GetUuidString()
      models.AddSessionTicket(*session_ticket)

      refresh_ticket.AccessToken = session_ticket.Id
      models.UpdateRefreshTicket(refresh_ticket.Id, refresh_ticket.AccessToken)

      *reply = *refresh_ticket
      return nil
    }
  } else {
    reply = nil
    return nil
  }
}


func (t *Auth) CreateCode(args *models.CreateCodeArgs, reply *models.CreateCodeReply) error {
  fmt.Println("Auth.CreateCode: %d", args)

  timestamp := models.GetTimestamp()
  code := models.FindVerifyCode(args.Id)
  if (code != nil) {
    if (code.ExpiresAt > timestamp) {
      var rs = &models.CreateCodeReply{}
      rs.Code = code.Code
      rs.Status = 409
      *reply = *rs
      return nil
    }
  }

  // login := models.FindAuthLogin(args.Id)
  // fmt.Println("Auth.login: %d", login)
  // if (login == nil || login.Id == "") {
  //   var rs = &models.CreateCodeReply{}
  //   rs.Status = 404
  //   *reply = *rs
  //   return nil
  // }

  // create verify_code for register/lost-pwd/...
  var verify_code = &models.VerifyCode{}
  verify_code.Id = args.Id
  // verify_code.AccountId = login.AccountId
  verify_code.ExpiresAt = timestamp + 300000 // 5mins
  verify_code.Code = models.RandomString(6, "a0")

  models.AddVerifyCode(*verify_code)
  var rs = &models.CreateCodeReply{}
  rs.Code = verify_code.Code
  rs.Status = 200
  *reply = *rs
  return nil
}


func (t *Auth) RetrieveCode(args *models.RetrieveCodeArgs, reply *models.VerifyCode) error {
  fmt.Println("Auth.RetrieveCode: %d", args)

  code := models.FindVerifyCode(args.Id)
  if code != nil {
    if (code.ExpiresAt < models.GetTimestamp()) {
      reply = nil
      return nil
    } else {
      *reply = *code
      return nil
    }
  } else {
    reply = nil
    return nil
  }
}


func main() {
    auth := new(Auth)
    rpc.Register(auth)

    tcpAddr, err := net.ResolveTCPAddr("tcp", ":12345")
    if err != nil {
        fmt.Println("Fatal error:", err)
        os.Exit(1)
    }

    listener, err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        fmt.Println("Fatal error:", err)
        os.Exit(1)
    }

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        rpc.ServeConn(conn)
    }

    // for {
    //     time.Sleep(1 * time.Second)
    // }
}
