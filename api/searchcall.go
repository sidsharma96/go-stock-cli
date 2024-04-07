package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sidsharma96/go-stock-cli/util"
)


func GetSearchResult(company string) {
	avclient := alphaVantageClient
	godotenv.Load(".env")

	apiKey := os.Getenv("AV_API_KEY")
	if apiKey == ""{
		log.Fatalln(errorTextFont("Api Key is missing!"))
	}

	baseUrl := os.Getenv("AV_BASE_URL")
	if baseUrl == ""{
		log.Fatalln(errorTextFont("Alpha vantage base url is missing!"))
	}

	reqSymbolSearch, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Fatalln("Failed to create request to API - ", errorTextFont(err))
	}

	reqTimeSeries, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Fatalln("Failed to create request to API - ", errorTextFont(err))
	}

	queryParams := reqSymbolSearch.URL.Query()
	queryParams.Add(util.Function, util.SymbolSearch)
	queryParams.Add(util.Keywords, company)
	queryParams.Add(util.Apikey, apiKey)
	reqSymbolSearch.URL.RawQuery = queryParams.Encode()

	respSymbolSearch, err := avclient.Do(reqSymbolSearch)
	if err != nil {
		log.Fatalln("Errors getting data for symbol search: ", errorTextFont(err))
	}
	defer respSymbolSearch.Body.Close()

	if respSymbolSearch.StatusCode != http.StatusOK {
		log.Fatalln("Symbol search API returned status code: ", strconv.Itoa(respSymbolSearch.StatusCode))
	}

	symbolSearchBody, ioError := io.ReadAll(respSymbolSearch.Body)
	if ioError != nil {
		log.Fatalln("IO error while reading body for symbol search response: ", err)
	}

	symbolSearchResponseData, err:= toSymbolSearchResponse(symbolSearchBody)
	if err != nil {
		log.Fatalln("Failed to parse response: ", err)
	}
	
	if len(symbolSearchResponseData.BestMatches) == 0 {
		log.Fatalln("Could not find a ticker for the company name provided")
	}
	symbol := symbolSearchResponseData.BestMatches[0].Symbol
	matchScore := symbolSearchResponseData.BestMatches[0].MatchScore

	fmt.Printf("Found ticker symbol %s for company %s with match score of %s \n",
		symbol, company, matchScore,
	)
	time.Sleep(time.Millisecond * 500)

	timeSeriesQueryParams := reqTimeSeries.URL.Query()
	timeSeriesQueryParams.Add(util.Function, util.TimeSeriesDaily)
	timeSeriesQueryParams.Add(util.Symbol, symbol)
	timeSeriesQueryParams.Add(util.Apikey, apiKey)
	reqTimeSeries.URL.RawQuery = timeSeriesQueryParams.Encode()

	respTimeSeries, err := avclient.Do(reqTimeSeries)
	if err != nil {
		log.Fatalln("Errors getting data for time series: ", err)
	}
	defer respTimeSeries.Body.Close()

	if respTimeSeries.StatusCode != http.StatusOK {
		log.Fatalln("Time series API returned status code: ", strconv.Itoa(respTimeSeries.StatusCode))
	}

	timeSeriesBody, ioError := io.ReadAll(respTimeSeries.Body)
	if ioError != nil {
		log.Fatalln("IO error while reading body: ", err)
	}

	timeSeriesResponseData, err := toTimeSeriesResponse(timeSeriesBody)
	if err != nil {
		log.Fatalln("Failed to parse response: ", err)
	}

	lastRefreshedDate := timeSeriesResponseData.Metadata.LastRefreshed
	openPrice := timeSeriesResponseData.TimeSeriesDaily[lastRefreshedDate].Open
	closePrice := timeSeriesResponseData.TimeSeriesDaily[lastRefreshedDate].Close
	high := timeSeriesResponseData.TimeSeriesDaily[lastRefreshedDate].High
	low := timeSeriesResponseData.TimeSeriesDaily[lastRefreshedDate].Low
	volume := timeSeriesResponseData.TimeSeriesDaily[lastRefreshedDate].Volume


	fmt.Println("Retrieving company stock information...")
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("Date for retrieval: %s \n Opening price: %v \n Closing price: %v \n High: %v \n Low: %v \n Volume: %v \n",
		lastRefreshedDate, openPrice, closePrice, high, low, volume,
	)
}