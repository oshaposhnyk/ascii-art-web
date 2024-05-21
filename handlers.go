package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/oshaposhnyk/ascii-art/ascii"
)

type Data struct {
	Errors []string
	Result string
	Text   string
}

const (
	ContentTypeHTML = "text/html; charset=utf-8"
)

func rootHandler(ctx *gin.Context) {
	buf := renderTmpl(Data{}, "home.gotmpl")
	ctx.Data(http.StatusOK, ContentTypeHTML, buf.Bytes())
}

func convertStringHandler(ctx *gin.Context) {
	text := ctx.PostForm("text")
	template := ctx.PostForm("template")
	buf := &bytes.Buffer{}
	var data Data

	config := ascii.Config{Text: text, Template: template}
	err := ascii.Run(config, buf)
	if err != nil {
		data.Errors = []string{err.Error()}
	}

	data.Result = buf.String()
	data.Text = text
	fmt.Println(data.Result)

	home := renderTmpl(data, "home.gotmpl")

	ctx.Data(http.StatusOK, ContentTypeHTML, home.Bytes())

}

func renderTmpl(dataTmpl Data, nameTmpl string) bytes.Buffer {
	buf := &bytes.Buffer{}
	path := path.Join("templates", nameTmpl)
	tmpl := template.Must(template.ParseFiles(path))
	tmpl.Execute(buf, dataTmpl)
	return *buf
}
