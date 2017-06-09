package login

import (
	"github.com/gocraft/web"
	"fmt"
	"encoding/json"
	"net/http"
)

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
