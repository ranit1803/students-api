package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ranit1803/students-api/internal/config"
)


func main(){
	/*
	the main things in main to do is
	1) load the config
	2) setup the database
	3) setup the router
	4) setup the server
	*/
	//Loading the config
	cfg:=config.MustLoad()

	//Setting up the router
	router:= http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to this project!"))
	})

	//Setup the Server
	server:= http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}
	slog.Info("Server Listening at",slog.String(": ",cfg.HTTPServer.Address))

	channel_shutdown:= make(chan os.Signal, 1)
	signal.Notify(channel_shutdown,os.Interrupt,syscall.SIGINT,syscall.SIGABRT,syscall.SIGTERM)
	go func ()  {
		err:= server.ListenAndServe() // without goroutine i cannot graceful shutdown
		if err!= nil{
			log.Fatal("Failed to Serve!")
	}
	}()
	<- channel_shutdown

	//Stopping the server
	slog.Info("Server Stopping..")
	ctx, cancel:= context.WithTimeout(context.Background(),time.Second * 5)
	err:= server.Shutdown(ctx)
	if err!= nil{
		slog.Error("Failed to Shut Down",slog.String("Error",err.Error()))
	}
	defer cancel()
	slog.Info("Server Shutdown Successfully")
}