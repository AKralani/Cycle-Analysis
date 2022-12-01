package primoIncrease

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/akralani/cycleanalysis/cors"
)

const primoIncreasesPath = "primoIncreases"

func handlePrimoIncreaseCalculate(w http.ResponseWriter, r *http.Request) {
	initiatedPriceString := r.URL.Query().Get("initiatedPrice")
	initiatedPrice, err := strconv.ParseFloat(initiatedPriceString, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hundredPercentIncreaseString := r.URL.Query().Get("hundredPercentIncrease")
	hundredPercentIncrease, err := strconv.ParseFloat(hundredPercentIncreaseString, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	priceIsAboveMA50String := r.URL.Query().Get("priceIsAboveMA50")
	priceIsAboveMA50, err := strconv.ParseBool(priceIsAboveMA50String)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	goldenCrossString := r.URL.Query().Get("goldenCross")
	goldenCross, err := strconv.ParseBool(goldenCrossString)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	twoHundredPercentIncrease := 2 * hundredPercentIncrease
	threeHundredPercentIncrease := 3 * hundredPercentIncrease
	fourHundredPercentIncrease := 4 * hundredPercentIncrease
	fiveHundredPercentIncrease := 5 * hundredPercentIncrease
	triggerCondition := initiatedPrice + (hundredPercentIncrease*initiatedPrice)/100
	entryPoint := (initiatedPrice + (hundredPercentIncrease*initiatedPrice)/100) + (initiatedPrice * 0.0006111)
	exitPoint1 := initiatedPrice + (twoHundredPercentIncrease*initiatedPrice)/100
	exitPoint2 := initiatedPrice + (threeHundredPercentIncrease*initiatedPrice)/100
	exitPoint3 := initiatedPrice + (fourHundredPercentIncrease*initiatedPrice)/100
	exitPoint4 := initiatedPrice + (fiveHundredPercentIncrease*initiatedPrice)/100
	average1 := (exitPoint1 + exitPoint2) / 2
	average2 := (exitPoint2 + exitPoint3) / 2
	average3 := (exitPoint3 + exitPoint4) / 2
	twoHundredProfit := exitPoint1 - entryPoint
	threeHundredProfit := exitPoint2 - entryPoint
	fourHundredProfit := exitPoint3 - entryPoint
	fiveHundredProfit := exitPoint4 - entryPoint
	average1Profit := average1 - entryPoint
	average2Profit := average2 - entryPoint
	average3Profit := average3 - entryPoint
	result := PrimoIncrease{
		InitiatedPrice:              initiatedPrice,
		HundredPercentIncrease:      hundredPercentIncrease,
		TwoHundredPercentIncrease:   twoHundredPercentIncrease,
		ThreeHundredPercentIncrease: threeHundredPercentIncrease,
		FourHundredPercentIncrease:  fourHundredPercentIncrease,
		FiveHundredPercentIncrease:  fiveHundredPercentIncrease,
		TriggerCondition:            triggerCondition,
		EntryPoint:                  entryPoint,
		ExitPoint1:                  exitPoint1,
		ExitPoint2:                  exitPoint2,
		ExitPoint3:                  exitPoint3,
		ExitPoint4:                  exitPoint4,
		Average1:                    average1,
		Average2:                    average2,
		Average3:                    average3,
		TwoHundredProfit:            twoHundredProfit,
		ThreeHundredProfit:          threeHundredProfit,
		FourHundredProfit:           fourHundredProfit,
		FiveHundredProfit:           fiveHundredProfit,
		Average1Profit:              average1Profit,
		Average2Profit:              average2Profit,
		Average3Profit:              average3Profit,
		PriceIsAboveMA50:            priceIsAboveMA50,
		GoldenCross:                 goldenCross,
	}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = w.Write(jsonResult)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	calculateHandler := http.HandlerFunc(handlePrimoIncreaseCalculate)
	http.Handle(fmt.Sprintf("%s/%s/calculate", apiBasePath, primoIncreasesPath), cors.Middleware(calculateHandler))
}
