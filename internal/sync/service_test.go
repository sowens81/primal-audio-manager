package sync

import (
	"bytes"
	"errors"
	"testing"

	"github.com/sowens81/primal-audio-manager/pkg/discogs/models"
)

//
// Mock Client
//

type mockCollectionClient struct {
	GetFoldersFunc       func(string) (*models.CollectionFolders, error)
	GetFolderByIdFunc    func(string, int) (*models.CollectionFolder, error)
	GetItemsByFolderFunc func(string, int, models.PageSettings) (*models.CollectionReleases, error)
	AddFolderFunc        func(string, string) (*models.CollectionFolder, error)
}

func (m *mockCollectionClient) GetFolders(username string) (*models.CollectionFolders, error) {
	if m.GetFoldersFunc != nil {
		return m.GetFoldersFunc(username)
	}
	return nil, nil
}

func (m *mockCollectionClient) GetFolderById(username string, folderID int) (*models.CollectionFolder, error) {
	if m.GetFolderByIdFunc != nil {
		return m.GetFolderByIdFunc(username, folderID)
	}
	return nil, nil
}

func (m *mockCollectionClient) GetItemsByFolder(username string, folderID int, ps models.PageSettings) (*models.CollectionReleases, error) {
	if m.GetItemsByFolderFunc != nil {
		return m.GetItemsByFolderFunc(username, folderID, ps)
	}
	return nil, nil
}

func (m *mockCollectionClient) AddFolder(username string, name string) (*models.CollectionFolder, error) {
	if m.AddFolderFunc != nil {
		return m.AddFolderFunc(username, name)
	}
	return nil, nil
}

//
// SyncCollection
//

func TestService_SyncCollection(t *testing.T) {
	buf := &bytes.Buffer{}

	callCount := 0

	mock := &mockCollectionClient{
		GetItemsByFolderFunc: func(username string, folderID int, ps models.PageSettings) (*models.CollectionReleases, error) {
			// First call: return pagination info
			if callCount == 0 {
				callCount++
				return &models.CollectionReleases{
					Pagination: models.Pagination{
						Pages: 2,
					},
				}, nil
			}

			// Subsequent calls: return actual releases
			callCount++

			return &models.CollectionReleases{
				Releases: []models.CollectionItem{
					{
						ID: 1,
						BasicInfo: models.BasicInformation{
							Title: "Album",
							Year:  2020,
							Artists: []models.Artist{
								{Name: "Artist"},
							},
							Genres: []string{"Rock"},
						},
					},
				},
			}, nil
		},
	}

	service := NewService(mock, "user", buf)

	err := service.SyncCollection()
	if err != nil {
		t.Fatal(err)
	}

	output := buf.String()

	if !contains(output, "Getting All Collections folder") {
		t.Fatal("expected initial message")
	}

	if !contains(output, "Getting page 1") {
		t.Fatal("expected page 1 output")
	}

	if !contains(output, "Getting page 2") {
		t.Fatal("expected page 2 output")
	}

	if !contains(output, "Album") {
		t.Fatal("expected release output")
	}
}

func TestService_SyncCollection_Error(t *testing.T) {
	mock := &mockCollectionClient{
		GetItemsByFolderFunc: func(username string, folderID int, ps models.PageSettings) (*models.CollectionReleases, error) {
			return nil, errors.New("fail")
		},
	}

	service := NewService(mock, "user", &bytes.Buffer{})

	err := service.SyncCollection()
	if err == nil {
		t.Fatal("expected error")
	}
}

//
// AddFolder
//

func TestService_AddFolder(t *testing.T) {
	buf := &bytes.Buffer{}

	mock := &mockCollectionClient{
		AddFolderFunc: func(username, name string) (*models.CollectionFolder, error) {
			return &models.CollectionFolder{
				ID:    1,
				Name:  name,
				Count: 0,
			}, nil
		},
	}

	service := NewService(mock, "user", buf)

	err := service.AddFolder("Hardcore")
	if err != nil {
		t.Fatal(err)
	}

	if !contains(buf.String(), "Added folder") {
		t.Fatal("expected output")
	}
}

func TestService_AddFolder_Error(t *testing.T) {
	mock := &mockCollectionClient{
		AddFolderFunc: func(username, name string) (*models.CollectionFolder, error) {
			return nil, errors.New("fail")
		},
	}

	service := NewService(mock, "user", &bytes.Buffer{})

	err := service.AddFolder("Hardcore")
	if err == nil {
		t.Fatal("expected error")
	}
}

//
// GetFolderByID
//

func TestService_GetFolderByID(t *testing.T) {
	buf := &bytes.Buffer{}

	mock := &mockCollectionClient{
		GetFolderByIdFunc: func(username string, id int) (*models.CollectionFolder, error) {
			return &models.CollectionFolder{
				ID:    id,
				Name:  "Test",
				Count: 5,
			}, nil
		},
	}

	service := NewService(mock, "user", buf)

	err := service.GetFolderByID(1)
	if err != nil {
		t.Fatal(err)
	}

	if !contains(buf.String(), "Retrieved folder") {
		t.Fatal("expected output")
	}
}

//
// GetItemsByFolder
//

func TestService_GetItemsByFolder(t *testing.T) {
	buf := &bytes.Buffer{}

	mock := &mockCollectionClient{
		GetItemsByFolderFunc: func(username string, id int, ps models.PageSettings) (*models.CollectionReleases, error) {
			return &models.CollectionReleases{
				Releases: []models.CollectionItem{
					{
						ID: 1,
						BasicInfo: models.BasicInformation{
							Title: "Album",
							Year:  2021,
							Artists: []models.Artist{
								{Name: "Artist"},
							},
							Genres: []string{"Electronic"},
						},
					},
				},
			}, nil
		},
	}

	service := NewService(mock, "user", buf)

	err := service.GetItemsByFolder(0, models.NewPageSettings())
	if err != nil {
		t.Fatal(err)
	}

	output := buf.String()

	if !contains(output, "Album") {
		t.Fatal("expected release output")
	}
}

func TestService_GetItemsByFolder_Error(t *testing.T) {
	mock := &mockCollectionClient{
		GetItemsByFolderFunc: func(username string, id int, ps models.PageSettings) (*models.CollectionReleases, error) {
			return nil, errors.New("fail")
		},
	}

	service := NewService(mock, "user", &bytes.Buffer{})

	err := service.GetItemsByFolder(0, models.NewPageSettings())
	if err == nil {
		t.Fatal("expected error")
	}
}

//
// Helper
//

func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
