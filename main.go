package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Status int    `json:"status"`
	Body   string `json:"body"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	router := mux.NewRouter()
	fmt.Println("Listening on port 8888...")
	router.HandleFunc("/", RootHandler).Methods("GET")
	router.HandleFunc("/isgomuxup", HealthCheck).Methods("GET")
	router.HandleFunc("/sendUsers", ProcessUsers).Methods("POST")
	router.HandleFunc("/upload", UploadHandler).Methods("POST")
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
	host, _ := os.Hostname()
	response := &Response{
		Status: http.StatusOK,
		Body:   fmt.Sprintf("go-mux is up and running on %v!", host),
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

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// limit upload size to 10 MB
	r.ParseMultipartForm(10 << 20)
	//var fileBuff bytes.Buffer
	file, header, fileErr := r.FormFile("file")
	if fileErr != nil {
		err := fmt.Sprintf("Error reading form file: [%v]", fileErr)
		fmt.Printf("%v\n", err)
		createResponse(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	fmt.Printf("Received file: [%v]\n", header.Filename)
	dst, _ := os.Create("received.zip")
	defer dst.Close()
	io.Copy(dst, file)

	_, readErr := ioutil.ReadFile("./received.zip")
	if readErr != nil {
		err := fmt.Sprintf("Error saving uploaded file: [%v]", readErr)
		fmt.Printf("%v\n", err)
		createResponse(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Printf("Uploaded File Length: [%d]", header.Size)
	createResponse(w, http.StatusAccepted, "")
	return
}

func createResponse(w http.ResponseWriter, status int, body string) {
	//data := bytes(body)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}
