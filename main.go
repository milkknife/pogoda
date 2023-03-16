package main

import (
	"fmt"
	"log"
	"flag"
	"strings"
	"encoding/json"
	"github.com/gocolly/colly"
)

func main() {
	// zachmurzenie duże, pochmurno, zachmurzenie umiarkowane, bezchmurnie
	url := flag.String("url", "https://pogoda.interia.pl/polska", "wybierz url: https://pogoda.interia.pl/lista-wojewodztw")
	flag.Parse()
	c := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(0),
		colly.MaxBodySize(72000),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	// goquery jest na podstawie https://www.w3schools.com/jquery/trysel.asp
	wc := ".weather-currently-"
	c.OnHTML(wc+"middle", func(e *colly.HTMLElement) {
		data := map[string]interface{}{
			"ikon": e.ChildText(wc+"icon-description"),
			"temp": e.ChildText(wc+"temp-strict"),
			"odczuwalna": e.ChildText(wc+"details-item.feelTemperature span"),
			"wiatr_": e.ChildAttr("img", "alt"),
			"wiatr": strings.Join(strings.Fields(e.ChildText(wc+"details-item.wind span")), " "),
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}
		fmt.Print(string(jsonData))
	})
	c.OnRequest(func(r *colly.Request) {
		// log.Println("LETZ GOOOOO", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("ITS OVERRRR", err)
	})
	c.Visit(*url)
	c.Wait()
}
