package main

import (
	"fmt"
	"tutor/go-sonar-simple/pkg"
)

func main() {
	module := pkg.NewUserModule()
	user, err := module.GetUser("mrandiiw@gmail.com")
	if err != nil {
		fmt.Printf("error when calling GetUser: %v", err)
	}

	fmt.Println(user)
}
