package sync

import "github.com/sowens81/primal-audio-manager/pkg/discogs/models"

type CollectionClient interface {
	GetFolders(username string) (*models.CollectionFolders, error)
	GetFolderById(username string, folderID int) (*models.CollectionFolder, error)
	GetFolderReleases(username string, folderID int, pageOpts models.PageSettings) (*models.CollectionReleases, error)
	AddFolder(username string, name string) (*models.CollectionFolder, error)
}
