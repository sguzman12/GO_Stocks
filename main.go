package main

import (
	entity "Main/entity"
	utils "Main/utility"
	"database/sql"

	// api "Main/api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Stock entity.Stock

var Stocks []Stock

func handleRequests(d *sql.DB) {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage(d)).Methods("GET", "OPTIONS")
	router.HandleFunc("/stocks/", returnAllStocks(d)).Methods("GET", "OPTIONS")
	router.HandleFunc("/stocks/{id}", returnSingleStock).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":10000", router))

}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

func homePage(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the HomePage!")
		fmt.Println(db)
		fmt.Println("Endpoint Hit: homePage")
	}
}

func returnAllStocks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Endpoint Hit: returnAllStocks")
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")

		rows, err := db.Query("SELECT * FROM test_table")

		if err != nil {
			w.Write([]byte(err.Error()))
		}
		defer rows.Close()

		var stks []Stock

		for rows.Next() {
			var s Stock

			err := rows.Scan(&s.AlphaID, &s.CompanyName, &s.Value)
			if err != nil {
				w.Write([]byte(err.Error()))
			}

			stks = append(stks, s)
		}
		if err = rows.Err(); err != nil {
			w.Write([]byte(err.Error()))
		}

		json.NewEncoder(w).Encode(stks)
	}

}

func returnSingleStock(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleStock")
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")

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
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to DB
	dbConnection := utils.ConnectDB()

	// // Test Data
	// Stocks = []Stock{
	// 	Stock{AlphaID: "goog", CompanyName: "Google", Value: float32(342.44)},
	// 	Stock{AlphaID: "appl", CompanyName: "Apple", Value: float32(288.21)},
	// 	Stock{AlphaID: "amz", CompanyName: "Amazon", Value: float32(333.23)},
	// 	Stock{AlphaID: "wmt", CompanyName: "Walmart", Value: float32(129.90)},
	// 	Stock{AlphaID: "ldos", CompanyName: "Leidos", Value: float32(98.67)},
	// 	Stock{AlphaID: "dyn", CompanyName: "Dynetics", Value: float32(43.78)},
	// 	Stock{AlphaID: "wdc", CompanyName: "Walt Disney Corporation", Value: float32(114.22)},
	// 	Stock{AlphaID: "msf", CompanyName: "Microsoft", Value: float32(399.02)},
	// 	Stock{AlphaID: "sam", CompanyName: "Samsung", Value: float32(78.43)},
	// }

	handleRequests(dbConnection)
}
