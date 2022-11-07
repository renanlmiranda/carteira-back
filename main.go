package main

import (
	"fmt"

	"github.com/renanlmiranda/carteira-back/routes"
)

func main() {
	fmt.Println("Iniciando server")
	routes.HandleRequest()
}
