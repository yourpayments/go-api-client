package ypmn

import(
	"os"
)

type Merchant struct {
	MerchantCode string
	MerchantSecret string
}

func GetMerchantFromEnv () Merchant{
	merchantCode := os.Getenv("MERCHANT_CODE")
	merchantSecret := os.Getenv("MERCHANT_SECRET")
	return Merchant{
		merchantCode,
		merchantSecret,
	}
}