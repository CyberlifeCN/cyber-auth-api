package models

import (
  "encoding/json"
  "fmt"

  "github.com/bradfitz/gomemcache/memcache"
)


func FindSessionTicket(access_token string) *SessionTicket {
  // Connect to our memcache instance
  mc := memcache.New("127.0.0.1:11211")

  fmt.Println("access_token:", access_token)
  // Get a single value
  val, err := mc.Get(access_token)
  if err != nil {
    return nil
  }

  fmt.Println("Item:", val)
  fmt.Println("Key:", val.Key)
  fmt.Println("Value:", val.Value)

  var ticket = &SessionTicket{}
  json.Unmarshal(val.Value, &ticket)
  fmt.Println("SessionTicket:", ticket)

	return ticket
}


func AddSessionTicket(ticket SessionTicket) {
  // Connect to our memcache instance
  mc := memcache.New("127.0.0.1:11211")

  fmt.Println("SessionTicket:", ticket)
  jsonBytes, err := json.Marshal(ticket)
  if err != nil {
    panic(err)
  }
  fmt.Println("Bytes:", jsonBytes)

  // Set some values
  mc.Set(&memcache.Item{Key: ticket.Id, Value: []byte(jsonBytes)})
}
