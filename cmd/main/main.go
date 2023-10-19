package main

import (
	"log"
	"net/http"

	"github.com/akshatphumbhra/device-tracker/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func init() {
	err := godotenv.Load("../../pkg/config/.env")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := mux.NewRouter()
	routes.RegisterDeviceRoutes(router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"}, // Replace with your Vue.js frontend URL
		AllowedMethods:   []string{"GET", "PATCH"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe("localhost:3000", handler))
}
