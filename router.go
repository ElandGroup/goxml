package main

import (
	"goxml/service"
)

func RegisterUserApi() {
	api := e.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.GET("", service.GetUser)

}
