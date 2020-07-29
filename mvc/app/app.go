package app

import (
	"fmt"
	"github.com/uwaifo/introserver/mvc/controllers"
	"net/http"
)


func StartApp()  {
	
	http.HandleFunc("/users", controllers.GetUser)
	
	if err := http.ListenAndServe(":8800", nil); err != nil {
		panic(err)
	}
	fmt.Println("Server started on")
	
}