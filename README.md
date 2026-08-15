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

| Folder        | File extensions                                                                 |
|---------------|---------------------------------------------------------------------------------|
| Images        | `jpg`, `jpeg`, `png`, `gif`, `webp`, `svg`, `bmp`, `tiff`, `tif`, `heic`, `heif`, `ico`, `raw` |
| Videos        | `mp4`, `mov`, `avi`, `mkv`, `webm`, `wmv`, `flv`, `mpeg`, `mpg`, `m4v`, `3gp` |
| Audio         | `mp3`, `wav`, `flac`, `m4a`, `ogg`, `aac`, `wma`, `aiff`, `alac`, `mid`, `midi` |
| Documents     | `pdf`, `doc`, `docx`, `odt`, `rtf`, `tex`, `txt`, `md`, `csv` |
| Spreadsheets  | `xls`, `xlsx`, `ods`, `numbers` |
| Presentations | `ppt`, `pptx`, `odp`, `key` |
| Archives      | `zip`, `rar`, `7z`, `tar`, `gz`, `bz2`, `xz`, `iso` |
| Apps          | `app`, `exe`, `dmg`, `apk`, `ipa` |
| Installers    | `msi`, `pkg`, `deb`, `rpm`, `appimage` |
| Code          | `go`, `js`, `ts`, `jsx`, `tsx`, `html`, `css`, `scss`, `json`, `xml`, `yaml`, `yml`, `py`, `java`, `c`, `cpp`, `h`, `hpp`, `cs`, `php`, `rb`, `rs`, `kt`, `swift`, `sh` |
| Fonts         | `ttf`, `otf`, `woff`, `woff2`, `eot` |
| Ebooks        | `epub`, `mobi`, `azw`, `azw3`, `fb2` |
| Design        | `psd`, `ai`, `xd`, `fig`, `sketch`, `indd` |

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

