package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	router := newRouter()
	mockServer := httptest.NewServer(router)

	res, err := http.Get(mockServer.URL + "/welcome")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("status code should be 200, got %d", res.StatusCode)
	}

	checkReponseBody(res, t, "Welcome to Birdpedia!")
}

func TestRouterForNonExistingRoute(t *testing.T) {
	router := newRouter()
	mockServer := httptest.NewServer(router)

	res, err := http.Post(mockServer.URL+"/welcome", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("status code should be 405, got %d", res.StatusCode)
	}

	checkReponseBody(res, t, "")
}

func TestStaticFileServer(t *testing.T) {
	router := newRouter()
	mockServer := httptest.NewServer(router)

	res, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("status code should be 200, got %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if expectedContentType != contentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}

func checkReponseBody(res *http.Response, t *testing.T, expected string) {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	actual := string(body)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'", expected, actual)
	}
}
