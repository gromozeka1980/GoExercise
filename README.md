
# URL File Downloader

This project provides a Go-based tool for downloading files from URLs specified in a text file. It supports downloading with both a hierarchical directory structure ("tree" structure) and a flat directory structure.

## Features

- Concurrent downloading of files using Go routines.
- Option to choose between saving files in a hierarchical folder structure based on URL or in a flat structure.
- Error handling to manage common issues like network errors or write errors.

## Getting Started

### Prerequisites

- Go (version 1.15 or higher is recommended).
- Access to the internet to fetch URLs.
- A text file containing the URLs to download.

### Installation

Clone the repository to your local machine:

\```bash
git clone https://github.com/yourusername/url-file-downloader.git
cd url-file-downloader
\```

### Usage

This program accepts command line arguments to specify the input file, output directory, and the directory structure type.

#### Command Line Arguments

- `-urlfile`: Path to the text file containing the URLs (default is "ListOfAsciiSiteUrls.txt").
- `-dir`: Directory where files will be downloaded (default is "downloads").
- `-tree`: Boolean flag to choose between hierarchical (`true`) or flat (`false`) directory structure for saving files (default is `false`).
- `-workers`: Number of concurrent workers for downloading files (default is 10).


### Directory Structures

- **Hierarchical Structure**: Files are saved mimicking the URL path. For example, `https://www.example.com/path/file` will be saved as `downloads/www.example.com/path/file.html`.
- **Flat Structure**: Files are saved directly under the specified directory with names derived from the URL, converting special characters to underscores. For example, `https://www.example.com/path/file` becomes `downloads/www.example.com_path_file.html`.

### Examples

You can run appropriate usage_example script. 
You will get two folders:
- `downloads_tree`: Contains files downloaded using the hierarchical structure.
- `downloads_flat`: Contains files downloaded using the flat structure.

### License

Do whatever you want!
