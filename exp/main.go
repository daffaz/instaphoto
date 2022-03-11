package main

import (
	"fmt"
	"instaphoto/models"
)

func main() {
	dsn := "host=CONFIDENTAL user=CONFIDENTAL password=CONFIDENTAL dbname=CONFIDENTAL port=CONFIDENTAL sslmode=disable TimeZone=CONFIDENTAL"
	us, err := models.NewUserService(dsn)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	// // Create a user
	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@dundermifflin.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}
	// NOTE: You may need to update the query code a bit as well
	foundUser, err := us.ByEmail("michael@dundermifflin.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)
}
