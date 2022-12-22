package evidenceBook

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/akralani/cycleanalysis/cors"
	"golang.org/x/net/websocket"
)

const evidenceBooksPath = "evidenceBooks"

func handleEvidenceBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		evidenceBookList, err := getEvidenceBookList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(evidenceBookList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var evidenceBook EvidenceBook
		err := json.NewDecoder(r.Body).Decode(&evidenceBook)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		evidenceBookID, err := insertEvidenceBook(evidenceBook)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(fmt.Sprintf(`{"evidenceBookId":%d}`, evidenceBookID)))
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleEvidenceBook(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", evidenceBooksPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	evidenceBookID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		evidenceBook, err := getEvidenceBook(evidenceBookID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if evidenceBook == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		j, err := json.Marshal(evidenceBook)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPut:
		var evidenceBook EvidenceBook
		err := json.NewDecoder(r.Body).Decode(&evidenceBook)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if *evidenceBook.EvidenceBookID != evidenceBookID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateEvidenceBook(evidenceBook)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		err := removeEvidenceBook(evidenceBookID)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	evidenceBooksHandler := http.HandlerFunc(handleEvidenceBooks)
	evidenceBookHandler := http.HandlerFunc(handleEvidenceBook)
	http.Handle("/websocket", websocket.Handler(evidenceBookSocket))
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, evidenceBooksPath), cors.Middleware(evidenceBooksHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, evidenceBooksPath), cors.Middleware(evidenceBookHandler))
}
