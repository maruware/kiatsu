package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	simpleJson "github.com/bitly/go-simplejson"
)

func FetchApi(place string) (pressure int) {
	apiKey := os.Getenv("OPEN_WEATHER_MAP_API_KEY")

	query := url.Values{}
	query.Add("q", place)
	query.Add("APPID", apiKey)

	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather" + "?" + query.Encode())
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	js, err := simpleJson.NewJson([]byte(body))

	pressure, _ = js.Get("main").Get("pressure").Int()
	return pressure
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SavePlace(place string) {
	b := []byte(place)
	err := ioutil.WriteFile(".place", b, 0644)
	check(err)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func ReadPlace() (place string) {
	place = ""
	if FileExists(".place") {
		b, _ := ioutil.ReadFile(".place")
		place = string(b)
	}

	return
}

func RemovePlace() error {
	err := os.Remove(".place")
	return err
}

func main() {
	var saveFlag bool
	var resetFlag bool
	flag.BoolVar(&saveFlag, "save", false, "save settting option")
	flag.BoolVar(&resetFlag, "reset", false, "reset settting option")
	flag.Parse()

	if resetFlag {
		RemovePlace()
		return
	}

	place := "Tokyo"
	if flag.NArg() == 0 {
		savedPlace := ReadPlace()
		if savedPlace != "" {
			place = savedPlace
		}
	} else {
		place = flag.Arg(0)
		if saveFlag {
			SavePlace(place)
		}
	}

	pressure := FetchApi(place)
	text := fmt.Sprintf("%vhPa@%v", pressure, place)
	fmt.Println(text)
}
