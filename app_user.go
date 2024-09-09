package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"

	config "github/fadlinux/edot/common/util/config"

	grace "gopkg.in/paytm/grace.v1"

	productCmd "github/fadlinux/edot/cmd/product"
	shopCmd "github/fadlinux/edot/cmd/shop"
	userCmd "github/fadlinux/edot/cmd/user"

	cHttp "github/fadlinux/edot/common/http"
)

const (
	slashHealthcheck  = "/healthcheck"
	slashUserRegister = "/user/register"
	slashUserLogin    = "/user/login"

	slashProductSearch = "/product/search"
	slashProductAdd    = "/product/"

	slashShopSearch = "/shop/search"
	slashShopAdd    = "/shop/search"
)

func init() {
	var err error

	// init config
	filePath := "files/config.toml"

	err = config.NewConfigFromFile("user", "toml", filePath, config.NewConfigOptions{
		DefaultName: "user",
	})

	if err != nil {
		log.Fatalln("Failed to init config", err)
	}

	log.Println("Init config from ", filePath)

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
}

func initRoute(router *httprouter.Router) {
	// Healthcheck
	router.HEAD(slashHealthcheck, healthcheck)
	router.GET(slashHealthcheck, healthcheck)

	//User Handlers
	router.POST(slashUserRegister, userCmd.HTTPDelivery.HandleUserRegister)
	router.POST(slashUserLogin, userCmd.HTTPDelivery.HandleUserLogin)

	//Product Handlers
	router.GET(slashProductSearch, productCmd.HTTPDelivery.HandleGet)
	router.POST(slashProductAdd, productCmd.HTTPDelivery.HandleAddProduct)

	//Shop Handlers
	router.GET(slashShopSearch, shopCmd.HTTPDelivery.HandleGet)

}

func healthcheck(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cHttp.Render(w, "success", 0, req.FormValue("callback"))
}
