package main

import (
	"os"

	"github.com/ccasado/udemy-go-projects/cms"
)

func main() {
	p := &cms.Page{
		Title:   "Hello, world!",
		Content: "This is the body of your webpage",
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}
