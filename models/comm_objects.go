package models


type AuthLogin struct {
  Id            string    `json:"uid"`
  Salt          string    `json:"salt"`
  HashPwd       string    `json:"hash_pwd"`
  AccountId     string    `json:"account_id"`
  Ctime         int64     `json:"ctime"`
}

type SessionTicket struct {
  Id            string    `json:"access_token"`
  AccountId     string    `json:"account_id"`
  RefreshToken  string    `json:"refresh_token"`
  ExpiresAt     int64     `json:"expires_at"`
  TokenType     string    `json:"token_type"`
  Scope         string    `json:"scope"`
}

type RefreshTicket struct {
  Id            string    `json:"refresh_token"`
  AccessToken   string    `json:"access_token"`
  AccountId     string    `json:"account_id"`
  ExpiresAt     int64     `json:"expires_at"`
  TokenType     string    `json:"token_type"`
  Scope         string    `json:"scope"`
}


type RegisterReq struct {
  Username   string         `json:"username"`
  Password   string         `json:"password"`
}

type RegisterResp struct {
	Code   	int 					    `json:"err_code"`
	Msg    	string 						`json:"err_msg"`
  Rs   	 	string 	          `json:"rs"`
}

type LoginReq struct {
  Username   string         `json:"username"`
  Password   string         `json:"password"`
}

type LoginResp struct {
	Code   	int 					    `json:"err_code"`
	Msg    	string 						`json:"err_msg"`
  Rs   	 	SessionTicket 	  `json:"rs"`
}

type RetrieveSessionTicketResp struct {
	Code   	int 					    `json:"err_code"`
	Msg    	string 						`json:"err_msg"`
  Rs   	 	SessionTicket 	  `json:"rs"`
}

type RefreshSessionTicketResp struct {
	Code   	int 					    `json:"err_code"`
	Msg    	string 						`json:"err_msg"`
  Rs   	 	RefreshTicket 	  `json:"rs"`
}
