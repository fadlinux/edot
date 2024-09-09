package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"

	grace "gopkg.in/paytm/grace.v1"

	orderCmd "github/fadlinux/edot/cmd/order"
	productCmd "github/fadlinux/edot/cmd/product"
	shopCmd "github/fadlinux/edot/cmd/shop"
	userCmd "github/fadlinux/edot/cmd/user"

	cHttp "github/fadlinux/edot/common/http"
)

const (
	slashHealthcheck = "/healthcheck"
	slashUser        = "/user/"
	slashProduct     = "/product/"
	slashShop        = "/shop/"
	slashOrder       = "/order/"
	slashWarehouse   = "/warehouse/"
)

func init() {
	// var err error

	// // init config
	// filePath := "files/config.toml"

	// err = config.NewConfigFromFile("user", "toml", filePath, config.NewConfigOptions{
	// 	DefaultName: "user",
	// })

	// if err != nil {
	// 	log.Fatalln("Failed to init config", err)
	// }

	// log.Println("Init config from ", filePath)

}

func main() {
	initAllModule()
	router := httprouter.New()
	initRoute(router)
	grace.Serve(":8001", router)
}

func initAllModule() {
	userCmd.Initialize()
	productCmd.Initialize()
	shopCmd.Initialize()
	orderCmd.Initialize()
}

func initRoute(router *httprouter.Router) {
	// Healthcheck
	router.HEAD(slashHealthcheck, healthcheck)
	router.GET(slashHealthcheck, healthcheck)

	//User Handlers
	router.POST(slashUser+"register", userCmd.HTTPDelivery.HandleUserRegister)
	router.POST(slashUser+"login", userCmd.HTTPDelivery.HandleUserLogin)

	//Product Handlers
	router.GET(slashProduct+"search", productCmd.HTTPDelivery.HandleSearch)
	router.POST(slashProduct, productCmd.HTTPDelivery.HandleAddProduct)

	//Shop Handlers
	router.GET(slashShop+"search", shopCmd.HTTPDelivery.HandleSearch)
	router.GET(slashShop, shopCmd.HTTPDelivery.HandleAddShop)

	//Order Handler
	router.POST(slashOrder+"checkout", orderCmd.HTTPDelivery.HandleAddOrder)

}

func healthcheck(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cHttp.Render(w, "success", 0, req.FormValue("callback"))
}
