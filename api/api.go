package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type PageResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Date        string  `json:"date"`
		MetricValue float32 `json:"metric_value"`
	} `json:"results"`
}

type DayResult struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

func NewHandler() http.Handler {
	router := mux.NewRouter()
	router.Handle("/cases", http.HandlerFunc(casesHandler))
	return router
}

func casesHandler(w http.ResponseWriter, r *http.Request) {
	pageStart := r.URL.Query()["pageStart"][0]
	pageEnd := r.URL.Query()["pageEnd"][0]

	res, err := getCases(pageStart, pageEnd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	stringified, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling result: %s", err.Error())
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	_, err = w.Write(stringified)
	if err != nil {
		fmt.Printf("Error writing result: %s", err.Error())
		w.WriteHeader(500)
		return
	}
}

func getCases(pageStart, pageEnd string) ([]DayResult, error) {
	fmt.Printf("Getting covid cases for pages: %s to %s\n", pageStart, pageEnd)

	res, err := getCasesByPage("1")
	if err != nil {
		return res, err
	}

	return res, nil
}

func getCasesByPage(page string) ([]DayResult, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.ukhsa-dashboard.data.gov.uk/themes/infectious_disease/sub_themes/respiratory/topics/COVID-19/geography_types/Nation/geographies/England/metrics/COVID-19_testing_PCRcountByDay?page=%s", page))
	if err != nil {
		fmt.Printf("Error fetching dependencies: %s", err.Error())
		return []DayResult{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s", err.Error())
		return []DayResult{}, err
	}

	var parsed PageResponse
	err = json.Unmarshal(body, &parsed)
	if err != nil {
		fmt.Printf("Error parsing response body: %s", err.Error())
		return nil, err
	}

	var result []DayResult
	for _, r := range parsed.Results {
		result = append(result, DayResult{
			Date:  r.Date,
			Count: int(r.MetricValue),
		})
	}
	return result, nil
}
