package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

func main() {
	for {
		fmt.Println("Press Ctrl + C to quit")
		fmt.Print("Please enter your city name: ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		city := input.Text()

		if city == "" {
			city = "stockholm"
		}

		res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?q=" + city + "&key=1b12545642964a19a8a132630232005")
		panicErr(err)

		if res.StatusCode != 200 {
			panic("Weather api not available.")
		}

		body, err := io.ReadAll(res.Body)
		panicErr(err)

		var weather Weather

		err = json.Unmarshal(body, &weather)
		panicErr(err)

		location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

		fmt.Printf("\n%s, %s: %.0fC %s\n",
			location.Name,
			location.Country,
			current.TempC,
			current.Condition.Text)

		for _, hour := range hours {
			date := time.Unix(hour.TimeEpoch, 0)
			if date.Before(time.Now()) {
				continue
			}
			message := fmt.Sprintf("%s - %.0fC %.0f%% %s\n",
				date.Format("15.04"),
				hour.TempC,
				hour.ChanceOfRain,
				hour.Condition.Text)

			if hour.ChanceOfRain < 40 {
				fmt.Print(message)
			} else {
				color.Cyan(message)
			}
		}
		fmt.Println("")
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch    int64   `json:"time_epoch"`
				TempC        float64 `json:"temp_c"`
				ChanceOfRain float64 `json:"chance_of_rain"`
				Condition    struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}
