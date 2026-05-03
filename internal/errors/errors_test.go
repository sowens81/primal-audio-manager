package errors

import (
	"bytes"
	"errors"
	"testing"
)

// mock DiscogsError
type mockDiscogsError struct{}

func (m *mockDiscogsError) Error() string        { return "mock error" }
func (m *mockDiscogsError) Status() int          { return 400 }
func (m *mockDiscogsError) MessageText() string  { return "bad request" }
func (m *mockDiscogsError) Details() interface{} { return map[string]string{"field": "invalid"} }

func TestHandleError_DiscogsError(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	exitCalled := false
	exitFunc = func(code int) {
		exitCalled = true
		if code != 1 {
			t.Fatalf("expected exit code 1, got %d", code)
		}
	}

	err := &mockDiscogsError{}

	HandleError(err)

	output := buf.String()

	if !exitCalled {
		t.Fatal("expected exit to be called")
	}

	if !contains(output, "Discogs API error") {
		t.Fatal("expected Discogs error message")
	}

	if !contains(output, "bad request") {
		t.Fatal("expected message text")
	}

	if !contains(output, "field") {
		t.Fatal("expected details in output")
	}
}

func TestHandleError_GenericError(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	exitCalled := false
	exitFunc = func(code int) {
		exitCalled = true
	}

	err := errors.New("something went wrong")

	HandleError(err)

	output := buf.String()

	if !exitCalled {
		t.Fatal("expected exit to be called")
	}

	if !contains(output, "something went wrong") {
		t.Fatal("expected generic error output")
	}
}

func TestHandleError_Nil(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	exitCalled := false
	exitFunc = func(code int) {
		exitCalled = true
	}

	HandleError(nil)

	if exitCalled {
		t.Fatal("did not expect exit to be called")
	}

	if buf.Len() != 0 {
		t.Fatal("expected no output")
	}
}

// helper
func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
