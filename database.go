package main

type ClientProfile struct {
	Id    string
	Name  string
	Email string
	Token string
}

var database = map[string]ClientProfile{
	"user1": {
		Id:    "user1",
		Name:  "John Doe",
		Email: "tzirw@example.com",
		Token: "123",
	},
	"user2": {
		Id:    "user2",
		Name:  "Jane Doe",
		Email: "ejeyd@example.com",
		Token: "456",
	},
}
