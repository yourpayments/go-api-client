package ypmn

type Capture struct{
	PayuPaymentReference string `json:"payuPaymentReference"`
	OriginalAmount float64 `json:"originalAmount"`
	Amount float64 `json:"amount"`
	Currency float64 `json:"currency"`
}