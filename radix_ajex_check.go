package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var RADIX_RES_SUCSES uint32
var RADIX_RES_FAIL uint32

type RadixLoginRequest struct {
	Id    uint32
	Uname string
	Psw   string
}

type RadixBasicResponce struct {
	Resp uint32
}

// Load the index.html template.
var tmpl = template.Must(template.New("tmpl").ParseFiles("test_ajex.html"))

func main() {
	InitVariables()
	http.HandleFunc("/", HomePage) // serve / file
	http.HandleFunc("/login", ClientRadixLoginRequest)
	var testInc = 0

	// Serve callme with a text response.
	http.HandleFunc("/table", func(w http.ResponseWriter, r *http.Request) {
		t := strconv.Itoa(testInc)

		fmt.Fprintln(w, "Test Response! "+t)
		testInc++
	})

	// Start the server at http://localhost:8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func InitVariables() {
	RADIX_RES_FAIL = 0
	RADIX_RES_SUCSES = 1212
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	if err := tmpl.ExecuteTemplate(w, "test_ajex.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ClientRadixLoginRequest(w http.ResponseWriter, r *http.Request) {
	var rdxLoginReq RadixLoginRequest
	if r.Body == nil {
		fmt.Println("ERROR :: HTTP Empty Request")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&rdxLoginReq)
	if err != nil {
		fmt.Println("ERROR :: HTTP Jason Decode Failed")
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println("UseName :: " + rdxLoginReq.Uname + " Psw :: " + rdxLoginReq.Psw)

	resp := RadixBasicResponce{RADIX_RES_SUCSES}
	json.NewEncoder(w).Encode(resp)
}

func ClientRadixTableRequest(w http.ResponseWriter, r *http.Request) {

}
