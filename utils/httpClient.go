package utils

import (
	"net/http"
	"net/http/cookiejar"
)

// http client
func httpClient(rawurl string, cookie string) (res *http.Response, err error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		Fatal(err)
		return nil, err
	}

	client := http.Client{
		Jar: jar,
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		Error("Make a request fail")
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Cookie", cookie)
	resp, err := client.Do(req)
	if err != nil {
		Error("request fatal:")
		Error(err)
		return nil, err
	}
	return resp, nil
}
