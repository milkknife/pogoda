# pogoda
szybkie API do pogody w Polsce

## wykorzystanie

1. wybierz miejscowość na https://pogoda.interia.pl/lista-wojewodztw
2. skopiuj URL i użyj go w komendzie (przykłady w "przykłady")

## przykłady

`go run main.go`: (domyślnie Warszawa)

``` json
{"ikon":"Bezchmurnie","odczuwalna":"-5°C","temp":"-2°C","wiatr":"11 km/h","wiatr_":"W"}
```

`go run main.go -url "https://pogoda.interia.pl/prognoza-szczegolowa-gdansk,cId,8048" | jq`:

``` json
{
  "ikon": "Zachmurzenie małe",
  "odczuwalna": "-4°C",
  "temp": "-2°C",
  "wiatr": "8 km/h",
  "wiatr_": "W"
}
```

