package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/sowens81/primal-audio-manager/pkg/discogs/models"
)

type mockAPIClient struct {
	NewRequestFunc func(method, path string) (*http.Request, error)
	ExecuteFunc    func(req *http.Request, v interface{}) error
}

func (m *mockAPIClient) NewRequest(method, path string) (*http.Request, error) {
	return m.NewRequestFunc(method, path)
}

func (m *mockAPIClient) Execute(req *http.Request, v interface{}) error {
	return m.ExecuteFunc(req, v)
}

// Test GetFolders method of collection service
func TestCollectionService_GetFolders(t *testing.T) {
	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			expected := "/users/testuser/collection/folders"
			if path != expected {
				t.Fatalf("expected path %s, got %s", expected, path)
			}
			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			result, ok := v.(*models.CollectionFolders)
			if !ok {
				t.Fatal("type assertion failed")
			}
			result.Folders = []models.CollectionFolder{
				{
					ID:    0,
					Name:  "All",
					Count: 23,
				},
			}
			return nil
		},
	}

	service := NewCollectionService(mock)

	resp, err := service.GetFolders("testuser")
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Folders) != 1 {
		t.Fatalf("expected 1 folder, got %d", len(resp.Folders))
	}

	if resp.Folders[0].Name != "All" {
		t.Fatalf("unexpected folder name")
	}
}

func TestCollectionService_GetFolders_Error(t *testing.T) {
	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			return assertError()
		},
	}

	service := NewCollectionService(mock)

	_, err := service.GetFolders("testuser")
	if err == nil {
		t.Fatal("expected error")
	}
}

// Test AddFolder method of collection service
func TestCollectionService_AddFolder(t *testing.T) {
	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			expectedPath := "/users/testuser/collection/folders"

			if method != "POST" {
				t.Fatalf("expected POST method, got %s", method)
			}

			if path != expectedPath {
				t.Fatalf("expected path %s, got %s", expectedPath, path)
			}

			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			// ✅ Read body
			bodyBytes, err := io.ReadAll(req.Body)
			if err != nil {
				t.Fatal(err)
			}

			// ✅ Validate JSON structure
			var payload map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &payload); err != nil {
				t.Fatalf("invalid JSON body: %v", err)
			}

			// ✅ Validate "name" field
			name, ok := payload["name"].(string)
			if !ok {
				t.Fatal("expected 'name' field in body")
			}

			if name != "NewFolder" {
				t.Fatalf("expected name NewFolder, got %s", name)
			}

			// ✅ Validate Content-Type
			if req.Header.Get("Content-Type") != "application/json" {
				t.Fatalf("expected Content-Type application/json")
			}

			// ✅ Mock response
			result, ok := v.(*models.CollectionFolder)
			if !ok {
				t.Fatal("expected *models.CollectionFolder")
			}

			*result = models.CollectionFolder{
				ID:    1,
				Name:  "NewFolder",
				Count: 0,
			}

			return nil
		},
	}

	service := NewCollectionService(mock)

	resp, err := service.AddFolder("testuser", "NewFolder")
	if err != nil {
		t.Fatal(err)
	}

	if resp.ID != 1 {
		t.Fatalf("expected folder ID 1, got %d", resp.ID)
	}

	if resp.Name != "NewFolder" {
		t.Fatalf("unexpected folder name")
	}

	if resp.Count != 0 {
		t.Fatalf("expected folder count 0, got %d", resp.Count)
	}
}

func TestCollectionService_AddFolder_Error(t *testing.T) {
	expectedErr := errors.New("execute failed")

	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			return expectedErr
		},
	}

	service := NewCollectionService(mock)

	_, err := service.AddFolder("testuser", "NewFolder")

	if err == nil {
		t.Fatal("expected error")
	}

	if err != expectedErr {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}

// Test GetFolderById method of collection service
func TestCollectionService_GetFolderById(t *testing.T) {
	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			expected := "/users/testuser/collection/folders/1"
			if path != expected {
				t.Fatalf("expected path %s, got %s", expected, path)
			}
			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			result, ok := v.(*models.CollectionFolder)
			if !ok {
				t.Fatal("type assertion failed")
			}
			*result = models.CollectionFolder{
				ID:    1,
				Name:  "All",
				Count: 23,
			}
			return nil
		},
	}

	service := NewCollectionService(mock)

	resp, err := service.GetFolderById("testuser", 1)
	if err != nil {
		t.Fatal(err)
	}

	if resp.ID != 1 {
		t.Fatalf("expected folder ID 1, got %d", resp.ID)
	}

	if resp.Name != "All" {
		t.Fatalf("unexpected folder name")
	}
}

func TestCollectionService_GetFolderById_Error(t *testing.T) {
	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			return assertError()
		},
	}

	service := NewCollectionService(mock)

	_, err := service.GetFolderById("testuser", 1)
	if err == nil {
		t.Fatal("expected error")
	}
}

