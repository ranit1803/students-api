package main

import (
	"context"
	// "fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ranit1803/students-api/internal/config"
	"github.com/ranit1803/students-api/internal/http/handlers/student"
	"github.com/ranit1803/students-api/internal/storage/mysql"
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

	//setup the Database
	storage, err := mysql.New(cfg)
	if err!=nil {
		log.Fatal(err)
	}
	slog.Info("Storage Initialized",slog.String("env",cfg.Env))

	//Setting up the router
	router:= http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))

	//Setup the Server
	server:= http.Server{
		Addr: cfg.HTTPServer.Address,
		Handler: router,
	}
	slog.Info("Server Listening at",slog.String("address",cfg.HTTPServer.Address))
	// fmt.Printf("Server Listening at %s\n",server.Addr)

	channel_shutdown:= make(chan os.Signal, 1)
	signal.Notify(channel_shutdown,os.Interrupt,syscall.SIGINT,syscall.SIGABRT,syscall.SIGTERM)
	go func ()  {
		err:= server.ListenAndServe() // without goroutine i cannot graceful shutdown
		if err!= nil && err !=http.ErrServerClosed{
			log.Fatalf("Failed to Serve: %s",err)
	}
	}()
	<- channel_shutdown

	//Stopping the server
	slog.Info("Server Stopping..")
	ctx, cancel:= context.WithTimeout(context.Background(),time.Second * 5)
	err = server.Shutdown(ctx)
	if err!= nil{
		slog.Error("Failed to Shut Down",slog.String("Error",err.Error()))
	}
	defer cancel()
	slog.Info("Server Shutdown Successfully")
}