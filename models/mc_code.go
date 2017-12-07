package models

import (
  "encoding/json"
  "fmt"

  "github.com/bradfitz/gomemcache/memcache"
)


func FindVerifyCode(uid string) *VerifyCode {
  // Connect to our memcache instance
  mc := memcache.New("127.0.0.1:11211")

  fmt.Println("uid:", uid)
  // Get a single value
  val, err := mc.Get(uid)
  if err != nil {
    return nil
  }

  fmt.Println("Item:", val)
  fmt.Println("Key:", val.Key)
  fmt.Println("Value:", val.Value)

  var code = &VerifyCode{}
  json.Unmarshal(val.Value, &code)
  fmt.Println("VerifyCode:", code)

	return code
}


func AddVerifyCode(code VerifyCode) {
  // Connect to our memcache instance
  mc := memcache.New("127.0.0.1:11211")

  fmt.Println("VerifyCode:", code)
  jsonBytes, err := json.Marshal(code)
  if err != nil {
    panic(err)
  }
  fmt.Println("Bytes:", jsonBytes)

  // Set some values
  mc.Set(&memcache.Item{Key: code.Id, Value: []byte(jsonBytes)})
}
