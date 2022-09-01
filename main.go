package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/users/controllers"
	"github.com/users/middleware"
)

func main() {

	options := middleware.CORSOptions{Origin: "http://localhost:3333"}

	//creating a new gin engine instance
	engine := gin.New()

	//creating routing groups for ordinary users and admin
	user := engine.Group("/user")
	//admin := admin.Group("/admin")

	//using default recovery middleware
	engine.Use(gin.Recovery())

	//using custom CORS & Log middleware
	engine.Use(middleware.CORS(options))
	engine.Use(middleware.JSONLogMiddleware())

	//Endpoints

	user.POST("/register", controllers.Register)

	//admin.GET("/login", controllers.AdminLogin)
	srv := &http.Server{
		Addr:    "localhost:3333",
		Handler: engine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")

}
