package models

type DefMsg struct {
	Title   string `json:"title"`
	Comment string `json:"comment"`
}

type DefMsgResp struct {
	DefMsgId  string `json:"defmsg_id"`
	Title     string `json:"title"`
	Comment   string `json:"comment"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type DefMsgUpdate struct {
	Title    string `json:"title"`
	Comment  string `json:"comment"`
}

type DefMsgsList struct {
	DefMsgs []DefMsgResp `json:"defmsgs"`
	Count   int64        `json:"count"`
}
