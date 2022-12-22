package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"todo-api/views"
	"todo-api/model"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == http.MethodGet){
			data := views.Response{
				Code: http.StatusOK,
				Body: "Pong",
			}
			fmt.Println(data)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}

	}
}

func addTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if(r.Method == http.MethodPost) {
			data := views.AddTodo {}
			json.NewDecoder(r.Body).Decode(&data)

			if err := model.AddTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error!!"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}

func getTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if(r.Method == http.MethodGet) {
			
			data, err := model.GetTodo()
			fmt.Println("getData:", data)

			result := views.GetTodo{
				Code: http.StatusOK,
				List: data,
			}

			if err !=nil {
				w.Write([]byte("Some error!!"))
				return
			}
			
			
			json.NewEncoder(w).Encode(result)
		}
	}
}