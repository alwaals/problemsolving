package main

import (
	"encoding/json"
	"net/http"
)

type Ticket struct {
	FromLocation   string
	ToLocation     string
	Price          float64
	NoOfPassengers int
}

var tickets = []Ticket{}

func main() {
	http.Handle("/add", http.HandlerFunc(add))
	http.Handle("/get", http.HandlerFunc(get))
	http.ListenAndServe(":8080", nil)
}
func add(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		return
	}
	tickets = append(tickets, ticket)
	resp, _ := json.Marshal(tickets)
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
}
func get(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(tickets)
	w.Write(resp)
	w.WriteHeader(http.StatusOK)
}
