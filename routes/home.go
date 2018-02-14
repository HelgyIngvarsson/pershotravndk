package routes

import (
	"github.com/martini-contrib/render"
)

func IndexHandler(rnd render.Render) {
	rnd.HTML(200, "index", nil)
}
