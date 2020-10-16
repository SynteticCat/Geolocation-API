package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"geolocationAPI/api"
)

func InitHttp() error {
	port := "8081"
	log.Println("start http on port " + port)
	return http.ListenAndServe(":"+port, NewRouter())
}

/*http роутер*/
func NewRouter() http.Handler {
	router := chi.NewRouter()
	router.Mount("/api", api.NewRouter())
	return router
}

func main() {
	/*иициализация http сервера*/
	err := InitHttp()
	if err != nil {
		log.Fatal(err)
	}
}


