package models

import (
    "crypto/md5"
    "encoding/hex"
    "github.com/satori/go.uuid"
  	"strings"
)

//生成32位md5字串
func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}


//生成32位uuid字串
func GetUuidString() string {
  uid := strings.Replace(uuid.NewV4().String(), "-", "", -1)
  return uid
}
