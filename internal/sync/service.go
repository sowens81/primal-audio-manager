package sync

import (
	"fmt"
	"io"
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
	resp, err := s.client.GetFolders(s.username)
	if err != nil {
		return err
	}

	fmt.Fprintln(s.out, "Collection folders:")

	for _, folder := range resp.Folders {

		if err := s.GetFolderByID(folder.ID); err != nil {
			return err
		}

		if err := s.GetFolderReleases(folder.ID); err != nil {
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

func (s *Service) GetFolderReleases(folderID int) error {
	resp, err := s.client.GetFolderReleases(s.username, folderID)
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
