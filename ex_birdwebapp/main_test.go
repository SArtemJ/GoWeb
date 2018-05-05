package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHandler(t *testing.T)  {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	recoder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler)

	hf.ServeHTTP(recoder, req)

	if status := recoder.Code; status != http.StatusOK {
		t.Errorf("hanfler returned wrong status %v want %v", status, http.StatusOK)
	}

	expected := "Test app"
	actual := recoder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body %v want %v", actual, expected)
	}

}
