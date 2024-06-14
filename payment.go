package ypmn

type Payment struct {
	MerchantPaymentReference string `json:"merchantPaymentReference"`
	Currency string `json:"currency"`
	ReturnUrl string `json:"returnUrl"`
	Authorization Authorization `json:"authorization"`
	Сlient Client `json:"client"`
	Products []Product `json:"products"`	 
}

func (payment *Payment) AddProduct(product Product) {
    payment.Products = append(payment.Products, product)
}