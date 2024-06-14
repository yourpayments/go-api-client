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
)
	
const HOST	 = "https://secure.ypmn.ru"
const SANDBOX_HOST  = "https://sandbox.ypmn.ru" 

const AUTHORIZE_API = "/api/v4/payments/authorize"

type ApiRequest struct {
	Merchant Merchant
	SandboxModeIsOn bool
	DebugModeIsOn bool
}

func (request ApiRequest) SendAuthRequest(payment Payment) (map[string]interface{}, error) {
	return request.sendPostRequest(payment, AUTHORIZE_API)
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

	url := urlHost + AUTHORIZE_API


	req, err := http.NewRequest("POST", url, bytes.NewReader(json_data))
  
	if err != nil {
	  return nil, err
	}

	signature := getSignature(request.Merchant, date, url, "POST", hash)
	
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

func getSignature (merchant Merchant, date string, apiUrl string, httpMethod string, bodyHash string) string {

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