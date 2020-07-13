package rq

import (
	"encoding/json"
	"net/http"
	"strings"
)

// RQ (requester) implements HTTP requests operations
//
// If an operation within the pipeline fails,
// then all subsequent operations will silently forward the eror,
// and no operations are done.
type RQ struct {
	// URL is the request endpoint
	URL string
	// Body define a set of bytes to send
	Body []byte
	// Err stores all kind of errors that can occour on any operation.
	// The Err is overrited if there is a new error
	Err error
	// Client is a http client, used to make requests
	Client *http.Client
	// Header content for requests
	Header http.Header
	// res is a HTTP response
	res *http.Response
	// resBody stores all bytes from the response body
	resBody []byte
	// cookies to send within the request
	cookies []*http.Cookie
}

func fixEndpointPrefix(endpoint string) string {
	if strings.HasPrefix(endpoint, "http://") || strings.HasPrefix(endpoint, "https://") {
		return endpoint
	}
	return "http://" + endpoint
}

// Endpoint creates an RQ which will request the endpoint.
// endpoint must contain the prefix http://
func Endpoint(endpoint string) *RQ {
	return &RQ{
		URL:     fixEndpointPrefix(endpoint),
		Body:    nil,
		Err:     nil,
		Client:  &http.Client{},
		Header:  http.Header{},
		res:     nil,
		resBody: nil,
		cookies: nil,
	}
}

// SetClient defines the new http client used to make requests
func (rq *RQ) SetClient(client *http.Client) *RQ {
	rq.Client = client
	return rq
}

// JSON sets the JSON used to send within the request
func (rq *RQ) JSON(data interface{}) *RQ {
	rq.Body, rq.Err = json.Marshal(data)
	rq.Header.Set("Content-Type", "application/json")
	return rq
}

// Text sets the text used to send within the request
func (rq *RQ) Text(data string) *RQ {
	rq.Body = []byte(data)
	rq.Header.Set("Content-Type", "text/plain")
	return rq
}

// Cookies sets the Cookies which is used to send within the request
func (rq *RQ) Cookies(cookies []*http.Cookie) *RQ {
	rq.cookies = cookies
	return rq
}

// Cookie sets a cookie which is used to send within the request
func (rq *RQ) Cookie(cookie *http.Cookie) *RQ {
	rq.cookies = append(rq.cookies, cookie)
	return rq
}

// SetHeader sets a new field for HTTP header
func (rq *RQ) SetHeader(name, value string) *RQ {
	rq.Header.Set(name, value)
	return rq
}

// Endpoint sets a new value for endpoint
// endpoint must contain the prefix http://
func (rq *RQ) Endpoint(endpoint string) *RQ {
	rq.URL = fixEndpointPrefix(endpoint)
	return rq
}

// Post requests using the HTTP POST method
func (rq *RQ) Post() *RQ {
	return rq.fetch("POST")
}

// Get requests using the HTTP GET method
func (rq *RQ) Get() *RQ {
	return rq.fetch("GET")
}

// Put requests using the HTTP PUT method
func (rq *RQ) Put() *RQ {
	return rq.fetch("PUT")
}

// Delete requests using the HTTP DELETE method
func (rq *RQ) Delete() *RQ {
	return rq.fetch("DELETE")
}

// Head requests using the HTTP HEAD method
func (rq *RQ) Head() *RQ {
	return rq.fetch("HEAD")
}

// Patch requests using the HTTP PATCH method
func (rq *RQ) Patch() *RQ {
	return rq.fetch("PATCH")
}

// ToJSON unmarshal the response body to data
func (rq *RQ) ToJSON(data interface{}) *RQ {
	if rq.Err == nil {
		rq.Err = json.Unmarshal(rq.resBody, data)
	}
	return rq
}

// ToString convert the response body to string
func (rq *RQ) ToString(str *string) *RQ {
	*str = string(rq.resBody)
	return rq
}

// StatusCode returns the last status code response
func (rq *RQ) StatusCode() int {
	return rq.Response().StatusCode
}

// ResponseBody returns the response body of the latest request made
func (rq *RQ) ResponseBody() []byte {
	return rq.resBody
}

// Response returns a http response
func (rq *RQ) Response() *http.Response {
	return rq.res
}
