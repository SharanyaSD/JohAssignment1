package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var mutex sync.Mutex

type Websites struct {
	WebList []string `json:"websites"`
}

var websitesMap map[string]string = make(map[string]string)
var list []string

func main() {
	fmt.Println("server starting")
	r := mux.NewRouter()
	r.HandleFunc("/", genericHandler).Methods(http.MethodGet)
	r.HandleFunc("/input", inputWebsites).Methods("POST")
	r.HandleFunc("/check", handleRequest).Methods("GET")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("error starting server", err)
	}

}

func inputWebsites(w http.ResponseWriter, r *http.Request) {
	var webs Websites
	err := json.NewDecoder(r.Body).Decode(&webs)

	if err != nil {
		w.WriteHeader((http.StatusBadRequest))
		return
	}

	list = webs.WebList
	w.WriteHeader(http.StatusOK)

}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("input")

	for _, website := range list {
		go func(website string) {
			status := CheckStatusCode(website)
			websitesMap[website] = status
			fmt.Printf("Website %s is %s\n", website, status)
		}(website)
	}
	time.Sleep(5 * time.Second)

	responseJSON, err := json.Marshal(map[string]string{input: websitesMap[input]})
	// fmt.Printf(input, websitesMap[input])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
	return

}

func CheckStatusCode(website string) string {
	fmt.Print("Checking trafiic - Status Code")

	for {
		resp, err := http.Get(website)
		if err != nil {
			return "DOWN"
		}
		defer resp.Body.Close()

		mutex.Lock()
		defer mutex.Unlock()

		// status code =200?
		if resp.StatusCode == http.StatusOK {
			return "UP"
		}
		return "DOWN"
	}
	//time.Sleep(1*time.Minute)

}

func genericHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Page does not exist")
}
