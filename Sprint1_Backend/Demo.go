package main

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	Type     string `json:"type,omitempty"`
	Pid      string `json:"pid,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Profile struct {
	Type      string `json:"type,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

func main() {
	fmt.Println("Welcome new user!")
	fmt.Println("Please enter your first name")
	var fName string
	fmt.Scanln(&fName)

	fmt.Println("Please enter your last name")
	var lName string
	fmt.Scanln(&lName)

	fmt.Println("Please enter your current occupation")
	var job string
	fmt.Scanln(&job)
	Professional := Profile{
		Type:      job,
		Firstname: fName,
		Lastname:  lName,
	}
	jsonBytes, _ := json.Marshal(Professional)
	fmt.Println(string(jsonBytes))
}
