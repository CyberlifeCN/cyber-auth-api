package models

type CreateTicketArgs struct {
  Username      string
  Md5Password   string
}

type RetrieveTicketArgs struct {
  AccessToken   string
}

type RefreshTicketArgs struct {
  RefreshToken   string
}
