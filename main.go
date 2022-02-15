package main

import (
	entity "Main/entity"
	// utils "Main/utility"
	// api "Main/api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type Stock entity.Stock

var Stocks []Stock

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	// originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router.HandleFunc("/", homePage).Methods("GET", "OPTIONS")
	router.HandleFunc("/stocks", returnAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/stocks/{id}", returnSingleStock).Methods("GET", "OPTIONS")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:10000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":10000", handler))

	// log.Fatal(http.ListenAndServe(":10000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllStocks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllStocks")
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
	json.NewEncoder(w).Encode(Stocks)
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
	// utils.ConnectDB()

	// Test Data
	Stocks = []Stock{
		Stock{AlphaID: "goog", CompanyName: "Google", Value: float32(342.44)},
		Stock{AlphaID: "appl", CompanyName: "Apple", Value: float32(288.21)},
		Stock{AlphaID: "amz", CompanyName: "Amazon", Value: float32(333.23)},
		Stock{AlphaID: "wmt", CompanyName: "Walmart", Value: float32(129.90)},
		Stock{AlphaID: "ldos", CompanyName: "Leidos", Value: float32(98.67)},
		Stock{AlphaID: "dyn", CompanyName: "Dynetics", Value: float32(43.78)},
		Stock{AlphaID: "wdc", CompanyName: "Walt Disney Corporation", Value: float32(114.22)},
		Stock{AlphaID: "msf", CompanyName: "Microsoft", Value: float32(399.02)},
		Stock{AlphaID: "sam", CompanyName: "Samsung", Value: float32(78.43)},
	}

	handleRequests()
}
