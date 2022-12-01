package cycleAnalysis

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/akralani/cycleanalysis/cors"
)

const cycleAnalysisPath = "cycleAnalysis"

func handleCycleAnalysisCalculate(w http.ResponseWriter, r *http.Request) {
	cycleStartString := r.URL.Query().Get("cycleStart")
	cycleStart, err := strconv.ParseFloat(cycleStartString, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	higherHigh1String := r.URL.Query().Get("higherHigh1")
	higherHigh1, err := strconv.ParseFloat(higherHigh1String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeCSHH1 := math.Round(((higherHigh1-cycleStart)/cycleStart)*100*100) / 100

	higherLow1String := r.URL.Query().Get("higherLow1")
	higherLow1, err := strconv.ParseFloat(higherLow1String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeHH1HL1 := math.Round(((higherLow1-higherHigh1)/higherHigh1)*100*100) / 100

	higherHigh2String := r.URL.Query().Get("higherHigh2")
	higherHigh2, err := strconv.ParseFloat(higherHigh2String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeHL1HH2 := math.Round(((higherHigh2-higherLow1)/higherLow1)*100*100) / 100

	higherLow2String := r.URL.Query().Get("higherLow2")
	higherLow2, err := strconv.ParseFloat(higherLow2String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeHH2HL2 := math.Round(((higherLow2-higherHigh2)/higherHigh2)*100*100) / 100

	higherHigh3String := r.URL.Query().Get("higherHigh3")
	higherHigh3, err := strconv.ParseFloat(higherHigh3String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeHL2HH3 := math.Round(((higherHigh3-higherLow2)/higherLow2)*100*100) / 100

	higherLow3String := r.URL.Query().Get("higherLow3")
	higherLow3, err := strconv.ParseFloat(higherLow3String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeHH3HL3 := math.Round(((higherLow3-higherHigh3)/higherHigh3)*100*100) / 100

	higherHigh4String := r.URL.Query().Get("higherHigh4")
	higherHigh4, err := strconv.ParseFloat(higherHigh4String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeHL3HH4 := math.Round(((higherHigh4-higherLow3)/higherLow3)*100*100) / 100

	higherLow4String := r.URL.Query().Get("higherLow4")
	higherLow4, err := strconv.ParseFloat(higherLow4String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeHH4HL4 := math.Round(((higherLow4-higherHigh4)/higherHigh4)*100*100) / 100

	cycleHighString := r.URL.Query().Get("cycleHigh")
	cycleHigh, err := strconv.ParseFloat(cycleHighString, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeHL4CH := math.Round(((cycleHigh-higherLow4)/higherLow4)*100*100) / 100

	lowerLow1String := r.URL.Query().Get("lowerLow1")
	lowerLow1, err := strconv.ParseFloat(lowerLow1String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeCHLL1 := math.Round(((lowerLow1-cycleHigh)/cycleHigh)*100*100) / 100

	lowerHigh1String := r.URL.Query().Get("lowerHigh1")
	lowerHigh1, err := strconv.ParseFloat(lowerHigh1String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeLL1LH1 := math.Round(((lowerHigh1-lowerLow1)/lowerLow1)*100*100) / 100

	lowerLow2String := r.URL.Query().Get("lowerLow2")
	lowerLow2, err := strconv.ParseFloat(lowerLow2String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeLH1LL2 := math.Round(((lowerLow2-lowerHigh1)/lowerHigh1)*100*100) / 100

	lowerHigh2String := r.URL.Query().Get("lowerHigh2")
	lowerHigh2, err := strconv.ParseFloat(lowerHigh2String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeLL2LH2 := math.Round(((lowerHigh2-lowerLow2)/lowerLow2)*100*100) / 100

	lowerLow3String := r.URL.Query().Get("lowerLow3")
	lowerLow3, err := strconv.ParseFloat(lowerLow3String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeLH2LL3 := math.Round(((lowerLow3-lowerHigh2)/lowerHigh2)*100*100) / 100

	lowerHigh3String := r.URL.Query().Get("lowerHigh3")
	lowerHigh3, err := strconv.ParseFloat(lowerHigh3String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeLL3LH3 := math.Round(((lowerHigh3-lowerLow3)/lowerLow3)*100*100) / 100

	lowerLow4String := r.URL.Query().Get("lowerLow4")
	lowerLow4, err := strconv.ParseFloat(lowerLow4String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeLH3LL4 := math.Round(((lowerLow4-lowerHigh3)/lowerHigh3)*100*100) / 100

	lowerHigh4String := r.URL.Query().Get("lowerHigh4")
	lowerHigh4, err := strconv.ParseFloat(lowerHigh4String, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeLL4LH4 := math.Round(((lowerHigh4-lowerLow4)/lowerLow4)*100*100) / 100

	cycleEndString := r.URL.Query().Get("cycleEnd")
	cycleEnd, err := strconv.ParseFloat(cycleEndString, 64)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	changeLH4CE := math.Round(((cycleEnd-lowerHigh4)/lowerHigh4)*100*100) / 100

	result := CycleAnalysis{
		CycleStart:  cycleStart,
		HigherHigh1: higherHigh1,
		ChangeCSHH1: changeCSHH1,

		HigherLow1:   higherLow1,
		ChangeHH1HL1: changeHH1HL1,

		HigherHigh2:  higherHigh2,
		ChangeHL1HH2: changeHL1HH2,

		HigherLow2:   higherLow2,
		ChangeHH2HL2: changeHH2HL2,

		HigherHigh3:  higherHigh3,
		ChangeHL2HH3: changeHL2HH3,

		HigherLow3:   higherLow3,
		ChangeHH3HL3: changeHH3HL3,

		HigherHigh4:  higherHigh4,
		ChangeHL3HH4: changeHL3HH4,

		HigherLow4:   higherLow4,
		ChangeHH4HL4: changeHH4HL4,

		CycleHigh:   cycleHigh,
		ChangeHL4CH: changeHL4CH,

		LowerLow1:   lowerLow1,
		ChangeCHLL1: changeCHLL1,

		LowerHigh1:   lowerHigh1,
		ChangeLL1LH1: changeLL1LH1,

		LowerLow2:    lowerLow2,
		ChangeLH1LL2: changeLH1LL2,

		LowerHigh2:   lowerHigh2,
		ChangeLL2LH2: changeLL2LH2,

		LowerLow3:    lowerLow3,
		ChangeLH2LL3: changeLH2LL3,

		LowerHigh3:   lowerHigh3,
		ChangeLL3LH3: changeLL3LH3,

		LowerLow4:    lowerLow4,
		ChangeLH3LL4: changeLH3LL4,

		LowerHigh4:   lowerHigh4,
		ChangeLL4LH4: changeLL4LH4,

		CycleEnd:    cycleEnd,
		ChangeLH4CE: changeLH4CE,
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
	calculateHandler := http.HandlerFunc(handleCycleAnalysisCalculate)
	http.Handle(fmt.Sprintf("%s/%s/calculate", apiBasePath, cycleAnalysisPath), cors.Middleware(calculateHandler))
}
