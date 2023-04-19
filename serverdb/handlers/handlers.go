package handlers

import (
	"dacode/http/serverdb/repository"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func UsuarioHandler(w http.ResponseWriter, r *http.Request) {

	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		repository.UsuarioID(w, r, id)
	case r.Method == "GET":
		repository.ListarTodos(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpe! Este caminho não existe")
	}

}
