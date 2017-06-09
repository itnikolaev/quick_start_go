package http

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gocraft/web"
	"github.com/itnikolaev/quick_start_go"
	"github.com/itnikolaev/quick_start_go/handlers/index"
	"github.com/itnikolaev/quick_start_go/handlers/login"
)

func StartServer(port string) error {
	router := web.New(quickstart.HttpContext{})
	router.Get("/", index.Handle)
	router.Get("/login", login.HandleGet)
	router.Post("/login", login.HandlePost)

	serv := &http.Server{Addr: port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	serv.SetKeepAlivesEnabled(false)

	listener, err := net.Listen("tcp", serv.Addr)
	if err != nil {
		return err
	}

	go serv.Serve(listener)
	fmt.Println("server started")
	return nil
}
