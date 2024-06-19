package main

import(
	"fmt"
	"github.com/yourpayments/go-api-client"
)

func main(){

	merchant := ypmn.GetMerchantFromEnv() // Получить данные мерчанта из environment
	// merchant := ypmn.Merchant{MerchantCode:"CC1", MerchantSecret:"SECRET_KEY"}
	// Можно передать данные вручную

	capture := Capture{
		PayuPaymentReference: "123456"
		OriginalAmount: 100
		Amount: 50
		Currency: "RUB"
	}

	responce, err := apiRequest.SendCaptureRequest(capture)

	if(err != nil){
		panic("Error: " + err.Error())
	}
	
	code, _ := responce["code"].(float64)

	if(int(code) == 429){
		panic("Слишком много запросов")
	}
	if(int(code) != 200){
		fmt.Println(responce)
		panic("Bad responce")
	}

	fmt.Println("Списание прошло успешно")

}