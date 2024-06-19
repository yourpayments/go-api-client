package main

import(
	"fmt"
	"github.com/yourpayments/go-api-client"
)

func main(){

	merchant := ypmn.GetMerchantFromEnv() // Получить данные мерчанта из environment
	// merchant := ypmn.Merchant{MerchantCode:"CC1", MerchantSecret:"SECRET_KEY"}
	// Можно передать данные вручную

	refund = Refund{
		PayuPaymentReference: "123456"
		OriginalAmount: 100
		Amount: 100
		Currency: "RUB"	
	}

	responce, err := apiRequest.SendRefundRequest(refund)

	if(err != nil){
		panic("Error: " + err.Error())
	}
	
	code, _ := responce["code"].(float64)

	if(int(code) == 429){
		panic("Слишком много запросов")
	}
	if(int(code) == 202){
		panic("Операция возврата уже была совершена")
	}
	if(int(code) != 200){
		fmt.Println(responce)
		panic("Bad responce")
	}

	fmt.Println("Успешный возврат")

}