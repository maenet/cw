package chatwork

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	_, err := NewClient("https://example.com", "abc", nil)
	if err != nil {
		t.Errorf("failed to create client: %v", err)
	}
}
