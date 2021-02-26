package ipp

import "io"

type Adapter interface {
	SendRequest(url string, req *Request, additionalResponseData io.Writer) (*Response, error)
	GetHttpUri(namespace string, object interface{}) string
	TestConnection() error
}
