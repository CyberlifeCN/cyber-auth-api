package models

import (
  "log"
  // "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)


func FindLogin(uid string) *AuthLogin {
  session := GlobalMgoSession.Clone()
  defer session.Close()

  c := session.DB("legend_dev").C("auth_login")

  //查询数据
  var result = &AuthLogin{}
  query := c.Find(bson.M{"_id": uid})
  var err = query.One(result)
  if err != nil {
    log.Print(err)
    return nil
  }
  // If you must detect "not found" case:
  if result == nil {
      // No result
      return nil
  }

	return result
}
