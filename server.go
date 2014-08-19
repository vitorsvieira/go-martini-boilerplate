package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {

	m := martini.Classic()

	StaticOptions := martini.StaticOptions{Prefix: "public"}
	m.Use(martini.Static("public", StaticOptions))
	m.Use(martini.Static("public", StaticOptions))
	m.Use(martini.Static("public", StaticOptions))
	m.Use(martini.Static("public", StaticOptions))

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",       // Specify what path to load the templates from.
		Layout:     "layout",          // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl"}, // Specify extensions to load for templates.
		Charset:    "UTF-8",           // Sets encoding for json and html content-types. Default is "UTF-8".
	}))

	m.Get("/", IndexRouter)
	m.Get("/about", AboutRoute)
	m.Get("/contact", ContactRoute)
	m.Get("/signin", SigninRoute)
	m.Get("/signup", SignupRoute)

	m.Run()
}

func IndexRouter(r render.Render) {
	r.HTML(200, "home/index", nil)
}

func AboutRoute(r render.Render) {
	r.HTML(200, "home/about", nil)
}

func ContactRoute(r render.Render) {
	r.HTML(200, "home/contact", nil)
}

func SigninRoute(r render.Render) {
	r.HTML(200, "account/signin", nil)
}

func SignupRoute(r render.Render) {
	r.HTML(200, "account/signup", nil)
}
