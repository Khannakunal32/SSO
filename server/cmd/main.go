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
	http.HandleFunc("/home", controller.Home)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/refresh", controller.Refresh)

	fmt.Println("server started on port " + portAddr)
	http.ListenAndServe(portAddr, nil)
}
