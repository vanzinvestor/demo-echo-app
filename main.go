package main

import (
	"html/template"
	"io"

	"github.com/gtongy/demo-echo-app/handlers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.tpl")),
	}
	e.Renderer = renderer

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/css", "./assets/css")

	e.GET("/login", handlers.User.Login)
	e.GET("/register", handlers.User.Register)
	e.POST("/user/create", handlers.User.Create)

	e.Logger.Fatal(e.Start(":1323"))
}