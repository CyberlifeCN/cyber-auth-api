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

  login := models.FindAuthLogin(args.Id)
  if login != nil {
    var rs = &models.CreateLoginReply{}
    rs.Status = 409
    *reply = *rs
    return nil
  } else {
    timestamp := models.GetTimestamp()

    // check verify_code
    code := models.FindRegisterVerifyCode(args.Id)
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
    register.Id = args.Id
    register.AccountId = models.GetUuidString()
    register.Ctime = timestamp
    register.Salt = models.RandomString(10, "a0")
    md5_salt := models.GetMd5String(register.Salt)
    hash_pwd := models.GetMd5String(args.Md5pwd + md5_salt)
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

  login := models.FindAuthLogin(args.Id)
  if login != nil {
    fmt.Println("got login: %d", login)
    md5_salt := models.GetMd5String(login.Salt)
    hash_pwd := models.GetMd5String(args.Md5pwd + md5_salt)
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


func (t *Auth) CreateRegisterCode(args *models.CreateRegisterCodeArgs, reply *models.CreateRegisterCodeReply) error {
  fmt.Println("Auth.CreateRegisterCode: %d", args)

  login := models.FindAuthLogin(args.Id)
  fmt.Println("Auth.login: %d", login)
  if (login != nil) {
    var rs = &models.CreateRegisterCodeReply{}
    rs.Status = 412
    *reply = *rs
    return nil
  }

  timestamp := models.GetTimestamp()
  code := models.FindRegisterVerifyCode(args.Id)
  if (code != nil) {
    if (code.ExpiresAt > timestamp) {
      var rs = &models.CreateRegisterCodeReply{}
      rs.Code = code.Code
      rs.Status = 409
      *reply = *rs
      return nil
    }
  }

  // create verify_code for register/lost-pwd/...
  var verify_code = &models.RegisterVerifyCode{}
  verify_code.Id = args.Id
  // verify_code.AccountId = login.AccountId
  verify_code.ExpiresAt = timestamp + 300000 // 5mins
  verify_code.Code = models.RandomString(6, "a0")

  models.AddRegisterVerifyCode(*verify_code)
  var rs = &models.CreateRegisterCodeReply{}
  rs.Code = verify_code.Code
  rs.Status = 200
  *reply = *rs
  return nil
}


func (t *Auth) CreateLostpwdCode(args *models.CreateLostpwdCodeArgs, reply *models.CreateLostpwdCodeReply) error {
  fmt.Println("Auth.CreateLostpwdCode: %d", args)

  timestamp := models.GetTimestamp()
  code := models.FindLostpwdVerifyCode(args.Id)
  if (code != nil) {
    if (code.ExpiresAt > timestamp) {
      var rs = &models.CreateLostpwdCodeReply{}
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
  var verify_code = &models.LostpwdVerifyCode{}
  verify_code.Id = args.Id
  // verify_code.AccountId = login.AccountId
  verify_code.ExpiresAt = timestamp + 300000 // 5mins
  verify_code.Code = models.RandomString(6, "a0")

  models.AddLostpwdVerifyCode(*verify_code)
  var rs = &models.CreateLostpwdCodeReply{}
  rs.Code = verify_code.Code
  rs.Status = 200
  *reply = *rs
  return nil
}


func (t *Auth) Lostpwd(args *models.LostpwdArgs, reply *models.LostpwdReply) error {
  fmt.Println("Auth.Lostpwd: %d", args)

  login := models.FindAuthLogin(args.Id)
  if login == nil {
    var rs = &models.LostpwdReply{}
    rs.Status = 404
    *reply = *rs
    return nil
  } else {
    timestamp := models.GetTimestamp()

    // check verify_code
    code := models.FindLostpwdVerifyCode(args.Id)
    if (code == nil) {
      var rs = &models.LostpwdReply{}
      rs.Status = 412 // 服务器未满足请求者在请求中设置的其中一个前提条件。
      *reply = *rs
      return nil
    } else {
      if (code.Code != args.Code) {
        var rs = &models.LostpwdReply{}
        rs.Status = 412 // 服务器未满足请求者在请求中设置的其中一个前提条件。
        *reply = *rs
        return nil
      }
    }

    if (code.ExpiresAt < timestamp) {
      var rs = &models.LostpwdReply{}
      rs.Status = 408
      *reply = *rs
      return nil
    }

    // update salt,hash_pwd in mysql:auth_login
    salt := models.RandomString(10, "a0")
    md5_salt := models.GetMd5String(salt)
    hash_pwd := models.GetMd5String(args.Md5pwd + md5_salt)
    models.UpdateAuthLogin(args.Id, salt, hash_pwd)

    var rs = &models.LostpwdReply{}
    rs.Status = 200
    rs.Id = args.Id
    *reply = *rs
    return nil
  }
}


func (t *Auth) RetrieveRegisterCode(args *models.RetrieveRegisterCodeArgs, reply *models.RegisterVerifyCode) error {
  fmt.Println("Auth.RetrieveCode: %d", args)

  code := models.FindRegisterVerifyCode(args.Id)
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


func (t *Auth) Logout(args *models.LogoutArgs, reply *models.LogoutReply) error {
  fmt.Println("Auth.Logout: %d", args)

  ticket := models.FindSessionTicket(args.AccessToken)
  if ticket == nil {
    code := &models.LogoutReply{}
    code.Status = 404
    *reply = *code
    return nil
  }

  // TODO delete session_ticket by access_token from memcache
  models.DeleteSessionTicket(ticket.Id)
  // TODO delete session_ticket by refresh_token from mysql:auth_ticket
  models.DeleteRefreshTicket(ticket.RefreshToken)

  code := &models.LogoutReply{}
  code.Status = 200
  *reply = *code
  return nil
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
