package sync

import (
	"fmt"
	"io"
	"strconv"

	"github.com/sowens81/primal-audio-manager/internal/cli"
	"github.com/sowens81/primal-audio-manager/internal/helpers"
	models "github.com/sowens81/primal-audio-manager/internal/models"
	discogsmodels "github.com/sowens81/primal-audio-manager/pkg/discogs/models"
)

type Service struct {
	DiscogsCollectionClient DiscogsCollectionClient
	DiscogsReleaseClient    DiscogsReleaseClient
	username                string
	out                     io.Writer
}

func NewService(discogsCollectionClient DiscogsCollectionClient, discogsReleaseClient DiscogsReleaseClient, username string, out io.Writer) *Service {
	return &Service{
		DiscogsCollectionClient: discogsCollectionClient,
		DiscogsReleaseClient:    discogsReleaseClient,
		username:                username,
		out:                     out,
	}
}

func (s *Service) SyncCollection() error {
	// Get Collection Folders
	fmt.Fprintln(s.out, "Getting Collection Folders:")

	collectionFolders, err := s.DiscogsCollectionClient.GetFolders(s.username)
	if err != nil {
		return err
	}

	// for each folder, print out the folder id, name and number of items in the folder

	for _, folder := range collectionFolders.Folders {
		fmt.Fprintf(s.out, "Folder ID: %d, Name: %s, Items: %d\n",
			folder.ID, folder.Name, folder.Count)
	}

	// user input to get folder ID to sync

	userResponse := cli.PromptRequired("Enter folder Id to Sync: ")

	folderId, err := strconv.Atoi(userResponse)
	if err != nil {
		fmt.Println("invalid number:", err)
		return err
	}

	// get releases in a folder
	userCollection := []models.CollectionItem{}

	discogsCollections, err := s.DiscogsCollectionClient.GetItemsByFolder(s.username, folderId, discogsmodels.NewPageSettings())
	if err != nil {
		return err
	}

	totalPages := discogsCollections.Pagination.Pages
	fmt.Printf("Total Pages: %d\n", totalPages)

	for page := 1; page <= totalPages; page++ {
		pageOpts := discogsmodels.NewPageSettings()
		pageOpts.Page = page
		fmt.Fprintf(s.out, "Getting page %d of All Collections folder:\n", page)
		fmt.Fprintf(s.out, "------------------------------------------\n")

		discogsCollections, err := s.DiscogsCollectionClient.GetItemsByFolder(s.username, folderId, pageOpts)
		if err != nil {
			return err
		}

		for _, item := range discogsCollections.Releases {

			releaseId := item.BasicInfo.ID

			// get release details for each release in the collection
			release, err := s.DiscogsReleaseClient.GetReleaseById(releaseId)
			if err != nil {
				return err
			}

			collectionItem := models.CollectionItem{
				ID:                  1, // need a function to generate a unique ID for the collection item
				CollectionID:        item.ID,
				InstanceID:          item.InstanceID,
				CatalogNumber:       helpers.ConvertDiscogsCatNoToListString(item.BasicInfo.Labels),
				Artists:             helpers.ConvertDiscogsArtistToString(item.BasicInfo.Artists),
				Title:               item.BasicInfo.Title,
				Year:                item.BasicInfo.Year,
				Genre:               append(item.BasicInfo.Genres, item.BasicInfo.Styles...),
				Labels:              helpers.ConvertDiscogsLabelsToListString(item.BasicInfo.Labels),
				Rating:              item.Rating,
				FolderID:            item.FolderID,
				Notes:               item.Notes,
				CoverImage:          item.BasicInfo.CoverImage,
				CoverImageThumbnail: release.Thumb,
				TrackList:           release.Tracklist,
			}

			userCollection = append(userCollection, collectionItem)

		}
	}

	for _, item := range userCollection {
		fmt.Fprintf(s.out, "Title: %s, Artist: %s, Year: %d\n",
			item.Title, item.Artists, item.Year)
	}

	return nil
}
