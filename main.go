package main

import (
	"fmt"
	"flag"
	"strings"
	"strconv"
	"encoding/json"
	"go.uber.org/zap"
	"github.com/gocolly/colly"
)

func main() {
	// zachmurzenie duże, pochmurno, zachmurzenie umiarkowane, bezchmurnie, częściowo słonecznie
	// parametry
	url := flag.String("url", "https://pogoda.interia.pl/polska", "wybierz url: https://pogoda.interia.pl/lista-wojewodztw")
	flag.Parse()
	// init logging
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	// init scraping
	c := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(0),
		colly.MaxBodySize(72000),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	// goquery jest na podstawie https://www.w3schools.com/jquery/trysel.asp
	wc := ".weather-currently"
	c.OnHTML(wc+wc, func(e *colly.HTMLElement) {
		wiat_, _ := strconv.Atoi(strings.Split(e.ChildText(wc+"-details-item.wind span"), " ")[0])
		temp_, _ := strconv.Atoi(e.ChildText(wc+"-temp-strict")[:1])
		odcz_, _ := strconv.Atoi(e.ChildText(wc+"-details-item.feelTemperature span")[:1])

		data := map[string]interface{}{
			"miejscowosc":	e.ChildText(wc+"-city"),
			"opis":			e.ChildText(wc+"-icon-description"),
			"temperatura":	temp_,
			"odczuwalna":	odcz_,
			"w_kiurenek":	e.ChildAttr("img", "alt"),
			"wiatr":		wiat_,
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			logger.Fatal("", zap.Error(err))
		}
		fmt.Print(string(jsonData))
	})
	// c.OnRequest(func(r *colly.Request) {
		// logger.Info("LETZ GOOOOO", r.URL)
	// })
	c.OnError(func(_ *colly.Response, err error) {
		logger.Fatal("ITS OVERRRR", zap.Error(err))
	})
	c.Visit(*url)
	c.Wait()
}
