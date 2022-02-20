package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//router.Handle("POST", "/contact", PostContact)

	err := http.ListenAndServe("0.0.0.0:8080", router)

	if err != nil {
		log.Fatal(err)
	}
}
