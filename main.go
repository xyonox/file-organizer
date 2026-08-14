package main

import (
	"fmt"
	"os"
	"strings"
)

type Folder struct {
	Name      string
	FileTypes []string
}

func match(folders []Folder, filetype string) string {
	for _, f := range folders {
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

	var folders = make([]Folder, 30)
	folders[0] = Folder{"Images", []string{"jpg", "png"}}
	folders[1] = Folder{"Videos", []string{"mp4"}}
	folders[2] = Folder{"Documents", []string{"pdf"}}

	fmt.Println(folders)

	dir, err := os.ReadDir("/Users/cfa/Downloads")
	if err != nil {
		return
	}

	fmt.Println(dir)

	for _, file := range dir {
		fmt.Println(file.Name())
		fmt.Println(getFileType(file.Name()))
		fmt.Println(match(folders, getFileType(file.Name())))

	}
}
