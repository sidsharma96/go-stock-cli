package api

import (
	"net/http"
	"time"
)

var (
	alphaVantageClient = http.Client {
		Timeout: time.Second * 3,
	}
)