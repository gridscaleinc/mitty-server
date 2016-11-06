package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/justinas/alice"
	_ "github.com/lib/pq"
	"mitty.co/mitty-server/app"
	"mitty.co/mitty-server/app/filters"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
	"mitty.co/mitty-server/config"
)

func getEnvVar(key, defaultVal string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		val = defaultVal
	}
	return val
}

func main() {
	envName := getEnvVar("GO_ENV", config.Development)

	err := config.SetEnvironment(envName)
	if err != nil {
		logrus.Fatal("Failed to load environment configuration")
		return
	}

	// setup databases
	postgresDB := helpers.SetupDatabase(config.CurrentSet)
	defer postgresDB.Close()
	dbmap := helpers.GetPostgres()
	models.AddTableWithName(dbmap)

	// setup filters
	mws := []alice.Constructor{
		helpers.NewHandler,
		func(handler http.Handler) http.Handler { return filters.RenderSetupHandler(envName, handler) },
		//func(handler http.Handler) http.Handler { return filters.AuthHandler(handler) },
	}

	appHandler := alice.New(mws...).Then(app.BuildRouter())

	http.Handle("/", appHandler)

	port := getEnvVar("PORT", config.DefaultServerPort)

	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)

}
