package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	connectionredis "weather-api/ConnectionRedis"
	getweather "weather-api/GetWeather"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"

	"github.com/joho/godotenv"
)

// -----------------------------------------------+
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file") //|
	}
}

// -----------------------------------------------+

func main() {
	LoadEnv()
	reidsAddr := os.Getenv("REDIS_ADDR")
	ctx := context.Background()

	router := chi.NewRouter()
	rdb := connectionredis.RedisConnection(reidsAddr, ctx)

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/weather/{location}", func(w http.ResponseWriter, r *http.Request) {
		location := chi.URLParam(r, "location")
		cacheKey := strings.ToLower(location)

		cacheData, err := rdb.Get(ctx, cacheKey).Result()

		if err == redis.Nil {
			data, fetchErr := getweather.GetWeather(location)

			if fetchErr != nil {
				http.Error(w, "failed fetch weather data", http.StatusInternalServerError)
				log.Println("Error #block main", err)
				return
			}

			err = rdb.Set(ctx, cacheKey, data, 24*time.Hour).Err()
			if err != nil {
				log.Fatal(err)
				http.Error(w, "Some error in #block main2", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(data)

		} else if err != nil {
			http.Error(w, "Redis error", http.StatusInternalServerError)
			log.Println("Redis error:", err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(cacheData))
		}
	})

	http.ListenAndServe(":3000", router)

}
