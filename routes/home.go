package routes

import (
	"github.com/martini-contrib/render"
)

func IndexHandler(rnd render.Render) {
	rnd.HTML(200, "index", nil)
}

func SignUp(rnd render.Render) {
	rnd.HTML(200, "signUp", nil)
}
