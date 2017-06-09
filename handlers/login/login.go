package login

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gocraft/web"
	"github.com/itnikolaev/quick_start_go"
	"github.com/itnikolaev/quick_start_go/sessions"
)

var sStorage quickstart.SessionStorage

func init() {
	sStorage = sessions.NewStorage()
}

func HandleGet(rw web.ResponseWriter, r *web.Request) {
	c, err := r.Cookie("SESSID")
	if err != nil {
		http.Redirect(rw, r.Request, "/", 301)
		return
	}
	sess := sStorage.Get(c.Value)
	if sess == nil {
		rw.Header().Add("Content-type", "text/plain")
		fmt.Fprintf(rw, "session incorrect")
		return
	}
	ba, err := json.Marshal(sess)
	if err != nil {
		rw.Header().Add("Content-type", "text/plain")
		fmt.Fprintf(rw, "json error: %s", err)
		return
	}
	rw.Header().Add("Content-type", "application/json")
	fmt.Fprintf(rw, "%s", ba)
}

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
