package main

import(
	"fmt"
	"github.com/yourpayments/go-api-client"
)

func main(){

	merchant := ypmn.GetMerchantFromEnv() // Получить данные мерчанта из environment
	// merchant := ypmn.Merchant{MerchantCode:"CC1", MerchantSecret:"SECRET_KEY"}
	// Можно передать данные вручную

	externalUrl := "https://example.com/webhook" // Url, по которому обращается ypmn, нужно точное совпадение для проверки signature

	ypmn.WebhookHandler("/webhook", 8000, merchant, externalUrl, func(data map[string]interface{}) {
		fmt.Println(data)
	}, true)
	// Можно отключить проверку подписи для отладки - последний параметр должен быть false

}