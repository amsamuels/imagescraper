package utils

import (
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// download image
func Download(imgUrl string, directory string, wg *sync.WaitGroup, cookie string) {
	defer wg.Done()
	fileName, err := getUrlFileName(imgUrl)
	if err != nil {
		return
	}
	path := filepath.Join(directory, fileName)
	file, err := os.Create(path)
	if err != nil {
		Warning(err)
		return
	}
	defer file.Close()
	resp, err := httpClient(imgUrl, cookie)
	if err != nil {
		Error(err)
		return
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		Warning(err)
		return
	}
	Infof("Downloaded a file %s with size %s ", fileName, ByteCount(size, "KB"))
}

// get file name by image url
func getUrlFileName(imgUrl string) (string, error) {
	fileUrl, err := url.Parse(imgUrl)
	if err != nil {
		Warning(err)
		return "", err
	}
	path := fileUrl.Path
	segments := strings.Split(path, "/")
	return segments[len(segments)-1], nil
}
