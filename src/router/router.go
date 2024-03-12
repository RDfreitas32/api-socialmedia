package router

import "github.com/gorilla/mux"

//Gerar estabelacele um Router para main
func Gerar() *mux.Router {
	return mux.NewRouter()
}
