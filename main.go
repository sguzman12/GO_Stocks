package main

import (
	"Main/Structs"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Stock Structs.Stock

var Stocks []Stock

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stocks", returnAllStocks)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllStocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllStocks")
	json.NewEncoder(w).Encode(Stocks)
}

func main() {
	Stocks = []Stock{
		Stock{AlphaID: "goog", CompanyName: "Google", Value: float32(342.44)},
		Stock{AlphaID: "appl", CompanyName: "Apple", Value: float32(288.21)},
	}
	handleRequests()
}
