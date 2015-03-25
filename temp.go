package gotemp

import (
	"encoding/json"
	"fmt"
	"net/url"
	"net/http"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func Now(city string) int {

	baseURL := "https://query.yahooapis.com/v1/public/yql?"
	location := strings.ToLower(city)
	
	yqlQuery := fmt.Sprintf("select item.condition.temp from weather.forecast where woeid in (select woeid from geo.places(1) where text=\"%s\")", location)

	yqlString := baseURL + "q=" + url.QueryEscape(yqlQuery) + "&format=json"

	res, err := http.Get(yqlString)
	HandleError(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(res.Body, 1048576))
	HandleError(err)

	var dat map[string]interface{}

	if err = json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	return GetTemp(dat)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetTemp(dat map[string]interface{}) int {

	temp := dat["query"].(map[string]interface{})["results"].
		(map[string]interface{})["channel"].
		(map[string]interface{})["item"].
		(map[string]interface{})["condition"].
		(map[string]interface{})["temp"].
		(string)

	result, err := strconv.Atoi(temp)
	HandleError(err)
	
	return result
}

