package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Id int
	NameProject  string
	StartDate  string
	EndDate string
	Duration string
	PostDate time.Time
	Description string
	Html string
	Css string
	ReactJs string
	JavaScript string
}

var dataProject = []Project{
	{
		NameProject: "Game Web Store",
		StartDate:   "07-06-2023",
		EndDate:     "08-06-2023",
		Duration:    "1 month",
		Description: "Masuki alam semesta virtual yang memikat di mana para pemain game dari segala usia dan preferensi dapat menemukan pelarian yang sempurna",
		// Html: `<i class="fa-brands fa-html5"></i>`,
		// Css: `<i class="fa-brands fa-css3-alt"></i>`,
		// ReactJs: `<i class="fa-brands fa-react"></i></i>`,
		// JavaScript: `<i class="fa-brands fa-square-js"></i>`,

	},
	{
		NameProject: "NFT Web",
		StartDate:   "07-06-2023",
		EndDate:     "08-06-2023",
		Duration:    "1 month",
		Description: "Masuki alam semesta virtual yang memikat di mana para pemain game dari segala usia dan preferensi dapat menemukan pelarian yang sempurna",
	},
	{
		NameProject: "Portfolio Web",
		StartDate:   "07-06-2023",
		EndDate:     "09-06-2023",
		Duration:    "3 days",
		Description: "Masuki alam semesta virtual yang memikat di mana para pemain game dari segala usia dan preferensi dapat menemukan pelarian yang sempurna",
	},
}

func main() {
	e := echo.New()

	e.Static("/public", "public")
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/form-add-project", formAddProject)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/testimonial", testimonial)
	e.GET("/update-project/:id", updateProjectEdit)

	e.POST("/add-project", addProject)
	e.POST("/update-project/:id", updateProject)
	e.POST("/delete-project/:id", deleteProject)
	
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{}{
		"Projects": dataProject,
	}

	return tmpl.Execute(c.Response(), projects)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func formAddProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func calculateDuration(startDate, endDate string) string {
	startTime, _ := time.Parse("2006-01-02", startDate)
	endTime, _ := time.Parse("2006-01-02", endDate)

	durationTime := int(endTime.Sub(startTime).Hours())
	durationDays := durationTime / 24
	durationWeeks := durationDays / 7
	durationMonths := durationWeeks / 4
	durationYears := durationMonths / 12

	var duration string

	if durationYears > 1 {
		duration = strconv.Itoa(durationYears) + " years"
	} else if durationYears > 0 {
		duration = strconv.Itoa(durationYears) + " year"
	} else {
		if durationMonths > 1 {
			duration = strconv.Itoa(durationMonths) + " months"
		} else if durationMonths > 0 {
			duration = strconv.Itoa(durationMonths) + " month"
		} else {
			if durationWeeks > 1 {
				duration = strconv.Itoa(durationWeeks) + " weeks"
			} else if durationWeeks > 0 {
				duration = strconv.Itoa(durationWeeks) + " week"
			} else {
				if durationDays > 1 {
					duration = strconv.Itoa(durationDays) + " days"
				} else {
					duration = strconv.Itoa(durationDays) + " day"
				}
			}
		}
	}

	return duration
}

