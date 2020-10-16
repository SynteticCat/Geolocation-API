package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"googlemaps.github.io/maps"

	"geolocationAPI/handlers"
)

type GetRouteResponse struct {
	Summary string `json:"summary"`
	Steps []*maps.Step `json:"steps"`
	Distance string `json:"distance"`
	Duration string `json:"duration"`
}

func GetRoute(w http.ResponseWriter, r *http.Request) {
	/*получаем точку отправления из параметра запроса*/
	origin := r.URL.Query().Get("origin")
	/*получаем точку прибытия из параметра запроса*/
	destination := r.URL.Query().Get("destination")

	/*создаем клиента с помощью API_KEY*/
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil{
		log.Fatal(err)
	}

	/*данные для получения маршрута*/
	points := &maps.DirectionsRequest{
		Origin:      origin,
		Destination: destination,
	}

	/*расчитываем маршрут с помощью google maps directions*/
	resp, _, err := c.Directions(context.Background(), points)
	if err != nil{
		log.Fatal(err)
	}

	/*записываем полученные данные в ответ*/
	response := GetRouteResponse{
		Summary:  resp[0].Summary,
		Steps:    resp[0].Legs[0].Steps,
		Distance: resp[0].Legs[0].Distance.HumanReadable,
		Duration: handlers.HumanizeDuration(resp[0].Legs[0].Duration),
	}

	handlers.JsonResponse(w, response, http.StatusOK)
}