// Test GetItemsByReleaseId method of collection service
func TestCollectionService_GetItemsByReleaseId(t *testing.T) {
	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			expected := "/users/testuser/collection/releases/2464521"
			if path != expected {
				t.Fatalf("expected path %s, got %s", expected, path)
			}
			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			result, ok := v.(*models.CollectionReleases)
			if !ok {
				t.Fatal("type assertion failed")
			}
			*result = models.CollectionReleases{
				Releases: []models.CollectionItem{
					{
						ID:         2464521,
						InstanceID: 1,
						FolderID:   1,
						Rating:     0,
						BasicInfo: models.BasicInformation{
							Title: "Information Chase",
							Year:  2006,
							Artists: []models.Artist{
								{Name: "Bit Shifter"},
							},
						},
					},
				},
			}
			return nil
		},
	}

	service := NewCollectionService(mock)

	resp, err := service.GetItemsByReleaseId("testuser", 2464521)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Releases[0].ID != 2464521 {
		t.Fatalf("expected release ID 2464521, got %d", resp.Releases[0].ID)
	}

	if resp.Releases[0].BasicInfo.Title != "Information Chase" {
		t.Fatalf("unexpected title")
	}
}

func TestCollectionService_GetItemsByReleaseId_Error(t *testing.T) {
	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			return assertError()
		},
	}

	service := NewCollectionService(mock)

	_, err := service.GetItemsByReleaseId("testuser", 2464521)
	if err == nil {
		t.Fatal("expected error")
	}
}

// Test GetItemsByFolder method of collection service
func TestCollectionService_GetItemsByFolder(t *testing.T) {
	pageOpts := models.NewPageSettings(
		models.WithPage(2),
		models.WithPerPage(50),
	)

	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			expected := "/users/testuser/collection/folders/1/releases?page=2&per_page=50"

			if method != "GET" {
				t.Fatalf("expected GET method")
			}

			if path != expected {
				t.Fatalf("expected %s, got %s", expected, path)
			}

			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			result, ok := v.(*models.CollectionReleases)
			if !ok {
				t.Fatal("expected *models.CollectionReleases")
			}

			result.Releases = []models.CollectionItem{
				{
					ID:         2464521,
					InstanceID: 1,
					FolderID:   1,
					Rating:     0,
					BasicInfo: models.BasicInformation{
						Title: "Information Chase",
						Year:  2006,
						Artists: []models.Artist{
							{Name: "Bit Shifter"},
						},
					},
				},
			}

			return nil
		},
	}

	service := NewCollectionService(mock)

	resp, err := service.GetItemsByFolder("testuser", 1, pageOpts)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Releases) != 1 {
		t.Fatalf("expected 1 release")
	}

	if resp.Releases[0].BasicInfo.Title != "Information Chase" {
		t.Fatalf("unexpected title")
	}
}

func TestCollectionService_GetItemsByFolder_Error(t *testing.T) {
	pageOpts := models.NewPageSettings() // defaults

	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			return assertError()
		},
	}

	service := NewCollectionService(mock)

	_, err := service.GetItemsByFolder("testuser", 1, pageOpts)
	if err == nil {
		t.Fatal("expected error")
	}
}

// Test GetCustomFields method of collection service
func TestCollectionService_GetCustomFields(t *testing.T) {
	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			expected := "/users/testuser/collection/fields"
			if path != expected {
				t.Fatalf("expected path %s, got %s", expected, path)
			}
			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			result, ok := v.(*models.CollectionFields)
			if !ok {
				t.Fatal("type assertion failed")
			}
			result.Fields = []models.CollectionField{
				{
					ID:       1,
					Name:     "Condition",
					Type:     "dropdown",
					Position: 1,
					Public:   true,
					Options:  []string{"Mint", "Near Mint", "Very Good Plus"},
				},
			}
			return nil
		},
	}

	service := NewCollectionService(mock)

	resp, err := service.GetCustomFields("testuser")
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Fields) != 1 {
		t.Fatalf("expected 1 field, got %d", len(resp.Fields))
	}

	if resp.Fields[0].Name != "Condition" {
		t.Fatalf("unexpected field name")
	}
}

func TestCollectionService_GetCustomFields_Error(t *testing.T) {
	mock := &mockAPIClient{
		NewRequestFunc: func(method, path string) (*http.Request, error) {
			return http.NewRequest(method, "http://example.com"+path, nil)
		},
		ExecuteFunc: func(req *http.Request, v interface{}) error {
			return assertError()
		},
	}

	service := NewCollectionService(mock)

	_, err := service.GetCustomFields("testuser")
	if err == nil {
		t.Fatal("expected error")
	}
}

// Helper function to return a consistent error for testing
func assertError() error {
	return &mockError{}
}

type mockError struct{}

func (e *mockError) Error() string {
	return "mock error"
}
