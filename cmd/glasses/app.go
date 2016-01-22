package main

import (
	"flag"
	"log"

	. "github.com/ahamidi/glasses"
)

var (
	file  = flag.String("f", "", "Local file to analyze.")
	url   = flag.String("u", "", "URL for image to analyze.")
	face  = flag.Bool("face", false, "Face detection.")
	label = flag.Bool("label", false, "Label detection.")
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

	reqs := &CloudVisionRequest{
		Requests: []*AnnotateRequest{},
		User:     "ahamidi",
	}

	if *face {
		req := &AnnotateRequest{
			Image: f,
			Features: []Feature{
				{"FACE_DETECTION", 10},
			},
		}
		reqs.Requests = append(reqs.Requests, req)
	}

	if *label {
		req := &AnnotateRequest{
			Image: f,
			Features: []Feature{
				{"LABEL_DETECTION", 10},
			},
		}
		reqs.Requests = append(reqs.Requests, req)
	}

	gl, err := NewGlasses()

	resp, err := gl.Do(reqs)
	if err != nil {
		log.Println("Error:", err)
	}

	log.Println("Response:", resp)
}
