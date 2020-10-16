package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"googlemaps.github.io/maps"

	"geolocationAPI/handlers"
)

type GetLocationResponse struct {
	Address string `json:"address"`
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	var (
		err 	error
	 	latitude float64
	 	longitude float64
	)

	/*достаем широту из параметра запроса*/
	lat := r.URL.Query().Get("lat")
	/*конвертируем в float64*/
	latitude, err = strconv.ParseFloat(lat, 64)

	/*достаем долготу из параметра запроса*/
	long := r.URL.Query().Get("long")
	/*конвертируем в float64*/
	longitude, err = strconv.ParseFloat(long, 64)

	/*создаем клиента с помощью API_KEY*/
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	/*получение адреса по координатам*/
	res, err := c.Geocode(context.Background(), &maps.GeocodingRequest{
		LatLng:       &maps.LatLng{
			Lat:   latitude,
			Lng:   longitude,
		},
	})
	if err != nil{
		log.Fatal(err)
	}

	/*записываем полученные данные в ответ*/
	response := GetLocationResponse{
		Address: res[0].FormattedAddress,
	}
	handlers.JsonResponse(w, response, http.StatusOK)
}
