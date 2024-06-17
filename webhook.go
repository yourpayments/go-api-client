package ypmn

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "log"
 "net/http"
 "crypto/md5"
 "encoding/hex"
 "strconv"
)

type CallbackFunc func(map[string]interface{})

func WebhookHandler(path string, port int, merchant Merchant, externalUrl string, callback CallbackFunc, validateSignature bool) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if validateSignature {
		date := r.Header.Get("X-Header-Date")
		signature := r.Header.Get("X-Header-Signature") 
		h := md5.New() 
		h.Write([]byte(body))
		hash := hex.EncodeToString(h.Sum(nil))

		calculatedSignature := GetSignature(merchant, date, externalUrl, "POST", hash)
		if(calculatedSignature != signature){
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	callback(data)

	fmt.Fprintf(w, "OK")
 })

 log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}
