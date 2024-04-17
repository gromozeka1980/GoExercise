package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	var urlFilePath, downloadDirectory, useTreeStr string
	var numWorkers int

	flag.StringVar(&urlFilePath, "urlfile", "urls.txt", "Path to the text file containing the URLs")
	flag.StringVar(&downloadDirectory, "dir", "downloads", "Directory to download the files")
	flag.StringVar(&useTreeStr, "tree", "false", "Use tree structure for directories ('true' or 'false')")
	flag.IntVar(&numWorkers, "workers", 10, "Number of concurrent workers for downloading files")
	flag.Parse()

	useTree := (useTreeStr == "true") // Convert the string flag to boolean

	fmt.Println("URL File Path:", urlFilePath)
	fmt.Println("Download Directory:", downloadDirectory)
	fmt.Println("Use Tree Structure:", useTree)
	fmt.Println("Number of Workers:", numWorkers)

	downloadFilesFromUrls(urlFilePath, downloadDirectory, useTree, numWorkers)
}

func downloadFilesFromUrls(urlFilePath, downloadDirectory string, useTree bool, numWorkers int) {
	urls, err := readLines(urlFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	urlChan := make(chan string, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			for url := range urlChan {
				downloadFile(url, downloadDirectory, useTree)
			}
			wg.Done()
		}()
	}

	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan)
	wg.Wait()
}

func readLines(path string) ([]string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	for i, line := range lines {
		lines[i] = strings.Replace(line, "\r", "", -1) //Windows...
	}
	return lines, nil
}

func downloadFile(urlString, basePath string, useTree bool) {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error parsing URL:", urlString, err)
		return
	}

	var localPath string
	if useTree {
		localPath = filepath.Join(basePath, parsedUrl.Host, parsedUrl.Path)
		if strings.HasSuffix(localPath, "/") {
			localPath = filepath.Join(localPath, "index.html")
		} else {
			localPath = filepath.Join(filepath.Dir(localPath), filepath.Base(localPath), "index.html")
		}
	} else {
		safeFileName := urlToFileName(parsedUrl.Host + parsedUrl.Path)
		localPath = filepath.Join(basePath, safeFileName+".html")
	}

	err = os.MkdirAll(filepath.Dir(localPath), 0755)
	if err != nil {
		fmt.Println("Error creating directories:", localPath, err)
		return
	}

	resp, err := http.Get(urlString)
	if err != nil {
		fmt.Println("Error getting URL:", urlString, err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(localPath)
	if err != nil {
		fmt.Println("Error creating file:", localPath, err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error writing to file:", localPath, err)
	}
}

func urlToFileName(urlPart string) string {
	return strings.NewReplacer("http://", "", "https://", "", "/", "_", "?", "_", "&", "_", "=", "_").Replace(urlPart)
}
