package glasses

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	CLOUD_VISION_ENDPOINT = "https://vision.googleapis.com/v1alpha1/images:annotate"
)

type Glasses struct {
	Client *http.Client
}

type CloudVisionRequest struct {
	Requests []*AnnotateRequest `json:"requests"`
	User     string             `json:"user"`
}

type AnnotateRequest struct {
	Image        *Image        `json:"image"`
	Features     []Feature     `json:"features"`
	ImageContext *ImageContext `json:"imageContext,omitempty"`
}

type Feature struct {
	Type       string `json:"type"`
	MaxResults int    `json:"maxResults"`
}

type ImageContext struct {
	LatLongRect                 interface{} `json:"latLongRect"`
	ImageContextSearchExtension interface{} `imageContextSearchExtension`
}

//type AnnotateResponse struct {
//FaceAnnotations      []FaceAnnotation     `json:"faceAnnotations"`
//LandmarkAnnotations  []LandmarkAnnotation `json:"landmarkAnnotation"`
//LogoAnnotations      []LogoAnnotation     `json:"logoAnnotations"`
//LabelAnnotations     []LabelAnnotation    `json:"labelAnnotations"`
//TextAnnotations      []TextAnnotation     `json:"textAnnotations"`
//SafeSearchAnnotation SafeSearchAnnotation `json:"safeSearchAnnotation"`
//SuggestAnnotations   []SuggestAnnotation  `json:"suggestAnnotations"`
//QueryAnnotation      QueryAnnotation      `json:"queryAnnotation"`
//Error                Status               `json:"error"`
//}

type AnnotateResponses struct {
	Responses []AnnotateResponse `json:"responses"`
}

type AnnotateResponse struct {
	FaceAnnotations      []interface{} `json:"faceAnnotations"`
	LandmarkAnnotations  []interface{} `json:"landmarkAnnotation"`
	LogoAnnotations      []interface{} `json:"logoAnnotations"`
	LabelAnnotations     []interface{} `json:"labelAnnotations"`
	TextAnnotations      []interface{} `json:"textAnnotations"`
	SafeSearchAnnotation interface{}   `json:"safeSearchAnnotation"`
	SuggestAnnotations   []interface{} `json:"suggestAnnotations"`
	QueryAnnotation      interface{}   `json:"queryAnnotation"`
	Error                interface{}   `json:"error"`
}

func NewGlasses() (*Glasses, error) {
	client, err := google.DefaultClient(oauth2.NoContext, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return nil, err
	}

	return &Glasses{client}, nil
}

func (g *Glasses) Do(r *CloudVisionRequest) (*AnnotateResponses, error) {

	jR, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	rawResp, err := g.Client.Post(CLOUD_VISION_ENDPOINT, "application/json", bytes.NewBuffer(jR))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return nil, err
	}

	//log.Println("Body:", string(body))

	var resp *AnnotateResponses

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
