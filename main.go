package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

type User struct {
	Name	string 	`json:"name"`
	Age		int 	`json:"age"`
}

func main() {
	router := mux.NewRouter()
	fmt.Println("Listening on port 8888...")
	router.HandleFunc("/", RootHandler).Methods("GET")
	router.HandleFunc("/isgomuxup", HealthCheck).Methods("GET")
	router.HandleFunc("/sendUsers", ProcessUsers).Methods("POST")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	response := &Response{
		Status: http.StatusOK,
		Body:   "Welcome to go-mux!",
	}
	data, jsonErr := json.Marshal(response)
	if jsonErr != nil {
		fmt.Printf("Error creating response")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsonErr)
	} else {
		w.WriteHeader(response.Status)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		//json.NewEncoder(w).Encode("Welcome to go-mux!")
		//fmt.Fprint(w, "Welcome to go-mux!")
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := &Response{
		Status: http.StatusOK,
		Body:   "go-mux is up!",
	}
	data, jsonErr := json.Marshal(response)
	if jsonErr != nil {
		fmt.Printf("Error creating response")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsonErr)
	} else {
		w.WriteHeader(response.Status)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		//json.NewEncoder(w).Encode("Welcome to go-mux!")
		//fmt.Fprint(w, "go-mux is up!")
	}
}

//decode json array into array struct
func ProcessUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	response := &Response{
		Status: http.StatusOK,
		Body:   "Processed!",
	}
	data, _ := json.Marshal(response)
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		fmt.Printf("womp womp\n")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		fmt.Printf("%+v\n", users)
		w.WriteHeader(response.Status)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}
