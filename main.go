package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type OrganizedFolder struct {
	Name      string
	FileTypes []string
}

func match(organizedFolders []OrganizedFolder, filetype string) string {
	for _, f := range organizedFolders {
		for _, fileType := range f.FileTypes {
			if fileType == filetype {
				return f.Name
			}
		}
	}
	return "NONE"
}

func getFileType(path string) string {

	split := strings.Split(path, ".")
	fmt.Println(len(split))

	return split[len(split)-1]
}

func main() {

	dirPath := "/Users/cfa/Downloads"

	var organizedFolder = make([]OrganizedFolder, 30)
	organizedFolder[0] = OrganizedFolder{"Images", []string{"jpg", "png"}}
	organizedFolder[1] = OrganizedFolder{"Videos", []string{"mp4"}}
	organizedFolder[2] = OrganizedFolder{"Documents", []string{"pdf"}}

	fmt.Println(organizedFolder)

	dir, err := os.ReadDir(dirPath)
	if err != nil {
		return
	}

	fmt.Println(dir)

	var dirFolders []string

	for _, file := range dir {
		fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() {
			dirFolders = append(dirFolders, file.Name())
		}
	}

	fmt.Println(dirFolders)

	for _, organizedF := range organizedFolder {
		folderPath := filepath.Join(dirPath, organizedF.Name)

		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			return
		}
	}
}
