package sync

import (
	"fmt"
	"io"

	"github.com/sowens81/primal-audio-manager/pkg/discogs/models"
)

type Service struct {
	client   CollectionClient
	username string
	out      io.Writer
}

func NewService(client CollectionClient, username string, out io.Writer) *Service {
	return &Service{
		client:   client,
		username: username,
		out:      out,
	}
}

func (s *Service) SyncCollection() error {
	fmt.Fprintln(s.out, "Getting All Collections folder:")

	collectionReleases, err := s.client.GetItemsByFolder(s.username, 0, models.NewPageSettings())
	if err != nil {
		return err
	}

	fmt.Printf("Number of Pages: %d\n", collectionReleases.Pagination.Pages)

	for page := 1; page <= collectionReleases.Pagination.Pages; page++ {
		pageOpts := models.NewPageSettings()
		pageOpts.Page = page
		fmt.Fprintf(s.out, "Getting page %d of All Collections folder:\n", page)
		fmt.Fprintf(s.out, "------------------------------------------\n")
		fmt.Fprintf(s.out, "------------------------------------------\n")
		fmt.Fprintf(s.out, "------------------------------------------\n")

		err := s.GetItemsByFolder(0, pageOpts)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) AddFolder(name string) error {
	resp, err := s.client.AddFolder(s.username, name)
	if err != nil {
		return err
	}

	fmt.Fprintf(s.out, "Added folder: [%d] %s (%d items)\n", resp.ID, resp.Name, resp.Count)
	return nil
}

func (s *Service) GetFolderByID(folderID int) error {
	resp, err := s.client.GetFolderById(s.username, folderID)
	if err != nil {
		return err
	}

	fmt.Fprintf(s.out, "Retrieved folder: [%d] %s (%d items)\n", resp.ID, resp.Name, resp.Count)
	return nil
}

func (s *Service) GetItemsByFolder(folderID int, pageOpts models.PageSettings) error {
	resp, err := s.client.GetItemsByFolder(s.username, folderID, pageOpts)
	if err != nil {
		return err
	}

	fmt.Fprintf(s.out, "Releases in folder [%d]:\n", folderID)

	for _, release := range resp.Releases {

		artist := "Unknown"
		if len(release.BasicInfo.Artists) > 0 {
			artist = release.BasicInfo.Artists[0].Name
		}

		genre := "Unknown"
		if len(release.BasicInfo.Genres) > 0 {
			genre = release.BasicInfo.Genres[0]
		}

		fmt.Fprintf(
			s.out,
			"- %s by %s (ID: %d, Year: %d, Genre: %s)\n",
			release.BasicInfo.Title,
			artist,
			release.ID,
			release.BasicInfo.Year,
			genre,
		)
	}

	return nil
}
