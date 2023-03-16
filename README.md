# pogoda
szybkie API do pogody w Polsce

<img alt="cirno - touhou" src="https://user-images.githubusercontent.com/111794344/225698854-6ac9b4eb-1f15-4ae0-b86f-8528052faf37.png" width="250">

## wykorzystanie

1. wybierz miejscowość na https://pogoda.interia.pl/lista-wojewodztw
2. skopiuj URL i użyj go w komendzie:

## przykłady

`go run main.go`: (domyślnie Warszawa)

``` json
{"miejscowosc":"Warszawa","odczuwalna":0,"opis":"Przeważnie słonecznie","temperatura":5,"w_kiurenek":"WNW","wiatr":18}
```

`go run main.go -url "https://pogoda.interia.pl/prognoza-szczegolowa-gdansk,cId,8048" | jq`:

``` json
{
  "miejscowosc": "Gdańsk",
  "odczuwalna": 0,
  "opis": "Przeważnie słonecznie",
  "temperatura": 2,
  "w_kiurenek": "NNW",
  "wiatr": 12
}
```
