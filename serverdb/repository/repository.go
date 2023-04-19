package repository

import (
	"dacode/http/serverdb/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func UsuarioID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:rootroot@/cursogo")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var usuario model.Usuario

	db.QueryRow("SELECT id, nome FROM usuarios where id = ?", id).Scan(&usuario.ID, &usuario.Nome)
	json, _ := json.Marshal(usuario)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
func ListarTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:rootroot@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("SELECT id, nome FROM usuarios ")
	defer db.Close()
	var usuarios []model.Usuario
	for rows.Next() {
		var usuario model.Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)

	}
	json, _ := json.Marshal(usuarios)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}
