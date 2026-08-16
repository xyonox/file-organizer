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
- Loads folder and extension mappings from `organizedFolders.json`.
- Supports a dry-run mode that shows planned moves without changing files.
- Supports optional recursive scanning with the `-r` flag.
- Leaves macOS `.app` bundles and configured category folders untouched.
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

To preview the changes without moving any files, use `--dry-run`:

```bash
go run . -d /path/to/directory --dry-run
```

To scan subdirectories as well, add the `-r` flag:

```bash
go run . -d /path/to/directory -r
```

The flags can also be combined for a recursive preview:

```bash
go run . -d /path/to/directory -r --dry-run
```

The `-d`, `-r` and `--dry-run` flags can also be used with a built binary.

## Building a binary

Build the application with:

```bash
go build -o file-organizer .
```

## How it works

By default, the application reads files directly inside the selected directory and
matches their file extensions against the `organizedFolders` list. With the `-r`
flag, it also scans nested subdirectories. Files found recursively are moved into
the category folders in the selected directory.

Configured category folders are not scanned again, and macOS application bundles
whose names end in `.app` are skipped because they are directories containing an
application. If no directory is provided with `-d`, the application prompts for
one.

On startup, the application loads `organizedFolders.json` from the current working
directory. If the file does not exist or contains invalid JSON, the built-in default
categories are used and saved to that file. You can customize the categories by
editing the JSON configuration. Each entry has a `name` and a `fileTypes` array:

```json
{
  "organizedFolders": [
    {
      "name": "Images",
      "fileTypes": ["jpg", "jpeg", "png"]
    }
  ]
}
```

The following folders are currently supported:

| Folder        | File extensions                                                                                                                                                         |
|---------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Images        | `jpg`, `jpeg`, `png`, `gif`, `webp`, `svg`, `bmp`, `tiff`, `tif`, `heic`, `heif`, `ico`, `raw`                                                                          |
| Videos        | `mp4`, `mov`, `avi`, `mkv`, `webm`, `wmv`, `flv`, `mpeg`, `mpg`, `m4v`, `3gp`                                                                                           |
| Audio         | `mp3`, `wav`, `flac`, `m4a`, `ogg`, `aac`, `wma`, `aiff`, `alac`, `mid`, `midi`                                                                                         |
| Documents     | `pdf`, `doc`, `docx`, `odt`, `rtf`, `tex`, `txt`, `md`, `csv`                                                                                                           |
| Spreadsheets  | `xls`, `xlsx`, `ods`, `numbers`                                                                                                                                         |
| Presentations | `ppt`, `pptx`, `odp`, `key`                                                                                                                                             |
| Archives      | `zip`, `rar`, `7z`, `tar`, `gz`, `bz2`, `xz`, `iso`                                                                                                                     |
| Apps          | `app`, `exe`, `dmg`, `apk`, `ipa`                                                                                                                                       |
| Installers    | `msi`, `pkg`, `deb`, `rpm`, `appimage`                                                                                                                                  |
| Code          | `go`, `js`, `ts`, `jsx`, `tsx`, `html`, `css`, `scss`, `json`, `xml`, `yaml`, `yml`, `py`, `java`, `c`, `cpp`, `h`, `hpp`, `cs`, `php`, `rb`, `rs`, `kt`, `swift`, `sh` |
| Fonts         | `ttf`, `otf`, `woff`, `woff2`, `eot`                                                                                                                                    |
| Ebooks        | `epub`, `mobi`, `azw`, `azw3`, `fb2`                                                                                                                                    |
| Design        | `psd`, `ai`, `xd`, `fig`, `sketch`, `indd`                                                                                                                              |

Matching files are moved into the corresponding folder. Files without a matching
extension remain in their source directory. In dry-run mode, the planned moves are
printed but no files are changed.

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
├── main.go               # Application source code
├── organizedFolders.json # Folder and extension configuration
├── go.mod                # Go module definition
└── README.md             # Project documentation
```

## Current limitations

- Recursive scanning must be explicitly enabled with the `-r` flag.
- Files with unsupported extensions are not moved.
- Recursive moves place files from nested directories into the category folders in
  the selected directory; the original subdirectory structure is not preserved.
- The configuration file is loaded from the current working directory rather than
  from the directory being organized.
- Existing files with the same name in a destination folder may cause a move error.

## Next steps

- Improve error handling and reporting.
