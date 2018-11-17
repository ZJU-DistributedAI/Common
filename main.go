//go:generate goagen bootstrap -d github.com/ZJU-DistributedAI/Common/design
//go:generate goagen -d github.com/ZJU-DistributedAI/Common/design swagger -o public
//go:generate goagen -d github.com/ZJU-DistributedAI/Common/design schema -o public
//go:generate go-bindata -ignore 'bindata.go' -pkg swagger -o public/swagger/bindata.go ./public/swagger/...
//go:generate echo Start copying the swagger-ui dist resources
//go:generate cp -a swagger-ui-dist/. public/swagger
//go:generate echo Finished copying the swagger-ui dist resources

package main

import (
	"github.com/ZJU-DistributedAI/Common/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("common services")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount controllers
	storageController := NewStorageController(service)
	app.MountStorageController(service, storageController)

	schemaController := NewSchemaController(service)
	app.MountSchemaController(service, schemaController)

	swaggerController := NewSwaggerController(service)
	app.MountSwaggerController(service, swaggerController)

	// Start service
	if err := service.ListenAndServe(":3001"); err != nil {
		service.LogError("startup", "err", err)
	}

}
