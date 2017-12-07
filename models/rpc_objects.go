package models

type CreateLoginArgs struct {
  Username      string
  Md5Password   string
  Code          string
}

type CreateLoginReply struct {
  Id            string
  Status        int
}

type CreateTicketArgs struct {
  Username      string
  Md5Password   string
}

type RetrieveTicketArgs struct {
  AccessToken   string
}

type RefreshTicketArgs struct {
  RefreshToken  string
}

type CreateCodeArgs struct {
  Id            string
}

type CreateCodeReply struct {
  Code          string
  Status        int
}

type RetrieveCodeArgs struct {
  Id            string
  Type          string
}
