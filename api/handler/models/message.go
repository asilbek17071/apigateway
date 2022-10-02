package models

type Message struct {
	Title              string `json:"title"`
	Message            string `json:"message"`
	SenderId           string `json:"sender_id"`
	ReceiverId         string `json:"receiver_id"`
	PermissionSender   string `json:"permission_sender"`
	PermissionReceiver string `json:"permission_receiver"`
}

type MessageResp struct {
	MessageId  string `json:"message_id"`
	Title      string `json:"title"`
	Message    string `json:"message"`
	SenderId   string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
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
