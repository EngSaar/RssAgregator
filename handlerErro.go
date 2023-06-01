package main

import "net/http"

func handlerErro(responseWriter http.ResponseWriter, r *http.Request) {
	respondWithErro(responseWriter, 400, "Erro interno de servidor...")
}
