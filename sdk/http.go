package sdk

import (
	"io"
	"net/http"
)

var (
	httpGet  = http.Get
	httpPost = http.Post
)

// HTTPFunctions is an interface that represents io library in golang sdk
//go:generate mockery -name=HTTPFunctions
type HTTPFunctions interface {
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

type httpFunctionsImpl struct{}

func (h *httpFunctionsImpl) Get(url string) (resp *http.Response, err error) {
	return httpGet(url)
}

func (h *httpFunctionsImpl) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	return httpPost(url, contentType, body)
}

// ProvideHTTPFunctions ...
func ProvideHTTPFunctions() HTTPFunctions {
	return &httpFunctionsImpl{}
}
