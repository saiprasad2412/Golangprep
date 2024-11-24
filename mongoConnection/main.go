package main

import (
	"fmt"
	"log"
	router "main/routers"
	"net/http"
)

func main(){
	fmt.Println("MongoDB API")
	r:=router.Router()
	log.Fatal(http.ListenAndServe(":8000",r))
	fmt.Println("server started on localhost:8000")
}