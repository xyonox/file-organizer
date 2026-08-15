package main

import (
	"bufio"
	"flag"
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

	dirPathFlag := flag.String("d", "", "set the directory path")

	flag.Parse()

	dirPath := *dirPathFlag

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

	organizedFolder := []OrganizedFolder{
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
	}

	dir, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
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
			fmt.Printf("Error moving %v: %v\n", file.Name(), err)
			continue
		}

		fmt.Println("Moved " + file.Name() + " to " + matchStr)
	}

	fmt.Println("Done!")
}
