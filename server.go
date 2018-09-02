package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
  Name  string `json:"name" xml:"name"`
  Email string `json:"email" xml:"email"`
}

type Message struct {
  Body  string `json:"body" xml:"body"`
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		msg := &Message {
			Body:  "HELLO!",
		}
		return c.JSON(http.StatusOK, msg)
	})

	e.GET("/user", func(c echo.Context) error {
		msg := &User{
			Name:  "John Mitsch",
			Email: "johnpmitsch@gmail.com",
		}
		return c.JSON(http.StatusOK, msg)
	})

  e.Logger.Fatal(e.Start(":1323"))
}
