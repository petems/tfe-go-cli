package cmd

import (
  "fmt"
  "testing"

  "net/http"
  "net/http/httptest"
)

type TestHttpMock struct {
  server *httptest.Server
}

type urlValid struct {
  url                string
  validUrl           bool 
}

type urlTestResponse struct {
  url        string
  expectedResponse   int
}

var urlTestResponses = []urlTestResponse {
  {"meta_200.txt", 200},
  {"meta_404.txt", 404},
  {"anything_else.txt", 501},
}

const dynamicTestURL = "%s/meta_%d.txt"

func TestGetResponseFromURL(t *testing.T) {
  for _, tt := range urlTestResponses {
    testHttpMock := setUpMockHttpServer()

    defer testHttpMock.server.Close()

    actual, err := CheckIf200FromURL(fmt.Sprintf(dynamicTestURL, testHttpMock.server.URL, tt.expectedResponse))
    if actual.StatusCode != tt.expectedResponse {
      t.Errorf("CheckIf200FromURL(%s): expected %d, actual %d, error if present was: %s", tt.url, tt.expectedResponse, actual.StatusCode, err)
    }
  }
}

func setUpMockHttpServer() *TestHttpMock {
  Server := httptest.NewServer(
    http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

      w.Header().Set("Content-Type", "text/plain")
      if r.URL.Path == "/meta_200.txt" {
        w.WriteHeader(http.StatusOK)
      } else if r.URL.Path == "/meta_404.txt" {
        w.WriteHeader(http.StatusNotFound)
      } else {
        w.WriteHeader(http.StatusNotImplemented)
      }
    }),
  )

  return &TestHttpMock{
    server: Server,
  }
}