package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	log "github/fadlinux/edot/common/util/log"

	json "github.com/json-iterator/go"
)

const (
	contentType              = "Content-Type"
	applicationJSON          = "application/json"
	accessControlAllowOrigin = "Access-Control-Allow-Origin"
)

// Render render http response with cache (in minutes) and callback param if exist
func Render(w http.ResponseWriter, response interface{}, cache int, callback string) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.WithField("error", err).Error("Failed to marshal response")
	}

	if callback != "" {
		w.Header().Add(contentType, "text/javascript")
	} else {
		w.Header().Add(contentType, applicationJSON)

	}

	if cache > 0 {
		now := time.Now()
		t := now.Add(time.Duration(cache) * time.Minute)

		w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(cache*60))
		w.Header().Set("Expires", t.Format(time.RFC1123))
		w.Header().Set("Date", now.Format(time.RFC1123))
	}

	if callback != "" {
		fmt.Fprintf(w, "%s(%s)", callback, jsonResponse)
	} else {
		w.Write(jsonResponse)
	}
}

// RenderHTTPJSON middleware to write json to http
func RenderHTTPJSON(w http.ResponseWriter, data interface{}, code int, callback string) {
	b, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "StatusInternalServerError", http.StatusInternalServerError)
		return
	}

	if callback != "" {
		w.Header().Set(contentType, "text/javascript")
		fmt.Fprintf(w, "%s(%s)", callback, b)
		return
	}

	w.Header().Set(accessControlAllowOrigin, "*")

	w.Header().Set(contentType, applicationJSON)
	w.WriteHeader(code)
	w.Write(b)
}

// RenderHTTPCSV middleware to write CSV to http
func RenderHTTPCSV(w http.ResponseWriter, data string, code int, callback string) {
	var output = []byte(data)
	w.Header().Set(accessControlAllowOrigin, "*")

	w.Header().Set(contentType, "application/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=feedback.csv")

	w.Write(output)
}

// RenderHTTPCode middle to write http header code to header
func RenderHTTPCode(w http.ResponseWriter, message string, code int) {
	type response struct {
		M string `json:"message"`
	}

	b, err := json.Marshal(response{M: message})
	if err != nil {
		http.Error(w, "StatusInternalServerError", http.StatusInternalServerError)
		return
	}

	w.Header().Set(accessControlAllowOrigin, "*")

	w.Header().Set(contentType, applicationJSON)
	w.WriteHeader(code)
	w.Write(b)
}

// RenderHTTPText middleware to write plain text to http
func RenderHTTPText(w http.ResponseWriter, message string, code int) {
	b := []byte(message)

	w.Header().Set(accessControlAllowOrigin, "*")

	w.Header().Set(contentType, "text/plain")
	w.WriteHeader(code)
	w.Write(b)
}

// RenderHTTPTextWithHeader middleware to write plain text to http
func RenderHTTPTextWithHeader(w http.ResponseWriter, message string, code int, additionalHeader map[string]string) {
	b := []byte(message)

	w.Header().Set(accessControlAllowOrigin, "*")

	for k := range additionalHeader {
		w.Header().Set(k, additionalHeader[k])
	}

	w.Header().Set(contentType, "text/plain")
	w.WriteHeader(code)
	w.Write(b)
}

// GetJSON parse http response from the given url to targeted struct
// Returned response must be in json format
func GetJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// CheckURLCode try to hit and check URL given to find out the http code
func CheckURLCode(url string) (httpCode int) {
	// default value - if the result is -1, there is problem with the code
	httpCode = -1
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()

	req = req.WithContext(ctx)
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64; rv:80.0) Gecko/20100101 Firefox/80.0`)
	req.Header.Add("Origin", "https://www.tokopedia.com")

	// Sends request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err == nil {
		httpCode = resp.StatusCode
	}
	return
}
