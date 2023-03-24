package main

import (
	"app/config"
	"app/infrastructures"
	userservice "app/internal/core/services/user"
	userHandler "app/internal/handlers/user"
	"app/internal/repositories"
	"app/cmd/http"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	
)

func main() {
	config.Init()
	mongo := infrastructures.NewMongoDB()
	redis := infrastructures.NewRedis()
	cache := repositories.NewCacheRepository(redis)
	userRepo := repositories.NewUserRepository(mongo,config.Get().Mongo.Database)
	userService := userservice.NewUserService(userRepo,cache)
	userHdl := userHandler.NewAuthHanderler(userService)
	hs := httphandler.NewHTTPServer(userHdl)
	
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	addr := fmt.Sprintf("%s:%d", config.Get().HTTPServer.Host, config.Get().HTTPServer.Port)
	fmt.Println(addr)
	srv := &http.Server{
		Addr:    addr,
		Handler: hs,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
