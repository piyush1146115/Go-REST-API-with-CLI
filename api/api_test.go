package api

import (
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/cast"

	"github.com/stretchr/testify/assert"
)

type Request struct {
	Method             string
	URL                string
	Body               io.Reader
	ExpectedStatusCode int
}

func TestSomething(t *testing.T) {
	assert.True(t, true, "True is true!")
}

func TestGetArticle(t *testing.T) {
	requests := make([]Request, 3)

	requests[0] = Request{
		"GET",
		"http://localhost:10000/article/1",
		nil,
		200,
	}
	requests[1] = Request{
		"GET",
		"http://localhost:10000/article/2",
		nil,
		200,
	}
	requests[2] = Request{
		"GET",
		"http://localhost:10000/article/4",
		nil,
		404,
	}

	processRequest(t, requests)
}

func TestGetArticles(t *testing.T) {
	requests := make([]Request, 2)

	requests[0] = Request{
		"GET",
		"http://localhost:10000/articles",
		nil,
		200,
	}
	requests[1] = Request{
		"GET",
		"http://localhost:10000/article",
		nil,
		404,
	}

	processRequest(t, requests)
}

func processRequest(t *testing.T, reqs []Request) {
	for _, req := range reqs {
		r, _ := http.NewRequest(req.Method, req.URL, req.Body)
		r.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("user:secret")))
		w := httptest.NewRecorder()
		MyRouter.ServeHTTP(w, r)
		if w.Code != req.ExpectedStatusCode {
			t.Error("\nExpected Status Code\t= " + cast.ToString(req.ExpectedStatusCode) + "\nFound Status Code\t\t= " + cast.ToString(w.Code) + "\n")
		}
	}
}
