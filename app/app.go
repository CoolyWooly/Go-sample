package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sample_rest/model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"sample_rest/config"
	"sample_rest/handler"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("sqlite3", dir+"/gorm.db")
	//db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err.Error())
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	/*dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, dir)*/

	var files []string

	root := "/app"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Fprintf(w, file)
	}
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/", homeLink)
	a.Get("/employees", a.GetAllEmployees)
	a.Post("/employees", a.CreateEmployee)
	a.Get("/employees/{title}", a.GetEmployee)
	a.Put("/employees/{title}", a.UpdateEmployee)
	a.Delete("/employees/{title}", a.DeleteEmployee)
	a.Put("/employees/{title}/disable", a.DisableEmployee)
	a.Put("/employees/{title}/enable", a.EnableEmployee)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage Employee Data
func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handler.GetAllEmployees(a.DB, w, r)
}

func (a *App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.CreateEmployee(a.DB, w, r)
}

func (a *App) GetEmployee(w http.ResponseWriter, r *http.Request) {
	handler.GetEmployee(a.DB, w, r)
}

func (a *App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.UpdateEmployee(a.DB, w, r)
}

func (a *App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DeleteEmployee(a.DB, w, r)
}

func (a *App) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DisableEmployee(a.DB, w, r)
}

func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.EnableEmployee(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
