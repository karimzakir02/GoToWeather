package main

import(
  "net/http"
  "html/template"
  "log"
  // "fmt"
  "github.com/antchfx/htmlquery"
  "strings"
  "strconv"
  "encoding/csv"
  "os"
  "sort"
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
  // WindDirection string
  // ChanceRain int8
  // ChanceSnow int8
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
  // windDirectionArray []string
  // chanceRainArray []int8
  // chanceSnowArray []int8
  humidityArray []int8
  visibilityArray []int16
  sunriseArray []string
  sunsetArray []string
}

func (w *weatherData) weatherChannel(link string) {
  doc, _ := htmlquery.LoadURL(link)

  var degrees int
  temperatureNodes, _ := htmlquery.QueryAll(doc, "//*[@id='WxuCurrentConditions-main-b3094163-ef75-4558-8d9a-e35e6b9b1034']/div/section/div/div[2]/div[1]/span")
  temperatureNode := temperatureNodes[0]
  degreesString := htmlquery.InnerText(temperatureNode)
  degreesTrimmed := strings.Trim(degreesString, "°")
  degrees, _ = strconv.Atoi(degreesTrimmed)
  w.temperatureArray = append(w.temperatureArray, float32(degrees))

  var highDegrees int
  highTemperatureNodes, _ := htmlquery.QueryAll(doc, "//*[@id='todayDetails']/section/div[2]/div[1]/div[2]/span[1]")
  highTemperatureNode := highTemperatureNodes[0]
  highDegreesString := htmlquery.InnerText(highTemperatureNode)
  highDegreesTrimmed := strings.Trim(highDegreesString, "°")
  highDegrees, _ = strconv.Atoi(highDegreesTrimmed)
  w.highTemperatureArray = append(w.highTemperatureArray, float32(highDegrees))

  var lowDegrees int
  lowTemperatureNodes, _ := htmlquery.QueryAll(doc, "//*[@id='todayDetails']/section/div[2]/div[1]/div[2]/span[2]")
  lowTemperatureNode := lowTemperatureNodes[0]
  lowDegreesString := htmlquery.InnerText(lowTemperatureNode)
  lowDegreesTrimmed := strings.Trim(lowDegreesString, "°")
  lowDegrees, _ = strconv.Atoi(lowDegreesTrimmed)
  w.lowTemperatureArray = append(w.lowTemperatureArray, float32(lowDegrees))

  var condition string
  conditionNodes, _ := htmlquery.QueryAll(doc, "//*[@id='WxuCurrentConditions-main-b3094163-ef75-4558-8d9a-e35e6b9b1034']/div/section/div/div[2]/div[1]/div")
  conditionNode := conditionNodes[0]
  condition = htmlquery.InnerText(conditionNode)
  w.weatherConditionArray = append(w.weatherConditionArray, condition)

  var windSpeed int
  windSpeedNodes, _ := htmlquery.QueryAll(doc, "//*[@id='todayDetails']/section/div[2]/div[2]/div[2]/span/text()")
  windSpeedNode := windSpeedNodes[0]
  windSpeedString := htmlquery.InnerText(windSpeedNode)
  windSpeedTrimmed := strings.Trim(windSpeedString, " mph")
  windSpeedTrimmed2 := strings.Trim(windSpeedTrimmed, "Wind Direction")
  windSpeed, _ = strconv.Atoi(windSpeedTrimmed2)
  w.windSpeedArray = append(w.windSpeedArray, float32(windSpeed))
  // var windDirection string
  //
  // var chanceRain int8
  //
  // var chanceSnow int8
  //
  var humidity int
  humidityNodes, _ := htmlquery.QueryAll(doc, "//*[@id='todayDetails']/section/div[2]/div[3]/div[2]/span")
  humidityNode := humidityNodes[0]
  humidityString := htmlquery.InnerText(humidityNode)
  humidityTrimmed := strings.Trim(humidityString, "%")
  humidity, _ = strconv.Atoi(humidityTrimmed)
  w.humidityArray = append(w.humidityArray, int8(humidity))

  var visibility int
  visibilityNodes, _ := htmlquery.QueryAll(doc, "//*[@id='todayDetails']/section/div[2]/div[7]/div[2]/span")
  visibilityNode := visibilityNodes[0]
  visibilityString := htmlquery.InnerText(visibilityNode)
  visibilityTrimmed := strings.Trim(visibilityString, " mi")
  visibility, _ = strconv.Atoi(visibilityTrimmed)
  w.visibilityArray = append(w.visibilityArray, int16(visibility))

  var sunrise string
  sunriseNodes, _ := htmlquery.QueryAll(doc, "//*[@id='SunriseSunsetContainer-fd88de85-7aa1-455f-832a-eacb037c140a']/div/div/div/div[1]/p")
  sunriseNode := sunriseNodes[0]
  sunrise = htmlquery.InnerText(sunriseNode)
  w.sunriseArray = append(w.sunriseArray, sunrise)

  var sunset string
  sunsetNodes, _ := htmlquery.QueryAll(doc, "//*[@id='SunriseSunsetContainer-fd88de85-7aa1-455f-832a-eacb037c140a']/div/div/div/div[2]/p")
  sunsetNode := sunsetNodes[0]
  sunset = htmlquery.InnerText(sunsetNode)
  w.sunsetArray = append(w.sunsetArray, sunset)
}

