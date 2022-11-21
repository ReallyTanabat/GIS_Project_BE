package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gis-project-backend/cmd/api/core/api"
	"gis-project-backend/cmd/api/handler"
	"gis-project-backend/docs"
	"gis-project-backend/pkg/core"
	"gis-project-backend/pkg/core/client"
	"gis-project-backend/pkg/core/utils"

	swagger "github.com/arsmn/fiber-swagger/v2"

	appLogger "gis-project-backend/pkg/core/logger"

	"github.com/gofiber/fiber/v2"

	//api
	quest4 "gis-project-backend/pkg/module/quest4/usecase"
	quest5 "gis-project-backend/pkg/module/quest5/usecase"

	_ "github.com/microsoft/go-mssqldb"
)

func Run() {

	//Log
	log := appLogger.NewLogger("./logs/app.log")
	//connectionString := fmt.Sprintf("user id=%s;password=%s;port=%s;database=%s", "SA", "@reallyStrongPwd123", "1433", "SQLDB")
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", "SA", "@reallyStrongPwd123", "172.17.0.2", "1433", "SQLDB")
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	log.Info("Connect DB Success")
	app := fiber.New(api.ServerConfig())

	//Swagger Docs
	if utils.IsNotEmpty(os.Getenv("SWAGGER_HOST")) {
		docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
		log.Info("Set Swagger host")
		fmt.Println(">>>>>>>>> Set Swagger host: ", docs.SwaggerInfo.Host)
	}
	if utils.IsNotEmpty(os.Getenv("BASE_PATH")) {
		docs.SwaggerInfo.BasePath = os.Getenv("BASE_PATH")
		log.Info("Set Swagger BasePath")
		fmt.Println("aster-data>>>>>>>>> Set Swagger BasePath: ", docs.SwaggerInfo.BasePath)
	}

	//Swagger Path
	app.Get("/swagger", swagger.HandlerDefault)
	app.Get("/swagger/*", swagger.HandlerDefault)

	//Core
	coreRegistry := &core.CoreRegistry{
		Logger:     log,
		RestClient: client.NewRestClient(client.ClientOption{Debug: false, Insecure: true, Timeout: 30 * time.Second}),
		DB:         db,
	}
	//Middleware
	api.RegisterMiddleware(app, coreRegistry)

	//Service

	quest4UseCase := quest4.NewQuest4UseCase(coreRegistry)
	quest5UseCase := quest5.NewQuest5UseCase(coreRegistry)

	handler.NewQuest4APIHandler(app, coreRegistry, quest4UseCase).Init()
	handler.NewQuest5APIHandler(app, coreRegistry, quest5UseCase).Init()

	hostname := "localhost"
	port := "9090"

	log.Info("Starting...")
	log.Info("Swagger URL : http://" + hostname + ":" + port + "/swagger/")

	err = app.Listen(":" + port)
	if err != nil {
		log.Error(err.Error())
		panic(err.Error())
	}

}
