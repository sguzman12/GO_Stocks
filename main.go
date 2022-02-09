package main

import (
	structs "Main/Structs"
	utilities "Main/Utilities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Stock structs.Stock

var Stocks []Stock

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/stocks", returnAllStocks)
	router.HandleFunc("/stocks/{id}", returnSingleStock)
	log.Fatal(http.ListenAndServe(":10000", router))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllStocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllStocks")
	json.NewEncoder(w).Encode(Stocks)
}

func returnSingleStock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleStock")

	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintln(w, "Key: "+key)

	for _, stock := range Stocks {
		if stock.AlphaID == key {
			json.NewEncoder(w).Encode(stock)
		}
	}
}

func main() {
	// Connect to DB
	utilities.ConnectDB()
	Stocks = []Stock{
		Stock{AlphaID: "goog", CompanyName: "Google", Value: float32(342.44)},
		Stock{AlphaID: "appl", CompanyName: "Apple", Value: float32(288.21)},
	}
	handleRequests()
}
