package main

import (
	"log"
	"net/http"

	"github.com/akralani/cycleanalysis/database"
	"github.com/akralani/cycleanalysis/cycleAnalysis"
	"github.com/akralani/cycleanalysis/primoDecrease"
	"github.com/akralani/cycleanalysis/primoIncrease"
	"github.com/akralani/cycleanalysis/evidenceBook"
	_ "github.com/go-sql-driver/mysql"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	primoIncrease.SetupRoutes(basePath)
	primoDecrease.SetupRoutes(basePath)
	cycleAnalysis.SetupRoutes(basePath)
	evidenceBook.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
