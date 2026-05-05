package sync

import "github.com/sowens81/primal-audio-manager/pkg/discogs/models"

type DiscogsCollectionClient interface {
	GetFolders(username string) (models.CollectionFoldersResponse, error)
	GetFolderById(username string, folderID int) (models.CollectionFolder, error)
	GetItemsByFolder(username string, folderID int, pageOpts models.PageSettings) (models.CollectionReleases, error)
	AddFolder(username string, name string) (models.CollectionFolder, error)
}

type DiscogsReleaseClient interface {
	GetReleaseById(releaseID int) (models.Release, error)
}
