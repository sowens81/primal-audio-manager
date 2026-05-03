package discogs

import "testing"

func TestNewClient(t *testing.T) {
	token := "test-token"

	client := NewClient(token)

	if client.Token != token {
		t.Fatalf("expected token %s, got %s", token, client.Token)
	}

	if client.BaseURL == "" {
		t.Fatal("expected BaseURL to be set")
	}

	if client.HTTPClient == nil {
		t.Fatal("expected HTTPClient to be set")
	}

	if client.Collection == nil {
		t.Fatal("expected Collection service to be initialized")
	}
}
