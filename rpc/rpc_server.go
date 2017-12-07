package main

import (
    "fmt"
    "net"
    "net/rpc"
    "os"
    // "log"
    "time"

    "cyber-auth-api/models"
)

type Mgo int


func (t *Mgo) CreateTicket(args *models.CreateTicketArgs, reply *models.SessionTicket) error {
  fmt.Println("Mgo.CreateTicket: %d", args)

  login := models.FindLogin(args.Username)
  if login != nil {
    fmt.Println("got login: %d", login)
    md5_salt := models.GetMd5String(login.Salt)
    hash_pwd := models.GetMd5String(args.Md5Password + md5_salt)
    fmt.Println("got hash_pwd: %d", hash_pwd)
    if (hash_pwd == login.HashPwd) {
      fmt.Println("got loginname & password are correct")

      timestamp := time.Now().UnixNano() / 1000000 // 毫秒

      // create access_token for session_ticket
      var ticket = &models.SessionTicket{}
      ticket.AccountId = login.AccountId
      ticket.AccessToken = models.GetUuidString()
      ticket.ExpiresAt = timestamp + 7200000 // 2hours
      ticket.RefreshToken = models.GetUuidString()
      ticket.TokenType = "Bearer"
      ticket.Scope = "all"
      ticket.Id = ticket.AccessToken
      models.AddAccessToken(*ticket)

      // create refresh_token for session_ticket
      var refresh_ticket = &models.SessionTicket{}
      refresh_ticket.AccountId = login.AccountId
      refresh_ticket.AccessToken = ticket.AccessToken
      refresh_ticket.ExpiresAt = timestamp + 108000000 // 30days
      refresh_ticket.RefreshToken = ticket.RefreshToken
      refresh_ticket.TokenType = "Bearer"
      refresh_ticket.Scope = "ticket"
      refresh_ticket.Id = ticket.RefreshToken
      models.AddRefreshToken(*refresh_ticket)

      *reply = *ticket
      return nil
    }
  }

  reply = nil
  return nil
}

func (t *Mgo) RetrieveTicket(args *models.RetrieveTicketArgs, reply *models.SessionTicket) error {
  fmt.Println("Mgo.RetrieveTicket: %d", args)

  ticket := models.FindAccessToken(args.AccessToken)
  if ticket != nil {
    *reply = *ticket
    return nil
  }

  reply = nil
  return nil
}

func (t *Mgo) RefreshTicket(args *models.RefreshTicketArgs, reply *models.SessionTicket) error {
  fmt.Println("Mgo.RefreshTicket: %d", args)

  ticket := models.FindRefreshToken(args.RefreshToken)
  if ticket != nil {
    *reply = *ticket
    return nil
  }

  reply = nil
  return nil
}

func main() {
    mgo := new(Mgo)
    rpc.Register(mgo)

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
