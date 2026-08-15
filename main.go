package main

import (
	"bufio"
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
	filetype = strings.ToLower(strings.TrimPrefix(filetype, "."))

	for _, f := range organizedFolders {
		for _, fileType := range f.FileTypes {
			if strings.ToLower(strings.TrimPrefix(fileType, ".")) == filetype {
				return f.Name
			}
		}
	}
	return "NONE"
}

func getFileType(path string) string {
	extension := filepath.Ext(path)
	if extension == "" {
		return ""
	}

	return strings.ToLower(strings.TrimPrefix(extension, "."))
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		fmt.Println("No input provided")
		return
	}

	dirPath := scanner.Text()

	organizedFolder := []OrganizedFolder{
		{Name: "Images", FileTypes: []string{"jpg", "jpeg", "png", "gif", "webp", "svg"}},
		{Name: "Videos", FileTypes: []string{"mp4", "mov", "avi", "mkv", "webm"}},
		{Name: "Audio", FileTypes: []string{"mp3", "wav", "flac", "m4a", "ogg"}},
		{Name: "Documents", FileTypes: []string{"pdf", "doc", "docx", "odt", "xls", "xlsx", "ppt", "pptx", "txt", "csv"}},
		{Name: "Archives", FileTypes: []string{"zip", "rar", "7z", "tar", "gz"}},
		{Name: "Apps", FileTypes: []string{"app", "exe", "dmg", "apk", "ipa"}},
		{Name: "Installers", FileTypes: []string{"msi", "pkg", "deb", "rpm"}},
	}

	dir, err := os.ReadDir(dirPath)
	if err != nil {
		return
	}

	fmt.Println(dir)

	var dirFolders []string

	for _, file := range dir {
		if file.IsDir() {
			dirFolders = append(dirFolders, file.Name())
		}
	}

	fmt.Printf("%v entities found in %v\n", len(dir)-len(dirFolders), dirPath)
	fmt.Println("Starting to organize...")

	for _, organizedF := range organizedFolder {
		folderPath := filepath.Join(dirPath, organizedF.Name)

		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			// Folder already exists
		}
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		matchStr := match(organizedFolder, getFileType(file.Name()))

		if matchStr == "NONE" {
			continue
		}

		err := os.Rename(filepath.Join(dirPath, file.Name()), filepath.Join(dirPath, matchStr, file.Name()))
		if err != nil {
			return
		}

		fmt.Println("Moved " + file.Name() + " to " + matchStr)
	}

	fmt.Println("Done!")
}
