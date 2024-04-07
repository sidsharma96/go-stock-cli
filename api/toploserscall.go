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


func GetTopLosers() {
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

	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		log.Fatalln("Failed to create request to API - ", errorTextFont(err))
	}

	queryParams := req.URL.Query()
	queryParams.Add(util.Function, util.TopGainersLosers)
	queryParams.Add(util.Apikey, apiKey)
	req.URL.RawQuery = queryParams.Encode()

	resp, err := avclient.Do(req)
	if err != nil {
		log.Fatalln("Errors getting data for top gainers: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Top Gainers API returned status code: ", strconv.Itoa(resp.StatusCode))
	}

	body, ioError := io.ReadAll(resp.Body)
	if ioError != nil {
		log.Fatalln("IO error while reading body: ", err)
	}

	responseData, err:= toTopGainersLosersResponse(body)
	if err != nil {
		log.Fatalln("Failed to parse response: ", err)
	}

	if len(responseData.TopGainers) == 0 {
		log.Fatalln("No top losers found")
	}

	fmt.Println("Top 5 losers in the market are - ")
	for i := range 5 {
		time.Sleep(time.Millisecond * 300)
		topLoser := responseData.TopLosers[i]
		fmt.Printf("Ticker: %s Price: %s Change Amount: %s Change Percentage: %s \n", 
			topLoser.Ticker, topLoser.Price, topLoser.ChangeAmount, topLoser.ChangePercentage,
		)
	}
}