package discogs

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestParseError_WithMessage(t *testing.T) {
	resp := &http.Response{
		StatusCode: 404,
		Body:       io.NopCloser(strings.NewReader(`{"message":"Not found"}`)),
	}

	err := parseError(resp)

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected APIError type")
	}

	if apiErr.StatusCode != 404 {
		t.Fatalf("expected status 404, got %d", apiErr.StatusCode)
	}

	if apiErr.Message != "Not found" {
		t.Fatalf("unexpected message: %s", apiErr.Message)
	}
}

func TestParseError_WithoutMessage(t *testing.T) {
	resp := &http.Response{
		StatusCode: 500,
		Body:       io.NopCloser(strings.NewReader(`{}`)),
	}

	err := parseError(resp)
	if err == nil {
		t.Fatal("expected error")
	}

	expected := "discogs error: status=500 body={}"
	if err.Error() != expected {
		t.Fatalf("unexpected error message:\nexpected: %s\ngot:      %s", expected, err.Error())
	}
}
