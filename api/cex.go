package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"cryptomasters.com/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate,error) {
	if len(currency) == 0 {
		return nil, fmt.Errorf("3 letter currency code is required; %d received", len(currency))
	}
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err 
	} 
	
	var response CEXresponse 

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)	
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil
}