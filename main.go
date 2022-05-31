package main

import (
	"WebService/db"
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

	store, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	configService := postServer{
		store: store,
	}

	router.HandleFunc("/config/", configService.createConfigHandler).Methods("POST")
	router.HandleFunc("/config/", configService.getAllConfigHandler).Methods("GET")
	router.HandleFunc("/config/{id}/{version}", configService.getConfigHandler).Methods("GET")
	router.HandleFunc("/config/{id}/{version}", configService.delConfigHandler).Methods("DELETE")

	//configGroupService := db.ConfigGroupService{
	//	Data: map[string]*db.ConfigGroup{},
	//}

	//	router.HandleFunc("/configgroup/", configGroupService.createConfigGroupHandler).Methods("POST")
	//	router.HandleFunc("/configgroup/", configGroupService.getAllConfigGroupHandler).Methods("GET")
	//	router.HandleFunc("/configgroup/{id}", configGroupService.getConfigGroupHandler).Methods("GET")
	//	router.HandleFunc("/configgroup/{id}", configGroupService.delConfigGroupHandler).Methods("DELETE")

	// start server

	srv := &http.Server{Addr: "0.0.0.0:5000", Handler: router}
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
