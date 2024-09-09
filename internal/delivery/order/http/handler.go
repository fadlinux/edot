package http

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	cHttp "github/fadlinux/edot/common/http"

	mOrder "github/fadlinux/edot/internal/model/order"

	"github.com/julienschmidt/httprouter"
)

// HandleSearch : handler for order search
func (d Delivery) HandleSearch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	startTime := time.Now()
	var err error
	var result mOrder.Response
	var dataReq mOrder.Order

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&dataReq)
	if err != nil {
		result.Header.Message = err.Error()
		result.Header.StatusCode = 400
	}

	result.Data = make([]mOrder.Data, 0)

	if err != nil {
		result.Header.Message = err.Error()
		result.Header.StatusCode = 400
	}

	result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
	result.Header.Message = "Success, Create order!"
	result.Header.StatusCode = 200

	cHttp.Render(w, result, 0, req.FormValue("callback"))

}

// HandleAddOrder : handler for order add
func (d Delivery) HandleAddOrder(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	startTime := time.Now()
	var err error
	var result mOrder.Response
	var dataReq mOrder.Order
	var lastID int64

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&dataReq)
	if err != nil {
		result.Header.Message = err.Error()
		result.Header.StatusCode = 400
	}

	lastID, err = d.orderUC.AddOrder(context.Background(), dataReq)

	if err != nil {
		result.Header.Message = err.Error()
		result.Header.StatusCode = 400
	}

	result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
	result.Header.Message = "Success, Create order !" + string(lastID)
	result.Header.StatusCode = 200

	cHttp.Render(w, result, 0, req.FormValue("callback"))

}
