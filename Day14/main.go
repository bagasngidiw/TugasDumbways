package main

import (
	"context"
	"net/http"
	"personal-web/connection"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

// kalo isi struct penulisan kecil, itu private, klo huruf gede packagenya bisa di export
type project struct {
	Id           int
	ProjectName  string
	StartDate    time.Time
	EndDate      time.Time
	Duration     string
	Description  string
	Technologies []string
	NodeJs       bool
	ReactJs      bool
	Python       bool
	JavaScript   bool
}

// var dataProjects = []project{ //project -> struct biasa, kalo []project -> slice of struct
// 	// {
// 	// 	ProjectName: "Project Name 1",
// 	// 	StartDate:   "2022-01-20",
// 	// 	EndDate:     "2022-01-23",
// 	// 	Duration:    duration("2022-01-20", "2022-01-23"),
// 	// 	NodeJs:      true,
// 	// 	ReactJs:     false,
// 	// 	Python:      false,
// 	// 	JavaScript:  false,
// 	// 	Description: "Ini Deskripsi Projek 1",
// 	// },
// 	// {
// 	// 	ProjectName: "Project Name 2",
// 	// 	StartDate:   "2022-01-20",
// 	// 	EndDate:     "2022-01-27",
// 	// 	Duration:    duration("2022-01-20", "2022-01-27"),
// 	// 	NodeJs:      true,
// 	// 	ReactJs:     true,
// 	// 	Python:      false,
// 	// 	JavaScript:  false,
// 	// 	Description: "Ini Deskripsi Projek 2",
// 	// },
// 	// {
// 	// 	ProjectName: "Project Name 3",
// 	// 	StartDate:   "2022-01-20",
// 	// 	EndDate:     "2022-02-20",
// 	// 	Duration:    duration("2022-01-20", "2022-02-20"),
// 	// 	NodeJs:      true,
// 	// 	ReactJs:     true,
// 	// 	Python:      true,
// 	// 	JavaScript:  false,
// 	// 	Description: "Ini Deskripsi Projek 3",
// 	// },

// 	// {
// 	// 	ProjectName: "Project Name 4",
// 	// 	StartDate:   "2022-01-20",
// 	// 	EndDate:     "2023-01-20",
// 	// 	Duration:    duration("2022-01-20", "2023-01-20"),
// 	// 	NodeJs:      true,
// 	// 	ReactJs:     true,
// 	// 	Python:      true,
// 	// 	JavaScript:  true,
// 	// 	Description: "Ini Deskripsi Projek 4",
// 	// },
// }

func main() {
	e := echo.New()

	connection.DataBaseConnect()

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

	data, _ := connection.Conn.Query(context.Background(), "SELECT id, name, start_date, end_date, description, technologies FROM tb_project")

	var projectData []project

	for data.Next() {
		var each = project{}

		err := data.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technologies)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		each.Duration = duration(each.StartDate, each.EndDate)
		if checkValue(each.Technologies, "NodeJs") {
			each.NodeJs = true
		}
		if checkValue(each.Technologies, "ReactJs") {
			each.ReactJs = true
		}
		if checkValue(each.Technologies, "Python") {
			each.Python = true
		}
		if checkValue(each.Technologies, "JavaScript") {
			each.JavaScript = true
		}

		projectData = append(projectData, each)
	}

	projects := map[string]interface{}{
		"projects": projectData,
	}

	var tmpl, err = template.ParseFiles("./index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), projects)
	// data, _ := connection.Conn.Query(context.Background(), (`SELECT id, name, "start_date", "end_date", "description", "technologies" FROM "tb_project" order by id asc`))
	// // data, _ := connection.Conn.Query(context.Background(), ("SELECT * FROM public.\"TB_PROJECT\""))

	// // "SELECT ID, NAME, START_DATE, END_DATE, DESCRIPTION, TECHNOLOGIES, IMAGE FROM public.\"TB_PROJECT\""
	// var dataProjects []project

	// for data.Next() {
	// 	values, err := data.Values()
	// 	if err != nil {
	// 		fmt.Println("error while iterating dataset")
	// 	}

	// 	// convert DB types to Go types
	// 	id := values[0].(int32)
	// 	projectName := values[1].(string)
	// 	StartDate := values[2].(time.Time)
	// 	EndDate := values[3].(time.Time)
	// 	description := values[4].(string)
	// 	technologies := values[5].([]interface{})
	// 	// fmt.Println(projectNameEx)
	// 	// convert slice interface ke slice string
	// 	technologiesString := make([]string, len(technologies))
	// 	for i, v := range technologies {
	// 		technologiesString[i] = v.(string)
	// 	}
	// 	start, _ := time.Parse("2006-01-02", StartDate.Format("2006-01-02"))
	// 	end, _ := time.Parse("2006-01-02", EndDate.Format("2006-01-02"))
	// 	// fmt.Println(technologiesString)
	// 	newProject := project{
	// 		Id:          int(id),
	// 		ProjectName: projectName,
	// 		StartDate:   start,
	// 		EndDate:     end,
	// 		Description: description,
	// 		Duration:    duration(start, end),
	// 		NodeJs:      (checkValue(technologiesString, "NodeJs")),
	// 		Python:      (checkValue(technologiesString, "Python")),
	// 		ReactJs:     (checkValue(technologiesString, "ReactJs")),
	// 		JavaScript:  (checkValue(technologiesString, "JavaScript")),
	// 	}

	// 	isDataFound := false
	// 	for _, item := range dataProjects {
	// 		if newProject.Id == item.Id {
	// 			isDataFound = true
	// 			break
	// 		}
	// 	}
	// 	if isDataFound {
	// 		dataProjects = append(dataProjects, newProject)
	// 	}

	// 	// fmt.Println("[id:", id, ", first_name:", firstName, ", last_name:", lastName, ", date_of_birth:", dateOfBirth, "]")
	// }
	// fmt.Println(dataProjects)
	// for data.Next() {
	// 	var each = project{}

	// 	// err := data.Scan(&each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technologies)
	// 	err := data.Scan(&each.ProjectName)

	// 	if err != nil {
	// 		// return c.JSON(http.StatusInternalServerError, err.Error())
	// 		fmt.Println(err.Error())
	// 	}
	// 	// each.Duration = duration(each.StartDate, each.EndDate)
	// 	// if checkValue(each.Technologies, "nodejs") {
	// 	// 	each.NodeJs = true
	// 	// }
	// 	// if checkValue(each.Technologies, "reactjs") {
	// 	// 	each.ReactJs = true
	// 	// }
	// 	// if checkValue(each.Technologies, "python") {
	// 	// 	each.Python = true
	// 	// }
	// 	// if checkValue(each.Technologies, "javascript") {
	// 	// 	each.JavaScript = true
	// 	// }
	// 	// dataProjects = append(dataProjects, each)
	// 	fmt.Println(each.ProjectName)

	// }

	
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

	ProjectDetail := project{}

	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, description, technologies FROM tb_project WHERE id=$1", id).Scan(&ProjectDetail.Id,
		&ProjectDetail.ProjectName, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Description, &ProjectDetail.Technologies,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ProjectDetail.Duration = duration(ProjectDetail.StartDate, ProjectDetail.EndDate)

	if checkValue(ProjectDetail.Technologies, "NodeJs") {
		ProjectDetail.NodeJs = true
	}
	if checkValue(ProjectDetail.Technologies, "ReactJs") {
		ProjectDetail.ReactJs = true
	}
	if checkValue(ProjectDetail.Technologies, "Python") {
		ProjectDetail.Python = true
	}
	if checkValue(ProjectDetail.Technologies, "JavaScript") {
		ProjectDetail.JavaScript = true
	}

	// fmt.Println(dataProjects)
	template, err := template.ParseFiles("./projectDetail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// for i, data := range dataProjects {
	// 	if i == id {
	// 		ProjectDetail = project{
	// 			ProjectName: data.ProjectName,
	// 			StartDate:   data.StartDate,
	// 			EndDate:     data.EndDate,
	// 			Description: data.Description,
	// 			NodeJs:      data.NodeJs,
	// 			ReactJs:     data.ReactJs,
	// 			Python:      data.Python,
	// 			JavaScript:  data.JavaScript,
	// 			Duration:    data.Duration,
	// 		}
	// 	}
	// }

	projectDetailData := map[string]interface{}{
		"DetailProjects":  ProjectDetail,
		"startDateString": ProjectDetail.StartDate.Format("2006-01-02"),
		"endDateString":   ProjectDetail.EndDate.Format("2006-01-02"),
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

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT id, name, start_date, end_date, description, technologies FROM tb_project WHERE id=$1", id).Scan(&ProjectDetail.Id,
		&ProjectDetail.ProjectName, &ProjectDetail.StartDate, &ProjectDetail.EndDate, &ProjectDetail.Description, &ProjectDetail.Technologies,
	)

	if errQuery != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	ProjectDetail.Duration = duration(ProjectDetail.StartDate, ProjectDetail.EndDate)

	if checkValue(ProjectDetail.Technologies, "NodeJs") {
		ProjectDetail.NodeJs = true
	}
	if checkValue(ProjectDetail.Technologies, "ReactJs") {
		ProjectDetail.ReactJs = true
	}
	if checkValue(ProjectDetail.Technologies, "Python") {
		ProjectDetail.Python = true
	}
	if checkValue(ProjectDetail.Technologies, "JavaScript") {
		ProjectDetail.JavaScript = true
	}

	start := ProjectDetail.StartDate.Format("2006-01-02")
	end := ProjectDetail.EndDate.Format("2006-01-02")
	// for i, data := range dataProjects {
	// 	if i == id {
	// 		ProjectDetail = project{
	// 			ProjectName: data.ProjectName,
	// 			StartDate:   data.StartDate,
	// 			EndDate:     data.EndDate,
	// 			Description: data.Description,
	// 			NodeJs:      data.NodeJs,
	// 			ReactJs:     data.ReactJs,
	// 			Python:      data.Python,
	// 			JavaScript:  data.JavaScript,
	// 			Duration:    data.Duration,
	// 		}
	// 	}
	// }

	projectDetailData := map[string]interface{}{
		"DetailProjects": ProjectDetail,
		"StartDate":      start,
		"EndDate":        end,
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
	javacript := c.FormValue("input-javascript")
	technologies := []string{reactjs, nodejs, python, javacript}

	// start, _ := time.Parse("2006-01-02", startDate)
	// end, _ := time.Parse("2006-01-02", endDate)

	// lastId := 0
	// for _, item := range dataProjects {
	// 	lastId = item.Id + 1
	// }

	_, err := connection.Conn.Exec(context.Background(), 
		"INSERT INTO tb_project ( name, start_date, end_date, description, technologies) VALUES($1, $2, $3, $4, $5)",
		projectName, startDate, endDate, description, technologies,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// newProject := project{
	// 	Id:          lastId,
	// 	ProjectName: projectName,
	// 	StartDate:   start,
	// 	EndDate:     end,
	// 	Description: description,
	// 	Duration:    duration(start, end),
	// 	NodeJs:      (nodejs == "nodejs"),
	// 	Python:      (python == "python"),
	// 	ReactJs:     (reactjs == "reactjs"),
	// 	JavaScript:  (javascript == "javascript"),
	// }
	// dataProjects = append(dataProjects, newProject)

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

	ReactJs := c.FormValue("input-reactjs")
	NodeJs := c.FormValue("input-nodejs")
	Python := c.FormValue("input-python")
	JavaScript := c.FormValue("input-javascript")
	technologies := []string{ReactJs, NodeJs, Python, JavaScript}


	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	_, err := connection.Conn.Exec(context.Background(),
		"UPDATE tb_project SET name=$1, start_date=$2, end_date=$3, description=$4, technologies=$5 WHERE id=$6",
		projectName, start, end, description, technologies, id,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	// editProject := project{
	// 	ProjectName: projectName,
	// 	StartDate:   start,
	// 	EndDate:     end,
	// 	Description: description,
	// 	Duration:    duration(start, end),
	// 	NodeJs:      (nodejs == "nodejs"),
	// 	Python:      (python == "python"),
	// 	ReactJs:     (reactjs == "reactjs"),
	// 	JavaScript:  (javascript == "javascript"),
	// }
	// dataProjects[id] = editProject

	// fmt.Println(c.Param("id"))
	// fmt.Println(editProject.ProjectName)

	return c.Redirect(http.StatusMovedPermanently, "/")

}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := connection.Conn.Exec(context.Background(), "DELETE FROM tb_project WHERE id=$1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func duration(date1 time.Time, date2 time.Time) string {
	//apa maksudnya 2006-01-02? jadi Go memiliki standar layout format yang cukup unik.
	//Go menggunakan 2006 untuk parsing tahun 4 digit, bukan YYYY
	//01 untuk parsing bulan 2 digit
	//02 untuk parsing hari 2 digit
	//selengkapnya bisa dibaca berikut : https://dasarpemrogramangolang.novalagung.com/A-time-parsing-format.html

	difference := date2.Sub(date1)       //sub adalah subtract yg artinya mengurangi
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

// berfungsi untuk memeriksa apakah suatu nilai (string) terdapat dalam sebuah slice
func checkValue(slice []string, object string) bool {
	for _, data := range slice {
		if data == object {
			return true
		}
	}
	return false
}
