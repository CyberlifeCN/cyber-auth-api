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
    *reply = *rs
    return nil
  } else {
    // create account_id for login
    var register = &models.AuthLogin{}
    register.Id = args.Username
    register.AccountId = models.GetUuidString()
    register.Ctime = models.GetTimestamp()
    register.Salt = models.RandomString(10, "a0")
    md5_salt := models.GetMd5String(register.Salt)
    hash_pwd := models.GetMd5String(args.Md5Password + md5_salt)
    register.HashPwd = hash_pwd

    models.AddAuthLogin(*register)
    var rs = &models.CreateLoginReply{}
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


func (t *Auth) RefreshTicket(args *models.RefreshTicketArgs, reply *models.SessionTicket) error {
  fmt.Println("Auth.RefreshTicket: %d", args)

  ticket := models.FindRefreshToken(args.RefreshToken)
  if ticket != nil {
    *reply = *ticket
    return nil
  }

  reply = nil
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
