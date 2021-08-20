package main
import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Request struct {
	Number1 int `json:"Number1"`
	Number2 int `json:"Number2"`
	ID      string `json:"id"`
	OperationType string `json:"OperationType"`
}

var requests []Request

func getRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}


func getRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range requests {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Request{})
}

func createRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request Request
	_ = json.NewDecoder(r.Body).Decode(&request)
	request.ID = strconv.Itoa(rand.Intn(1000000))
	requests = append(requests, request)
	json.NewEncoder(w).Encode(request)
}


func updateRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range requests {
		if item.ID == params["id"] {
			requests = append(requests[:index], requests[index+1:]...)
			var request Request
			_ = json.NewDecoder(r.Body).Decode(&request)
			request.ID = params["id"]
			requests = append(requests, request)
			json.NewEncoder(w).Encode(request)
			return
		}
	}
	json.NewEncoder(w).Encode(requests)
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range requests {
		if item.ID == params["id"] {
			requests = append(requests[:index], requests[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(requests)
}

func main() {
	r := mux.NewRouter()
	requests = append(requests , Request{ID: "1", Number1: 1, Number2: 2, OperationType: "plus"})
	requests = append(requests , Request{ID: "2", Number1: 4, Number2: 5, OperationType: "minus"})
	r.HandleFunc("/requests", getRequests).Methods("GET")
	r.HandleFunc("/request/{id}", getRequest).Methods("GET")
	r.HandleFunc("/requests", createRequest).Methods("POST")
	r.HandleFunc("/requests/{id}", updateRequest).Methods("PUT")
	r.HandleFunc("/requests/{id}", deleteRequest).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}