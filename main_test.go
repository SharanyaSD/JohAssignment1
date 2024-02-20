package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInputWebsites(t *testing.T) {
	var test = map[string][]string{
		"websites": {"www.google.com", "www.facebook.com", "www.fakewebsite1.com"},
	}
	expectedStatusCode := http.StatusOK

	requestObj, _ := json.Marshal(test)

	req, err := http.NewRequest(http.MethodPost, "/input", bytes.NewBuffer(requestObj))
	if err != nil {
		t.Errorf("Error occurred while making http request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(inputWebsites)
	handler.ServeHTTP(rr, req)

	if rr.Code != expectedStatusCode {
		t.Errorf("InputWebsites handler failed, expected status code %d but got %d", expectedStatusCode, rr.Code)
	}
}

func TestHandleRequest(t *testing.T) {
	test := []struct {
		name               string
		queryName          string
		websiteObj         map[string]string
		expectedStatusCode int
	}{
		{
			name:               "Success website",
			expectedStatusCode: http.StatusOK,
			queryName:          "www.google.com",
			websiteObj: map[string]string{
				"www.facebook.com":     "UP",
				"www.fakewebsite1.com": "DOWN",
				"www.google.com":       "UP",
			},
		},
		{
			name:               "Empty map - No websites",
			websiteObj:         map[string]string{},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "No key",
			websiteObj: map[string]string{
				"www.facebook.com":     "UP",
				"www.fakewebsite1.com": "DOWN",
				"www.google.com":       "UP",
			},
			expectedStatusCode: http.StatusBadRequest,
			queryName:          "www.something.com",
		},
		{
			name: "Status of all websites",
			websiteObj: map[string]string{
				"www.facebook.com":     "UP",
				"www.fakewebsite1.com": "DOWN",
				"www.google.com":       "UP",
			},
			expectedStatusCode: http.StatusOK,
		},
	}
	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", fmt.Sprintf("/check?input=%s", test.queryName), bytes.NewBuffer([]byte("")))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handleRequest(rr, req)
			if rr.Code != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Code)
			}
		})
	}
}
