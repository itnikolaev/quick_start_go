package login

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gocraft/web"
	"github.com/itnikolaev/quick_start_go"
)


func HandlePost(rw web.ResponseWriter, r *web.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	if email == "" {
		http.Redirect(rw, r.Request, "/", 301)
		return
	}
	key := fmt.Sprintf("%x", sha256.Sum256([]byte(email+time.Now().String())))

	sess := new(quickstart.Session)
	sess.Email = email
	sess.ID = rand.Int63()
	sess.Setting1 = 3
	sess.Setting2 = "golang"
	sess.Setting3 = rand.Float64()
	sess.Setting4 = nil

	sStorage.Set(key, sess)
	http.SetCookie(rw, &http.Cookie{Name: "SESSID", Value: key})
	http.Redirect(rw, r.Request, "/login", 301)
}
