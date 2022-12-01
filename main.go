package main

import (
	"log"
	"net/http"

	"github.com/akralani/cycleanalysis/cycleAnalysis"
	"github.com/akralani/cycleanalysis/primoDecrease"
	"github.com/akralani/cycleanalysis/primoIncrease"
)

const basePath = "/api"

func main() {
	primoIncrease.SetupRoutes(basePath)
	primoDecrease.SetupRoutes(basePath)
	cycleAnalysis.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
