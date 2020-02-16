package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	"sample_rest/config"
	"sample_rest/handler"
	"sample_rest/model"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {

	db, err := gorm.Open("postgres", "postgres://tqurgkblyqvifi:a2f89ec18e50140bf497bf3da06b4254481e93af5ad4217335de5e6b5ae25fdc@ec2-54-195-252-243.eu-west-1.compute.amazonaws.com:5432/d1bb15eo4hprtf")
	//db, err := gorm.Open("sqlite3", dir+"/gorm.db")
	//db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err.Error())
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wake up, Neo")
}

// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/", homeLink)
	a.Get("/tasks", a.GetAllTasks)
	a.Post("/task", a.CreateTask)
	a.Get("/task/{title}", a.GetTask)
	a.Put("/task/{title}", a.UpdateTask)
	a.Delete("/task/{title}", a.DeleteTask)
	a.Put("/task/{title}/disable", a.DisableTask)
	a.Put("/task/{title}/enable", a.EnableTask)
	a.Get("/info", a.GetInfo)
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
func (a *App) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	handler.GetAllTasks(a.DB, w, r)
}

func (a *App) CreateTask(w http.ResponseWriter, r *http.Request) {
	handler.CreateTask(a.DB, w, r)
}

func (a *App) GetTask(w http.ResponseWriter, r *http.Request) {
	handler.GetTask(a.DB, w, r)
}

func (a *App) UpdateTask(w http.ResponseWriter, r *http.Request) {
	handler.UpdateTask(a.DB, w, r)
}

func (a *App) DeleteTask(w http.ResponseWriter, r *http.Request) {
	handler.DeleteTask(a.DB, w, r)
}

func (a *App) DisableTask(w http.ResponseWriter, r *http.Request) {
	handler.DisableTask(a.DB, w, r)
}

func (a *App) EnableTask(w http.ResponseWriter, r *http.Request) {
	handler.EnableTask(a.DB, w, r)
}

func (a *App) GetInfo(w http.ResponseWriter, r *http.Request) {
	handler.GetInfo(a.DB, w, r)
}
