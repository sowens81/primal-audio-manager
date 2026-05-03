package discogs

import "testing"

func TestNewRequest(t *testing.T) {
	client := NewClient("test-token")

	req, err := client.NewRequest("GET", "/users/test/collection/folders")
	if err != nil {
		t.Fatal(err)
	}

	if req.Method != "GET" {
		t.Fatalf("expected GET method, got %s", req.Method)
	}

	if req.URL.Path != "/users/test/collection/folders" {
		t.Fatalf("unexpected path: %s", req.URL.Path)
	}

	auth := req.Header.Get("Authorization")
	if auth == "" {
		t.Fatal("expected Authorization header")
	}

	if req.Header.Get("User-Agent") == "" {
		t.Fatal("expected User-Agent header")
	}
}
