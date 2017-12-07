package models

import (
  "log"
  // "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)


func FindRefreshToken(uid string) *SessionTicket {
  session := GlobalMgoSession.Clone()
  defer session.Close()

  c := session.DB("legend_dev").C("auth_refresh_token")

  //查询数据
  var result = &SessionTicket{}
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


func AddRefreshToken(s SessionTicket) {
	//插入数据
  session := GlobalMgoSession.Clone()
  defer session.Close()

  c := session.DB("legend_dev").C("auth_refresh_token")

  var err = c.Insert(s)
  if err != nil {
    panic(err)
  }
}
