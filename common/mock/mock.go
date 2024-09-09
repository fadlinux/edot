package mock

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	es "gopkg.in/olivere/elastic.v6"
)

type (
	// ElasticMock : Represent Elastic Search Mock Data
	ElasticMock struct {
		ElasticClient *es.Client
		Server        *httptest.Server
	}
)

// NewMockElastic : Make new elastic mock connection
func NewMockElastic(bodyJSON *string, wantedResult string) (mockedElastic ElasticMock, err error) {
	mockedServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		body, _ := ioutil.ReadAll(req.Body)
		*bodyJSON = string(body)
		if wantedResult == "" {
			wantedResult = `{"result" : "created"}`
		}
		io.WriteString(rw, wantedResult)
	}))
	mockedElastic.Server = mockedServer
	mockedElastic.ElasticClient, err = es.NewSimpleClient(es.SetURL(mockedServer.URL))

	return
}

// NewHTTPMock : Make new http connection
func NewHTTPMock(bodyJSON *string, expectedResult string) (mockedServer *httptest.Server) {
	mockedServer = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if expectedResult == "error" {
			rw.Header().Set("Content-Length", "1")
		} else {
			rw.Header().Set("Content-Type", "application/json")
			body, _ := ioutil.ReadAll(req.Body)
			*bodyJSON = string(body)
			io.WriteString(rw, expectedResult)
		}
	}))
	return
}
