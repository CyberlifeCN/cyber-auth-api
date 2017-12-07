package models

import (
  "encoding/json"
  "fmt"

  "github.com/bradfitz/gomemcache/memcache"
)


func FindRegisterVerifyCode(uid string) *RegisterVerifyCode {
  // Connect to our memcache instance
  mc := memcache.New("127.0.0.1:11211")

  fmt.Println("uid:", uid)
  // Get a single value
  val, err := mc.Get("register/" + uid)
  if err != nil {
    return nil
  }

  fmt.Println("Item:", val)
  fmt.Println("Key:", val.Key)
  fmt.Println("Value:", val.Value)

  var code = &RegisterVerifyCode{}
  json.Unmarshal(val.Value, &code)
  fmt.Println("RegisterVerifyCode:", code)

	return code
}


func AddRegisterVerifyCode(code RegisterVerifyCode) {
  // Connect to our memcache instance
  mc := memcache.New("127.0.0.1:11211")

  fmt.Println("RegisterVerifyCode:", code)
  jsonBytes, err := json.Marshal(code)
  if err != nil {
    panic(err)
  }
  fmt.Println("Bytes:", jsonBytes)

  // Set some values
  mc.Set(&memcache.Item{Key: "register/"+code.Id, Value: []byte(jsonBytes)})
}
