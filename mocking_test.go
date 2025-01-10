package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status": "good"}`)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

type ExpectedResponseBody struct {
	Status string `json:"status"`
}

func TestDownloadMocking(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Download content from remote URL")
	{
		t.Logf("\tTest 0:\tWhen checking valid URL %q for status code. %d", server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to make the Get call", succeeded)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t%s\tShould receive a %d status code.", succeeded, statusCode)
			} else {
				t.Errorf("\t%s\tShould receive a %d status code got %d", failed, statusCode, resp.StatusCode)
			}

			body := ExpectedResponseBody{}
			err = json.NewDecoder(resp.Body).Decode(&body)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to decode the response : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to decode the response", succeeded)

			if body.Status == "good" {
				t.Logf("\t%s\tShould receive a \"good\" status.", succeeded)
			} else {
				t.Errorf("\t%s\tShould receive a \"good\" status, got %s", failed, body.Status)
			}
		}
	}
}
