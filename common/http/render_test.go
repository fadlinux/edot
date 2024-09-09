package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestRender(t *testing.T) {
	type args struct {
		response interface{}
		cache    int
		callback string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Complete Argument",
			args: args{
				response: `{"alive": true}`,
				callback: "callback",
				cache:    0,
			},
		},
		{
			name: "Empty Callback",
			args: args{
				response: `{"alive": true}`,
				callback: "",
				cache:    0,
			},
		},
		{
			name: "Set Cache",
			args: args{
				response: `{"alive": true}`,
				callback: "",
				cache:    60,
			},
		},
		{
			name: "Fail Json Marshal",
			args: args{
				response: `fail`,
				callback: "",
				cache:    60,
			},
		},
	}
	for _, tt := range tests {
		var ts *httptest.Server
		if tt.args.response == "fail" {
			ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Render(w, make(chan int), tt.args.cache, tt.args.callback)
			}))
		} else {
			ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Render(w, tt.args.response, tt.args.cache, tt.args.callback)
			}))
		}
		defer ts.Close()
		http.Get(ts.URL)
	}
}

func TestCheckURLCode(t *testing.T) {
	firstURL := "https://tokopedia/test/success"
	secondURL := "https://tokopedia/test/failed"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", firstURL,
		httpmock.NewStringResponder(200, `{}`))

	httpmock.RegisterResponder("GET", secondURL,
		httpmock.NewStringResponder(500, `{}`))

	type args struct {
		url string
	}
	tests := []struct {
		name         string
		args         args
		wantHttpCode int
	}{
		{
			name: "case 1: 200",
			args: args{
				url: firstURL,
			},
			wantHttpCode: 200,
		},
		{
			name: "case 2: 500",
			args: args{
				url: secondURL,
			},
			wantHttpCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHttpCode := CheckURLCode(tt.args.url); gotHttpCode != tt.wantHttpCode {
				t.Errorf("CheckURLCode() = %v, want %v", gotHttpCode, tt.wantHttpCode)
			}
		})
	}
}
