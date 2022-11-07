package routes

import (
	"log"
	"net/http"

	"github.com/renanlmiranda/carteira-back/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}