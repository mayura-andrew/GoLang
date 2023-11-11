package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mayura-andrew/GoLang/CRYTPOMASTERS/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error){
	if len(currency) == 0 {
		return nil, errors.New("curreny empty")
	}
	res, err := http.Get(fmt.Sprintf(apiUrl, currency))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK{
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		} 
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("statud code received %v", res.StatusCode)
	}
	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil

}