func (w *weatherData) bbcWeather(link string) {
  doc, _ := htmlquery.LoadURL(link)

  var degrees int
  temperatureNodes, _ := htmlquery.QueryAll(doc, "//*[@id='wr-forecast']/div[4]/div/div[1]/div[2]/div/div/div/div[2]/ol/li[1]/button/div[1]/div[2]/div[3]/div[2]/div/div/div[2]/span/span[3]")
  temperatureNode := temperatureNodes[0]
  degreesString := htmlquery.InnerText(temperatureNode)
  degreesTrimmed := strings.Trim(degreesString, "°")
  degrees, _ = strconv.Atoi(degreesTrimmed)
  w.temperatureArray = append(w.temperatureArray, float32(degrees))

  var highDegrees int
  highTemperatureNodes, _ := htmlquery.QueryAll(doc, "//*[@id='daylink-0']/div[4]/div[1]/div/div[4]/div/div[1]/span[2]/span/span[3]")
  highTemperatureNode := highTemperatureNodes[0]
  highDegreesString := htmlquery.InnerText(highTemperatureNode)
  highDegreesTrimmed := strings.Trim(highDegreesString, "°")
  highDegrees, _ = strconv.Atoi(highDegreesTrimmed)
  w.highTemperatureArray = append(w.highTemperatureArray, float32(highDegrees))

  var lowDegrees int
  lowTemperatureNodes, _ := htmlquery.QueryAll(doc, "//*[@id='daylink-0']/div[4]/div[1]/div/div[4]/div/div[2]/span[2]/span/span[3]")
  lowTemperatureNode := lowTemperatureNodes[0]
  lowDegreesString := htmlquery.InnerText(lowTemperatureNode)
  lowDegreesTrimmed := strings.Trim(lowDegreesString, "°")
  lowDegrees, _ = strconv.Atoi(lowDegreesTrimmed)
  w.lowTemperatureArray = append(w.lowTemperatureArray, float32(lowDegrees))

  var condition string
  conditionNodes, _ := htmlquery.QueryAll(doc, "//*[@id='daylink-0']/div[4]/div[2]/div")
  conditionNode := conditionNodes[0]
  condition = htmlquery.InnerText(conditionNode)
  w.weatherConditionArray = append(w.weatherConditionArray, condition)

  var humidity int
  humidityNodes, _ := htmlquery.QueryAll(doc, "//*[@id='wr-forecast']/div[4]/div/div[1]/div[2]/div/div/div/div[2]/ol/li[1]/button/div[2]/div/div/div[1]/dl/dd[1]")
  humidityNode := humidityNodes[0]
  humidityString := htmlquery.InnerText(humidityNode)
  humidityTrimmed := strings.Trim(humidityString, "%")
  humidity, _ = strconv.Atoi(humidityTrimmed)
  w.humidityArray = append(w.humidityArray, int8(humidity))

  var sunrise string
  sunriseNodes, _ := htmlquery.QueryAll(doc, "//*[@id='wr-forecast']/div[4]/div/div[1]/div[4]/div/div[1]/div[1]/span[1]/span[2]")
  sunriseNode := sunriseNodes[0]
  sunrise = htmlquery.InnerText(sunriseNode)
  w.sunriseArray = append(w.sunriseArray, sunrise)

  var sunset string
  sunsetNodes, _ := htmlquery.QueryAll(doc, "//*[@id='wr-forecast']/div[4]/div/div[1]/div[4]/div/div[1]/div[1]/span[2]/span[2]")
  sunsetNode := sunsetNodes[0]
  sunset = htmlquery.InnerText(sunsetNode)
  w.sunsetArray = append(w.sunsetArray, sunset)

}

