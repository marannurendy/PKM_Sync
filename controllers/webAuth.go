package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// "PKM_mekaar/conf"
// "strings"

type WebAuth struct {
	Refresh_token_expires_in string `json:"refresh_token_expires_in"`
	Api_product_list         string `json:"api_product_list"`
	Api_product_list_json    string `json:"api_product_list_json"`
	Organization_name        string `json:"organization_name"`
	DeveloperEmail           string `json:"developer.email"`
	Token_type               string `json:"token_type"`
	Issued_at                string `json:"issued_at"`
	Client_id                string `json:"client_id"`
	Access_token             string `json:"access_token"`
	Application_name         string `json:"application_name"`
	Scope                    string `json:"scope"`
	Expires_in               string `json:"expires_in"`
	Refresh_count            string `json:"refresh_count"`
	Status                   string `json:"status"`
}

type GetWeb struct {
	SellerId       string `json:"sellerid"`
	Name           string `json:"name"`
	BusinessUnit   string `json:"businessUnit"`
	BusinessUnitId string `json:"businessUnitId"`
	Title          string `json:"title"`
}

func (c *MainController) WebAuth() {

	// decoder := json.NewDecoder(c.Ctx.Request.Body)
	// var fuck GetWeb
	// err := decoder.Decode(&fuck)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bytes, err := json.Marshal(fuck)

	// fmt.Println(string(bytes))

	decoder := json.NewDecoder(c.Ctx.Request.Body)
	var fuck GetWeb
	err := decoder.Decode(&fuck)

	lastPoint := "https://sandbox.partner.api.bri.co.id/v1.0/umi/webauth"

	var Anotherdata = []byte(`{
		"sellerId": "` + fuck.SellerId + `",
		"name": "` + fuck.Name + `",
		"businessUnit": "` + fuck.BusinessUnit + `",
		"businessUnitId": "` + fuck.BusinessUnitId + `",
		"title": "` + fuck.Title + `"
	}`)

	token := GetToken()
	method := "POST"
	path := "/v1.0/umi/webauth"
	timestamp := (time.Now().UTC().Format("2006-01-02T15:04:05.000Z"))
	body := Anotherdata
	signature := GetSignature(token, method, path, string(body), timestamp)

	Otherclient := &http.Client{}
	s, err := http.NewRequest(method, lastPoint, bytes.NewBuffer(Anotherdata)) // URL-encoded payload

	if err != nil {
		fmt.Println(err)
	}

	var bearer = "Bearer " + token

	s.Header.Add("Content-Type", "application/json")
	s.Header.Add("BRI-Signature", signature)
	s.Header.Add("BRI-Timestamp", timestamp)
	s.Header.Add("Authorization", bearer)

	resul, err := Otherclient.Do(s)
	if err != nil {
		fmt.Println(err)
	}

	// log.Println(resul.Status)
	defer resul.Body.Close()

	bodyOther, err := ioutil.ReadAll(resul.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bodyOther))

	// var dt = (body)

	// var dataShit WebAuth

	// err = json.Unmarshal([]byte(dt), &dataShit)

	// fmt.Println(dataShit.Client_id)

	c.ServeJSON()
}
