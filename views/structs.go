package views 

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

func PrintData() string {
	return "Hi!!!"
}

type AddTodo struct {
	ID   primitive.ObjectID 	`bson:"_id" json:"id,omitempty"`
	Name string 				`json:"name"`
	Todo string 				`json:"todo"`
}

type GetTodo struct {
	Code int `json:code`
	List []AddTodo `json:list`
}