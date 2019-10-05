package main

import (
	"github.com/ourapi/routers"
	"log"
)

func main()  {
	e := routers.NewRouter()
	log.Fatal(e.Start(":1234"))
}