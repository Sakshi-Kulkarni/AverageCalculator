package main

import (
 
	"log"
	"net/http"
    
	"GO-TRAINING/handler"

)

func main(){
	http.HandleFunc("/average", handler.AverageHandler)
	log.Println("server started at :9901")
	log.Fatal(http.ListenAndServe(":9901",nil))
}

 
 