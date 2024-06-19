package main

import(
	"fmt"	
	"github.com/yourpayments/go-api-client"
)

func main(){

	merchant := ypmn.GetMerchantFromEnv() // Получить данные мерчанта из environment
	// merchant := ypmn.Merchant{MerchantCode:"CC1", MerchantSecret:"SECRET_KEY"}
	// Можно передать данные вручную


	apiRequest := ypmn.ApiRequest{
		Merchant: merchant,
		SandboxModeIsOn: true,
		DebugModeIsOn: true,
	}

	responce, _ := apiRequest.SendOrderStatusRequest("60d19e2c-90b1-4677-90ce-f4a42fcbcff9")

	fmt.Println(responce)

}