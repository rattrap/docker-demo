package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"os"
	"strings"
)

type Env struct {
	Key   string
	Value string
}

func main() {
	app := iris.New(iris.Configuration{Charset: "UTF-8"})
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())
	app.Adapt(view.HTML("./templates", ".html"))

	app.HandleFunc("GET", "/", func(ctx *iris.Context) {

		hostname, err := os.Hostname()
		if err != nil {
			hostname = ""
		}

		name, ok := os.LookupEnv("NAME")
		if !ok {
			name = "Docker"
		}

		envs := []Env{}
		for _, e := range os.Environ() {
			pair := strings.Split(e, "=")
			envs = append(envs, Env{Key: pair[0], Value: pair[1]})
		}

		ctx.Render("index.html", iris.Map{
			"Name":     name,
			"Hostname": hostname,
			"Envs":     envs,
		})
	})

	app.StaticWeb("/static", "./static")
	app.Listen(":8080")
}
