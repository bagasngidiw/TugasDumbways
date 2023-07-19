package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/css", "css")
	e.Static("/images", "images")
	e.Static("/js", "js")

	e.GET("/home", home)
	e.GET("/contact", contact)
	e.GET("/addProject", addProject)
	e.GET("/projectDetail/:id", projectDetail)
	e.GET("/testimonials", testimonials)

	e.POST("/addFormProject", addFormProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	template, err := template.ParseFiles("./index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return template.Execute(c.Response(), nil)
}
func contact(c echo.Context) error {
	template, err := template.ParseFiles("./contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return template.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	template, err := template.ParseFiles("addProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return template.Execute(c.Response(), nil)
}

// yg ini belum bisa
func projectDetail(c echo.Context) error {
	id := c.Param("id")

	template, err := template.ParseFiles("projectDetail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	projectDetail := map[string]interface{}{
		"Id":          id,
		"Title":       "Contoh Title Penerapan Golang",
		"Description": "Contoh Description Penerapan Golang, Keren euy!",
	}

	return template.Execute(c.Response(), projectDetail)
}

func testimonials(c echo.Context) error {
	template, err := template.ParseFiles("./testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return template.Execute(c.Response(), nil)
}

func addFormProject(c echo.Context) error {
	name := c.FormValue("input-name-project")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	description := c.FormValue("input-description")

	fmt.Println("nama :", name)
	fmt.Println("start-date :", startDate)
	fmt.Println("end-date :", endDate)
	fmt.Println("text area :", description)

	return c.Redirect(http.StatusMovedPermanently, "/addProject")
}
