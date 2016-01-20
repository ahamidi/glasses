package glasses

import (
	b64 "encoding/base64"
	"io/ioutil"
	"net/http"
)

type Image struct {
	Content string `json:"content"`
}

func NewImageFromURL(url string) (*Image, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	enc := b64.StdEncoding.EncodeToString(body)

	return &Image{enc}, nil
}

func NewImageFromFile(file string) (*Image, error) {

	// Read File
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Encode File
	enc := b64.StdEncoding.EncodeToString(dat)

	img := &Image{enc}

	return img, nil
}
