package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var preOrganizedFolders OrganizedFolders = OrganizedFolders{OrganizedFolder: []OrganizedFolder{
	{Name: "Images", FileTypes: []string{
		"jpg", "jpeg", "png", "gif", "webp", "svg", "bmp", "tiff", "tif", "heic", "heif", "ico", "raw",
	}},
	{Name: "Videos", FileTypes: []string{
		"mp4", "mov", "avi", "mkv", "webm", "wmv", "flv", "mpeg", "mpg", "m4v", "3gp",
	}},
	{Name: "Audio", FileTypes: []string{
		"mp3", "wav", "flac", "m4a", "ogg", "aac", "wma", "aiff", "alac", "mid", "midi",
	}},
	{Name: "Documents", FileTypes: []string{
		"pdf", "doc", "docx", "odt", "rtf", "tex", "txt", "md", "csv",
	}},
	{Name: "Spreadsheets", FileTypes: []string{
		"xls", "xlsx", "ods", "numbers",
	}},
	{Name: "Presentations", FileTypes: []string{
		"ppt", "pptx", "odp", "key",
	}},
	{Name: "Archives", FileTypes: []string{
		"zip", "rar", "7z", "tar", "gz", "bz2", "xz", "iso",
	}},
	{Name: "Apps", FileTypes: []string{
		"app", "exe", "dmg", "apk", "ipa",
	}},
	{Name: "Installers", FileTypes: []string{
		"msi", "pkg", "deb", "rpm", "appimage",
	}},
	{Name: "Code", FileTypes: []string{
		"go", "js", "ts", "jsx", "tsx", "html", "css", "scss", "json", "xml", "yaml", "yml",
		"py", "java", "c", "cpp", "h", "hpp", "cs", "php", "rb", "rs", "kt", "swift", "sh",
	}},
	{Name: "Fonts", FileTypes: []string{
		"ttf", "otf", "woff", "woff2", "eot",
	}},
	{Name: "Ebooks", FileTypes: []string{
		"epub", "mobi", "azw", "azw3", "fb2",
	}},
	{Name: "Design", FileTypes: []string{
		"psd", "ai", "xd", "fig", "sketch", "indd",
	}},
}}

type OrganizedFolders struct {
	OrganizedFolder []OrganizedFolder `json:"organizedFolders"`
}

type OrganizedFolder struct {
	Name      string   `json:"name"`
	FileTypes []string `json:"fileTypes"`
}

func loadConfig() (OrganizedFolders, error) {

	var folders OrganizedFolders

	file, err := os.ReadFile("organizedFolders.json")
	if err != nil {
		folders = preOrganizedFolders
		err := saveConfig(folders)
		if err != nil {
			return folders, err
		}
		return folders, nil
	}

	err = json.Unmarshal(file, &folders)
	if err != nil {
		folders = preOrganizedFolders
		err := saveConfig(folders)
		if err != nil {
			return folders, err
		}
		return folders, nil
	}

	return folders, nil
}

func saveConfig(folders OrganizedFolders) error {
	out, err := json.MarshalIndent(folders, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("organizedFolders.json", out, 0644)
	if err != nil {
		return err
	}

	return nil
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

func isAppDirectory(entry os.DirEntry) bool {
	return entry.IsDir() && strings.EqualFold(filepath.Ext(entry.Name()), ".app")
}

// TODO: Error handling. As an example duplicates in a OrganizedFolder
func main() {

	dryRunFlag := flag.Bool("dry-run", false, "dry run")

	dirPathFlag := flag.String("d", "", "set the directory path")

	recursiveFlag := flag.Bool("r", false, "recursive directory scanning")

	flag.Parse()

	dirPath := *dirPathFlag

	fmt.Println(*recursiveFlag)

	if dirPath == "" {
		fmt.Println("Enter the path to the directory you want to organize:")
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)

		if !scanner.Scan() {
			fmt.Println("No input provided")
			return
		}

		dirPath = scanner.Text()
	}

	organizedFolders, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading organized folders config (organizedFolders.json):", err)
		return
	}

	organizedFolder := organizedFolders.OrganizedFolder

	dir, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	var dirFolders []string

	for _, file := range dir {
		if !file.IsDir() || isAppDirectory(file) {
			continue
		}

		founded := false

		for _, organizedFolder := range organizedFolder {
			if organizedFolder.Name == file.Name() {
				founded = true
			}
		}

		if founded {
			continue
		}

		dirFolders = append(dirFolders, file.Name())
	}

	for i := 0; i < len(dirFolders); i++ {
		dirFolder := dirFolders[i]

		readDir, err := os.ReadDir(filepath.Join(dirPath, dirFolder))
		if err != nil {
			continue
		}

		for _, file := range readDir {
			if !file.IsDir() || isAppDirectory(file) {
				continue
			}

			nestedFolder := filepath.Join(dirFolder, file.Name())
			dirFolders = append(dirFolders, nestedFolder)

			fmt.Println(nestedFolder)
		}
	}

	dirFolders = append(dirFolders, "")

	fmt.Println(dirFolders)

	if *dryRunFlag {
		fmt.Println("Dry run mode enabled. No files will be moved.")

		if *recursiveFlag {
			for _, dirFolder := range dirFolders {
				readDir, err := os.ReadDir(filepath.Join(dirPath, dirFolder))
				if err != nil {
					return
				}
				for _, file := range readDir {
					if file.IsDir() {
						continue
					}
					fmt.Printf("File %v/%v -> %v\n", dirFolder, file.Name(), match(organizedFolder, getFileType(file.Name())))
				}
			}
		} else {
			for _, file := range dir {
				if file.IsDir() {
					continue
				}

				matchStr := match(organizedFolder, getFileType(file.Name()))

				if matchStr == "NONE" {
					fmt.Printf("File %v would stay. Unknown file type\n", file.Name())
					continue
				}

				fmt.Printf("File %v -> %v\n", file.Name(), matchStr)
			}
		}

		return
	}

	//fmt.Printf("%v entities found in %v\n", len(dir)-len(dirFolders), dirPath)
	fmt.Println("Starting to organize...")

	for _, organizedF := range organizedFolder {
		folderPath := filepath.Join(dirPath, organizedF.Name)

		err := os.Mkdir(folderPath, 0755)
		if err != nil {
			// Folder already exists
		}
	}

	if *recursiveFlag {
		for _, dirFolder := range dirFolders {
			readDir, err := os.ReadDir(filepath.Join(dirPath, dirFolder))
			if err != nil {
				return
			}
			for _, file := range readDir {
				if file.IsDir() {
					continue
				}

				matchStr := match(organizedFolder, getFileType(file.Name()))

				if matchStr == "NONE" {
					continue
				}

				err := os.Rename(filepath.Join(dirPath, dirFolder, file.Name()), filepath.Join(dirPath, matchStr, file.Name()))
				if err != nil {
					fmt.Printf("Error moving %v: %v\n", file.Name(), err)
					continue
				}

				fmt.Println("Moved " + file.Name() + " to " + matchStr)
			}
		}
	} else {
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
				fmt.Printf("Error moving %v: %v\n", file.Name(), err)
				continue
			}

			fmt.Println("Moved " + file.Name() + " to " + matchStr)
		}
	}

	fmt.Println("Done!")
}
