package main

import (
	"log"
	"net/http"

	"github.com/rlgino/monda-todo/src/api/context/task/infraestructure/controller"
)

func main() {
	controller.Run()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
