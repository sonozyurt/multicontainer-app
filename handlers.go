package main

import (
	"context"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type templateData struct {
	Result         int
	PostgresResult []int
	RedisResult    []string
}
type Template struct {
	templates *template.Template
}

var data = templateData{Result: 0, PostgresResult: []int{}, RedisResult: []string{}}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func init() {
	t := &Template{
		template.Must(template.ParseGlob("templates/*"))}
	e.Renderer = t
}
func mainPage(e echo.Context) error {
	return e.Render(http.StatusOK, "mainpage.gohtml", data)
}
func postValue(e echo.Context) error {
	value := e.FormValue("value")
	valueInt, _ := strconv.Atoi(value)
	fib := calcFib(valueInt)
	cmd := client.Set(context.Background(), value, fib, 0)
	dbData := &dbData{GivenNumber: valueInt, CalcNumber: fib}

	if result := db.Where("given_number = ?", valueInt).First(&dbData); result.Error != nil {
		db.Create(&dbData)
		data.PostgresResult = append(data.PostgresResult, valueInt)
	}
	data.RedisResult = append(data.RedisResult, cmd.String())
	data.Result = fib
	return e.Redirect(http.StatusSeeOther, "/")
}
