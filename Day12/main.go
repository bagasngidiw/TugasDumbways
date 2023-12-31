package main

import (
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

// kalo isi struct penulisan kecil, itu private, klo huruf gede packagenya bisa di pake packagenya sm org lain
type project struct {
	ProjectName string
	StartDate   string
	EndDate     string
	Duration    string
	Description string
	NodeJs      bool
	ReactJs     bool
	Python      bool
	JavaScript  bool
}

var dataProjects = []project{ //project -> struct biasa, kalo []project -> slice of struct
	{
		ProjectName: "Project Name 1",
		StartDate:   "2022-01-20",
		EndDate:     "2022-01-23",
		Duration:    duration("2022-01-20", "2022-01-23"),
		NodeJs:      true,
		ReactJs:     false,
		Python:      false,
		JavaScript:  false,
		Description: "Ini Deskripsi Projek 1",
	},
	{
		ProjectName: "Project Name 2",
		StartDate:   "2022-01-20",
		EndDate:     "2022-01-27",
		Duration:    duration("2022-01-20", "2022-01-27"),
		NodeJs:      true,
		ReactJs:     true,
		Python:      false,
		JavaScript:  false,
		Description: "Ini Deskripsi Projek 2",
	},
	{
		ProjectName: "Project Name 3",
		StartDate:   "2022-01-20",
		EndDate:     "2022-02-20",
		Duration:    duration("2022-01-20", "2022-02-20"),
		NodeJs:      true,
		ReactJs:     true,
		Python:      true,
		JavaScript:  false,
		Description: "Ini Deskripsi Projek 3",
	},

	{
		ProjectName: "Project Name 4",
		StartDate:   "2022-01-20",
		EndDate:     "2023-01-20",
		Duration:    duration("2022-01-20", "2023-01-20"),
		NodeJs:      true,
		ReactJs:     true,
		Python:      true,
		JavaScript:  true,
		Description: "Ini Deskripsi Projek 4",
	},
}

func main() {
	e := echo.New()

	// static biar tampil css, images
	e.Static("/css", "css")
	e.Static("/images", "images")
	e.Static("/js", "js")

	// get
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/addProject", addProject)
	e.GET("/projectDetail/:id", projectDetail)
	e.GET("/testimonials", testimonials)
	e.GET("/editProject/:id", projectEdit)

	//  post
	e.POST("/", addFormProject) //ini yg di mau add project
	e.POST("/editProject/:id", editFormProject)
	e.POST("/deleteProject/:id", deleteProject)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	template, err := template.ParseFiles("./index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	projects := map[string]interface{}{
		"projects": dataProjects,
	}

	return template.Execute(c.Response(), projects)
}

func contact(c echo.Context) error {
	template, err := template.ParseFiles("./contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return template.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	template, err := template.ParseFiles("./addProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return template.Execute(c.Response(), nil)
}

func testimonials(c echo.Context) error {
	template, err := template.ParseFiles("./testimonials.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return template.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	template, err := template.ParseFiles("./projectDetail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ProjectDetail := project{}

	for i, data := range dataProjects {
		if i == id {
			ProjectDetail = project{
				ProjectName: data.ProjectName,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Description: data.Description,
				NodeJs:      data.NodeJs,
				ReactJs:     data.ReactJs,
				Python:      data.Python,
				JavaScript:  data.JavaScript,
				Duration:    data.Duration,
			}
		}
	}

	projectDetailData := map[string]interface{}{
		"DetailProjects": ProjectDetail,
	}

	return template.Execute(c.Response(), projectDetailData)
}

func projectEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	template, err := template.ParseFiles("./editProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ProjectDetail := project{}

	for i, data := range dataProjects {
		if i == id {
			ProjectDetail = project{
				ProjectName: data.ProjectName,
				StartDate:   data.StartDate,
				EndDate:     data.EndDate,
				Description: data.Description,
				NodeJs:      data.NodeJs,
				ReactJs:     data.ReactJs,
				Python:      data.Python,
				JavaScript:  data.JavaScript,
				Duration:    data.Duration,
			}
		}
	}

	projectDetailData := map[string]interface{}{
		"DetailProjects": ProjectDetail,
		"Id":             id,
	}

	// fmt.Println(id)

	return template.Execute(c.Response(), projectDetailData)
}

func addFormProject(c echo.Context) error {
	//didalem kurung formValue adalah input	name
	projectName := c.FormValue("input-name-project")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	description := c.FormValue("input-description")

	reactjs := c.FormValue("input-reactjs")
	nodejs := c.FormValue("input-nodejs")
	python := c.FormValue("input-python")
	javascript := c.FormValue("input-javascript")
	// myCars := c.Request().Form["technologies"]
	// fmt.Println( c.Request().Form["technologies"])

	// var nodejs bool //
	// if c.FormValue("nodejs") == "nodejs" {
	// 	nodejs = true
	// }
	// var python bool
	// if c.FormValue("python") == "python" {
	// 	python = true
	// }
	// var reactjs bool
	// if c.FormValue("reactjs") == "reactjs" {
	// 	reactjs = true
	// }
	// var javascript bool
	// if c.FormValue("javaScript") == "javascript" {
	// 	javascript = true
	// }

	// fmt.Println(c.FormValue("nodejs"))

	// fmt.Println("Nodejs Check  :" + nodejs	)

	// untuk checkbox blm tau caranya

	// setelah == itu adalah value dari si input
	newProject := project{
		ProjectName: projectName,
		StartDate:   startDate,
		EndDate:     endDate,
		Description: description,
		Duration:    duration(startDate, endDate),
		NodeJs:      (nodejs == "nodejs"),
		Python:      (python == "python"),
		ReactJs:     (reactjs == "reactjs"),
		JavaScript:  (javascript == "javascript"),
	}
	dataProjects = append(dataProjects, newProject)

	// fmt.Println(newProject.NodeJs)

	return c.Redirect(http.StatusMovedPermanently, "/")

}

func editFormProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	//didalem kurung formValue adalah input	name
	projectName := c.FormValue("input-name-project")
	startDate := c.FormValue("input-start-date")
	endDate := c.FormValue("input-end-date")
	description := c.FormValue("input-description")

	reactjs := c.FormValue("input-reactjs")
	nodejs := c.FormValue("input-nodejs")
	python := c.FormValue("input-python")
	javascript := c.FormValue("input-javascript")

	editProject := project{
		ProjectName: projectName,
		StartDate:   startDate,
		EndDate:     endDate,
		Description: description,
		Duration:    duration(startDate, endDate),
		NodeJs:      (nodejs == "nodejs"),
		Python:      (python == "python"),
		ReactJs:     (reactjs == "reactjs"),
		JavaScript:  (javascript == "javascript"),
	}
	dataProjects[id] = editProject

	// fmt.Println(c.Param("id"))
	// fmt.Println(editProject.ProjectName)

	return c.Redirect(http.StatusMovedPermanently, "/")

}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	dataProjects = append(dataProjects[:id], dataProjects[id+1:]...)
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func duration(date1 string, date2 string) string {
	//apa maksudnya 2006-01-02? jadi Go memiliki standar layout format yang cukup unik.
	//Go menggunakan 2006 untuk parsing tahun 4 digit, bukan YYYY
	//01 untuk parsing bulan 2 digit
	//02 untuk parsing hari 2 digit
	//selengkapnya bisa dibaca berikut : https://dasarpemrogramangolang.novalagung.com/A-time-parsing-format.html
	tgl1, _ := time.Parse("2006-01-02", date1)
	tgl2, _ := time.Parse("2006-01-02", date2)

	difference := tgl2.Sub(tgl1)         //sub adalah subtract yg artinya mengurangi
	days := int(difference.Hours() / 24) //maksudnya adalah, seteleah difference di kurangi, misal hasil pengurangannya 3 hari, maka 3 hari di convert jadi jam, yaitu 72 jam / 24 = 1 hari
	weeks := days / 7                    //misal value dari days adalah 7, maka 7/7 = 1
	months := days / 30                  //misal value dari days adalah 20, 20/30 = 0,6

	//kenapa return nya ga ditaro atas dluan? karena nnti yg bawah gabisa kebaca

	if months >= 12 {
		return strconv.Itoa(months/12) + " Year"
	}
	if months > 0 {
		return strconv.Itoa(months) + " Month"
	}
	if weeks > 0 {
		return strconv.Itoa(weeks) + " Week"
	}
	return strconv.Itoa(days) + " Day"

}