func addProject(c echo.Context) error {
	nameProject := c.FormValue("inputProjectName")
	startDate := c.FormValue("inputStartDate")
	endDate := c.FormValue("inputEndDate")
	description := c.FormValue("inputDesc")
	duration := calculateDuration(startDate, endDate)
	// html := c.FormValue("inputCheckboxHtml")
	// css := c.FormValue("inputCheckboxCss")
	// reactJs := c.FormValue("inputCheckboxReact")
	// js := c.FormValue("inputCheckboxJs")


	fmt.Println("NameProject 	:", nameProject)
	fmt.Println("Duration 		:", duration)
	fmt.Println("Description :", description)
	// fmt.Println(html)
	// fmt.Println(css)
	// fmt.Println(reactJs)
	// fmt.Println(js)

	// if html != "" {
	// 	html = `<i class="fa-brands fa-html5"></i>`
	// }
	// if css != "" {
	// 	css = `<i class="fa-brands fa-css3-alt"></i>`
	// } 
	// if reactJs != "" {
	// 	reactJs = `<i class="fa-brands fa-react"></i></i>`
	// } 
	// if js != "" {
	// 	js = `<i class="fa-brands fa-square-js"></i>`
	// } 

	var newProject = Project{
		NameProject	: nameProject,
		StartDate	: startDate,
		EndDate		: endDate,
		Duration	: duration,
		Description	: description,
		// Html		: html,
		// Css			: css,
		// ReactJs		: reactJs,
		// JavaScript	: js,

	}

	dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}



func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				NameProject	: data.NameProject,
				StartDate	: data.StartDate,
				EndDate		: data.EndDate,
				Duration	: data.Duration,
				Description	: data.Description,
				// Html		: data.Html,
				// Css			: data.Css,
				// ReactJs		: data.ReactJs,
				// JavaScript	: data.JavaScript,

			}
		}
	}

	data := map[string]interface{}{
		"Project": ProjectDetail,
	}

	// data := map[string]interface{}{
	// 	"Id":      id,
	// 	"Title":   "Game Web Store",
	// 	"Duration" : "1 bulan",
	// 	"Content": "Masuki alam semesta virtual yang memikat di mana para pemain game dari segala usia dan preferensi dapat menemukan pelarian yang sempurna. Toko web kami adalah gerbang ke koleksi game yang dirangkai dengan teliti, mulai dari petualangan penuh aksi yang mendebarkan hingga teka-teki yang membingungkan, simulasi olahraga yang menggetarkan hati, dan epik bermain peran yang mendalam. Dengan katalog yang luas yang mencakup permata indie dan judul blockbuster, toko web kami menawarkan beragam keajaiban gaming yang akan membuat Anda ketagihan berjam-jam.",
	// }

	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}


func deleteProject(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("index : ", i)

	dataProject = append(dataProject[:i], dataProject[i+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}






func updateProjectEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var ProjectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				Id: id,
				NameProject: data.NameProject,
				Description: data.Description,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Duration: data.Duration,
				// Html: data.Html,
				// Css: data.Css,
				// ReactJs: data.ReactJs,
				// JavaScript: data.JavaScript,
			}
		}
	}

	data := map[string]interface{} {
		"Project": ProjectDetail,
	}

	var tmpl, err = template.ParseFiles("views/update-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}


func updateProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index :", id)

	nameProject := c.FormValue("inputProjectName")
	startDate := c.FormValue("inputStartDate")
	endDate := c.FormValue("inputEndDate")
	description := c.FormValue("inputDesc")
	duration := calculateDuration(startDate, endDate)
	// html := c.FormValue("inputCheckboxHtml")
	// css := c.FormValue("inputCheckboxCss")
	// reactJs := c.FormValue("inputCheckboxReact")
	// js := c.FormValue("inputCheckboxJs")

	// if html != "" {
	// 	html = `<i class="fa-brands fa-html5"></i>`
	// }
	// if css != "" {
	// 	css = `<i class="fa-brands fa-css3-alt"></i>`
	// } 
	// if reactJs != "" {
	// 	reactJs = `<i class="fa-brands fa-react"></i></i>`
	// } 
	// if js != "" {
	// 	js = `<i class="fa-brands fa-square-js"></i>`
	// } 
		
		dataProject[id].NameProject = nameProject
		dataProject[id].Description = description
		dataProject[id].StartDate = startDate
		dataProject[id].EndDate = endDate
		dataProject[id].Duration = duration
		// dataProject[id].Html = html
		// dataProject[id].Css = css
		// dataProject[id].ReactJs = reactJs
		// dataProject[id].JavaScript = js

	return c.Redirect(http.StatusMovedPermanently, "/#my-project")
}	