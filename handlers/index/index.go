package index

import (
	"fmt"

	"github.com/gocraft/web"
)

func Handle(rw web.ResponseWriter, r *web.Request) {
	rw.Header().Add("Content-type", "text/html")

	fmt.Fprint(rw, `<form action="/login" method="post"><input type="text" name="email"><input type="submit"></form>`)
}
