package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() { //see python script for comments. Everything is the same but modified to work with Golang
	uid := getFormUID()
	referral := getReferralCode(uid)
	getMeterID(referral)
	time.Sleep(30 * time.Second)
	meterid := getMeterID(referral)
	fmt.Println("Meter ID: ", meterid)
	activateMeter(meterid)
	getMeterStatus(meterid)
	time.Sleep(60 * time.Second)
	getMeterStatus(meterid)
	getBill(meterid)

}
func getFormUID() string {
	var url = "https://utilityapi.com/api/v2/forms"
	fmt.Println(url)
	client := &http.Client{}
	req, err1 := http.NewRequest("POST", url, nil)
	if err1 != nil {
		fmt.Println("err1", err1)
	}
	req.Header.Add("Authorization", "Bearer 76201cfd80a04c279a92662a07d0b887")
	req.Header.Add("Content-Type", "application/json")
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("err2", err2)
	}
	body, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		fmt.Println("err3", err3)
	}
	jsonRes := make(map[string]interface{})
	jsonErr := json.Unmarshal(body, &jsonRes)
	if jsonErr != nil {
		fmt.Println("jsonErr", jsonErr)
	}
	fmt.Println("Form UID: ", jsonRes["uid"])
	return fmt.Sprintf("%v", jsonRes["uid"])
}
func getReferralCode(uid string) string {
	var url = "https://utilityapi.com/api/v2/forms/" + uid + "/test-submit"
	fmt.Println(url)
	payload, err := json.Marshal(map[string]string{
		"utility":  "DEMO",
		"scenario": "residential",
	})
	if err != nil {
		fmt.Println("err", err)
	}
	client := &http.Client{}
	req, err1 := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err1 != nil {
		fmt.Println("err1", err1)
	}
	req.Header.Add("Authorization", "Bearer 76201cfd80a04c279a92662a07d0b887")
	req.Header.Add("Content-Type", "application/json")
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("err2", err2)
	}
	body, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		fmt.Println("err3", err3)
	}
	jsonRes := make(map[string]interface{})
	jsonErr := json.Unmarshal(body, &jsonRes)
	if jsonErr != nil {
		fmt.Println("jsonErr", jsonErr)
	}
	fmt.Println("Referral: ", jsonRes["referral"])
	return fmt.Sprintf("%v", jsonRes["referral"])
}
func getMeterID(referral string) string {
	var url = "https://utilityapi.com/api/v2/authorizations?referrals=" + referral + "&include=meters"
	fmt.Println(url)
	client := &http.Client{}
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		fmt.Println("err1", err1)
	}
	req.Header.Add("Authorization", "Bearer 76201cfd80a04c279a92662a07d0b887")
	req.Header.Add("Content-Type", "application/json")
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("err2", err2)
	}
	body, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		fmt.Println("err3", err3)
	}
	jsonRes := make(map[string]interface{})
	jsonErr := json.Unmarshal(body, &jsonRes)
	if jsonErr != nil {
		fmt.Println("jsonErr", jsonErr)
	}
	jsonStr := fmt.Sprintf("%v", jsonRes["authorizations"])
	//fmt.Println("jsonStr", jsonStr)
	i := strings.Index(jsonStr, " uid") + 1
	retStr := strings.Replace(jsonStr, jsonStr[0:i], "", 1)
	i2 := strings.Index(retStr, " ")
	id := retStr[4:i2]
	return id
}
func activateMeter(meterid string) string {
	var url = "https://utilityapi.com/api/v2/meters/historical-collection"
	fmt.Println(url)
	list := [1]string{
		meterid,
	}
	payload, err := json.Marshal(map[string]interface{}{
		"meters": list,
	})
	if err != nil {
		fmt.Println("err", err)
	}
	client := &http.Client{}
	req, err1 := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err1 != nil {
		fmt.Println("err1", err1)
	}
	req.Header.Add("Authorization", "Bearer 76201cfd80a04c279a92662a07d0b887")
	req.Header.Add("Content-Type", "application/json")
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("err2", err2)
	}
	body, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		fmt.Println("err3", err3)
	}
	jsonRes := make(map[string]interface{})
	jsonErr := json.Unmarshal(body, &jsonRes)
	if jsonErr != nil {
		fmt.Println("jsonErr", jsonErr)
	}
	fmt.Println("Meter activated?: ", jsonRes["success"])
	return fmt.Sprintf("%v", jsonRes)
}
func getMeterStatus(meterid string) string {
	var url = "https://utilityapi.com/api/v2/meters/" + meterid
	fmt.Println(url)
	client := &http.Client{}
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		fmt.Println("err1", err1)
	}
	req.Header.Add("Authorization", "Bearer 76201cfd80a04c279a92662a07d0b887")
	req.Header.Add("Content-Type", "application/json")
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("err2", err2)
	}
	body, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		fmt.Println("err3", err3)
	}
	jsonRes := make(map[string]interface{})
	jsonErr := json.Unmarshal(body, &jsonRes)
	if jsonErr != nil {
		fmt.Println("jsonErr", jsonErr)
	}
	fmt.Println("Meter status: ", jsonRes["status"])
	return fmt.Sprintf("%v", jsonRes["status"])
}
func getBill(meterid string) string {
	var url = "https://utilityapi.com/api/v2/bills?meters=" + meterid
	fmt.Println(url)
	client := &http.Client{}
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		fmt.Println("err1", err1)
	}
	req.Header.Add("Authorization", "Bearer 76201cfd80a04c279a92662a07d0b887")
	req.Header.Add("Content-Type", "application/json")
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("err2", err2)
	}
	body, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		fmt.Println("err3", err3)
	}
	jsonRes := make(map[string]interface{})
	jsonErr := json.Unmarshal(body, &jsonRes)
	if jsonErr != nil {
		fmt.Println("jsonErr", jsonErr)
	}
	fmt.Println("Bill: ", jsonRes)
	return fmt.Sprintf("%v", jsonRes)
}