func (w *weatherData) timeAndDateWeather(link string) {
  doc, _ := htmlquery.LoadURL(link)

  // Need to convert from celsius to fahrenheit
  var degrees int
  temperatureNodes, _ := htmlquery.QueryAll(doc, "//*[@id='qlook']/div[2]")
  temperatureNode := temperatureNodes[0]
  degreesString := htmlquery.InnerText(temperatureNode)
  degreesTrimmed := strings.Trim(degreesString, " °C")
  degrees, _ = strconv.Atoi(degreesTrimmed)
  w.temperatureArray = append(w.temperatureArray, float32(degrees))

  forecastNodes, _ := htmlquery.QueryAll(doc, "//*[@id='qlook']/p[2]/span[1]")
  forecastNode := forecastNodes[0]
  forecastString := htmlquery.InnerText(forecastNode)
  forecastTrimmed := strings.Trim(forecastString, "Forecast: ")
  forecastTrimmed = strings.Trim(forecastTrimmed, " °C")

  var highDegrees int
  highDegreesString := forecastTrimmed[0:2]
  highDegrees, _ = strconv.Atoi(highDegreesString)
  w.highTemperatureArray = append(w.highTemperatureArray, float32(highDegrees))

  var lowDegrees int
  lowDegreesString := forecastTrimmed[len(forecastTrimmed)-4:len(forecastTrimmed)-2]
  lowDegrees, _ = strconv.Atoi(lowDegreesString)
  w.lowTemperatureArray = append(w.lowTemperatureArray, float32(lowDegrees))

  var condition string
  conditionNodes, _ := htmlquery.QueryAll(doc, "//*[@id='qlook']/p[1]")
  conditionNode := conditionNodes[0]
  condition = htmlquery.InnerText(conditionNode)
  condition = strings.Trim(condition, ".")
  w.weatherConditionArray = append(w.weatherConditionArray, condition)

  var humidity int
  humidityNodes, _ := htmlquery.QueryAll(doc, "/html/body/div[6]/main/article/section[1]/div[2]/table/tbody/tr[6]/td")
  humidityNode := humidityNodes[0]
  humidityString := htmlquery.InnerText(humidityNode)
  humidityTrimmed := strings.Trim(humidityString, "%")
  humidity, _ = strconv.Atoi(humidityTrimmed)
  w.humidityArray = append(w.humidityArray, int8(humidity))

  // Convert from miles to km
  var visibility int
  visibilityNodes, _ := htmlquery.QueryAll(doc, "/html/body/div[6]/main/article/section[1]/div[2]/table/tbody/tr[4]/td")
  visibilityNode := visibilityNodes[0]
  visibilityString := htmlquery.InnerText(visibilityNode)
  visibilityTrimmed := strings.Trim(visibilityString, " km")
  visibility, _ = strconv.Atoi(visibilityTrimmed)
  w.visibilityArray = append(w.visibilityArray, int16(visibility))

}


func (w weatherData) getResults(city string) resultData {
  var result resultData

  result.City = city
  result.Temperature = average(w.temperatureArray)
  result.HighTemperature = average(w.highTemperatureArray)
  result.LowTemperature = average(w.lowTemperatureArray)
  result.WeatherCondition = mode(w.weatherConditionArray)
  result.WindSpeed = average(w.windSpeedArray)

  var humidityTotal int16
  humidityTotal = 0
  for _, value := range w.humidityArray {
    humidityTotal += int16(value)
  }
  humidityAverage := humidityTotal / int16(len(w.humidityArray))
  result.Humidity = int8(humidityAverage)

  var visibilityTotal int16
  visibilityTotal = 0
  for _, value := range w.visibilityArray {
    visibilityTotal += value
  }
  visibilityAverage := visibilityTotal / int16(len(w.visibilityArray))
  result.Visibility = int16(visibilityAverage)

  result.Sunrise = mode(w.sunriseArray)
  result.Sunset = mode(w.sunsetArray)

  return result
}

func average(array []float32) float32 {
  var total float32
  total = 0
  for _, value := range array {
    total += value
  }
  average := total / float32(len(array))
  return float32(average)
}

func mode(array []string) string {
  sort.Strings(array)
  max := 0
  count := 0
  current := array[0]
  currentMax := array[0]
  for _, value := range array {
    if value == current {
      count += 1
    } else {
      if count > max {
        currentMax = current
        max = count
      }
      current = value
      count = 1
    }
  }
  return currentMax
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
  weather.bbcWeather(links[2])
  weather.timeAndDateWeather(links[3])

  result := weather.getResults(city)
  return result
}
