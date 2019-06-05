package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address,omitempty"`
}

func main() {

	//jsonToStruct()

	//jsonToMap()

	//structToJson()

	//jsonToStructs()

	//jsonToMultiStructs()

	multiStructsToJson()
}


func jsonToStruct() {

	data := []byte (`{"id":1, "name":"Bob", "email":"bob@gmail.com"}`)

	var user User
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(user)
}

func jsonToMap() {
	data := []byte (`{"id":1, "name":"Bob", "email":"bob@gmail.com"}`)

	var usermap map[string]interface{}

	err := json.Unmarshal(data, &usermap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(usermap["id"])
	fmt.Println(usermap["name"])
	fmt.Println(usermap["email"])
}

func structToJson() {
	user := User{Id:2,Name:"super name",Email:"test@gmail.com"}

	userData,err := json.Marshal(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	print(string(userData))
}

func jsonToStructs() {
	usersData := []byte(`[{"id":1, "username":"Bob", "email":"bob@gmail.com"}, {"id":1, "username":"Bob", "email":"bob@gmail.com"}]`)

	var users []User

	err := json.Unmarshal(usersData, &users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(users)
}

type Err struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type AppError struct {
	Error Err `json:"error"`
}
func jsonToMultiStructs() {

	// Unmarshalling JSON into struct
	errData := []byte(`{"error":{"code":200, "message":"oops, something went wrong"}}`)

	var appError AppError

	err := json.Unmarshal(errData, &appError)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(appError)
}

func multiStructsToJson() {
	appErr := AppError{
		Error: Err{
			Code:    200,
			Message: "Some error message",
		},
	}
	errData, err := json.Marshal(appErr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(errData))

}