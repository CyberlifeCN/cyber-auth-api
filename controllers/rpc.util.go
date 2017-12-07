package controllers

import (
  "log"
  "net/rpc"
)

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
