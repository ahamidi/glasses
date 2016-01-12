package glasses

import (
	b64 "encoding/base64"
	"io/ioutil"
)

type Image struct {
	Content string `json:"content"`
}

func NewImageFromURL(url string) (*Image, error) {

	return nil, nil
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
