package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/RamazanZholdas/APIWithGin/databaseConn"
	"github.com/RamazanZholdas/APIWithGin/ginLogs"
	"github.com/RamazanZholdas/APIWithGin/middleware"
	"github.com/RamazanZholdas/APIWithGin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load .env file:\n", err)
	}
	databaseConn.ConnectToDB()
	databaseConn.MigrateModelToDB()
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	file := ginLogs.SetupLogOutput()
	defer file.Close()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(ginLogs.Logger())
	router.Use(middleware.SetCorsMiddleware())

	router.GET("/getAllSongs", routes.GetAllSongs)
	router.GET("/getSongById/:id", routes.GetSongById)
	router.POST("/createSong", routes.CreateSong)
	router.PUT("/updateSong/:id", routes.UpdateSong)
	router.DELETE("/deleteSong/:id", routes.DeleteSong)

	srv := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
