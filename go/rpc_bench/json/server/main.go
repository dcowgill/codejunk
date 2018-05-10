package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	customer "../customer"
)

type server struct {
	savedCustomers []*customer.CustomerRequest
	mu             sync.Mutex
}

func (s *server) handleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var cr *customer.CustomerRequest
	if err := json.Unmarshal(body, &cr); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	s.mu.Lock()
	s.savedCustomers = append(s.savedCustomers, cr)
	s.mu.Unlock()
}

func (s *server) handleGetCustomers(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var filter *customer.CustomerFilter
	if err := json.Unmarshal(body, &filter); err != nil {
		http.Error(w, "bad request", 400)
		return
	}
	s.mu.Lock()
	customers := s.savedCustomers
	s.mu.Unlock()
	var rsp []*customer.CustomerRequest
	for _, customer := range customers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		rsp = append(rsp, customer)
	}
	data, err := json.Marshal(rsp)
	if err != nil {
		panic(err)
	}
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.Write(data)
	w.Write([]byte{'\n'})
}

func main() {
	s := &server{}
	h := http.NewServeMux()
	h.HandleFunc("/customers/create", s.handleCreateCustomer)
	h.HandleFunc("/customers/get", s.handleGetCustomers)
	hs := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", 50051),
		Handler: h,
	}
	log.Fatal(hs.ListenAndServe())

}
