package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)

	configService := ConfigService{
		Data: map[string]*Config{},
	}

	router.HandleFunc("/config/", configService.createConfigHandler).Methods("POST")
	router.HandleFunc("/config/", configService.getAllConfigHandler).Methods("GET")
	router.HandleFunc("/config/{id}", configService.getConfigHandler).Methods("GET")
	router.HandleFunc("/config/{id}", configService.delConfigHandler).Methods("DELETE")

	configGroupService := ConfigGroupService{
		Data: map[string]*ConfigGroup{},
	}

	router.HandleFunc("/configgroup/", configGroupService.createConfigGroupHandler).Methods("POST")
	router.HandleFunc("/configgroup/", configGroupService.getAllConfigGroupHandler).Methods("GET")
	router.HandleFunc("/configgroup/{id}", configGroupService.getConfigGroupHandler).Methods("GET")
	router.HandleFunc("/configgroup/{id}", configGroupService.delConfigGroupHandler).Methods("DELETE")

	// start server
	srv := &http.Server{Addr: "127.0.0.1:5000", Handler: router}
	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("configService shutting down ...")

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")
}
