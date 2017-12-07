package models

import (
	"time"
	"gopkg.in/mgo.v2"
  // "gopkg.in/mgo.v2/bson"
)

//global
var GlobalMgoSession *mgo.Session

const (
  URL = "mongodb://legend_dev:uWx-nJs-8J3-vA9@123.56.165.59:3717/legend_dev"
)

func init() {
	// init mongodb connection pool
	globalMgoSession, err := mgo.DialWithTimeout(URL, 10 * time.Second)
	if err != nil {
		panic(err)
	}
	GlobalMgoSession = globalMgoSession
	// Optional. Switch the session to a monotonic behavior.
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	//default is 4096
	GlobalMgoSession.SetPoolLimit(30)
}
