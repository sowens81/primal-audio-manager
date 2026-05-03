package sync

import (
	"fmt"

	"github.com/sowens81/primal-audio-manager/pkg/discogs"
)

func SyncCollection(client *discogs.Client, username string) error {
	resp, err := client.Collection.GetFolders(username)
	if err != nil {
		return err
	}

	fmt.Println("Collection folders:")
	for _, folder := range resp.Folders {
		err = GetCollectionFolderById(client, username, folder.ID)
		if err != nil {
			return err
		}

		err = GetCollectionFolderReleases(client, username, folder.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func AddCollectionFolder(client *discogs.Client, username string, folderName string) error {
	resp, err := client.Collection.AddFolder(username, folderName)
	if err != nil {
		fmt.Printf("Error adding folder: %v\n", err)
		return err
	}

	fmt.Printf("Added folder: [%d] %s (%d items)\n", resp.ID, resp.Name, resp.Count)
	return nil
}

func GetCollectionFolderById(client *discogs.Client, username string, folderId int) error {
	resp, err := client.Collection.GetFolderById(username, folderId)
	if err != nil {
		fmt.Printf("Error getting folder: %v\n", err)
		return err
	}

	fmt.Printf("Retrieved folder: [%d] %s (%d items)\n", resp.ID, resp.Name, resp.Count)
	return nil
}

func GetCollectionFolderReleases(client *discogs.Client, username string, folderId int) error {
	resp, err := client.Collection.GetFolderReleases(username, folderId)
	if err != nil {
		fmt.Printf("Error getting folder releases: %v\n", err)
		return err
	}

	fmt.Printf("Releases in folder [%d]:\n", folderId)
	for _, release := range resp.Releases {
		artist := release.BasicInfo.Artists[0].Name
		fmt.Printf("- %s by %s (ID: %d, Year: %d, Genre: %s)\n", release.BasicInfo.Title, artist, release.ID, release.BasicInfo.Year, release.BasicInfo.Genres[0])
	}

	return nil
}
