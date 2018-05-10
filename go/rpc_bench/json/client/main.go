package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	customer "../customer"
)

const address = "http://0.0.0.0:50051"

type client struct {
	hc *http.Client
}

func (c *client) createCustomer(cust *customer.CustomerRequest) {
	data, err := json.Marshal(cust)
	if err != nil {
		panic(err)
	}
	_, err = c.hc.Post(address+"/customers/create", "application/json; charset=utf-8", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
}

func (c *client) getCustomers(filter *customer.CustomerFilter) {
	data, err := json.Marshal(filter)
	if err != nil {
		panic(err)
	}
	rsp, err := c.hc.Post(address+"/customers/get", "application/json; charset=utf-8", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}
	var results []*customer.CustomerRequest
	if err := json.Unmarshal(body, &results); err != nil {
		panic(err)
	}
	for _, cust := range results {
		fmt.Printf("read customer id=%d\n", cust.Id)
	}
}

func main() {
	c := &client{hc: &http.Client{}}

	// Create 10 identical customers.
	for i := 0; i < 10; i++ {
		cust := &customer.CustomerRequest{
			Id:    int32(101 + i),
			Name:  "Shiju Varghese",
			Email: "shiju@xyz.com",
			Phone: "732-757-2923",
			Addresses: []*customer.CustomerRequest_Address{
				&customer.CustomerRequest_Address{
					Street:            "1 Mission Street",
					City:              "San Francisco",
					State:             "CA",
					Zip:               "94105",
					IsShippingAddress: false,
				},
				&customer.CustomerRequest_Address{
					Street:            "Greenfield",
					City:              "Kochi",
					State:             "KL",
					Zip:               "68356",
					IsShippingAddress: true,
				},
			},
		}
		c.createCustomer(cust)
	}

	const numTrials = 100
	filter := &customer.CustomerFilter{Keyword: "Shiju"}
	t1 := time.Now()
	for i := 0; i < numTrials; i++ {
		c.getCustomers(filter)
	}
	elapsed := time.Since(t1)
	fmt.Printf("trials=%d elapsed=%v per_call=%v\n", numTrials, elapsed, elapsed/time.Duration(numTrials))
}
