package routes

import (
	"net/http/httptest"
	"testing"
	"net/http"
	"github.com/gorilla/mux"
	
)

func TestGetTodo(t *testing.T) {
	req, err := http.NewRequest("GET", "/todo/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	Routes(router)
	router.ServeHTTP(rr, req)

  // Check the status code is what we expect.
  if status := rr.Code; status != http.StatusOK {
      t.Errorf("handler returned wrong status code: got %v want %v",
          status, http.StatusOK)
	}
	
	expected := `[{"ID":46,"CreatedAt":"2022-02-25T11:17:56.93215+09:00","UpdatedAt":"2022-02-25T11:17:56.93215+09:00","DeletedAt":null,"title":"test-title-console-check","description":"test-description-console-check","conditions":true}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
