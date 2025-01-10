package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInternalEndpoint(t *testing.T) {
	t.Log("Testing internal endpoint")
	{
		r := httptest.NewRequest("GET", "/internal", nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r) // this makes the call to the internal endpoint
		t.Logf("\tTest 0:\tWhen checking the internal endpoint")
		{

			if w.Code != http.StatusNotFound {
				t.Fatalf("\t%s\tShould receive a 404 status code.", failed)
			}
			t.Logf("\t%s\tShould receive a 404 status code, got %d", succeeded, w.Code)

			// more assertions here
			// like checking w.Body
		}
	}
}
