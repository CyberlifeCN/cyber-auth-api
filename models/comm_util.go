package models

import (
    "crypto/md5"
    "encoding/hex"
    "github.com/satori/go.uuid"
  	"strings"
    "time"
    "math/rand"
    "bytes"
)

//生成32位md5字串
func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}


//生成32位uuid字串
func GetUuidString() string {
//  uid := strings.Replace(uuid.NewV4().String(), "-", "", -1)
  suid,_ := uuid.NewV4()
  sha1 := suid.String()
  uid := strings.Replace(sha1, "-", "", -1)    
  return uid
}


//生成 timestamp like Unix
func GetTimestamp() int64 {
    timestamp := time.Now().UnixNano() / 1000000 // 毫秒
    return timestamp
}


// 生成随机字符串
func RandomString(randLength int, randType string) (result string) {
    var num string = "0123456789"
    var lower string = "abcdefghijklmnopqrstuvwxyz"
    var upper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

    b := bytes.Buffer{}
    if strings.Contains(randType, "0") {
        b.WriteString(num)
    }
    if strings.Contains(randType, "a") {
        b.WriteString(lower)
    }
    if strings.Contains(randType, "A") {
        b.WriteString(upper)
    }
    var str = b.String()
    var strLen = len(str)
    if strLen == 0 {
        result = ""
        return
    }

    rand.Seed(time.Now().UnixNano())
    b = bytes.Buffer{}
    for i := 0; i < randLength; i++ {
        b.WriteByte(str[rand.Intn(strLen)])
    }
    result = b.String()
    return
}
