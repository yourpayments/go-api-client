package ypmn

import (
    "crypto/md5"
    "encoding/json"
	"encoding/hex"
	"crypto/hmac"
 	"crypto/sha256"
	"crypto/tls"
	"net/http"
	"net/url"
	"io/ioutil"
	"time"
	"bytes"
	"github.com/google/go-querystring/query"
)
	
const HOST	 = "https://secure.ypmn.ru"
// const SANDBOX_HOST  = "https://sandbox.ypmn.ru" 
const SANDBOX_HOST  = "https://gap.ru.payu.local" 

const AUTHORIZE_API = "/api/v4/payments/authorize"
const CAPTURE_API = "/api/v4/payments/capture"
const REFUND_API = "/api/v4/payments/refund"
const REPORT_GENERAL_API = "/api/v4/reports/general"
const REPORT_CHART_API = "/api/v4/reports/chart"
const REPORT_ORDER_API = "/api/v4/reports/orders"

type ApiRequest struct {
	Merchant Merchant
	SandboxModeIsOn bool
	DebugModeIsOn bool
}

func (request ApiRequest) SendAuthRequest(payment Payment) (map[string]interface{}, error) {
	return request.sendPostRequest(payment, AUTHORIZE_API)
}

func (request ApiRequest) SendCaptureRequest(capture Capture) (map[string]interface{}, error) {
	return request.sendPostRequest(capture, CAPTURE_API)
}

func (request ApiRequest) SendRefundRequest(capture Capture) (map[string]interface{}, error) {
	return request.sendPostRequest(capture, REFUND_API)
}

func (request ApiRequest) sendPostRequest(data interface{}, api string) (map[string]interface{}, error) {

	json_data, err := json.Marshal(data)

	h := md5.New() 
	h.Write([]byte(json_data))
	hash := hex.EncodeToString(h.Sum(nil))
	date := time.Now().Format("2006-01-02T15:04:05+00:00")
	

	var tr *http.Transport

	if request.DebugModeIsOn {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		tr = &http.Transport{}
	}

	client := &http.Client {
		Transport: tr,
	}


	urlHost := ""

	if request.SandboxModeIsOn {
        urlHost = SANDBOX_HOST
    } else {
        urlHost = HOST
    } 

	url := urlHost + api


	req, err := http.NewRequest("POST", url, bytes.NewReader(json_data))
  
	if err != nil {
	  return nil, err
	}

	signature := GetSignature(request.Merchant, date, url, "POST", hash)
	
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Header-Signature", signature)
	req.Header.Add("X-Header-Merchant", request.Merchant.MerchantCode)
	req.Header.Add("X-Header-Date", date)
	req.Header.Add("Content-Type", "application/json")
  
	res, err := client.Do(req)
	if err != nil {
	  return nil, err
	}
	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var jsonMap map[string]interface{}
    err = json.Unmarshal(body, &jsonMap)

	return jsonMap, err
}

func (request ApiRequest) SendReportGeneralRequest(dateStart time.Time, dateEnd time.Time, periodLength string) (map[string]interface{}, error) {
	type GeneralReportRequest struct {
		DateStart string `url:"dateStart"`
		DateEnd string `url:"dateEnd"`
		PeriodLength string `url:"periodDate"`
		Merchant string `url:"merchant"`
	} 

	if(periodLength == ""){
		periodLength = DAY
	}

	if(dateEnd.IsZero()){
		dateEnd = time.Now()
	}

	data := GeneralReportRequest{
		DateStart: dateStart.Format("2006-01-02"),
		DateEnd: dateStart.Format("2006-01-02"),
		PeriodLength: periodLength,
	}
	return request.sendGetRequest(data, REPORT_GENERAL_API)	
}

func (request ApiRequest) SendReportChartRequest(dateStart time.Time, dateEnd time.Time, periodLength string) (map[string]interface{}, error){
	type ChartReportRequest struct{
		DateStart string `url:"dateStart"`
		DateEnd string `url:"dateEnd"`
		PeriodLength string `url:"periodDate"`
		Merchant string `url:"merchant"`
	} 

	if(periodLength == ""){
		periodLength = DAY
	}

	if(dateEnd.IsZero()){
		dateEnd = time.Now()
	}

	data := ChartReportRequest{
		DateStart: dateStart.Format("2006-01-02"),
		DateEnd: dateStart.Format("2006-01-02"),
		PeriodLength: periodLength,
	}
	return request.sendGetRequest(data, REPORT_CHART_API)	
}

func (request ApiRequest) SendReportOrderRequest(dateStart time.Time, dateEnd time.Time, periodLength string) (map[string]interface{}, error){
	type ChartReportRequest struct{
		DateStart string `url:"dateStart"`
		DateEnd string `url:"dateEnd"`
		PeriodLength string `url:"periodDate"`
		Merchant string `url:"merchant"`
	} 

	if(periodLength == ""){
		periodLength = DAY
	}

	if(dateEnd.IsZero()){
		dateEnd = time.Now()
	}

	data := ChartReportRequest{
		DateStart: dateStart.Format("2006-01-02"),
		DateEnd: dateStart.Format("2006-01-02"),
		PeriodLength: periodLength,
	}
	return request.sendGetRequest(data, REPORT_ORDER_API)	
}

func (request ApiRequest) sendGetRequest(data interface{}, api string) (map[string]interface{}, error) {


	h := md5.New() 
	hash := hex.EncodeToString(h.Sum(nil))
	date := time.Now().Format("2006-01-02T15:04:05+00:00")
	

	var tr *http.Transport

	if request.DebugModeIsOn {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		tr = &http.Transport{}
	}

	client := &http.Client {
		Transport: tr,
	}


	urlHost := ""

	if request.SandboxModeIsOn {
        urlHost = SANDBOX_HOST
    } else {
        urlHost = HOST
    } 

	vals, _ := query.Values(data)
	params := vals.Encode()
	
	url := urlHost + api + "/?" + params

	req, err := http.NewRequest("GET", url, nil)


	if err != nil {
	  return nil, err
	}

	signature := GetSignature(request.Merchant, date, url, "GET", hash)
	
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Header-Signature", signature)
	req.Header.Add("X-Header-Merchant", request.Merchant.MerchantCode)
	req.Header.Add("X-Header-Date", date)
	req.Header.Add("Content-Type", "application/json")
  
	res, err := client.Do(req)
	if err != nil {
	  return nil, err
	}
	defer res.Body.Close()
  
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var jsonMap map[string]interface{}
    err = json.Unmarshal(body, &jsonMap)

	return jsonMap, err
}

func GetSignature (merchant Merchant, date string, apiUrl string, httpMethod string, bodyHash string) string {

	urlParts, _ := url.Parse(apiUrl)

	urlHashableParts := httpMethod + urlParts.Path

	if urlParts.RawQuery != "" {
		urlHashableParts += urlParts.RawQuery
	}

	
	hashableString := merchant.MerchantCode + date + urlHashableParts + bodyHash


	h := hmac.New(sha256.New, []byte(merchant.MerchantSecret))
	h.Write([]byte(hashableString))
	return hex.EncodeToString(h.Sum(nil))

	
}