package services

import "net/http"

type APIClient interface {
	NewRequest(method, path string) (*http.Request, error)
	Execute(req *http.Request, v interface{}) error
}
