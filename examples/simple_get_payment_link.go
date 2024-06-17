package main

import(
	"fmt"
	"github.com/yourpayments/go-api-client"
	"github.com/google/uuid"
)

func main(){

	merchant := ypmn.GetMerchantFromEnv() // Получить данные мерчанта из environment
	// merchant := ypmn.Merchant{MerchantCode:"CC1", MerchantSecret:"SECRET_KEY"}
	// Можно передать данные вручную

	merchantPaymentReference := fmt.Sprintf("%s",uuid.New())

	product := ypmn.Product{
		Name:"Заказ № " + merchantPaymentReference,
		SKU: merchantPaymentReference,
		UnitPrice: 200.42,
		Quantity: 1,
	}

	billing := ypmn.Billing{
		CountryCode:"RU",
		FirstName:"Иван",
		LastName:"Петров",
		Email:"test1@ypmn.ru",
		Phone:"+7-800-555-35-35",
		City:"Москва",
	}

	client := ypmn.Client{
		Billing: billing,
	}

	authorization := ypmn.Authorization{
		PaymentMethod: ypmn.CCVISAMC,
		UsePaymentPage: "YES",
	}

	payment := ypmn.Payment{
		Currency:"RUB",
		MerchantPaymentReference: merchantPaymentReference,
		ReturnUrl:"https://example.com/return",
		Сlient: client,
		Authorization: authorization,
	}

	payment.AddProduct(product)

	apiRequest := ypmn.ApiRequest{
		Merchant: merchant,
		SandboxModeIsOn: true,
		DebugModeIsOn: true,
	}

	responce, err := apiRequest.SendAuthRequest(payment)

	if(err != nil){
		panic("Error: " + err.Error())

	}
	

	code, _ := responce["code"].(int64)

	if(responce["code"] == 429){
		panic("Слишком много запросов")
	}
	if(responce["code"] != 200){
		panic("Bad responce code: " + string(code))
	}

	paymentResult, _ := responce["paymentResult"].(map[string]interface{})
	link, _ := paymentResult["url"].(string)

	fmt.Println("Перейдите по ссылке, чтобы оплатить: " + link)

}