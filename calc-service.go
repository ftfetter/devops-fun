package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var history []string

func historyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Buscando histórico de cálculos")
	err := json.NewEncoder(w).Encode(history)
	if err != nil {
		fmt.Println(w, "Erro no parse para JSON")
	}
}

func additionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Parametros: ", vars)
	num1, _ := strconv.ParseFloat(vars["num1"], 32)
	num2, _ := strconv.ParseFloat(vars["num2"], 32)
	total := num1 + num2
	operation := fmt.Sprintf("%.f + %.f = %.f", num1, num2, total)
	history = append(history, operation)
	log.Println("Response: ", operation)
	fmt.Fprintf(w, operation)
}

func subtractionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Parametros: ", vars)
	num1, _ := strconv.ParseFloat(vars["num1"], 32)
	num2, _ := strconv.ParseFloat(vars["num2"], 32)
	total := num1 - num2
	operation := fmt.Sprintf("%.f - %.f = %.f", num1, num2, total)
	history = append(history, operation)
	log.Println("Response: ", operation)
	fmt.Fprintf(w, operation)
}

func multiplicationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Parametros: ", vars)
	num1, _ := strconv.ParseFloat(vars["num1"], 32)
	num2, _ := strconv.ParseFloat(vars["num2"], 32)
	total := num1 * num2
	operation := fmt.Sprintf("%.f * %.f = %.f", num1, num2, total)
	history = append(history, operation)
	log.Println("Response: ", operation)
	fmt.Fprintf(w, operation)
}

func divisionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println("Parametros: ", vars)
	num1, _ := strconv.ParseFloat(vars["num1"], 32)
	num2, _ := strconv.ParseFloat(vars["num2"], 32)
	total := num1 / num2
	operation := fmt.Sprintf("%.f / %.f = %.f", num1, num2, total)
	history = append(history, operation)
	log.Println("Response: ", operation)
	fmt.Fprintf(w, operation)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Working clean")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/calc/{num1}/add/{num2}", additionHandler).Methods("POST")
	router.HandleFunc("/api/v1/calc/{num1}/sub/{num2}", subtractionHandler).Methods("POST")
	router.HandleFunc("/api/v1/calc/{num1}/mul/{num2}", multiplicationHandler).Methods("POST")
	router.HandleFunc("/api/v1/calc/{num1}/div/{num2}", divisionHandler).Methods("POST")
	router.HandleFunc("/api/v1/calc/history", historyHandler).Methods("GET")
	router.HandleFunc("/", defaultHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
