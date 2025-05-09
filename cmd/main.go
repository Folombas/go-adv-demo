package main

import (
	"fmt"
	"net/http"
	"go/adv-demo/internal/hello"
)

func main() {
	router := http.NewServeMux()
	hello.NewHelloHandler(router)
	
	server := http.Server {
		Addr: ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
