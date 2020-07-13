package rq

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

// fetch requests a url using the http protocol.
func (rq *RQ) fetch(method string) *RQ {
	if rq.Err != nil {
		return rq
	}
	if rq.Client == nil {
		rq.Client = &http.Client{}
	}
	var req *http.Request
	if rq.Body != nil {
		req, rq.Err = http.NewRequest(method, rq.URL, bytes.NewBuffer(rq.Body))
	} else {
		req, rq.Err = http.NewRequest(method, rq.URL, nil)
	}
	if rq.Err != nil {
		return rq
	}
	for k, v := range rq.Header {
		req.Header.Set(k, strings.Join(v, ","))
	}
	for _, cookie := range rq.cookies {
		req.AddCookie(cookie)
	}
	rq.res, rq.Err = rq.Client.Do(req)
	if rq.Err != nil {
		return rq
	}
	defer rq.res.Body.Close()
	rq.resBody, rq.Err = ioutil.ReadAll(rq.res.Body)
	return rq
}
