package main

import(
	"fmt"
	"time"
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

	responce, err := apiRequest.SendReportGeneralRequest(time.Now().AddDate(0, 0, -7),time.Now(), ypmn.DAY)

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

	fmt.Println(responce)

}