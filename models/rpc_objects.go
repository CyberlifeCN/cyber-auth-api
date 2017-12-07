package models

type CreateLoginArgs struct {
  Id      string
  Md5pwd  string
  Code          string
}

type CreateLoginReply struct {
  Id            string
  Status        int
}

type CreateTicketArgs struct {
  Id      string
  Md5pwd   string
}

type RetrieveTicketArgs struct {
  AccessToken   string
}

type RefreshTicketArgs struct {
  RefreshToken  string
}

type CreateRegisterCodeArgs struct {
  Id            string
}

type CreateRegisterCodeReply struct {
  Code          string
  Status        int
}

type CreateLostpwdCodeArgs struct {
  Id            string
}

type CreateLostpwdCodeReply struct {
  Code          string
  Status        int
}

type LostpwdArgs struct {
  Id      string
  Md5pwd  string
  Code          string
}

type LostpwdReply struct {
  Id            string
  Status        int
}

type RetrieveRegisterCodeArgs struct {
  Id            string
}

type LogoutArgs struct {
  AccessToken   string
}

type LogoutReply struct {
  Status        int
}
