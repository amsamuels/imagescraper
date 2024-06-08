A page image downloader


Logic

- parse the crawle url and download directory from command args
- get the body of response 
- use regex match image url with suffix
- use go concurrency download the image to your directory
- exit

Usage

`go run cmd/run.go -d images -uri https://github.com -cookie "cookie string"`