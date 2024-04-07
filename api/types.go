package api

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/fatih/color"
)

type SymbolSearchResponse struct {
	BestMatches []MatchData `json:"bestMatches"`
}

type TopGainersLosersResponse struct {
	Metadata string    	 		 `json:"metadata"`
	LastUpdated string 	 		 `json:"last_updated"`
	TopGainers []TopData 		 `json:"top_gainers"`
	TopLosers []TopData  		 `json:"top_losers"`
	MostActivelyTraded []TopData `json:"most_actively_traded"`
}

type TimeSeriesResponse struct {
	Metadata         TimeSeriesMetadata        `json:"Meta Data"`
	TimeSeriesDaily  map[string]TimeSeriesData `json:"Time Series (Daily)"`
}

type TimeSeriesMetadata struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"4. Time Zone"`
}

type TimeSeriesData struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume uint64  `json:"5. volume,string"`
}

type MatchData struct {
	Symbol string		`json:"1. symbol"`
	Name string			`json:"2. name"`
	Type string			`json:"3. type"`
	Region string		`json:"4. region"`
	MarketOpen string	`json:"5. marketOpen"`
	MarketClose string  `json:"6. marketClose"`
	Timezone string	   	`json:"7. timezone"`
	Currency string		`json:"8. currency"`
	MatchScore string	`json:"9. matchScore"`
}

type TopData struct {
	Ticker string			`json:"ticker"`
	Price string			`json:"price"`
	ChangeAmount string     `json:"change_amount"`
	ChangePercentage string `json:"change_percentage"`
	Volume string           `json:"volume"`
}

var errorTextFont = color.New(color.Bold, color.FgRed).SprintFunc()

func toSymbolSearchResponse(body []byte) (SymbolSearchResponse, error) {
	check, val := checkForRateLimiting(body)
	if check || (!check && val != "") {
		return SymbolSearchResponse{}, errors.New(val)
	}
	response := SymbolSearchResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return SymbolSearchResponse{}, err
	}
	return response, nil
}

func toTimeSeriesResponse(body []byte) (TimeSeriesResponse, error) {
	check, val := checkForRateLimiting(body)
	if check || (!check && val != "") {
		return TimeSeriesResponse{}, errors.New(val)
	}
	response := TimeSeriesResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return TimeSeriesResponse{}, err
	}
	return response, nil
}

func toTopGainersLosersResponse(body []byte) (TopGainersLosersResponse, error) {
	check, val := checkForRateLimiting(body)
	if check || (!check && val != "") {
		return TopGainersLosersResponse{}, errors.New(val)
	}
	response := TopGainersLosersResponse{}
	err := json.Unmarshal(body, &response)
	if err != nil {
		return TopGainersLosersResponse{}, err
	}
	return response, nil
}

func checkForRateLimiting(body []byte) (bool, string) {
	var infoResponse map[string]interface{}
	err := json.Unmarshal(body, &infoResponse)
	if err != nil {
		return false, errorTextFont(err.Error())
	}
	infoMessage, ok := infoResponse["Information"]
	if ok{
		infoMessageStr := infoMessage.(string)
		if strings.Contains(infoMessageStr, "Thank you for using Alpha Vantage!") {
			return true, errorTextFont("429 response: " + infoMessageStr)
		}
	}
	return false, ""
}