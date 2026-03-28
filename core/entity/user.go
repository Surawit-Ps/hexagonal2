package entity

type User struct{
	Id string 
	Name string 
	LastName string 
	Age int 
	Email string 
	Tel string 
	Password string
	Role string // Admin, User
}

type UserRes struct{
	Name string 
	LastName string 
	Email string 
	Tel string 
}