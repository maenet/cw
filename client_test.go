package chatwork

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("https://example.com", "abc", nil)
	if err != nil {
		t.Errorf("failed to create client: %v", err)
	}
}

func TestRequestHeader(t *testing.T) {
	token := "abc123"

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctwant := "application/x-www-form-urlencoded"
		ctgot := r.Header.Get("Content-Type")
		if ctgot != ctwant {
			t.Errorf("Content-Type = %v, want %v", ctgot, ctwant)
		}

		tokenwant := token
		tokengot := r.Header.Get("X-ChatWorkToken")
		if tokengot != tokenwant {
			t.Errorf("X-ChatWorkToken = %v, want %v", tokengot, tokenwant)
		}
	})

	s := httptest.NewServer(h)
	defer s.Close()

	client, _ := NewClient(s.URL, token, nil)
	ctx := context.Background()
	client.GetAccount(ctx)
}

func TestGetAccount(t *testing.T) {
	called := false
	r := chi.NewRouter()
	r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(200)
	})
	s := httptest.NewServer(r)
	defer s.Close()

	c, _ := NewClient(s.URL, "token", nil)
	ctx := context.Background()
	c.GetAccount(ctx)

	if !called {
		t.Fatal("api was not called")
	}
}
