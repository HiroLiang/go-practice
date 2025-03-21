package model

type User struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Hiro"`
	Age  int    `json:"age" example:"25"`
}
