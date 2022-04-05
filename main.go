package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	routeOperator := RouteOperator{ensOperator: NewEnsOperator()}
	router := httprouter.New()
	router.GET("/get-name/:address", routeOperator.getName)
	router.GET("/get-address/:name", routeOperator.getAddress)
	router.GET("/get-avatar/:name", routeOperator.getAvatar)
	router.GET("/get-all/:param", routeOperator.getAll)
	log.Println("Serving now on http://localhost:1488.")
	log.Fatal(http.ListenAndServe(":1488", router))
}
