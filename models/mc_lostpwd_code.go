package models

import (
  "encoding/json"
  "fmt"

  "github.com/bradfitz/gomemcache/memcache"
)


func FindLostpwdVerifyCode(uid string) *LostpwdVerifyCode {
  // Connect to our memcache instance
  mc := memcache.New("127.0.0.1:11211")

  fmt.Println("uid:", uid)
  // Get a single value
  val, err := mc.Get("lostpwd/" + uid)
  if err != nil {
    return nil
  }

  fmt.Println("Item:", val)
  fmt.Println("Key:", val.Key)
  fmt.Println("Value:", val.Value)

  var code = &LostpwdVerifyCode{}
  json.Unmarshal(val.Value, &code)
  fmt.Println("LostpwdVerifyCode:", code)

	return code
}


func AddLostpwdVerifyCode(code LostpwdVerifyCode) {
  // Connect to our memcache instance
  mc := memcache.New("127.0.0.1:11211")

  fmt.Println("LostpwdVerifyCode:", code)
  jsonBytes, err := json.Marshal(code)
  if err != nil {
    panic(err)
  }
  fmt.Println("Bytes:", jsonBytes)

  // Set some values
  mc.Set(&memcache.Item{Key: "lostpwd/"+code.Id, Value: []byte(jsonBytes)})
}
