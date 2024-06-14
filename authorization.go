package ypmn

const (
    CCVISAMC string = "CCVISAMC"
    FASTER_PAYMENTS string = "FASTER_PAYMENTS"
    SOM string = "SOM"
)

type Authorization struct {
	UsePaymentPage string `json:"usePaymentPage"`
	PaymentMethod string `json:"paymentMethod"`
}