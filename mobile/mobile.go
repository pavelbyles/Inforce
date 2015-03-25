package hello

import (
	"html/template"
	"net/http"

	"appengine"
	//"appengine/datastore"
	//"appengine/memcache"
	"github.com/gorilla/mux"
	//"github.com/gorilla/schema"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles(
		"mobile/views/_layout.tmpl",
		"mobile/views/home/index.tmpl"))

	r := mux.NewRouter()

	s := r.PathPrefix("/mobile").Subrouter()
	s.HandleFunc("/", rootHandler)

	http.Handle("/mobile/", s)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	err := templates.ExecuteTemplate(w, "main", nil)

	if err != nil {
		c.Infof("Error: %v", err)
	}
}
