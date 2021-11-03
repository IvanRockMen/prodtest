package main

type User struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
	City string `json: "city"`
}
