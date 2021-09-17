package funcs

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Wgeter struct {
	url      string
	fileName string
}

func InitWget(url, fileName string) (*Wgeter, error) {
	if url == "" {
		return nil, errors.New("empty text")
	}
	return &Wgeter{url, fileName}, nil
}

func (w *Wgeter) Start() error {
	resp, err := getResponse(w.url)
	if err != nil {
		return err
	}
	fileName, err := createFile(w.url, w.fileName)
	if err != nil {
		return err
	}
	w.fileName = fileName

	f, err := os.OpenFile(w.fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	size, err := io.Copy(f, resp.Body)

	defer f.Close()
	println("Downloaded a file " + w.fileName + " with size " + strconv.FormatInt(size, 10))
	return nil
}

func getResponse(url string) (*http.Response, error) {
	client := &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func createFile(addr, filename string) (string, error) {
	if filename == "" {
		u, err := url.Parse(addr)
		if err != nil {
			panic(err)
		}
		filename = u.Path
		segments := strings.Split(filename, "/")
		filename = segments[len(segments)-1]
		if filename == "" {
			filename = "wget_" + u.Host
		}
	}
	file, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return filename, nil
}
