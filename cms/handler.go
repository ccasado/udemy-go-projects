package cms

import (

)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:	"Go Project CMS",
		Content: "Welcome to our page",
		Posts: []*Post{
			&Post{
				Title: "Hello, World",
				Content: "Thanks for coming to the site",
				DataPublished: time.Now(),
			}
		}

	}
}