package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/itnikolaev/quick_start_go/http"
)

var port = flag.String("port", ":8881", "listen port")

func main() {
	var err = http.StartServer(*port)
	if err != nil {
		panic(err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	for {
		select {
		case killSignal := <-interrupt:
			fmt.Println("Got signal:", killSignal)
			fmt.Println("Stoping listening on ", *port)
			if killSignal == os.Interrupt {
				fmt.Println("Daemon was interruped by system signal")
				return
			}
			return
		}
	}
}
