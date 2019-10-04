package main

import (
	"log"
	"go-rest-gin/rest"
)

func main() {
	log.Println("Main log....")
	log.Fatal(rest.RunAPI("127.0.0.1:8000"))
}
