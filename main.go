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
	//	gin.SetMode(gin.ReleaseMode)
	options := middleware.CORSOptions{Origin: "localhost:3333"}

	//creating a new gin engine instance
	engine := gin.New()

	//serving static files
	engine.LoadHTMLGlob("static/templates/*")

	//creating routing groups for ordinary users and admin
	user := engine.Group("/user")
	user.Static("/static", "./static")

	//using default recovery middleware
	engine.Use(gin.Recovery())

	//using custom CORS & Log middleware
	engine.Use(middleware.CORS(options))
	engine.Use(middleware.JSONLogMiddleware())

	//Endpoints' definitions
	user.GET("/", controllers.Homepage)
	user.POST("/register", controllers.Register)
	user.POST("/login", controllers.Login)
	user.GET("/logout", controllers.Logout)

	//Using Auth Middleware to protect viewing users' records from un-authorized users
	user.GET("/users", middleware.AuthMiddleware(), controllers.GetUsers)

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
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the requests it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")

}
