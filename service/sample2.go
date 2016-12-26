package service

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type User2 struct {
	Name string `xml:"name"`
	Sex  string `xml:"sex"`
}

func GetUser2(c echo.Context) error {
	jason := User{Name: "jason", Sex: "male"}
	buf, _ := xml.MarshalIndent(jason, "", "  ")
	fmt.Println(string(buf))

	var u User
	xml.Unmarshal(buf, &u)
	fmt.Printf("%+v", u)

	return c.JSON(http.StatusOK, APIResult{Success: true, Result: u})

}
