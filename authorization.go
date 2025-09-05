package ypmn

const (
    CCVISAMC string = "CCVISAMC"
    FASTER_PAYMENTS string = "FASTER_PAYMENTS"
    BNPL string = "BNPL"
    INTCARD string = "INTCARD" 
    ALFAPAY string = "ALFAPAY" 
    SBERPAY string = "SBERPAY" 
    MIRPAY string = "MIRPAY" 
    TPAY  string = "TPAY" 
)

type Authorization struct {
	UsePaymentPage string `json:"usePaymentPage"`
	PaymentMethod string `json:"paymentMethod"`
}