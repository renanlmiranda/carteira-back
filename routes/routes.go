package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renanlmiranda/carteira-back/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", r))
}
