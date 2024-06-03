package main

import (
	"io"
	"text/template"
  "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
  templates * template.Template
}

func (t *Templates) Render (w io.Writer, name string, page interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, page)
}

func NewTemplate() *Templates {
  return &Templates{
    templates: template.Must(template.ParseGlob("views/*.html")),
  }
}

type PageData struct {
}

func NewPageData() (PageData) {
  return PageData{
  }
}

func main() {
  e := echo.New()
  e.Use(middleware.Logger())

  page := NewPageData()
  e.Renderer = NewTemplate()

  e.GET("/", func(c echo.Context) error {
    return c.Render(http.StatusOK, "index", page)
  })

  e.Logger.Fatal(e.Start(":42069"))
}
