package models


type AuthLogin struct {
  Id          string    `bson:"_id,omitempty"`
  Salt        string    `bson:"salt,omitempty"`
  CreateTime  int       `bson:"create_time,omitempty"`
  AccountId   string    `bson:"account_id,omitempty"`
  HashPwd     string    `bson:"hash_pwd,omitempty"`
}

type SessionTicket struct {
  Id            string    `bson:"_id,omitempty"`
  AccessToken   string    `bson:"access_token"`
  AccountId     string    `bson:"account_id"`
  RefreshToken  string    `bson:"refresh_token"`
  ExpiresAt     int64     `bson:"expires_at,omitempty"`
  TokenType     string    `bson:"token_type"`
  Scope         string    `bson:"scope"`
}


type LoginReq struct {
  // Id bson.ObjectId `bson:"_id,omitempty"`
  Username   string         `json:"username"`
  Password   string         `json:"password"`
}

type LoginResp struct {
	Code   	int 					     `json:"err_code"`
	Msg    	string 						 `json:"err_msg"`
  Rs   	 	SessionTicket 	   `json:"rs"`
}


type RetrieveSessionTicketResp struct {
	Code   	int 					     `json:"err_code"`
	Msg    	string 						 `json:"err_msg"`
  Rs   	 	SessionTicket 	   `json:"rs"`
}
