package wahanthok

import (
	"log"
	"errors"
	"strings"
	"net/http"
	"html/template"
)

var (
	db Was
	p *SearchPage
	templates *template.Template
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	if err := templates.ExecuteTemplate(w, tmpl+".html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "search")
	// Reset details for the next request
	p = &SearchPage{}
	p.BackgroundImage = "images/default.jpg"
}

func search(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("word")
	if word == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	log.Println("Search for word:", word)
	if v, ok := db[strings.ToLower(word)]; ok { // The key is all in lower case
		p.W = v
		p.W.Word = word // Display the word in the same cases as user entered
		p.Msg = ""
		log.Println("Found meaning:", v)
	} else {
		p.W = Wa{word, nil}
		p.Msg = "Atoppa wahei ama thibiyo!"
		log.Println("Not found")
	}
	// Redirect to /
	http.Redirect(w, r, "/", http.StatusFound)
}

//func main() {
func init() {
	var err error
	var d *Was
	// Load the templates
	templates = template.Must(template.ParseFiles("html/search.html"))
	if templates == nil {
		log.Println("Unable to parse templates")
		panic(errors.New("Exiting application"))
	}
	// Load the db
	if d, err = LoadDB(); err != nil {
		log.Println("Unable to load db", err)
		panic(errors.New("Exiting application"))
	}
	db = *d
	// Set up default page data
	p = &SearchPage{}
	p.BackgroundImage = "images/default.jpg"
	
	// Routing
	http.HandleFunc("/", index)
	http.HandleFunc("/search", search)
	// Static files server
	http.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
    	http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/css/", func(w http.ResponseWriter, r *http.Request) {
    	http.ServeFile(w, r, r.URL.Path[1:])
	})

	//log.Println("Running wahei at port 9000")
	//http.ListenAndServe(":9000", nil)
}

