package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/paytabs/bank/delivery"
)

func RunCURLScript() {

	log.Println("Create account with ID 10 Balance 10 Location test AccountStatus Active")

	m := delivery.CreateAccountRequest{
		ID:            10,
		Balance:       10,
		Location:      "test",
		AccountStatus: 1,
	}
	reqBody, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/create", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		log.Fatalln(err)
	}
	cl := http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	b, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%+v\n", string(b))

	log.Println("Create account with ID 11 Balance 20 Location test AccountStatus Active")

	m = delivery.CreateAccountRequest{
		ID:            11,
		Balance:       20,
		Location:      "test",
		AccountStatus: 1,
	}
	reqBody, err = json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/create", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err = cl.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	b, _ = ioutil.ReadAll(resp.Body)
	log.Printf("%+v\n", string(b))

	log.Println("Fetch Account details for ID 10")

	f := delivery.AccountDetailsRequest{
		ID: 10,
	}
	reqBody, err = json.Marshal(f)
	if err != nil {
		log.Fatalln(err)
	}
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/fetch", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err = cl.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	b, _ = ioutil.ReadAll(resp.Body)
	log.Printf("%+v\n", string(b))

	log.Println("Fetch Account details for ID 11")

	f = delivery.AccountDetailsRequest{
		ID: 11,
	}
	reqBody, err = json.Marshal(f)
	if err != nil {
		log.Fatalln(err)
	}
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/fetch", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err = cl.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	b, _ = ioutil.ReadAll(resp.Body)
	log.Printf("%+v\n", string(b))

	log.Println("Send money from account ID 10 to account ID 11 with amount 2")

	s := delivery.SendMoneyRequest{
		SourceAccountID:      10,
		DestinationAccountID: 11,
		Amount:               2,
	}
	reqBody, err = json.Marshal(s)
	if err != nil {
		log.Fatalln(err)
	}
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/send", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err = cl.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	b, _ = ioutil.ReadAll(resp.Body)
	log.Printf("%+v\n", string(b))

	log.Println("Send big money from account ID 10 to account ID 11 with amount 11111")

	s = delivery.SendMoneyRequest{
		SourceAccountID:      10,
		DestinationAccountID: 11,
		Amount:               11111,
	}
	reqBody, err = json.Marshal(s)
	if err != nil {
		log.Fatalln(err)
	}
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/send", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err = cl.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	b, _ = ioutil.ReadAll(resp.Body)
	log.Printf("%+v\n", string(b))

	log.Println("Send money from invalid account ID 99 to account ID 11 with amount 2")

	s = delivery.SendMoneyRequest{
		SourceAccountID:      99,
		DestinationAccountID: 11,
		Amount:               2,
	}
	reqBody, err = json.Marshal(s)
	if err != nil {
		log.Fatalln(err)
	}
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/send", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err = cl.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	b, _ = ioutil.ReadAll(resp.Body)
	log.Printf("%+v\n", string(b))

	log.Println("Send money from account ID 10 to invalid destination account ID 98 with amount 2")

	s = delivery.SendMoneyRequest{
		SourceAccountID:      10,
		DestinationAccountID: 98,
		Amount:               2,
	}
	reqBody, err = json.Marshal(s)
	if err != nil {
		log.Fatalln(err)
	}
	req, err = http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/send", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err = cl.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	b, _ = ioutil.ReadAll(resp.Body)
	log.Printf("%+v\n", string(b))

}
