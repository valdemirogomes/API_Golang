package main

import (
	"dacode/http/serverdb/handlers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/usuarios/", handlers.UsuarioHandler)
	log.Println("Port 4000")
	log.Fatal(http.ListenAndServe(":4000", nil))

}
