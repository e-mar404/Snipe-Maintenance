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

type Device struct {
  AssetTag int
  SerialNumber string 
  Name string
  Company string
  Title string
  Notes string
  StartDate string
  CompletionDate string
}

func NewDevice(tag int, sn, name, company, title, notes, start, completion string) Device {
  return Device {
    AssetTag: tag,
    SerialNumber: sn,
    Name: name,
    Company: company,
    Title: title,
    Notes: notes,
    StartDate: start,
    CompletionDate: completion,
  }
}

type Devices = []Device
type MaintenanceTableData struct {
  Devices Devices 
}

func NewMaintenanceTableData() MaintenanceTableData {
  return MaintenanceTableData {
    Devices: []Device {
      NewDevice(
        1234,
        "123asd0p",
        "Test CB",
        "Campus 001",
        "Broken camera",
        "camera is blurry", 
        "2024-06-01",
        ""),

      NewDevice(
        5678, 
        "567hjk0p",
        "Test Laptop",
        "Campus 002",
        "Broken bottom cover",
        "right corner broke from hinge issue",
        "2024-06-01",
        ""),
    },
  }
}

type PageData struct {
  MaintenanceTableData MaintenanceTableData
}

func NewPageData() (PageData) {
  return PageData {
    MaintenanceTableData: NewMaintenanceTableData(),  
  }
}

func main() {
  e := echo.New()
  e.Use(middleware.Logger())

  e.Static("/css", "css")

  page := NewPageData()
  e.Renderer = NewTemplate()

  e.GET("/", func(c echo.Context) error {
    return c.Render(http.StatusOK, "index", page)
  })

  e.Logger.Fatal(e.Start(":42069"))
}
