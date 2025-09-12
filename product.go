package ypmn

type Product struct {
	Name string `json:"name"`
	SKU string `json:"sku"`
	Quantity int64 `json:"quantity"`
	VAT float64 `json:"vat"`
	UnitPrice float64 `json:"unitPrice"`
	Marketplace *MarketplaceProductDetails `json:"marketplace,omitempty"`
}


type MarketplaceProductDetails struct {
	Version int `json:"version"`
	MerchantCode string `json:"merchantCode"`
}