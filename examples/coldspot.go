package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jochasinga/gotemp"
	mux "github.com/julienschmidt/httprouter"

)

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	
	// "/" route
	
	fmt.Fprintf(w, "<h1>Welcome to Haunt Thermos!</h1>")
}

func IsHaunt(w http.ResponseWriter, r *http.Request, ps mux.Params) {

	// "/:city/:temp/" route
	
	w.WriteHeader(http.StatusOK)

	city := strings.Replace(ps.ByName("city"), "_", " ", -1)
	temp := ps.ByName("temp")
	
	t, err := strconv.Atoi(temp)
	HandleError(err)

	_, err = io.WriteString(w, strconv.FormatBool(ColdSpot(t, city)))
	HandleError(err)

}

func ColdSpot(currentTemp int, city string) bool {
	
	// Fictitious threshold to determine cold spot
	threshold := 5

	baseTemp := gotemp.Now(city)

	if baseTemp - currentTemp >= threshold {
		return true
	}

	return false
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	router := mux.New()
	router.Handle("GET", "/", Index)
	router.Handle("GET", "/:city/:temp", IsHaunt)

	log.Fatal(http.ListenAndServe(":8080", router))
}
