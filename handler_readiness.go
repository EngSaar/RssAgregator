package main

import "net/http"

func handlerReadiness(responseWriter http.ResponseWriter, response *http.Request) {
	respondWithJson(responseWriter, 200, struct{}{})
}
