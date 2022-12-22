package main

import (
	"net/http"
	"todo-api/controller"
	"todo-api/model"
)

func main () {
	//Mux register
	mux := controller.Register()

	//MongoDB connector
	client := model.MongoConnector()
	//Close MongoDB
	defer model.MongoClose(client)

	//Listen to the port
	http.ListenAndServe("localhost:8060", mux)

}