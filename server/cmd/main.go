package main

import (
	"fmt"
	"net/http"

	controller "github.com/khannakunal32/sso/controller"
)

const portAddr = "localhost:8080"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(portAddr)
	})
	http.HandleFunc("/home", controller.HomeController)
	http.HandleFunc("/login", controller.LoginController)
	// http.HandleFunc("/refresh", controller.RefreshController)

	fmt.Println("server started on port " + portAddr)
	http.ListenAndServe(portAddr, nil)
}
