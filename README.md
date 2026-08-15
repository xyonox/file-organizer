# File Organizer

File Organizer is a small command-line application written in Go.
The project was created as a learning project for:

- working with files and directories in Go
- organizing files by their file extension
- using the standard library

## Features

- Organizes files in the `/Users/cfa/Downloads` directory.
- Creates category folders automatically when they do not exist.
- Supports images, videos, audio files, documents, archives, apps and installers.
- Recognizes file extensions regardless of uppercase or lowercase letters.
- Leaves unsupported files and existing directories untouched.

## Requirements

- Go 1.25 or newer

## Running the application

Run the application with:

```bash
go run .
```

The application currently uses `/Users/cfa/Downloads` as its source directory.

## Building a binary

Build the application with:

```bash
go build -o file-organizer .
```

## How it works

The application reads the files in the configured download directory and matches
their file extensions against the `organizedFolder` list in `main.go`.

The following folders are currently supported:

| Folder     | File extensions                                                         |
|------------|-------------------------------------------------------------------------|
| Images     | `jpg`, `jpeg`, `png`, `gif`, `webp`, `svg`                              |
| Videos     | `mp4`, `mov`, `avi`, `mkv`, `webm`                                      |
| Audio      | `mp3`, `wav`, `flac`, `m4a`, `ogg`                                      |
| Documents  | `pdf`, `doc`, `docx`, `odt`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `csv` |
| Archives   | `zip`, `rar`, `7z`, `tar`, `gz`                                         |
| Apps       | `app`, `exe`, `dmg`, `apk`, `ipa`                                       |
| Installers | `msi`, `pkg`, `deb`, `rpm`                                              |

Matching files are moved into the corresponding folder. Files without a
matching extension remain in the source directory.

## Example output

```text
18 entities found in /Users/cfa/Downloads
Starting to organize...
Moved setup.exe to Apps
Moved archive.zip to Archives
Moved report.pdf to Documents
Done!
```

## Project structure

```text
.
├── main.go       # 
├── go.mod        # Go module definition
└── README.md     # Project documentation
```

## Current limitations

- The source directory is currently hard-coded in `main.go`.
- Only files directly inside the source directory are organized; subdirectories
  are not scanned recursively.
- Files with unsupported extensions are not moved.
- Errors while creating folders or moving files are not reported in detail.

## Next steps

- support for more file extensions
- support for more file types