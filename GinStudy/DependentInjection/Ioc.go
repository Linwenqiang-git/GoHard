package dependentinjection

import (
	"log"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
)

func ConfigureService(app *gin.Engine) {
	//controller declare
	var index controller.Index
	//inject declare
	db := datasource.Db{}
	//Injection
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: &index},
		&inject.Object{Value: &db},
		&inject.Object{Value: &repository.StartRepo{}},
		&inject.Object{Value: &service.StartService{}},
	)
	if err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("inject fatal: ", err)
	}

	//database connect
	err = db.Connect()
	if err != nil {
		log.Fatal("db fatal:", err)
	}
	v1 := app.Group("/")
	{
		v1.GET("/get/:msg", index.GetName)
	}
}
