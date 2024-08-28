package main

import (
	"cartoonydesu/api"
	"cartoonydesu/database"
	"cartoonydesu/item"
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// res, err := http.Get("https://api.upcitemdb.com/prod/trial/lookup?upc=011152263373")
	// if err != nil {
	// 	log.Print(err, "\n")
	// 	os.Exit(1)
	// }
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Print(err, "\n")
	// 	os.Exit(1)
	// }
	// fmt.Print("Data ====== ", string(data))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	db := database.NewPostgres()
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Panic("Can not Ping database")
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8081"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
	}))
	item.SetRouter(r, db)
	api.SetRouter(r)
	srv := http.Server{
		Addr:        ":" + os.Getenv("PORT"),
		Handler:     r,
		ReadTimeout: 3 * time.Second,
	}
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Panic(err)
		}
	}
	slog.Info("Server shutting down...")

}

// func enableCors(w *http.ResponseWriter) {
// 	(*c).Header().Set("Access-Control-Allow-Origin", "*")
// 	// (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
// }
