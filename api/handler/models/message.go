package models

type Message struct {
	Title    string `json:"title"`
	Message  string `json:"message"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

type MessageResp struct {
	MessageId string `json:"message_id"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type StatusUpdate struct {
	Id                 string `json:"id"`
	MessageId          string `json:"message_id"`
	PermissionSender   string `json:"permission_sender"`
	PermissionReceiver string `json:"permission_receiver"`
}

type MessageList struct {
	Messages []MessageResp `json:"message"`
	Count    int64         `json:"count"`
}
