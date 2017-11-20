package service

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}

	// configure for "/"
	mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")
	// configure for "/login"
	mx.HandleFunc("/login", loginHandler(formatter)).Methods("POST")
	// configure for "/api/test"
	mx.HandleFunc("/api/test", jsHandler(formatter)).Methods("GET")
	// configure for "/assets"
	mx.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(webRoot+"/assets/"))))
	// configure for "/unknown"
	mx.NotFoundHandler = developHandler()

}
