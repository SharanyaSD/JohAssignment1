//Sharanya Datrange

package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

//table driven testing

//call the function under test

// Assertion

func TestInputWebsites(t *testing.T) {
	inputJSOn := `{"websites": ["https://google.com","https://facebook.com","xtz.com"]}`
	req, err := http.NewRequest("POST", "/input", bytes.NewBuffer([]byte(inputJSOn)))
	if err != nil {
		t.Fatal(err)
	}
	// expStatusCode := 400

	// if req.Response.StatusCode != expStatusCode {
	// 	fmt.Printf("error")
	// 	t.Error("failed , err %d ", err.Error())
	// }

	w := httptest.NewRecorder()
	inputWebsites(w, req)

	//check 400 bad request
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got sttus %d ", http.StatusOK, w.Code)
	}
}


func TestHandleRequest(t *testing.T) {
	list = []string{"https://google.com", "https://facebook.com", "xtz.com"}
	//sample request
	req, err := http.NewRequest("GET", "/check?input=https://facebook.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	// expStatusCode := 204
	// if req.Response.StatusCode != expStatusCode {

	// 	fmt.Printf("error")
	// 	t.Errorf("failed, error %d ", err.Error())
	// }

	w := httptest.NewRecorder()

	handleRequest(w, req)

	//to check response code
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d got %d ", http.StatusOK, w.Code)

	}

}
