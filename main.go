package main

import(
  "net/http"
  "html/template"
  "log"
  "fmt"
)

func main() {
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/city", cityHandler)
  http.ListenAndServe(":8000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  p := someData{Number: 10, Text: "Enter your city!"}
  t, _ := template.ParseFiles("static/home.html")
  t.Execute(w, p)
}

type someData struct {
  Number int32
  Text string
}

func cityHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    log.Fatal("Error")
  }
  // fmt.Fprintln(w, "city: ", r.PostForm.Get("city"))
  // I don't think this actually works/runs lmao
  city := r.PostForm.Get("city")
  avgTemp := getWeather(city)
  fmt.Fprintln(w, "Temperature: ", avgTemp)
}

func getWeather(city string) int {
  temperature = weatherChannel(city)
  return temperature
}

func weatherChannel(city string) int {
  const weatherLink string = "https://weather.com/weather/today/l/62e0efebee1ac0e8fa9b21fd17d57a6a0001753ab6be8a4874bb78bbb52eda02"
  resp, _ := http.Get(weatherLink)
  doc, _ := goquery.NewDocumentFromReader(resp.body)
  var degrees int
  doc.Find("span").Each(func (i int, s *goquery.Selection) {
    class, _ := s.Attr("class")
    if class == "CurrentConditions--tempValue--3KcTQ" {
      degstring := s.Text()
      trimmed := strings.Trim(degstring, "°")
      degrees, _ = strconv.Atoi(trimmed)
    }
  })
  return degrees
}
