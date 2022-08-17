package chatwork

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostMessage(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/rooms/123/messages", func(w http.ResponseWriter, r *http.Request) {
		builder := new(strings.Builder)
		io.Copy(builder, r.Body)
		got := builder.String()
		want := "body=Hello+Chatwork%21&self_unread=0"
		if got != want {
			t.Errorf("want: %v, got: %v", want, got)
		}

		got = r.Header.Get("Content-Type")
		want = "application/x-www-form-urlencoded"
		if got != want {
			t.Errorf("want: %v, got: %v", want, got)
		}

		got = r.Header.Get("X-ChatWorkToken")
		want = "foo"
		if got != want {
			t.Errorf("want: %v, got: %v", want, got)
		}

		var body PostMessageResponseBody
		body.MessageID = "1234"
		json.NewEncoder(w).Encode(&body)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	client, err := NewClient(server.URL, "foo", nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	form := &PostMessageForm{
		Body:       "Hello Chatwork!",
		SelfUnread: false,
	}
	resp, err := client.PostMessage(ctx, "123", form)
	if err != nil {
		t.Fatal(err)
	}

	want := "1234"
	if resp.MessageID != want {
		t.Errorf("want: %v, got: %v", want, resp.MessageID)
	}
}
