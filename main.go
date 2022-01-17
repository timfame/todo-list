package main

import (
	"github.com/timfame/todo-list/internal/delivery/http"
	"github.com/timfame/todo-list/internal/service"
	"log"
)

func main() {
	todoListService := service.New()
	httpServer, err := http.NewServer(todoListService)
	if err != nil {
		log.Fatalln("Creating http server failed", err)
	}
	log.Println("Starting server...")
	log.Println("Error:", httpServer.Run())
}
