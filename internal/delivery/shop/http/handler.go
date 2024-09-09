package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	cHttp "github/fadlinux/edot/common/http"

	mShop "github/fadlinux/edot/internal/model/shop"

	"github.com/julienschmidt/httprouter"
)

// HandleSearch : handler for shop search
func (d Delivery) HandleSearch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	startTime := time.Now()
	var err error
	var result mShop.Response
	var dataReq mShop.Shop

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&dataReq)

	result.Data = make([]mShop.Data, 0)

	if err != nil {
		result.Header.Message = err.Error()
		result.Header.StatusCode = 400
	}

	if strings.TrimSpace(dataReq.Name) == "" {
		result.Header.Message = "Name must be filled out"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return

	}

	_, _ = d.shopUC.AddShop(context.Background(), mShop.Shop{
		Name: dataReq.Name,
	})

	result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
	result.Header.Message = "Success, Create shop!"
	result.Header.StatusCode = 200

	cHttp.Render(w, result, 0, req.FormValue("callback"))

}

// HandleAddShop : handler for shop add
func (d Delivery) HandleAddShop(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	startTime := time.Now()
	var err error
	var result mShop.Response
	var dataReq mShop.Shop

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&dataReq)

	result.Data = make([]mShop.Data, 0)

	if err != nil {
		result.Header.Message = err.Error()
		result.Header.StatusCode = 400
	}

	if strings.TrimSpace(dataReq.Name) == "" {
		result.Header.Message = "Name must be filled out"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return

	}

	_, _ = d.shopUC.AddShop(context.Background(), mShop.Shop{
		Name: dataReq.Name,
	})

	result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
	result.Header.Message = "Success, Create shop!"
	result.Header.StatusCode = 200

	cHttp.Render(w, result, 0, req.FormValue("callback"))

}
