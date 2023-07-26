package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
	"time"
)

func Handler(c echo.Context) error {
	data := time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local)
	duration := int(time.Until(data).Hours() / 24)
	responseMessage := fmt.Sprintf("Days left until January 1, 2025: %d", duration)
	return c.String(http.StatusOK, responseMessage)
}

func AdminRoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		UserRole := c.Request().Header.Get("User-Role")
		if strings.Contains(strings.ToLower(UserRole), "admin") {
			log.Println("red button user detected")
		}
		err := next(c)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
}

func main() {
	fmt.Println("Server running")
	server := echo.New()
	server.GET("/", Handler, AdminRoleCheck)
	err := server.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
