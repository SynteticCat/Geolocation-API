package api

import (
	"context"
	"log"
	"net/http"
	"os"

	"googlemaps.github.io/maps"

	"geolocationAPI/handlers"
)

type GetCoordinatesResponse struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func GetCoordinates(w http.ResponseWriter, r *http.Request) {

	/*достаем адрес из url запроса*/
	address := r.URL.Query().Get("address")

	/*создаем клиента с помощью API_KEY*/
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil{
		log.Fatal(err)
	}

	/*Получаем координаты по заданному адресу*/
	res, err := c.Geocode(context.Background(), &maps.GeocodingRequest{
		Address:      address,
	})
	if err != nil{
		log.Fatal(err)
	}

	/*записываем полученные данные в ответ*/
	response := GetCoordinatesResponse{
		Latitude: res[0].Geometry.Location.Lat,
		Longitude: res[0].Geometry.Location.Lng,
	}
	handlers.JsonResponse(w, response, http.StatusOK)
}
