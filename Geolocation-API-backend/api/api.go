package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	/*GetCoordinates определяет координаты адреса*/
	r.Get("/coordinates", GetCoordinates)
	/*GetLocation определяет местоположения по координатам адреса*/
	r.Get("/location", GetLocation)
	/*GetRoute определяет маршрут от точки А до точки Б*/
	r.Get("/route", GetRoute)
	return r
}
