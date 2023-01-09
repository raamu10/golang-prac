package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()

	// Ping_Poing api
	mux.HandleFunc("/ping", ping())

	//POST
	mux.HandleFunc("/addTodo", addTodo())

	mux.HandleFunc("/getTodo", getTodo())


	//Person with Multi TODO
	mux.HandleFunc("/addPersonTodo", addPersonTodo())
	mux.HandleFunc("/person/", getPersonDetails())
	mux.HandleFunc("/getAddress/", getAddressByPerson())
	mux.HandleFunc("/updateAddress/",updateAddressByPerson())


	return mux
}