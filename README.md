## Glasses
Go package for the [Google Cloud Vision API](https://cloud.google.com/vision/)

### TODO

- [x] Auth
- [x] Request
- [ ] Tests
- [ ] Docs
- [ ] CI

### Usage

#### CLI Tool

* Build: `go build ./cmd/glasses`
* Run: `./glasses -f image.jpg`

### Requirements

* Authentication:
    Authentication depends on [Application Default Credentials](https://developers.google.com/identity/protocols/application-default-credentials). Setting that up is beyond the scope of this project, but essentially requires the creation of a `Service Account`, download the JSON formatted credentials and setting an Env Var to point at it.

