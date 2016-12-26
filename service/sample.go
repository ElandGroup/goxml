package service

import (
	"goLog/logs"
	"net/http"

	"github.com/labstack/echo"

	"encoding/xml"
)

type (
	APIResult struct {
		Result  interface{} `json:"result"`
		Success bool        `json:"success"`
		Error   APIError    `json:"error"`
	}

	APIError struct {
		Code    int         `json:"code"`
		Details interface{} `json:"details"`
		Message string      `json:"message"`
	}
)

type User struct {
	Name string
	Sex  string
}

func GetUser(c echo.Context) error {
	u := User{Name: "xiaomiao", Sex: "Male"}
	// user.Name = "xiaomiao"
	// user.Sex = "Male"
	buf, _ := xml.MarshalIndent(u, "", " ")
	logs.Debug.Println(string(buf))

	u2 := User{}
	xml.Unmarshal(buf, &u2)
	logs.Debug.Printf("%+v\n", u2)

	return c.JSON(http.StatusOK, APIResult{Success: true, Result: u})

}

func (u User) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	uu := struct {
		Xmlns     string `xml:"xmlns,attr"`
		Name, Sex string
	}{
		"http://www.w3.org/2001/XMLSchema-instance",
		u.Name,
		u.Sex,
	}
	return e.EncodeElement(uu, start)
}

func (u *User) UmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	uu := struct {
		Xmlns     string `xml:"xmlns,attr`
		Name, Sex string
	}{}
	if err := d.DecodeElement(&uu, &start); err != nil {
		return err
	}
	*u = User{uu.Name, uu.Sex}
	return nil
}
