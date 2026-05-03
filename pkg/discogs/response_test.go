package discogs

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type mockHTTPClient struct {
	doFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.doFunc(req)
}

type mockError struct{}

func (e *mockError) Error() string {
	return "mock error"
}

func TestExecute_Success(t *testing.T) {
	mock := &mockHTTPClient{
		doFunc: func(req *http.Request) (*http.Response, error) {
			body := `{"name":"test"}`
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(body)),
				Header:     make(http.Header),
			}, nil
		},
	}

	client := NewClient("token")
	client.HTTPClient = mock

	req, _ := client.NewRequest("GET", "/test")

	var result map[string]interface{}
	err := client.Execute(req, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result["name"] != "test" {
		t.Fatalf("unexpected result")
	}
}

func TestExecute_HTTPError(t *testing.T) {
	mock := &mockHTTPClient{
		doFunc: func(req *http.Request) (*http.Response, error) {
			body := `{"message":"Not found"}`
			return &http.Response{
				StatusCode: 404,
				Body:       io.NopCloser(strings.NewReader(body)),
				Header:     make(http.Header),
			}, nil
		},
	}

	client := NewClient("token")
	client.HTTPClient = mock

	req, _ := client.NewRequest("GET", "/test")

	var result map[string]interface{}
	err := client.Execute(req, &result)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestExecute_NetworkError(t *testing.T) {
	mock := &mockHTTPClient{
		doFunc: func(req *http.Request) (*http.Response, error) {
			return nil, &mockError{}
		},
	}

	client := NewClient("token")
	client.HTTPClient = mock

	req, _ := client.NewRequest("GET", "/test")

	var result map[string]interface{}
	err := client.Execute(req, &result)

	if err == nil {
		t.Fatal("expected error")
	}
}
