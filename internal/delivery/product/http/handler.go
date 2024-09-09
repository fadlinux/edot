package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	cHttp "github/fadlinux/edot/common/http"

	mProduct "github/fadlinux/edot/internal/model/product"

	"github.com/julienschmidt/httprouter"
)

// HandleSearch : handler for product search
func (d Delivery) HandleSearch(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	startTime := time.Now()
	var err error
	var result mProduct.Response
	var dataReq mProduct.Product
	var count, page int64

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&dataReq)

	result.Data = make([]mProduct.Data, 0)

	idProduct := req.FormValue("id")
	keyword := req.FormValue("q")
	rows := req.FormValue("size")
	pageStr := req.FormValue("page")

	if rows == "" {
		count = 10
	} else {
		count, err = strconv.ParseInt(rows, 10, 64)
		if err != nil {
			count = 10
		}
	}

	if pageStr == "" {
		page = 1
	} else {
		page, err = strconv.ParseInt(pageStr, 10, 64)
		if err != nil {
			page = 1
			return
		}
	}

	if err != nil {
		result.Header.Message = err.Error()
		result.Header.StatusCode = 400
	}

	result, _ = d.productUC.FetchProduct(context.Background(), mProduct.SearchParam{
		Q:    keyword,
		ID:   idProduct,
		Size: count,
		Page: page,
	})

	result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
	result.Header.Message = "Success"
	result.Header.StatusCode = 200

	cHttp.Render(w, result, 0, req.FormValue("callback"))

}

// HandleAddProduct : handler for add product
func (d Delivery) HandleAddProduct(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	startTime := time.Now()
	var err error
	var result mProduct.Response
	var dataReq mProduct.Product

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&dataReq)

	result.Data = make([]mProduct.Data, 0)
	name := strings.TrimSpace(dataReq.Name)
	description := strings.TrimSpace(dataReq.Description)
	warehouse_id := strings.TrimSpace(dataReq.WarehouseId)

	if err != nil {
		result.Header.Message = err.Error()
		result.Header.StatusCode = 400
	}

	if strings.TrimSpace(name) == "" {
		result.Header.Message = "Name must be filled out"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return

	} else if dataReq.Price <= 0 {
		result.Header.Message = "Price must be filled out"
		result.Header.StatusCode = 400
		result.Header.ProcessTime = time.Since(startTime).Seconds() * 1000
		cHttp.Render(w, result, 0, req.FormValue("callback"))
		return
	}

	_, _ = d.productUC.AddProduct(context.Background(), mProduct.Product{
		Name:        name,
		Price:       dataReq.Price,
		Stock:       dataReq.Stock,
		Description: description,
		WarehouseId: warehouse_id,
	})

	result.Header.Message = "Success, Create product!"
	result.Header.StatusCode = 200

	cHttp.Render(w, result, 0, req.FormValue("callback"))

}
