package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	// Create a new instance of the application struct. For now, this just
	// contains a couple of mock loggers (which discard anything written to
	// them).
	app := &application{
		errorLog: log.New(ioutil.Discard, "", 0),
		infoLog:  log.New(ioutil.Discard, "", 0),
	}

	// Use the httptest.NewTLSServer() function to create a new test server,
	// passing in the value returned by our app.routes() method as the handler
	// for the server, This starts up a HTTPS server which listens on a
	// randomly-chosen port of your local machine for the duration of the
	// test. Defer a call to ts.Close() to shutdown the server when the
	// test finishes.
	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()

	rs, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	// Examine the http.Response to check that the status code
	// written by the ping handler was 200.
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	// Check the response body written by the ping handler equals "OK".
	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}
