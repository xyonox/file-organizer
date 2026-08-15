# File Organizer

File Organizer is a small command-line application written in Go.
The project was created as a learning project for:

- working with files and directories in Go
- organizing files by their file extension
- using the standard library

## Features

- Organizes files in a directory supplied with the `-d` flag or entered interactively.
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

When prompted, enter the path to the directory you want to organize. You can also
provide the directory directly with the `-d` flag:

```bash
go run . -d /path/to/directory
```

## Building a binary

Build the application with:

```bash
go build -o file-organizer .
```

## How it works

The application reads the files directly inside the selected directory and matches
their file extensions against the `organizedFolder` list in `main.go`. If no
directory is provided with `-d`, the application prompts for one.

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
18 entities found in /path/to/directory
Starting to organize...
Moved setup.exe to Apps
Moved archive.zip to Archives
Moved report.pdf to Documents
Done!
```

## Project structure

```text
.
├── main.go       # Application source code
├── go.mod        # Go module definition
└── README.md     # Project documentation
```

## Current limitations

- The application only organizes files directly inside the selected directory; it
  does not scan subdirectories recursively.
- Files with unsupported extensions are not moved.
- The application does not provide a dry-run mode or configurable categories.
- Errors while creating category folders are not reported in detail.

## Next steps

- Add recursive directory scanning.
- Add a dry-run option and configurable categories.
- Improve error handling and reporting.
