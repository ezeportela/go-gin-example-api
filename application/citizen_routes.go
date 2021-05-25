package application

import (
	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/controllers"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/gin-gonic/gin"
)

func RegisterCitizenController(r *gin.Engine, conf config.Config) {
	ctrl := controllers.CitizenController{}

	r.POST(shared.MakeRoute(conf.BasePath, "/citizen"), ctrl.CreateHandler())

	r.GET(shared.MakeRoute(conf.BasePath, "/citizen/:id"), ctrl.GetByIdHandler())

	r.POST(shared.MakeRoute(conf.BasePath, "/citizen/:id"), ctrl.UpdateHandler())

	r.POST(shared.MakeRoute(conf.BasePath, "/citizen/filter"), ctrl.FilterHandler())

	r.DELETE(shared.MakeRoute(conf.BasePath, "/citizen/:id"), ctrl.DeleteHandler())

	r.POST(shared.MakeRoute(conf.BasePath, "/citizen/batch"), ctrl.InsertManyHandler())
}
