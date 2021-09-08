package chatwork

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
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
