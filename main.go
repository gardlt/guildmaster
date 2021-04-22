package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	store "baiten.io/genbo/pkg/store"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}
type healthResponse struct {
	Pg string `json:"pg"`
}

func (a *App) getHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	healthResponse := healthResponse{Pg: "UP"}

	if err := a.DB.Ping(); err != nil {
		healthResponse.Pg = "DOWN"
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&healthResponse)
}

func (a *App) initRoutes() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/actuator/health", a.getHealth).Methods(http.MethodGet)
}

func (a *App) initDB(env *Env) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", env.DBHOST, env.DBPORT, env.DBUSER, env.DBPASSWORD, env.DBNAME)
	var err error
	a.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	driver, _ := postgres.WithInstance(a.DB, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
	m.Steps(2)
}

func (a *App) Initialize(env *Env) {
	a.initDB(env)
	a.initRoutes()
	store.StoreInit(a.DB, a.Router)
}

func (a *App) Run(addr string) {
	log.Info("Starting the server")
	http.ListenAndServe(addr, a.Router)
}

type Env struct {
	DBHOST     string
	DBUSER     string
	DBPASSWORD string
	DBNAME     string
	DBPORT     int
}

func main() {
	a := App{}

	env := Env{
		DBHOST:     "db",
		DBPORT:     5432,
		DBUSER:     os.Getenv("DBUSER"),
		DBPASSWORD: os.Getenv("DBPASSWORD"),
		DBNAME:     os.Getenv("DBNAME"),
	}

	a.Initialize(&env)
	a.Run(":8000")
}
