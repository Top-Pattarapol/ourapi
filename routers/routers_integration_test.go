package routers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllUsers(t *testing.T) {
	e := NewRouter()
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	e.ServeHTTP(resp, req)

	var uu []User
	json.NewDecoder(resp.Body).Decode(&uu)

	if http.StatusOK != resp.Code {
		t.Errorf("expected status code %d, but got %d", http.StatusOK, resp.Code)
	}

	if len(uu) < 1 {
		t.Error("should return at least 1 element")
	}

	if uu[0].Username == "" {
		t.Error("username should not empty")
	}
}