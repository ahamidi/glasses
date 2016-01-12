package main

import (
	"flag"
	"log"

	. "github.com/ahamidi/glasses"
)

var (
	file = flag.String("f", "", "Local file to analyze.")
	url  = flag.String("u", "", "URL for image to analyze.")
	face = flag.Bool("face", false, "Face detection.")
)

func main() {
	log.Println("Glasses - Google Cloud Vision CLI")
	flag.Parse()

	var f *Image
	var err error

	if *file != "" {
		f, err = NewImageFromFile(*file)
	}

	if *url != "" {
		f, err = NewImageFromURL(*url)
	}

	if err != nil {
		log.Println("Error:", err)
	}

	req := &AnnotateRequest{
		Image: f,
		Features: []Feature{
			{"FACE_DETECTION", 1},
		},
	}

	gl, err := NewGlasses()

	reqs := &CloudVisionRequest{
		Requests: []*AnnotateRequest{req},
		User:     "ahamidi",
	}

	resp, err := gl.Do(reqs)
	if err != nil {
		log.Println("Error:", err)
	}

	log.Println("Response:", resp)
}
