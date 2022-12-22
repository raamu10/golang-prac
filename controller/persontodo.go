package controller

import (
	"fmt"
	"net/http"
	"encoding/json"
	"todo-api/views"
	"todo-api/model"
	//"reflect"
)

func addPersonTodo() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == http.MethodPost) {
			//fmt.Println("R.body", r.Body)
			data := views.PersonInput {}
			json.NewDecoder(r.Body).Decode(&data)
			//fmt.Println("data1", data)

			if err := model.AddTodoList(data); err != nil {
				w.Write([]byte("Some error!!"))
				return
			}

			resp := views.Response{
				Code: http.StatusOK,
				Body: "Success",
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func getPersonDetails() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == http.MethodGet) {
			fmt.Println(r.URL.Query().Get("id"))

			var id string = r.URL.Query().Get("id")

			person, todoList, err := model.GetPersonData(id)

			if err != nil {
				w.Write([]byte("Some error on data fetch!!"))
				return
			}
			
			resp := views.PersonInput{
				ID: person.ID,
				Name: person.Name,
				Email: person.Email,
				TodoList: todoList,
			}

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(resp)

			fmt.Println("In getPerson----->")
			fmt.Println("person->", person)
			fmt.Println("todoList->", todoList)
			fmt.Println("person->", resp)
			
		}
	}
}