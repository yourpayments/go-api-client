package ypmn

type Client struct {
	Billing Billing `json:"billing"`
	Ip string `json:"clientIp"`
	Time string `json:"clientTime"`
	CommunicationLanguage string `json:"communicationLanguage"`
	Delivery Delivery `json:"delivery"`
}
