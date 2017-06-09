package login

import (
	"github.com/itnikolaev/quick_start_go"
	"github.com/itnikolaev/quick_start_go/sessions"
)

var sStorage quickstart.SessionStorage

func init() {
	sStorage = sessions.NewStorage()
}
