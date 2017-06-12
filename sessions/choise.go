package sessions

import (
	"fmt"
	"os"

	"github.com/itnikolaev/quick_start_go"
	"github.com/itnikolaev/quick_start_go/sessions/memory"
)

const storageType string = "memory" // maybe get it from config?

func NewStorage() quickstart.SessionStorage {
	switch storageType {
	case "memory", "mem":
		return memory.NewStorage()
	case "redis":
		return nil
	default:
		fmt.Println("storage incorrect!")
		os.Exit(1)
		return nil
	}
}
