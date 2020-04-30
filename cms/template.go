package cms

import (
	"html/template"
	"time"
)

//Tmpl is an exported variable
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

//Page struct
type Page struct {
	Title   string
	Content string
	Posts   []*Post
}

//Post struct
type Post struct {
	Title         string
	Content       string
	DatePublished time.Time
	Comments      []*Comment
}
