package entity

type Dogs struct{
	Id string 
	Name string
	Age uint
	Colour string
	UserID string
}

type DogRes struct{
	Name string `json:"name"`
	Age uint `json:"age"`
	UserID string `json:"user_id"`
}