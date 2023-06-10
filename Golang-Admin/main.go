package main

import (
	"context"
	"fmt"
	connection "gola1/conection"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	ID          int
	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	Author      string
	PostDate    string
	Image       string
}

var dataBlog = []Blog{
	{
		Title:       "Hallo Title 1",
		Description: "Halo Content 1",
		Author:      "Alex",
		Image:       "franky.jpg",
	},
	{
		Title:       "Hallo Title 2",
		Description: "Halo Content 2",
		Author:      "Alexis",
		Image:       "nami.jpg",
	},
}

func main() {
	connection.DatabaseConnection()

	e := echo.New()

	e.Static("/public", "public")

	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/form-blog", formAddBlog)
	e.GET("/blog-detail/:id", blogDetail)

	e.POST("/add-blog", addBlog)
	e.POST("/blog-edit/:id", editBlog)
	e.POST("/blog-delete/:id", deleteBlog)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, description, image, start_date, end_date FROM tb_blog")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.Description, &each.Image, &each.StartDate, &each.EndDate)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.Author = "Alex"

		result = append(result, each)
	}

	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	blogs := map[string]interface{}{
		"Blogs": result,
	}

	return tmpl.Execute(c.Response(), blogs)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, title, description, image, start_date, end_date FROM tb_blog")

	var result []Blog
	for data.Next() {
		var each = Blog{}

		err := data.Scan(&each.ID, &each.Title, &each.Description, &each.Image, &each.StartDate, &each.EndDate)
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"Message": err.Error()})
		}

		each.Author = "Alex"

		result = append(result, each)
	}

	var tmpl, err = template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	blogs := map[string]interface{}{
		"Blogs": result,
	}

	return tmpl.Execute(c.Response(), blogs)
}

func formAddBlog(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/form-blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func blogDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// data := map[string]interface{}{
	// 	"Id":      id,
	// 	"Title":   "Yamada is Pro Player in Game",
	// 	"Content": "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Saepe itaque omnis repellat aliquam repellendus et! Voluptatibus corrupti ratione cupiditate. Excepturi nulla soluta sed quasi omnis blanditiis inventore aliquam dicta quae.",
	// }

	var BlogDetail = Blog{}

	for i, data := range dataBlog {
		if id == i {
			BlogDetail = Blog{
				Title:       data.Title,
				Description: data.Description,
				Author:      data.Author,
				PostDate:    data.PostDate,
				Image:       data.Image,
			}
		}
	}

	data := map[string]interface{}{
		"Blog": BlogDetail,
	}

	var tmpl, err = template.ParseFiles("views/blog-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("input-tittle")
	description := c.FormValue("input-description")
	image := c.FormValue("input-image")
	timenow := time.Now()

	println("Title : " + title)
	println("Description : " + description)
	fmt.Println("time elapse in nanoseconds:", timenow)

	var newBlog = Blog{
		Title:       title,
		Description: description,
		Author:      "Alexandria",
		Image:       image,
	}

	dataBlog = append(dataBlog, newBlog)

	fmt.Println(dataBlog)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func deleteBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index : ", id)

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func editBlog(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index : ", id)

	dataBlog = append(dataBlog[:id], dataBlog[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/form-blog")

	title := c.FormValue("input-tittle")
	description := c.FormValue("input-description")
	image := c.FormValue("input-image")
	timenow := time.Now()

	println("Title : " + title)
	println("Description : " + description)
	fmt.Println("time elapse in nanoseconds:", timenow)

	var editBlog = Blog{
		Title:       title,
		Description: description,
		Author:      "Alexandria",
		Image:       image,
	}

	dataBlog = append(dataBlog, editBlog)

	fmt.Println(dataBlog)

	return c.Redirect(http.StatusMovedPermanently, "/form-blog")
}
