package concurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// false key
const apiKey = "6fad5235c96f6fc4443e48dd6a3a3c40"

func fetchWeatherWithoutConcurrency(city string) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error fetching weather. Err: %v", err))
		return data
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(fmt.Sprintf("Error decoding weather. Err: %v", err))
		return data
	}
	return data
}

func fetchWeatherWithConcurrency(city string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	defer wg.Done()

	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error fetching weather. Err: %v", err))
		return data
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(fmt.Sprintf("Error decoding weather. Err: %v", err))
		return data
	}

	ch <- fmt.Sprintf("This is the city %s", city)
	return data
}

func TryExample() {
	now := time.Now()

	cities := []string{"vadodara", "mumbai", "sydney", "canberra"}

	//for _, c := range cities {
	//	data := fetchWeatherWithoutConcurrency(c)
	//	fmt.Println(fmt.Sprintf("This is the weather data for city %s: %v", c, data))
	//}

	ch := make(chan string)
	wg := sync.WaitGroup{}

	for _, c := range cities {
		wg.Add(1)
		go fetchWeatherWithConcurrency(c, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		fmt.Println(res)
	}

	fmt.Println("Time elapsed: ", time.Since(now))
}
