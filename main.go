package main

import(
  "net/http"
  "html/template"
  "log"
  "fmt"
  "github.com/antchfx/htmlquery"
  "strings"
  "strconv"
  "encoding/csv"
  "os"
)

func main() {
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/weather", weatherHandler)
  http.ListenAndServe(":8000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  p := resultData{}
  t, _ := template.ParseFiles("static/home.html")
  t.Execute(w, p)
}

type resultData struct {
  City string
  Temperature float32
  HighTemperature float32
  LowTemperature float32
  WeatherCondition string
  WindSpeed float32
  WindDirection string
  ChanceRain int8
  ChanceSnow int8
  Humidity int8
  Visibility int16
  Sunrise string
  Sunset string
}

type weatherData struct {
  temperatureArray []float32
  highTemperatureArray []float32
  lowTemperatureArray []float32
  weatherConditionArray []string
  windSpeedArray []float32
  windDirectionArray []string
  chanceRainArray []int8
  chanceSnowArray []int8
  humidityArray []int8
  visibilityArray []int16
  sunriseArray []string
  sunsetArray []string
}

func (w *weatherData) weatherChannel(link string) {
  doc, _ := htmlquery.LoadURL(link)

  var degrees int
  temperatureNodes, _ := htmlquery.QueryAll(doc, "/html/body/div[1]/main/div[2]/main/div[1]/div/section/div/div[2]/div[1]/span")
  temperatureNode := temperatureNodes[0]
  degreesString := htmlquery.InnerText(temperatureNode)
  degreesTrimmed := strings.Trim(degreesString, "°")
  degrees, _ = strconv.Atoi(degreesTrimmed)
  weather.temperatureArray[0] = float32(degrees)

  var highDegrees int
  highTemperatureNodes, _ := htmlquery.QueryAll(doc, "/html/body/div[1]/main/div[2]/main/div[5]/section/div[2]/div[1]/div[2]/span[1]")
  highTemperatureNode := highTemperatureNodes[0]
  highDegreesString := htmlquery.InnerText(highTemperatureNode)
  highDegreesTrimmed := strings.Trim(highDegreesString, "°")
  highDegrees, _ = strconv.Atoi(highDegreesTrimmed)
  weather.highTemperatureArray[0] = float32(highDegrees)

  var lowDegrees int
  lowTemperatureNodes, _ := htmlquery.QueryAll(doc, "/html/body/div[1]/main/div[2]/main/div[5]/section/div[2]/div[1]/div[2]/span[2]")
  lowTemperatureNode := lowTemperatureNodes[0]
  lowDegreesString := htmlquery.InnerText(lowTemperatureNode)
  lowDegreesTrimmed := strings.Trim(lowDegreesString, "°")
  lowDegrees, _ = strconv.Atoi(lowDegreesTrimmed)
  weather.lowTemperatureArray[0] = float32(lowDegrees)

  var condition string
  conditionNodes, _ := htmlquery.QueryAll(doc, "/html/body/div[1]/main/div[2]/main/div[1]/div/section/div/div[2]/div[1]/div")
  conditionNode := conditionNodes[0]
  condition = htmlquery.InnerText(conditionNode)
  weather.weatherConditionArray[0] = condition

  var windSpeed int
  windSpeedNodes, _ := htmlquery.QueryAll(doc, "/html/body/div[1]/main/div[2]/main/div[5]/section/div[2]/div[2]/div[2]/span")
  windSpeedNode := windSpeedNodes[0]
  windSpeedString := htmlquery.InnerText(windSpeedNode)
  windSpeedTrimmed := strings.Trim(windSpeedString, " mph")
  windSpeedTrimmed2 := strings.Trim(windSpeedTrimmed, "Wind Direction")
  windSpeed, _ = strconv.Atoi(windSpeedTrimmed2)
  weather.windSpeedArray[0] = float32(windSpeed)
  // var windDirection string
  //
  // var chanceRain int8
  //
  // var chanceSnow int8
  //
  var humidity int
  humidityNodes, _ := htmlquery.QueryAll(doc, "/html/body/div[1]/main/div[2]/main/div[5]/section/div[2]/div[3]/div[2]/span")
  humidityNode := humidityNodes[0]
  humidityString := htmlquery.InnerText(humidityNode)
  humidityTrimmed := strings.Trim(humidityString, "%")
  humidity, _ = strconv.Atoi(humidityTrimmed)
  weather.humidityArray[0] = int8(humidity)

  var visibility int
  visibilityNodes, _ := htmlquery.QueryAll(doc, "/html/body/div[1]/main/div[2]/main/div[5]/section/div[2]/div[7]/div[2]/span")
  visibilityNode := visibilityNodes[0]
  visibilityString := htmlquery.InnerText(visibilityNode)
  visibilityTrimmed := strings.Trim(visibilityString, " mi")
  visibility, _ = strconv.Atoi(visibilityTrimmed)
  weather.visibility[0] = int16(visibility)

  var sunrise string
  sunriseNodes, _ := htmlquery.QueryAll(doc, "/html/body/div[1]/main/div[2]/main/div[5]/section/div[1]/div[2]/div/div/div/div[1]/p")
  sunriseNode := sunriseNodes[0]
  sunrise = htmlquery.InnerText(sunriseNode)
  weather.sunrise[0] = sunrise

  var sunset string
  sunsetNodes, _ := htmlquery.QueryAll(doc, "/html/body/div[1]/main/div[2]/main/div[5]/section/div[1]/div[2]/div/div/div/div[2]/p")
  sunsetNode := sunsetNodes[0]
  sunset = htmlquery.InnerText(sunsetNode)
  weather.sunset[0] = sunset
}

func (w weatherData) getResults(city string) resultData {
  var result resultData

  result.City = city
  // ...

  return result
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    log.Fatal("Error")
  }

  city := r.Form.Get("city")
  p := getWeather(city)
  t, _ := template.ParseFiles("static/weather_display.html")
  t.Execute(w, p)
}

func getWeather(city string) resultData {
  file, _ :=  os.Open("static/weather_links.csv");
  reader := csv.NewReader(file)
  records, _ := reader.ReadAll()
  var links []string
  for _, sublist := range records {
    if city == sublist[0] {
      links = sublist
    }
  }
  var weather weatherData
  weather.weatherChannel(links[1])
  // weather.accuWeather(links[2])
  // weather.weatherBug(links[3])

  result := weather.getResults()
  return result
}
