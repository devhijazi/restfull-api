package models

type User struct {
	FullName string
	Email    string
	Password string
	Phone    string
	Document string
	address  []Address
}

type Address struct {
	Number int
}
