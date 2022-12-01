package primoDecrease

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/akralani/cycleanalysis/cors"
)

const primoDecreasesPath = "primoDecreases"

func handlePrimoDecreaseCalculate(w http.ResponseWriter, r *http.Request) {
	initiatedPriceString := r.URL.Query().Get("initiatedPrice")
	initiatedPrice, err := strconv.ParseFloat(initiatedPriceString, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hundredPercentDecreaseString := r.URL.Query().Get("hundredPercentDecrease")
	hundredPercentDecrease, err := strconv.ParseFloat(hundredPercentDecreaseString, 64)
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
	twoHundredPercentDecrease := 2 * hundredPercentDecrease
	threeHundredPercentDecrease := 3 * hundredPercentDecrease
	fourHundredPercentDecrease := 4 * hundredPercentDecrease
	fiveHundredPercentDecrease := 5 * hundredPercentDecrease
	triggerCondition := initiatedPrice - (hundredPercentDecrease*initiatedPrice)/100
	exitPoint := triggerCondition - (triggerCondition * 0.0055)
	exitPoint1 := initiatedPrice - (twoHundredPercentDecrease*initiatedPrice)/100
	exitPoint2 := initiatedPrice - (threeHundredPercentDecrease*initiatedPrice)/100
	exitPoint3 := initiatedPrice - (fourHundredPercentDecrease*initiatedPrice)/100
	exitPoint4 := initiatedPrice - (fiveHundredPercentDecrease*initiatedPrice)/100
	average1 := (exitPoint1 + exitPoint2) / 2
	average2 := (exitPoint2 + exitPoint3) / 2
	average3 := (exitPoint3 + exitPoint4) / 2
	twoHundredLoses := exitPoint1 - exitPoint
	threeHundredLoses := exitPoint2 - exitPoint
	fourHundredLoses := exitPoint3 - exitPoint
	fiveHundredLoses := exitPoint4 - exitPoint
	average1Loses := average1 - exitPoint
	average2Loses := average2 - exitPoint
	average3Loses := average3 - exitPoint
	result := PrimoDecrease{
		InitiatedPrice:              initiatedPrice,
		HundredPercentIncrease:      hundredPercentDecrease,
		TwoHundredPercentIncrease:   twoHundredPercentDecrease,
		ThreeHundredPercentIncrease: threeHundredPercentDecrease,
		FourHundredPercentIncrease:  fourHundredPercentDecrease,
		FiveHundredPercentIncrease:  fiveHundredPercentDecrease,
		TriggerCondition:            triggerCondition,
		ExitPoint:                   exitPoint,
		ExitPoint1:                  exitPoint1,
		ExitPoint2:                  exitPoint2,
		ExitPoint3:                  exitPoint3,
		ExitPoint4:                  exitPoint4,
		Average1:                    average1,
		Average2:                    average2,
		Average3:                    average3,
		TwoHundredLoses:             twoHundredLoses,
		ThreeHundredLoses:           threeHundredLoses,
		FourHundredLoses:            fourHundredLoses,
		FiveHundredLoses:            fiveHundredLoses,
		Average1Loses:               average1Loses,
		Average2Loses:               average2Loses,
		Average3Loses:               average3Loses,
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
	calculateHandler := http.HandlerFunc(handlePrimoDecreaseCalculate)
	http.Handle(fmt.Sprintf("%s/%s/calculate", apiBasePath, primoDecreasesPath), cors.Middleware(calculateHandler))
}
