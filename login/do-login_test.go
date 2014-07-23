package login

import (
	_ "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func expectError(t *testing.T, URL string) {
	_, err := Do([]string{URL, "admin", "addmin"})
	if err == nil {
		t.Fail()
	}

}

func TestClosed(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	ts.Close()

	expectError(t, ts.URL[7:])
}

func TestUnauthorized(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)

	}))
	defer ts.Close()

	expectError(t, ts.URL[7:])
}

func TestNumber(t *testing.T) {
	_, err := Do([]string{"someurl.com:1234", "username", "password", "toomany"})
	if err == nil {
		// Expected error
		t.Fail()
	}
	_, err = Do([]string{"someurl.com:1234", "toofew"})
	if err == nil {
		// Expected error
		t.Fail()
	}

}
