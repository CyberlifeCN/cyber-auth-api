package main

import (
    "fmt"
    "log"
    "net/rpc"
    "cyber-auth-api/models"
)


func main() {
    service := "127.0.0.1:12345"

    client, err := rpc.Dial("tcp", service)
    if err != nil {
        log.Fatal("dialing:", err)
    }

    md5pwd := models.GetMd5String("md5pwd")
    createTicketArgs := models.CreateTicketArgs{"13910586316", md5pwd}
    var createTicketReply models.SessionTicket
    err = client.Call("Mgo.CreateTicket", createTicketArgs, &createTicketReply)
    if err != nil {
        log.Fatal("CreateTicket error :", err)
    }
    fmt.Println("CreateTicket:", createTicketArgs, createTicketReply)

    access_token := createTicketReply.AccessToken
    retrieveTicketArgs := models.RetrieveTicketArgs{access_token}
    var retrieveTicketReply models.SessionTicket
    err = client.Call("Mgo.RetrieveTicket", retrieveTicketArgs, &retrieveTicketReply)
    if err != nil {
        log.Fatal("RetrieveTicket error :", err)
    }
    fmt.Println("RetrieveTicket:", retrieveTicketArgs, retrieveTicketReply)

    refresh_token := retrieveTicketReply.RefreshToken
    refreshTicketArgs := models.RefreshTicketArgs{refresh_token}
    var refreshTicketReply models.SessionTicket
    err = client.Call("Mgo.RefreshTicket", refreshTicketArgs, &refreshTicketReply)
    if err != nil {
        log.Fatal("RefreshTicket error :", err)
    }
    fmt.Println("RefreshTicket:", refreshTicketArgs, refreshTicketReply)
}
