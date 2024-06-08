package utils

import (
	"fmt"
	"io"
	"net/url"
	"regexp"
	"strings"
)

// crawler page
func CrawlerPageImage(rawurl string, cookie string) ([]string, error) {
	//validate url
	// isvalid := validateUrl(rawurl, cookie)
	// if !isvalid {
	// 	Errorf("your url %s is not support", rawurl)
	// 	return nil, fmt.Errorf("your url %s is not support", rawurl)
	// }

	var urls []string
	Infof("start crawler %s page", rawurl)
	resp, err := httpClient(rawurl, cookie)
	if err != nil {
		Error(err)
		return nil, nil
	}
	defer resp.Body.Close()

	var bodyStr string

	if resp.StatusCode == 200 {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			Fatal(err)
		}
		bodyStr = string(bodyBytes)
	}
	reg := regexp.MustCompile(`<img.*src="(([a-zA-Z0-9]|\/|(https?)).*?)".*\/?>`)
	matchImgUrl := reg.FindAllStringSubmatch(bodyStr, -1)
	u, _ := url.Parse(rawurl)
	scheme := u.Scheme
	host := u.Host
	//use map extract duplicate
	var urlsMap = make(map[string]string, len(matchImgUrl))
	for _, item := range matchImgUrl {
		imgUrl := item[1]
		if strings.HasPrefix(imgUrl, "//") {
			imgUrl = fmt.Sprintf("%s:%s", scheme, imgUrl)
		} else if strings.HasPrefix(imgUrl, "/") {
			imgUrl = fmt.Sprintf("%s://%s%s", scheme, host, imgUrl)
		} else if !strings.HasPrefix(imgUrl, "http") {
			imgUrl = fmt.Sprintf("%s://%s%s%s", scheme, host, "/", imgUrl)
		}
		urlsMap[imgUrl] = ""
	}

	for imgUrl := range urlsMap {
		urls = append(urls, imgUrl)
	}
	return urls, nil
}

// validate crawler target is alive
// the method has timeout
// func validateUrl(rawurl string, cookie string) bool {
// 	if rawurl == "" {
// 		Error("Your url is empty")
// 		os.Exit(0)
// 		return false
// 	}

// 	Infof("dial %s", rawurl)
// 	resp, err := httpClient(rawurl, cookie)
// 	if err != nil {
// 		Error(err)
// 		return false
// 	}

// 	if resp.StatusCode != 200 {
// 		Infof("%s response code is %d.", rawurl, resp.StatusCode)
// 		return false
// 	}
// 	return true
// }
