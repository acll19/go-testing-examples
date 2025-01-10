package main

import (
	"net/http"
	"testing"
)

func TestDownloadSubTests(t *testing.T) {

	tests := []struct {
		name       string
		url        string
		statusCode int
	}{
		{"statusok", "http://www.google.com", http.StatusOK},
		{"statusnotfound", "http://www.google123.com", http.StatusNotFound},
	}

	t.Log("Download content from remote URL")
	{
		for i, tt := range tests {
			tf := func(t *testing.T) {
				t.Parallel() // this will run the subtests in parallel

				t.Logf("\tTest %d:\tWhen checking valid URL %q for status code. %d", i, tt.url, tt.statusCode)
				{
					resp, err := http.Get(tt.url)
					if err != nil {
						t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
					}
					t.Logf("\t%s\tShould be able to make the Get call", succeeded)

					defer resp.Body.Close()

					if resp.StatusCode == tt.statusCode {
						t.Logf("\t%s\tShould receive a %d status code.", succeeded, tt.statusCode)
					} else {
						t.Errorf("\t%s\tShould receive a %d status code got %d", failed, tt.statusCode, resp.StatusCode)
					}
				}
			}

			t.Run(tt.name, tf)
		}
	}
}
