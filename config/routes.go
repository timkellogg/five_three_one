package config

import (
	"github.com/timkellogg/five_three_one/app/controllers"
	"github.com/timkellogg/five_three_one/app/middlewares"
)

var routes = Routes{
	Route{"Info", "GET", "/api/info", middlewares.SetHeaders(controllers.InfoShow)},
}
