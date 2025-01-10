package main

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	url := "http://www.google.com"
	invalidURL := "http://www.google123.com"
	statusCode := 200
	notFoundCode := 404

	t.Log("Download content from remote URL")
	{
		t.Logf("\tTest 0:\tWhen checking valid URL %q for status code. %d", url, statusCode)
		{
			resp, err := http.Get(url)
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
		}

		t.Logf("\tTest 1:\tWhen checking invalid URL %q for status code. %d", invalidURL, notFoundCode)
		{
			resp, err := http.Get(invalidURL)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to make the Get call", succeeded)

			defer resp.Body.Close()

			if resp.StatusCode == notFoundCode {
				t.Logf("\t%s\tShould receive a %d status code.", succeeded, notFoundCode)
			} else {
				t.Errorf("\t%s\tShould receive a %d status code, got %d", failed, notFoundCode, resp.StatusCode)
			}
		}
	}
}
