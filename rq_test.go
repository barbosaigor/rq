package rq

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Product struct {
	Code     string  `json:"code"`
	Quantity float32 `json:"quantity"`
}

type Products []Product

func TestToJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"code": "1234", "quantity": 100}, {"code": "1235", "quantity": 200}]`))
	}))
	defer ts.Close()
	rq := Endpoint(ts.URL).Get()
	if rq.Err != nil {
		t.Errorf("ToJSON - Get: an error was found %v", rq.Err)
	} else {
		var p Products
		rq.ToJSON(&p)
		if rq.Err != nil {
			t.Errorf("ToJSON: an error was found %v", rq.Err)
		} else {
			t.Logf("ToJSON - body: %v", p)
		}
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("ToJSON Get: expected status code 200 but got %v", rq.StatusCode())
	}
}

func TestToString(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte(`bla bla bla`))
	}))
	defer ts.Close()
	rq := Endpoint(ts.URL).Get()
	if rq.Err != nil {
		t.Errorf("ToString - Get: an error was found %v", rq.Err)
	} else {
		var str string
		rq.ToString(&str)
		if rq.Err != nil {
			t.Errorf("ToString: an error was found %v", rq.Err)
		} else if str == "" {
			t.Error("ToString: expect non empty string but got an empty")
		} else {
			t.Logf("ToString - body: %v", str)
		}
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("ToString Get: expected status code 200 but got %v", rq.StatusCode())
	}
}

func TestGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
	}))
	defer ts.Close()
	rq := Endpoint(ts.URL).Get()
	if rq.Err != nil {
		t.Errorf("Get: an error was found %v", rq.Err)
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("Get: expected status code 200 but got %v", rq.StatusCode())
	}
}

func TestPut(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
	}))
	defer ts.Close()
	rq := Endpoint(ts.URL).Put()
	if rq.Err != nil {
		t.Errorf("Put: an error was found %v", rq.Err)
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("Put: expected status code 200 but got %v", rq.StatusCode())
	}
}

func TestPost(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
	}))
	defer ts.Close()
	rq := Endpoint(ts.URL).Post()
	if rq.Err != nil {
		t.Errorf("Post: an error was found %v", rq.Err)
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("Post: expected status code 200 but got %v", rq.StatusCode())
	}
}

func TestPatch(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
	}))
	defer ts.Close()
	rq := Endpoint(ts.URL).Patch()
	if rq.Err != nil {
		t.Errorf("Patch: an error was found %v", rq.Err)
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("Patch: expected status code 200 but got %v", rq.StatusCode())
	}
}

func TestDelete(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
	}))
	defer ts.Close()
	rq := Endpoint(ts.URL).Delete()
	if rq.Err != nil {
		t.Errorf("Delete: an error was found %v", rq.Err)
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("Delete: expected status code 200 but got %v", rq.StatusCode())
	}
}

func TestHead(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "HEAD" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
	}))
	defer ts.Close()
	rq := Endpoint(ts.URL).Head()
	if rq.Err != nil {
		t.Errorf("Head: an error was found %v", rq.Err)
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("Head: expected status code 200 but got %v", rq.StatusCode())
	}
}

func TestJSON(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" || r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}))
	defer ts.Close()
	payload := Products{
		Product{Code: "1234", Quantity: 112},
		Product{Code: "1111", Quantity: 20000.122},
	}
	rq := Endpoint(ts.URL).JSON(payload).Post()
	if rq.Err != nil {
		t.Errorf("JSON: an error was found %v", rq.Err)
	} else {
		var products Products
		rq.ToJSON(&products)
		if rq.Err != nil {
			t.Errorf("JSON: error to parse JSON %v", rq.Err)
		} else {
			t.Logf("JSON body: %v", products)
		}
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("Head: expected status code 200 but got %v", rq.StatusCode())
	}
}

func TestStatusCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
	}))
	defer ts.Close()
	rq := Endpoint(ts.URL).Get()
	if rq.Err != nil {
		t.Errorf("StatusCode: an error was found %v", rq.Err)
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("StatusCode: expected status code 200 but got %v", rq.StatusCode())
	}
	rq = Endpoint(ts.URL).Post()
	if rq.Err != nil {
		t.Errorf("StatusCode: an error was found %v", rq.Err)
	}
	if rq.StatusCode() != http.StatusNotImplemented {
		t.Errorf("StatusCode: expected status code %v but got %v", http.StatusNotImplemented, rq.StatusCode())
	}
}

func TestResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
	}))
	defer ts.Close()
	rq := Endpoint(ts.URL).Get()
	if rq.Err != nil {
		t.Errorf("Response: an error was found %v", rq.Err)
	}
	if rq.Response() == nil {
		t.Errorf("Response: expected response but got %v", rq.Response())
	}
}

func TestCookie(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "something", Value: "abc"})
		for _, cookie := range r.Cookies() {
			http.SetCookie(w, cookie)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"code": "1234", "quantity": 100}, {"code": "1235", "quantity": 200}]`))
	}))
	defer ts.Close()
	cookies := make([]*http.Cookie, 2)
	cookies[0] = &http.Cookie{Name: "field1", Value: "v1"}
	cookies[1] = &http.Cookie{Name: "field2", Value: "v2"}
	rq := Endpoint(ts.URL).Cookies(cookies).Get()
	if rq.Err != nil {
		t.Errorf("Cookie - Get: an error was found %v", rq.Err)
	} else {
		cks := rq.Response().Cookies()
		if len(cks) != 3 {
			t.Errorf("Cookie: expected len 3 but got %v", len(cks))
		} else {
			t.Logf("Cookie - Response Cookies: %v", cks)
		}
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("Cookie Get: expected status code 200 but got %v", rq.StatusCode())
	}
}

func TestSetHeader(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "wohooo"}`))
	}))
	defer ts.Close()
	rq := Endpoint(ts.URL).SetHeader("Content-Type", "application/json").Get()
	if rq.Err != nil {
		t.Errorf("SetHeader - Get: an error was found %v", rq.Err)
	}
	if rq.StatusCode() != http.StatusOK {
		t.Errorf("SetHeader Get: expected status code 200 but got %v", rq.StatusCode())
	}
	rq = Endpoint(ts.URL).SetHeader("Content-Type", "text/plain").Get()
	if rq.Err != nil {
		t.Errorf("SetHeader - Get: an error was found %v", rq.Err)
	}
	if rq.StatusCode() != http.StatusNotImplemented {
		t.Errorf("SetHeader Get: expected status code %v but got %v", http.StatusNotImplemented, rq.StatusCode())
	}
}
