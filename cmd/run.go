/*
* @Author: Jeffery
* @Date:   2020-04-21 18:29:00
* @Last Modified by:   Jeffery
* @Last Modified time: 2020-05-13 11:03:02
 */
package main

import (
	"flag"
	"fmt"
	"imagescraper/utils"
	"os"
	"sync"
)

var (
	version   string = "0.1.0"
	uri       string
	directory string
	cookie    string
	help      bool
)

func main() {
	utils.Init("image-downloader", true)
	parseFlag()

	//create directory
	createDownloadDir(directory)
	//crawler page
	imgUrls, _ := utils.CrawlerPageImage(uri, cookie)
	if len(imgUrls) == 0 {
		utils.Warning("empty images resource")
		return
	}
	utils.Infof("find %d images", len(imgUrls))
	var wg sync.WaitGroup
	for _, imgUrl := range imgUrls {
		wg.Add(1)
		go utils.Download(imgUrl, directory, &wg, cookie)
	}
	wg.Wait()
	utils.Info("download image task complete.")
}

// create save image directory
func createDownloadDir(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.FileMode(0755))
	}
	return path
}

// parse flag
func parseFlag() {
	defer flag.Parse()

	flag.BoolVar(&help, "h", false, "Show usage")
	flag.StringVar(&directory, "d", "images", "Save the directory of image")
	flag.StringVar(&cookie, "cookie", "", "Send cookies from string")
	flag.StringVar(&uri, "uri", "", "Your want download website url for image ")
	flag.Usage = usage
	if help {
		flag.Usage()
	}

}

// Print the tools usage
func usage() {
	fmt.Fprintf(os.Stderr, `download image tools for website version %s
Usage: download [-h] [-uri https://github.com] [-d] images

Options:
`, version)
	flag.PrintDefaults()
}
